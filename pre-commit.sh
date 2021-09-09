#!/usr/bin/bash

go mod tidy

# Lint
go fmt ./...

# Add changes
git add .
