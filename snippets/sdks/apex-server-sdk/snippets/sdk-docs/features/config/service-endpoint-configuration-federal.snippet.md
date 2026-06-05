---
id: apex-server-sdk/sdk-docs/features/config/service-endpoint-configuration-federal
sdk: apex-server-sdk
kind: reference
lang: bash
description: Service endpoint configuration example for Apex.

---

```bash
cd bridge && go build .

# other required export statements...

export LD_BASE_URI='https://sdk.launchdarkly.us'
export LD_EVENTS_URL='https://events.launchdarkly.us'

./bridge
```
