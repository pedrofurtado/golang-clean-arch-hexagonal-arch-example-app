# golang-clean-arch-hexagonal-arch-example-app

Example app with Clean Architecture and Hexagonal Architecture. Made in golang.

First setup at all

```bash
docker container run -w /app -v $(pwd):/app --rm -it golang:1.22.2 go mod init my-app
```

Localhost run

```bash
docker-compose up --build -d
```

Execute migrations

```bash
# Create a new SQL migration
docker-compose exec app goose create my_migration_name_here sql

# Check SQL file migrations status
docker-compose exec app goose status

# Check last SQL migration executed
docker-compose exec app goose version

# Run SQL migrations (with fix related to out-of-order migration)
docker-compose exec app goose -allow-missing up
```

Install new dependency

```bash
docker-compose exec app go get mysite.com/mypkg
or
docker container run -w /app -v $(pwd):/app --rm -it golang:1.22.2 go get mysite.com/mypkg
```
