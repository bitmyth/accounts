PWD := $(shell pwd)
BIND_DIR := dist
IMAGE := bitmyth/accounts
DEV_IMAGE := bitmyth/accounts:dev

## Create the "dist" directory
dist:
	mkdir $(BIND_DIR)

## Build goose Docker image
goose:
	docker build -t goose -f Dockerfile.goose .

dev-image: dist
	docker build -t "$(DEV_IMAGE)" -f Dockerfile.build .

## Build the linux binary
binary: dev-image
	docker run --rm -v $(PWD)/$(BIND_DIR):/go/src/github.com/bitmyth/accounts/$(BIND_DIR) $(DEV_IMAGE) ./script/make.sh

## Build a Docker image
image: binary
	docker build -t $(IMAGE) .

## Run Docker development image
run-dev:
	docker run --rm --net account-net --name accounts -v $(PWD):/go/src/github.com/bitmyth/accounts/ -p 8081:8081 $(DEV_IMAGE) go run src/server/main.go
dev:
	export $(cat .env.dev)
	docker-compose up
## Run Docker production image
run-prod:
	docker run --rm --net account-net --name accounts -v $(PWD)/config:/config -p 8081:8081 $(IMAGE)

## Run on local
run:
	go run src/server/main.go
stack:
	export $(cat .env.dev)
	docker stack deploy -c stack.yml accounts

