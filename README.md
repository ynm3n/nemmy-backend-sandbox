# Usage
1. create `.env` file
   1. `mv .env.example .env`
2. `docker compose up`
3. `curl localhost:${PORT}/api/users/:userId`
   1. IDは0以上100未満の整数のみ指定可能
