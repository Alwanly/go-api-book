run:
	go run main.go

test:
	go test -v ./domain/... ./infrastructure/...

dev:
	air

build:
	go build main.go	