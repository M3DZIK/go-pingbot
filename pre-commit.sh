#!/usr/bin/bash

go mod tidy

# Lint
go fmt ./...
