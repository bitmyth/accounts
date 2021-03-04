PWD := $(shell pwd)
BIND_DIR := dist
IMAGE := "bitmyth/accounts"
DEV_IMAGE := accounts-build$(if $(GIT_BRANCH),:$(subst /,-,$(GIT_BRANCH)))


## Create the "dist" directory
dist:
	mkdir $(BIND_DIR)

dev-image: dist
	docker build -t "$(DEV_IMAGE)" -f Dockerfile.build .

## Build the linux binary
binary: dev-image
	docker run --rm -v $(PWD)/$(BIND_DIR):/go/src/github.com/bitmyth/accounts/$(BIND_DIR) $(DEV_IMAGE) ./script/make.sh

## Build a Docker image
image: binary
	docker build -t $(IMAGE) .


## Run Docker image for development
serve:
	docker run --rm --net account-net --name accounts -v $(PWD)/config:/config -p 8081:8081 $(DEV_IMAGE) dist/accounts
## Build Dev Docker image
