dburl="postgresql://siavash:123@127.0.0.1:5432/bgmood_circles?sslmode=disable"

postgres:
	docker run --name postgres15 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123 -p 5432:5432 -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root bgmood_circles

mu:
	migrate -path internal/db/migration -database $(dburl) -verbose up

mu1:
	migrate -path internal/db/migration -database $(dburl) -verbose up 1

md:
	migrate -path internal/db/migration -database $(dburl) -verbose down

md1:
	migrate -path internal/db/migration -database $(dburl) -verbose down 1

mr:
	make md && make mu

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

proto:
	rm -rf internal/cpb/* 2>/dev/null
	protoc --go_out=internal/cpb --go_opt=paths=source_relative --go-grpc_out=internal/cpb --go-grpc_opt=paths=source_relative --grpc-gateway_out=internal/cpb --grpc-gateway_opt=paths=source_relative --proto_path=../bgmood-protos/circles-service ../bgmood-protos/circles-service/*.proto

.PHONY: postgres createdb mu mu1 md md1 mr sqlc test server proto