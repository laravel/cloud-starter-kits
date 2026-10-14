# FastAPI starter kit

A minimal [FastAPI](https://fastapi.tiangolo.com) app on Python 3.14, managed with [uv](https://docs.astral.sh/uv/). It has one route and one test. It's a starting point: build the app the user asks for on top of it.

## Commands

| Task | Command |
|---|---|
| Install dependencies | `uv sync` |
| Run locally, reloading on changes | `uv run uvicorn app.main:app --reload --port 3000` |
| Run the tests | `uv run pytest` |
| Lint | `uv run ruff check .` |
| Check formatting | `uv run ruff format --check .` |
| Fix lint and formatting | `uv run ruff check --fix . && uv run ruff format .` |

Run the tests, the linter, and the formatting check after every change. All three must pass.

## Structure

- `app/main.py` creates the FastAPI app as `app` and registers routes. Cloud starts this module, so keep the app object there and named `app`.
- `tests/` holds pytest tests, which call the app through FastAPI's `TestClient`.
- `pyproject.toml` lists dependencies, and `uv.lock` pins them.

## Conventions

- Type every route's parameters and return value. FastAPI uses the types for validation and the OpenAPI docs at `/docs`.
- Use Pydantic models for request and response bodies.
- When `app/main.py` grows, move groups of routes into `app/routers/` as `APIRouter` modules and include them in `app/main.py`.
- Add dependencies with `uv add` (or `uv add --dev` for tools) so `uv.lock` stays in sync, and commit the lockfile.
- Don't add a `[build-system]` to `pyproject.toml`. The app runs from the repository and isn't installed as a package.

## On Laravel Cloud

- Cloud detects FastAPI and starts the app with `uvicorn app.main:app --host '' --port $PORT`. Cloud installs the dependencies from `uv.lock` during the build.
- The Python version comes from `.python-version`, read on every deploy.
- Cloud sets `WEB_CONCURRENCY` to one worker per CPU core, and `APP_URL` to the environment's URL.

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

Use the Laravel Cloud MCP server to deploy, read logs, change environment variables, and attach resources. Cloud's documentation for agents is at https://laravel.com/cloud/docs/llms.txt, and runtime details are at https://laravel.com/cloud/docs/runtimes.
<!-- cloud:end -->
