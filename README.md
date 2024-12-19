### Local development
- Prerequisites:
    - `go` version 1.23.4
    - `make`
- Make a copy of [`.env.example`](.env.example) named `.env` and update with an API key.
- Run locally with `make run`.
- Perform adhoc tests with `make test-curl`. This POSTs the JSON in [example_request.json](example_request.json)

### Behaviour
This extends the core API by caching the data. This is required by low-latency clients.

Two flags can be specified to tune the behaviour depending on a client's needs:
- `ensurePresent` - This is ensures that a value is always returned. This means that if there is no value in the cache, then the API will wait for a response from the core API. Note, this will lead to higher latencies for calls that requiring blocking requests to the core API.
- `ensureNotStale` - This ensures that stale data is never returned to the client. If there is no data, empty data is returned.

**Summary or caching behaviour:**
| In cache? | Stale? | `ensurePresent`   | `ensureNotStale`  | Behaviour                                 |
| --------- | ------ | ----------------- | ----------------- | ----------------------------------------- |
| No        | N/A    | `false`           | `true` or `false` | Return Empty data                         |
| No        | N/A    | `true`            | `true` or `false` | Return good data (Blocking refresh)       |
| Yes       | No     | `true` or `false` | `true` or `false` | Return good data                          |
| Yes       | Yes    | `true` or `false` | `false`           | Return stale data (refresh in background) |
| Yes       | Yes    | `false`           | `true`            | Return empty data (refresh in background) |
| Yes       | Yes    | `true`            | `true`            | Return good data (Blocking refresh)       |

### Manual testing
TODO