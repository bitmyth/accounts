BIND_DIR := "dist"
IMAGE := "bitmyth/accounts"

## Create the "dist" directory
dist:
	mkdir dist

## Build the linux binary
binary:
	./script/make.sh

## Clean up static directory and build a Docker Traefik image
build-image: binary
	rm -rf static
	docker build -t $(IMAGE) .
