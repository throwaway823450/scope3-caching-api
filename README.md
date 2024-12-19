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
To test this behaviour you can run `make test-curl-routine`. This calls the API at different intervals so we can observe the behaviour.

Note: the expiry time is set to 10 seconds.

#### First call (0s)
```
$ curl ...
{"rows":[{"totalEmissions":0},{"totalEmissions":67.22831134537275},{"totalEmissions":0},{"totalEmissions":405.2662236466297}]}
```
```
Refreshing 2 item(s) in the main thread
Refreshed bbc.com
Refreshed cnn.com
Refreshing 2 item(s) in the background
Refreshed nytimes.com
Refreshed washingtonpost.com
```
bbc.com and cnn.com have `ensurePresent=true` so they are refreshed on the main thread. Others are refreshed in the background.
#### Second call (5s)
```
$ curl ...
{"rows":[{"totalEmissions":66.09248372608445},{"totalEmissions":67.22831134537275},{"totalEmissions":86.81333219738443},{"totalEmissions":405.2662236466297}]}
```
```
NO LOGS
```
This time we call within the expiry time. All items are now in the cache, including those that were refreshed in the background.
####  Third call (20s)
```
$ curl ...
{"rows":[{"totalEmissions":66.09248372608445},{"totalEmissions":67.22831134537275},{"totalEmissions":0},{"totalEmissions":405.2662236466297}]}
```
```
Refreshing 1 item(s) in the main thread
Refreshed cnn.com
Refreshing 3 item(s) in the background
Refreshed nytimes.com
Refreshed bbc.com
Refreshed washingtonpost.com
```
All of the items are now stale. cnn.com has `ensurePresent=true` and `ensureNotStale=true` so that is refreshed in the main thread.

washingtonpost.com is the only one that that returns empty data as it doesn't want stale data, but it doesn't mind if it is not present.