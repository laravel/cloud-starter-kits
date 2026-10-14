# Next.js starter kit

A minimal [Next.js](https://nextjs.org) app using the App Router, TypeScript, and Tailwind CSS. It has one page, a welcome screen, and one test. It's a starting point: build the app the user asks for on top of it.

<!-- BEGIN:nextjs-agent-rules -->

## This is NOT the Next.js you know

This version has breaking changes — APIs, conventions, and file structure may all differ from your training data. Read the relevant guide in `node_modules/next/dist/docs/` (resolved from this file's directory; in monorepos the `next` package may not be visible from the repo root) before writing any code. Heed deprecation notices.

This block is written and re-added by `next dev` — verify at `node_modules/next/dist/server/lib/generate-agent-files.js`. Removing it from a diff only re-creates the uncommitted change; committing it with your work keeps the tree clean.

<!-- END:nextjs-agent-rules -->

## Commands

| Task                                   | Command          |
| -------------------------------------- | ---------------- |
| Install dependencies                   | `npm install`    |
| Run locally                            | `npm run dev`    |
| Run the tests                          | `npm test`       |
| Type check, lint, and check formatting | `npm run lint`   |
| Fix formatting                         | `npm run format` |
| Build for production                   | `npm run build`  |
| Start in production                    | `npm start`      |

Run `npm test`, `npm run lint`, and `npm run build` after every change. All three must pass.

## Structure

- `app/` holds routes, using the App Router. `app/page.tsx` is the Laravel Cloud welcome screen, styled by `app/welcome.css`, and `app/layout.tsx` is the root layout. Replace the page and delete `app/welcome.css` when building the app's own home page.
- `app/globals.css` imports Tailwind CSS.
- `public/` holds static files served from the site root.
- `tests/` holds Vitest tests, run in jsdom with React Testing Library.

## Conventions

- Components are Server Components unless they need state or browser APIs. Add `'use client'` only to the components that do.
- Put route handlers in `app/**/route.ts` and server-side mutations in Server Actions.
- Import from the project root with the `@/` alias.
- Read secrets only on the server. Only variables prefixed `NEXT_PUBLIC_` reach the browser, and they're fixed at build time.
- Add dependencies with `npm install` so `package-lock.json` stays in sync, and commit the lockfile.

## On Laravel Cloud

- Cloud detects Next.js, builds with `npm run build`, and starts with `npm start`. `next start` reads `PORT` and listens on IPv4 and IPv6 by itself.
- The Node.js version comes from `engines.node` in `package.json` when the environment is created. After that it's set in the environment's settings in Cloud.
- Cloud sets `NEXT_PUBLIC_SITE_URL` to the environment's URL.

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
