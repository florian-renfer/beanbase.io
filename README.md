# beanbase

> [!IMPORTANT]
> beanbase is currently work in progress!

## 🚦 Running the Project

1. `git clone git@github.com:florian-renfer/beanbase.io.git "beanbase"`
2. `cd beanbase`
3. `go mod download`
4. `go run cmd/api/main.go`

## 🏎️ Roadmap

- [x] `docker-compose.yaml` for PostgreSQL
- [x] Database migrations for PostgreSQL using [migrate](https://github.com/golang-migrate/migrate)
- [ ] Basic REST API using `net/http` performing CRUD operations on `coffee_roasters` table

## 🧰 Technologies

- Go
- golang-migrate
- PostgreSQL
