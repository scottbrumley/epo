#!/usr/bin/env bash

## Build Package
go build ./src/github.com/scottbrumley/epo

## Install Package
go install ./src/github.com/scottbrumley/epo

## Install Hello Packages
go install ./src/github.com/scottbrumley/hello