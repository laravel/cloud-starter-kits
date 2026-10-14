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
