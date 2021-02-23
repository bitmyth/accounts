## Create the "dist" directory
dist:
	mkdir dist

build:
    rm -f dist/traefik
    CGO_ENABLED=0 GOGC=off  go build -a -installsuffix nocgo -o dist/accounts ./src/