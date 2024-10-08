doc:
	swag init
mock:
	go generate ./...
test:
	go test ./...
vendor:
	go mod vendor
build:vendor
	go build -o bin/challenge -mod=vendor main.go
run:build
	bin/challenge
compose/build:vendor
	docker compose build
compose/buildup:compose/build compose/up
compose/up:
	docker compose up
compose/down:
	docker compose down