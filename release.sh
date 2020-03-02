#!/bin/sh

go build  -ldflags="-s -w -X main.version=$(git describe --tags `git rev-list --tags --max-count=1`)-$(git rev-parse --short HEAD)" .
