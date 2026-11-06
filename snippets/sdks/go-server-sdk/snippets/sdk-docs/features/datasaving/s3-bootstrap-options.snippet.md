---
id: go-server-sdk/sdk-docs/features/datasaving/s3-bootstrap-options
sdk: go-server-sdk
kind: reference
lang: go
description: Options of the Amazon S3 data initializer for Go.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
import (
    "time"

    lds3 "github.com/launchdarkly/go-server-sdk-s3"
)

// Create the S3 client from a specific AWS configuration, and lower the
// default limits of 30 seconds for the fetch and 100 MiB for the object.
initializer := lds3.DataInitializerForObject("my-bucket", "flags/snapshot.json").
    AWSConfig(awsConfig).
    Timeout(10 * time.Second).
    MaxObjectSizeBytes(10 * 1024 * 1024)

// Or supply a preconfigured client, for example for an S3-compatible service.
initializer = lds3.DataInitializerForObject("my-bucket", "flags/snapshot.json").
    S3Client(myS3Client)
```
