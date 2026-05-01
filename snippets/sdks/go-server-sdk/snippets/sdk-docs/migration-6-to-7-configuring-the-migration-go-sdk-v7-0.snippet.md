---
id: go-server-sdk/sdk-docs/migration-6-to-7-configuring-the-migration-go-sdk-v7-0
sdk: go-server-sdk
kind: reference
lang: go
description: "Go SDK v7.0 in section \"Configuring the migration\""
---

```go

client, _ := ld.MakeClient("YOUR_SDK_KEY", 5*time.Second)

var comparison ld.MigrationComparisonFn
comparison = func(interface{}, interface{}) bool {
	// compare the two read values
	return true
}

migrator, err := ld.Migration(client).
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
