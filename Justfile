default:
    @just --list

test:
    go test -v -cover ./...

build:
    go build -ldflags="-s -w"
