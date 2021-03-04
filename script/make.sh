#!/usr/bin/env bash
set -e

if [ -z "$VERSION" ]; then
    VERSION=$(git rev-parse HEAD)
fi

if [ -z "$CODENAME" ]; then
    CODENAME=cheddar
fi

if [ -z "$DATE" ]; then
    DATE=$(date -u '+%Y-%m-%d_%I:%M:%S%p')
fi

GOOS=linux CGO_ENABLED=0 GOGC=off  go build -v -ldflags "\
-X github.com/bitmyth/accounts/src/app/version.Version=$VERSION \
-X github.com/bitmyth/accounts/src/app/version.Codename=$CODENAME \
-X github.com/bitmyth/accounts/src/app/version.BuildTime=$DATE " \
-a -installsuffix nocgo -o dist/accounts ./src/server
