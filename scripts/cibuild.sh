#!/usr/bin/env bash

## Build Package
#go build ./src/github.com/scottbrumley/epo
go get "golang.org/x/crypto/ssh/terminal"

go build epo.go

