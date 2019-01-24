build:
	go build cmd/main.go

run:
	go run cmd/main.go

.PHONY: test
test:
	go test -cover github.com/jdrouet/marionette/cmd/git
	go test -cover github.com/jdrouet/marionette/cmd/model
	go test -cover github.com/jdrouet/marionette/cmd/parser

vendor:
	dep ensure
