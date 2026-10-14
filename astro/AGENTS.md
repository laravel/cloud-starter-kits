# Astro starter kit

A minimal [Astro](https://astro.build) static site with Tailwind CSS. It has one page, a welcome screen, one layout, and one test, and builds to plain HTML in `dist/`. It's a starting point: build the site the user asks for on top of it.

## Commands

| Task                            | Command           |
| ------------------------------- | ----------------- |
| Install dependencies            | `npm install`     |
| Run locally                     | `npm run dev`     |
| Run the tests                   | `npm test`        |
| Type check and check formatting | `npm run lint`    |
| Fix formatting                  | `npm run format`  |
| Build the site into `dist/`     | `npm run build`   |
| Preview the built site          | `npm run preview` |

Run `npm test`, `npm run lint`, and `npm run build` after every change. All three must pass.

When you start the dev server yourself, use background mode, `npx astro dev --background`, and manage it with `npx astro dev stop`, `npx astro dev status`, and `npx astro dev logs`.

## Structure

- `src/pages/` holds pages. Each `.astro` or `.md` file becomes a route. `src/pages/index.astro` is the Laravel Cloud welcome screen; replace it with the site's own home page.
- `src/layouts/Layout.astro` is the HTML shell every page uses. It imports `src/styles/global.css`, which imports Tailwind CSS, compiled by the `@tailwindcss/vite` plugin in `astro.config.mjs`.
- `public/` holds files copied to the site root as they are.
- `tests/` holds Vitest tests, which render components with Astro's container API.

## Conventions

- Keep the site static. Don't add a server adapter or `output: 'server'`: Cloud serves this site as static files, and a server adapter would need a Node.js server instead.
- Put shared markup in components under `src/components/` and pass data through props.
- Use content collections in `src/content/` for Markdown content such as posts or docs.
- Add dependencies with `npm install` so `package-lock.json` stays in sync, and commit the lockfile.

Consult the Astro docs before working on related tasks:

- [Adding pages, dynamic routes, or middleware](https://docs.astro.build/en/guides/routing/)
- [Working with Astro components](https://docs.astro.build/en/basics/astro-components/)
- [Using React, Vue, Svelte, or other framework components](https://docs.astro.build/en/guides/framework-components/)
- [Adding or managing content](https://docs.astro.build/en/guides/content-collections/)
- [Adding styles or using Tailwind](https://docs.astro.build/en/guides/styling/)
- [Supporting multiple languages](https://docs.astro.build/en/guides/internationalization/)

## Laravel Cloud

This site deploys to [Laravel Cloud](https://laravel.com/cloud) with no settings changed. Cloud detects Astro, runs `npm run build`, and serves `dist/` as static files with nginx. No Node.js process runs in production. Keep the site inside these rules so that stays true.

- **Build time only:** environment variables are read during `npm run build` and baked into the HTML. Nothing runs on the server after that, so there are no runtime secrets, API routes, or form handlers. Use a separate Cloud application for anything that needs a server.
- **No Dockerfile:** Cloud builds the site from the repository. Don't add a `Dockerfile` or `docker-compose.yml`.
- **Output folder:** the site must build into `dist/`. Don't change `outDir`.
- **Versions:** the Node.js version used for the build comes from `engines.node` in `package.json` when the environment is created. After that it's set in the environment's settings in Cloud.

Use the [Laravel Cloud CLI](https://laravel.com/cloud/docs/api/cli) to deploy and read build logs. It needs PHP 8.3 or later and Composer. Install it with `composer global require laravel/cloud-cli`, then run `cloud skills:install` to add Laravel Cloud skills to your agent. Run `cloud auth`, or set `LARAVEL_CLOUD_TOKEN` where a browser login isn't possible, and pass `-n` to every command so it never waits for input. A Laravel Cloud MCP server is coming and will cover the same tasks.

Cloud's documentation for agents is at https://laravel.com/cloud/docs/llms.txt.
