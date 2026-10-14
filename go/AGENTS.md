# Go starter kit

A minimal Go web app using only the standard library's `net/http`, on Go 1.27. It serves a welcome screen at `/` and a JSON example at `/api`, with a test for each. It's a starting point: build the app the user asks for on top of it.

## Commands

| Task | Command |
|---|---|
| Run locally | `go run .` |
| Run the tests | `go test ./...` |
| Vet | `go vet ./...` |
| Check formatting (lists unformatted files) | `gofmt -l .` |
| Fix formatting | `gofmt -w .` |
| Build like Cloud does | `CGO_ENABLED=0 go build -o app .` |

Run the tests, `go vet`, and the formatting check after every change. All must pass, and `gofmt -l .` must print nothing.

## Structure

- `main.go` reads `PORT` and starts the server on host `::`. Keep it this small.
- `routes.go` registers routes on an `http.ServeMux` and holds their handlers.
- `welcome.html` is the Laravel Cloud welcome screen, embedded with `go:embed` and rendered by `html/template` at `/`. Replace it with the app's own home page.
- `routes_test.go` tests handlers through `routes()` with `net/http/httptest`.
- The module is called `app`. Rename it in `go.mod` to the repository's import path if the app grows into multiple packages that import each other.

## Conventions

- Register routes with method and path patterns, such as `mux.HandleFunc("GET /users/{id}", showUser)`, and read path values with `r.PathValue("id")`. `"GET /{$}"` matches only `/`.
- Keep the app at the repository root, in package `main`. When it grows, put internal packages under `internal/`.
- Prefer the standard library. Add a dependency only when it saves real work, with `go get`, and commit `go.mod` and `go.sum`.
- Don't use cgo unless a dependency needs it. Cloud builds with `CGO_ENABLED=0` unless it detects cgo.

## On Laravel Cloud

- Cloud detects Go, builds the repository root into a single binary with `CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o app .`, and starts it with `./app`.
- The Go version comes from the `go` directive in `go.mod`, read on every deploy. Cloud uses the lowest supported version that satisfies it, so keep it at `go 1.27` or later.
- Cloud sets `APP_URL` to the environment's URL.

<!-- cloud:start -->

## Laravel Cloud

This app deploys to [Laravel Cloud](https://laravel.com/cloud) with no settings changed. Cloud detects the language and framework from the repository and fills in the build and start commands. Keep the app inside these rules so that stays true.

### Rules

- **Port:** listen on the port in the `PORT` environment variable (3000 if it's unset), on host `::` so both IPv4 and IPv6 connect. Cloud runs nginx on port 8080 in front of the app. Never hard-code a port.
- **Configuration:** read every setting from environment variables. In production Cloud sets them in the process environment; there's no `.env` file. `.env.example` lists what the app reads, for local development only.
- **No Dockerfile:** Cloud builds the image from the repository. Don't add a `Dockerfile`, `docker-compose.yml`, or `Procfile` unless this file says the kit uses one.
- **Ephemeral filesystem:** files written at runtime are lost on every deploy and restart, and replicas don't share them. Store data in a database, cache, or object storage.
- **Scale to zero:** an environment can sleep when idle and wake on the next request. Don't rely on in-memory state or timers in the web process.
- **Versions:** the language version is declared in the repository (see below). Change it there, and only to a version Cloud supports. An unsupported version fails the deploy.

### Adding resources

Attach resources to the environment in Cloud. Cloud injects their connection details as environment variables on the next deploy:

| Resource                   | Variables                                                                                    |
| -------------------------- | -------------------------------------------------------------------------------------------- |
| Postgres or MySQL database | `DATABASE_URL`                                                                               |
| Valkey cache               | `REDIS_URL`                                                                                  |
| Object storage bucket      | `AWS_BUCKET`, `AWS_ENDPOINT_URL`, `AWS_REGION`, `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY` |

The Postgres `DATABASE_URL` ends in a `serverVersion` query parameter. Some drivers reject unknown parameters, so remove it before connecting if the driver complains.

Run background work as a custom worker background process on the environment, with its own command. Run migrations in the environment's deploy command, not at app startup.

### Managing the app

Use the [Laravel Cloud CLI](https://laravel.com/cloud/docs/api/cli) to deploy, read logs, change environment variables, and attach resources. It needs PHP 8.3 or later and Composer. Install it with `composer global require laravel/cloud-cli`, then run `cloud skills:install` to add Laravel Cloud skills to your agent. Run `cloud auth`, or set `LARAVEL_CLOUD_TOKEN` where a browser login isn't possible, and pass `-n` to every command so it never waits for input. A Laravel Cloud MCP server is coming and will cover the same tasks.

Cloud's documentation for agents is at https://laravel.com/cloud/docs/llms.txt, and runtime details are at https://laravel.com/cloud/docs/runtimes.
<!-- cloud:end -->
