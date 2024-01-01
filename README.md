# Requirement
- Go
- Docker
- Docker Compose
- [Task](https://taskfile.dev/)

# Usage
1. create `.env` file
   1. `cp .env.example .env`
2. `task up`
   - or `task up-a` (no detach option)
3. `curl localhost:${PORT}/api/users/:username`
