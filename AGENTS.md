# Laravel Cloud starter kits

This repository holds the starter kits, one per top-level folder. Each kit is a standalone app: Cloud copies a kit's folder into a new repository for the customer, so a kit must never depend on files outside its folder.

## Rules for every kit

- **No settings on Cloud.** A kit must deploy with Cloud's detected defaults: build command, start command, and version. If a kit needs a setting, treat it as a gap in Cloud's detection and fix the detection, not the kit.
- **Minimal.** One page or route and one test. No auth, database, queues, UI kit, CI workflows, or Docker files.
- **Cloud-ready.** Listen on `PORT` (default 3000) on host `::`, read configuration from environment variables, declare the language version in the repository, and commit the lockfile.
- **Standard commands.** Use each language's usual tools for running, testing, linting, and formatting, and list them in the kit's `AGENTS.md`.
- **Agent docs.** Each kit has an `AGENTS.md` with what it is, its commands, structure, conventions, and how it runs on Cloud, plus a `CLAUDE.md` containing `@AGENTS.md`, a short `README.md`, and a `.env.example`.

## The shared Cloud section

`tooling/cloud.md` is the Cloud section shared by every server kit. `tooling/sync-agents` copies it into each kit's `AGENTS.md` between the `<!-- cloud:start -->` and `<!-- cloud:end -->` markers. Edit `tooling/cloud.md`, never the copies, then run:

```sh
tooling/sync-agents          # update every kit
tooling/sync-agents --check  # fail if any kit is out of date
```

Format `tooling/cloud.md` with Prettier (the root `.prettierrc`) before syncing, so the JavaScript kits' format checks pass. The Astro kit is static and has its own Cloud section, without the markers.

## The welcome screen

`tooling/welcome/welcome.html` is the welcome screen every kit serves at `/`, written with Tailwind classes. `tooling/sync-welcome` renders it into each kit's templating:

- Kits that build Tailwind CSS on Cloud's defaults (Next.js, Nuxt, Astro, Rails) get the markup inside their own layout: JSX, a Vue page, an Astro page, and an ERB view.
- Kits with no frontend build on Cloud (Django, FastAPI, Go, Express) get a full HTML page that loads Tailwind's browser build from a CDN: a Django template, Go's `html/template`, or an HTML file.

Edit `tooling/welcome/welcome.html`, never the generated files, then run:

```sh
tooling/sync-welcome          # update every kit
tooling/sync-welcome --check  # fail if any kit is out of date
```

The generated files are excluded from Prettier in the JavaScript kits, so formatting doesn't fight the generator. Keep each text and its link on one line in the source: JSX and Astro drop whitespace at line breaks next to an expression.

## Publishing

Each kit is published to its own read-only template repository, `laravel/cloud-starter-<kit>`. On every push to `main`, `.github/workflows/split.yml` checks the generated files, runs each kit's tests and linters, and pushes the kit's folder history (`git subtree split`) to its repository. Pull requests run the same checks without publishing.

The workflow pushes with the `SPLIT_TOKEN` secret, which needs write access to the contents of every split repository. When adding a kit, create `laravel/cloud-starter-<kit>` as a public template repository with issues and pull requests off, then add the kit to the workflow.

Never put anything in a kit that only makes sense in this monorepo, such as contribution notes or CI workflows: the split repository is what customers copy.

## Checking a kit

Before committing, run the kit's test and lint commands from its `AGENTS.md` and start it the way Cloud does, checking that it answers on both `127.0.0.1` and `[::1]`.
