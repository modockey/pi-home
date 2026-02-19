# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Raspberry Pi home management system with 4 services: nginx reverse proxy, a home management API (Rust), an IP registration service (Rust), and a monthly photo contest app (Go backend + React frontend). Deployed directly on a Raspberry Pi using systemd and PM2.

## Build & Release Commands

```bash
# Release all services at once
make release-all

# Individual service releases (each builds + deploys via systemd/PM2)
make -C ./nginx release
make -C ./server/api release
make -C ./server/ip_register release
make -C ./photo-award release        # delegates to frontend
```

### Rust services (server/api, server/ip_register)
```bash
cargo build --release                 # build
cargo test                            # run tests
make generate                         # regenerate OpenAPI code (server/api only)
```

### Photo Award frontend (photo-award/frontend)
```bash
npm run dev                           # dev server
npm run build                         # production build (tsc + vite)
npm run lint                          # ESLint
npm run init-release                  # first deploy (PM2 start)
npm run release                       # subsequent deploys (build + PM2 restart)
```

### Photo Award backend (photo-award/server)
```bash
make -C ./photo-award/server run      # run on port 8000
make -C ./photo-award/server build    # build binary
swagger generate server -t gen -f ../swagger.yaml  # regenerate from OpenAPI spec
```

### Database
```bash
docker compose up -d                  # start PostgreSQL
make -C ./database impddl            # import DDL
make -C ./database impdml            # import DML
```

## Architecture

### Service Map

```
nginx (:50000/:50443)
  ├── /pi-home/api/v1/  → server/api (Rust, :50001) → PostgreSQL
  ├── /photo-award/     → photo-award/frontend (React, :50080 via PM2/serve)
  └── /                 → pukiwiki (:58080)

server/ip_register (Rust, systemd timer every 5min) → PostgreSQL, MYDNS API
photo-award/server (Go, :8000) ← swagger.yaml
```

### Key Service Details

- **server/api**: OpenAPI-generated Rust server using Hyper. Business logic in `src/usecase.rs`, DB access in `src/db.rs`, generated code in `src/server/` and `src/client/`. Uses Diesel ORM with PostgreSQL.
- **server/ip_register**: Standalone Rust service that checks external IP via inet-ip.info, updates MYDNS, and logs changes to `ipv4_history` table. Runs on a systemd timer.
- **photo-award/server**: Go server generated with go-swagger from `photo-award/swagger.yaml`. Custom handlers in `handler/`, auth in `authorizer/`, generated code in `gen/`.
- **photo-award/frontend**: React 18 + TypeScript + Vite + SWC. Base path is `/photo-award`. Served via `serve` managed by PM2.

### Database

PostgreSQL (user: `home-management`, db: `home-management`). Schema in `database/ddl.sql`. Rust services connect via `DATABASE_URL` env var (dotenv). Docker Compose runs the DB locally.

## Branch & PR Conventions

- Main branch: `main`
- Feature branches: `feature/#{Issue番号}` (e.g., `feature/#65`)
- Fix branches: `fix/` prefix
- All work tracked as GitHub Issues
- PR template requires: related Issue, PR summary, test details (Japanese)
