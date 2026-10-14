# Django starter kit

A minimal Django app on Python 3.14, ready to deploy to [Laravel Cloud](https://laravel.com/cloud) and built as a starting point for coding agents.

```sh
uv sync
uv run python manage.py migrate
uv run python manage.py runserver 3000
```

The app runs on http://localhost:3000. Run the tests with `uv run pytest`.

Agents: read [AGENTS.md](AGENTS.md) first.
