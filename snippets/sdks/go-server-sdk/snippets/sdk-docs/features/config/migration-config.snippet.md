---
id: go-server-sdk/sdk-docs/features/config/migration-config
sdk: go-server-sdk
kind: reference
lang: go
description: Migration configuration example for the Go SDK v7 — read/write methods, execution order, latency/error tracking.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go

client, _ := ld.MakeClient("YOUR_SDK_KEY", 5*time.Second)

var comparison ld.MigrationComparisonFn
comparison = func(interface{}, interface{}) bool {
	// compare the two read values
	return true
}

migrator, _ := ld.Migration(client).
	Read(
		func(interface{}) (interface{}, error) {
			return "old read", nil
		},
		func(interface{}) (interface{}, error) {
			return "new read", nil
		},
		&comparison,
	).
	ReadExecutionOrder(ldmigration.Random).
	Write(
		func(interface{}) (interface{}, error) {
			return "old write result", nil
		},
		func(interface{}) (interface{}, error) {
			return "new write result", nil
		},
	).
	TrackLatency(true).
	TrackErrors(true).
	Build()
```
