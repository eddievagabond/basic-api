### First go API

### Migrate

Requires golang migrate

```
migrate -path internal/storage/migrations -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up
```

```
migrate create -ext sql -dir internal/storage/migrations/ -seq products
migrate create -ext sql -dir internal/storage/migrations/ -seq products
```
