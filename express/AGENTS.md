# Express starter kit

A minimal [Express 5](https://expressjs.com) app written in TypeScript and run directly by Node.js 24, with no build step. It has one route and one test. It's a starting point: build the app the user asks for on top of it.

## Commands

| Task                                   | Command          |
| -------------------------------------- | ---------------- |
| Install dependencies                   | `npm install`    |
| Run locally, restarting on changes     | `npm run dev`    |
| Run the tests                          | `npm test`       |
| Type check, lint, and check formatting | `npm run lint`   |
| Fix formatting                         | `npm run format` |
| Start in production                    | `npm start`      |

Run `npm test` and `npm run lint` after every change. Both must pass.

## Structure

- `src/app.ts` builds the Express app and registers routes. It doesn't listen, so tests can import it.
- `src/server.ts` starts the app on `PORT` and host `::`. Keep it this small.
- `test/*.test.ts` holds the tests, run by the built-in `node:test` runner.

## Conventions

- Node runs the TypeScript files directly by stripping types, so use only erasable syntax: no `enum`, `namespace`, or parameter properties. Import local files with their `.ts` extension.
- Put routes in `src/app.ts` until it grows, then move groups of routes into `src/routes/` as `express.Router()` modules and mount them in `src/app.ts`.
- Test routes by starting the app on port 0 and calling it with `fetch`, as `test/app.test.ts` does.
- Express 5 passes errors from `async` handlers to the error handler, so there's no need to wrap them.
- Add dependencies with `npm install` so `package-lock.json` stays in sync, and commit the lockfile.

## On Laravel Cloud

- Cloud detects Express and starts the app with `node src/server.ts`, read from the `start` script. If you move the entry file, update the `start` script.
- The Node.js version comes from `engines.node` in `package.json` when the environment is created. After that it's set in the environment's settings in Cloud.
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

Use the Laravel Cloud MCP server to deploy, read logs, change environment variables, and attach resources. Cloud's documentation for agents is at https://laravel.com/cloud/docs/llms.txt, and runtime details are at https://laravel.com/cloud/docs/runtimes.
<!-- cloud:end -->
