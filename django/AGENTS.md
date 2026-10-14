# Django starter kit

A minimal [Django](https://www.djangoproject.com) app on Python 3.14, managed with [uv](https://docs.astral.sh/uv/). It has one project, one app with one page, and one test. It's a starting point: build the app the user asks for on top of it.

## Commands

| Task | Command |
|---|---|
| Install dependencies | `uv sync` |
| Run migrations locally | `uv run python manage.py migrate` |
| Run locally | `uv run python manage.py runserver 3000` |
| Run the tests | `uv run pytest` |
| Lint | `uv run ruff check .` |
| Check formatting | `uv run ruff format --check .` |
| Fix lint and formatting | `uv run ruff check --fix . && uv run ruff format .` |
| Run Django's system checks | `uv run python manage.py check` |

Run the tests, the linter, the formatting check, and `manage.py check` after every change. All must pass.

## Structure

- `config/` is the Django project: `settings.py`, `urls.py`, and the WSGI and ASGI entry points.
- `pages/` is the one app. `pages/views.py` holds the home view, and `pages/templates/pages/` its template.
- `tests/` holds pytest tests, run with pytest-django.
- `pyproject.toml` lists dependencies, and `uv.lock` pins them.

## Conventions

- Add features as new apps with `uv run python manage.py startapp <name>`, add them to `INSTALLED_APPS`, and include their URLs in `config/urls.py`.
- Keep templates in each app's `templates/<app>/` folder.
- Read any setting that differs between environments from `os.environ` in `config/settings.py`, with a default that works locally. Never commit real secrets.
- Create migrations with `uv run python manage.py makemigrations` after changing models, and commit them.
- Add dependencies with `uv add` (or `uv add --dev` for tools) so `uv.lock` stays in sync, and commit the lockfile.
- Don't add a `[build-system]` to `pyproject.toml`. The app runs from the repository and isn't installed as a package.

## On Laravel Cloud

- Cloud detects Django and starts the app with `gunicorn config.wsgi --bind [::]:$PORT`, using `WSGI_APPLICATION` from the settings. Cloud installs the dependencies from `uv.lock`, runs `collectstatic` during the build, and runs `python manage.py migrate` as the deploy command.
- The Python version comes from `.python-version`, read on every deploy.
- Cloud sets `SECRET_KEY`, `DEBUG=false`, `ALLOWED_HOSTS`, `CSRF_TRUSTED_ORIGINS`, `STATIC_ROOT`, `STATIC_URL`, `WEB_CONCURRENCY`, and `APP_URL`. The settings already read all of them.
- whitenoise serves static files, including the admin's.
- **The database is SQLite, and Cloud's filesystem is ephemeral.** Anything saved, including admin users, is lost on the next deploy or restart. Before the app stores real data, attach a Postgres or MySQL database in Cloud, add `dj-database-url` and a driver (`psycopg[binary]` or `mysqlclient`), and build `DATABASES` from `DATABASE_URL`.

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
