.PHONY: watch
watch:
	${GOPATH}/bin/CompileDaemon \
		-exclude-dir=".git" \
		-exclude-dir=data -exclude=".#*" \
		-recursive=true \
	 	-build="go build -o ./build/main ./cmd/main.go" \
	 -command="./build/main"

.PHONY: postgresup postgresdown createdb dropdb migrateup migratedown
postgresup:
	docker-compose up -d
postgresdown:
	docker-compose down
createdb:
	docker exec -it  basic-api-postres createdb --username=postgres --owner=postgres basic-api
dropdb:
	docker exec -it  basic-api-postres dropdb --username=postgres  basic-api
migrateup:
	migrate -path internal/storage/migrations -database "postgresql://postgres:postgres@localhost:5432/basic-api?sslmode=disable"   -verbose up 
migratedown:
	migrate -path  internal/storage/migrations -database "postgresql://postgres:postgres@localhost:5432/basic-api?sslmode=disable" -verbose down 

.PHONY: test
test:
	go test -v --race ./...