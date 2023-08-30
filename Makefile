create_migration:
	migrate create -ext=sql -dir=sql/migrations -seq $(name)

migrate-up:
		migrate -path=sql/migrations -database "postgresql://postgres:postgres@localhost:5432/db?sslmode=disable" -verbose up

migrate-down:
		migrate -path=sql/migrations -database "postgresql://postgres:postgres@localhost:5432/db?sslmode=disable" -verbose down

migrate-force:
		migrate -path=sql/migrations -database "postgresql://postgres:postgres@localhost:5432/db?sslmode=disable" -verbose force $(version)

migrate-goto:
		migrate -path=sql/migrations -database "postgresql://postgres:postgres@localhost:5432/db?sslmode=disable" -verbose goto $(version)

mock_repository: 
	mockgen -source=internal/infra/repository/repository_interface.go -destination=internal/infra/repository/mocks/repository.go -package=mock

air:
		docker-compose up air

test:
	GET_ENV_FILE=true go test -v -cover ./... -coverprofile=coverage.out -covermode=count && go tool cover -func=coverage.out

test_short:
	GET_ENV_FILE=true go test -short -v -cover ./... -coverprofile=coverage.out -covermode=count && go tool cover -func=coverage.out

.PHONY: create_migration migrateup migratedown mock_repository air test test_short