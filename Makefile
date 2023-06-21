infra-up:
	docker-compose up -d

infra-down:
	docker-compose stop

migrate-create:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up

migrate-destroy:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose down

run:
	cd cmd/app/ && go run main.go wire_gen.go

# PHONY is used to tell make that these are not files
.PHONY: infra-start infra-stop migrate-create migrate-destroy run