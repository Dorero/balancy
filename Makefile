include .env

BINARY_NAME=balancy

default: build

build:
	GOARCH=amd64 GOOS=linux go build -o ./target/$(BINARY_NAME) src/main.go

insert_tables:
	PGPASSWORD=${DB_PASSWORD} psql -h db -p 5432 -U ${DB_USER} -d ${DB_NAME} -f schema.sql

clean:
	go clean
	rm ./target/${BINARY_NAME}

run:
	./target/${BINARY_NAME}

