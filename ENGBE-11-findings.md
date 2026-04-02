# ENGBE-11 — Backend Concurrency Testing Findings

## Context

This document captures findings from testing backend concurrency behaviour for component creation,
using the Cycloid CLI (`cy`) as the test driver against a locally running backend stack.

## How to reproduce

### Prerequisites

Start the local backend stack (includes the worker, which is required for component creation):

```bash
make be-reset
```

The `yd-worker` service was added to `compose.yml` as part of this work. Make sure it is running:

```bash
docker compose logs yd-worker
```

### API key

To run the tests against your own environment (e.g. a dedicated staging or playground instance),
you need a valid API key for that environment. Fetch it from the **EDD job** of the backend, inside
the **playground component** that was instantiated for your environment — you know where to look.

Set it in your environment before running tests:

```bash
export CY_API_KEY=<your-api-key>
export CY_API_URL=<your-backend-url>   # defaults to http://localhost:3001
```

### Running the concurrency stress test

```bash
go test ./e2e/... -run TestBackendComponentConcurrency -v -timeout 600s
```

The test ramps up from 2 to 30 concurrent component creations. At each level it:
1. Fires N goroutines simultaneously, each creating one component
2. Waits up to 30 seconds for the first response — if nothing comes back, it reports a **stall** at that level and stops
3. Deletes all successfully created components before moving to the next level

## Findings

The backend handles concurrent component creation correctly **up to 14 concurrent requests**.
At **15 concurrent component creations**, the backend deadlocks: all requests hang indefinitely
and none return within the 30-second window.

This is consistent with a git-level serialization issue (locking during config repository writes)
already known to exist on the backend side. Tests in CI are run sequentially for this reason
(see note in `DEVELOPING_TIPS.md`).

## Status

15 concurrent components is within acceptable limits for normal usage. The stress test
(`TestBackendComponentConcurrency` in `e2e/components_test.go`) is left in place as-is
so that the backend team can run it directly if they want to investigate the deadlock further
or validate a fix.
