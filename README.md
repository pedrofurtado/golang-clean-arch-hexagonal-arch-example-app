# golang-clean-arch-hexagonal-arch-example-app

Example app with Clean Architecture and Hexagonal Architecture. Made in golang.

First setup at all

```bash
docker container run -w /app -v $(pwd):/app --rm -it golang:1.22.2 go mod init my-app
```

Localhost run

```bash
docker container run -p 3000:80 -w /app -v $(pwd):/app --rm -it golang:1.22.2 go run cmd/web/main.go
```

Install new dependency

```bash
docker container run -w /app -v $(pwd):/app --rm -it golang:1.22.2 go get mysite.com/mypkg
```
