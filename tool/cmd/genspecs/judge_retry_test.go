package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sync/atomic"
	"testing"
	"time"
)

func TestIsRetryable(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false},
		{"context cancelled", context.Canceled, false},
		{"context deadline", context.DeadlineExceeded, false},
		{"4xx http", &retryableHTTPError{StatusCode: 403, Status: "403 Forbidden"}, false},
		{"400 http", &retryableHTTPError{StatusCode: 400, Status: "400"}, false},
		{"408 http", &retryableHTTPError{StatusCode: 408, Status: "408"}, true},
		{"429 http", &retryableHTTPError{StatusCode: 429, Status: "429"}, true},
		{"500 http", &retryableHTTPError{StatusCode: 500, Status: "500"}, true},
		{"503 http", &retryableHTTPError{StatusCode: 503, Status: "503"}, true},
		{"plain io.EOF", io.EOF, true},
		{"unexpected eof", io.ErrUnexpectedEOF, true},
		{"connection reset string", errors.New(`Post "...": read tcp 1.2.3.4->5.6.7.8: read: connection reset by peer`), true},
		{"i/o timeout string", errors.New(`Post "...": dial tcp 1.2.3.4: i/o timeout`), true},
		{"no such host", errors.New(`Get "...": dial tcp: lookup foo.example.com: no such host`), true},
		{"random error", errors.New("something completely different"), false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := isRetryable(tc.err); got != tc.want {
				t.Fatalf("isRetryable(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}

func TestParseRetryAfter(t *testing.T) {
	if got := parseRetryAfter("3"); got != 3*time.Second {
		t.Fatalf("got %v, want 3s", got)
	}
	if got := parseRetryAfter(""); got != 0 {
		t.Fatalf("got %v, want 0", got)
	}
	if got := parseRetryAfter("not-a-number"); got != 0 {
		t.Fatalf("got %v, want 0", got)
	}
}

// flakyJudge errors `failBefore` times, then succeeds on attempt failBefore+1.
type flakyJudge struct {
	failBefore int32
	calls      atomic.Int32
	err        error
	resp       JudgeResponse
}

func (f *flakyJudge) Model() string    { return "flaky" }
func (f *flakyJudge) Describe() string { return "flaky" }
func (f *flakyJudge) Judge(_ context.Context, _ PromptPack) (JudgeResponse, error) {
	n := f.calls.Add(1)
	if n <= f.failBefore {
		return JudgeResponse{}, f.err
	}
	return f.resp, nil
}

func TestRetryingJudgeRecoversFromTransient(t *testing.T) {
	flaky := &flakyJudge{
		failBefore: 2,
		err:        &retryableHTTPError{StatusCode: 503, Status: "503"},
		resp:       JudgeResponse{State: "supported", Confidence: "medium"},
	}
	// Use a tight retry loop so the test is fast.
	r := &retryingJudge{inner: flaky, maxAttempts: 4}
	resp, err := r.Judge(context.Background(), PromptPack{})
	if err != nil {
		t.Fatalf("expected success after retries, got %v", err)
	}
	if resp.State != "supported" {
		t.Fatalf("got state %q, want supported", resp.State)
	}
	if got := flaky.calls.Load(); got != 3 {
		t.Fatalf("expected 3 calls (2 fails + 1 success), got %d", got)
	}
}

func TestRetryingJudgeStopsAfterMaxAttempts(t *testing.T) {
	flaky := &flakyJudge{
		failBefore: 99, // always fails
		err:        &retryableHTTPError{StatusCode: 500, Status: "500"},
	}
	r := &retryingJudge{inner: flaky, maxAttempts: 3}
	_, err := r.Judge(context.Background(), PromptPack{})
	if err == nil {
		t.Fatal("expected error after exhausting retries")
	}
	if got := flaky.calls.Load(); got != 3 {
		t.Fatalf("expected exactly 3 attempts, got %d", got)
	}
}

func TestRetryingJudgeDoesNotRetryNonRetryable(t *testing.T) {
	flaky := &flakyJudge{
		failBefore: 99,
		err:        &retryableHTTPError{StatusCode: 403, Status: "403"},
	}
	r := &retryingJudge{inner: flaky, maxAttempts: 5}
	_, err := r.Judge(context.Background(), PromptPack{})
	if err == nil {
		t.Fatal("expected error")
	}
	if got := flaky.calls.Load(); got != 1 {
		t.Fatalf("expected single attempt for non-retryable error, got %d", got)
	}
}

func TestRetryingJudgeRespectsContext(t *testing.T) {
	flaky := &flakyJudge{
		failBefore: 99,
		err:        &retryableHTTPError{StatusCode: 500, Status: "500"},
	}
	r := &retryingJudge{inner: flaky, maxAttempts: 10}
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel up front
	_, err := r.Judge(ctx, PromptPack{})
	if err == nil {
		t.Fatal("expected error from cancelled context")
	}
}

// Sanity: ensure errors.As still finds the typed error after wrap-and-unwrap.
func TestRetryableHTTPErrorIsStructuredViaErrorsAs(t *testing.T) {
	var raw error = fmt.Errorf("decoded: %w", &retryableHTTPError{StatusCode: 503})
	var hErr *retryableHTTPError
	if !errors.As(raw, &hErr) {
		t.Fatal("expected errors.As to unwrap the typed error")
	}
	if hErr.StatusCode != 503 {
		t.Fatalf("got status %d, want 503", hErr.StatusCode)
	}
	if !hErr.retryable() {
		t.Fatalf("503 should be retryable")
	}
}
