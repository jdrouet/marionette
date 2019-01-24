build:
	go build -o marionette cmd/main.go

build-all:
	GOOS=linux go build -o marionette-linux cmd/main.go
	GOOS=darwin go build -o marionette-darwin cmd/main.go
	GOOS=windows go build -o marionette.exe cmd/main.go

run:
	go run cmd/main.go

.PHONY: test
test:
	go test -cover github.com/jdrouet/marionette/cmd/git
	go test -cover github.com/jdrouet/marionette/cmd/model
	go test -cover github.com/jdrouet/marionette/cmd/parser

vendor:
	dep ensure
