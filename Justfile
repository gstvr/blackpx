default:
    @just --list

test:
    go test -v -cover ./...
