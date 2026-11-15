# reelhouse

Cinema seat booking: hold a seat for 2 minutes, then confirm or release it.
Go (`net/http`) + Redis, with a single-page frontend in `static/`.

## Run locally

```sh
docker compose up --build app redis
# open http://localhost:8080
```

Without Docker for the app (Redis still required):

```sh
docker compose up -d redis
go run ./cmd
```

## Configuration

| Variable    | Default                  | Notes                                   |
|-------------|--------------------------|-----------------------------------------|
| `PORT`      | `8080`                   | Set automatically by most hosts         |
| `REDIS_URL` | `redis://localhost:6379` | Use `rediss://` for TLS (e.g. Upstash)  |

`GET /healthz` returns 200 when the app can reach Redis.

## Tests

```sh
go test ./...   # the concurrency test is skipped if Redis isn't reachable
```
