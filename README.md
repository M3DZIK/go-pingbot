# [Pingbot](https://pingbot.cf/) - Backend

[![Repo Size](https://img.shields.io/github/repo-size/MedzikUser/go-pingbot)](https://github.com/MedzikUser/go-pingbot)
[![Build](https://img.shields.io/github/workflow/status/MedzikUser/go-pingbot/release/main)](https://github.com/MedzikUser/go-pingbot/actions/workflows/release.yml)

This application "pings" websites every few minutes (to be set in config). It can be used to keep the application alive on e.g. [glitch.me](https://glitch.com/) or [repl.it](https://replit.com/).

## Install Pre-Compile binary

* 💻 Linux amd64
  * [Download](https://github.com/MedzikUser/go-pingbot/releases) latest version
  * Unpack file `tar xzf pingbot_*_linux_amd64.tar.gz`
  * Done your binary is `pingbot.out`

## Compile from Source Code
Not recommended because automatic updates don't work

### Requirements

* [Go](https://golang.org/dl) (recommended latest version)

### Compile

* Download source code `git clone https://github.com/MedzikUser/go-pingbot.git --depth 1`
* Go to folder with source code `cd go-pingbot`
* Download dependencies `go mod tidy`
* Build `go build -o pingbot.out`
* Done your compined binary is `pingbot.out`

## Configurate

* Complete .env according to .env.schema
* And fill in config.toml according to config.schema.toml

## Run

* `./pingbot.out`
