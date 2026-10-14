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

## Checking a kit

Before committing, run the kit's test and lint commands from its `AGENTS.md` and start it the way Cloud does, checking that it answers on both `127.0.0.1` and `[::1]`.
