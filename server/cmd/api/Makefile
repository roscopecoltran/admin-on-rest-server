
APP_NAME := aors-api

all: deps build

deps:
	go get -v github.com/Masterminds/glide
	glide install --strip-vendor

build: deps
	go build -o ./bin/$(APP_NAME) .

quick:
	go run *.go --insecure-http

flags:
	go run *.go --help	