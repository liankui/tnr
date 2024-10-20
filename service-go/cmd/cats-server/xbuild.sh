#!/bin/bash
set -e

case "$1" in
"linux" | "windows" | "darwin") goos="$1" ;;
"mac") goos="darwin" ;;
"") goos="" ;;
*) echo "Unknown 1st option, GOOS:$1" && exit 1 ;;
esac

case "$2" in
"amd64" | "arm64") goarch="$2" ;;
"") goarch="" ;;
*) echo "Unknown 2nd option, GOARCH:$2" && exit 1 ;;
esac

set -x
env CGO_ENABLED=0 GOOS="$goos" GOARCH="$goarch" go build -ldflags "-s -w" -ldflags "-X main.gitHash=$(git rev-list -1 HEAD) -X main.gitBranch=$(git branch --show-current) -X 'main.buildTime=$(date "+%Y-%m-%d %H:%M:%S %Z")'"
