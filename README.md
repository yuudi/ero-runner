# Ero Runner

## Deploy

1. `sh docker/build.sh`
1. `docker network create ero-runner`
1. `cp .env.example .env`
1. Edit `.env`
1. `docker compose up -d`
1. (Optional) reverse proxy your port <8080>

## usage:

POST /api/v1/commands

Headers:
```headers
Content-Type: application/json
Authorization: Bearer your-secret-here
```

Body:
```json
{
  "user": "test001",
  "command": "run sh\nwhoami"
}
```
