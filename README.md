# HTTP Echo & Bridge Compatibility Microservice

This module exposes a minimal HTTP interface for echo testing and internal compatibility validation. It's designed to support evolving transport backends, including WebSocket bridges.

## Endpoints

- `POST /echo` — Echoes request body
- `GET /ping` — Returns a simple liveness probe (`200 OK`)
- `GET /ws` — WebSocket test endpoint (for future middleware compatibility)

## Usage

### Start the service

```bash
make run
