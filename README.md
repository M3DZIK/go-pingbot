# [Pingbot](https://pingbot.cf/) - Backend

[![Repo Size](https://img.shields.io/github/repo-size/MedzikUser/go-pingbot)](https://github.com/medzikuser/go-pingbot)
[![Build](https://img.shields.io/github/workflow/status/MedzikUser/go-pingbot/release/main)](https://github.com/medzikuser/go-pingbot/actions/workflows/release.yml)

This application "pings" websites every few minutes (to be set in config). It can be used to keep the application alive on e.g. [glitch.me](https://glitch.com/) or [repl.it](https://replit.com/).

## ⚡ Install Pre-Compile binary

* 💻 Linux amd64
  * [Download](https://github.com/medzikuser/go-pingbot/releases) latest version
  * Unpack file `tar xf pingbot_*_linux_amd64.tar.xz`
  * Done your binary is `pingbot.out`

## 👨‍💻 Compile from Source Code
‼️ Not recommended because automatic updates don't work

### 🖥️ Requirements

* [Go](https://golang.org/dl) (recommended latest version)

### ⭐ Compile

Method 2:
* `go install github.com/medzikuser/go-pingbot@latest`
* Output binary path
  * Check GOPATH `go env GOPATH`
  * Go to GOPATH and binary name is `pingbot`

Method 2:
* Download source code `git clone https://github.com/medzikuser/go-pingbot.git --depth 1`
* Go to folder with source code `cd go-pingbot`
* Build `make` or `go build -o pingbot.out`
* Done your compiled binary name is `pingbot.out`

### ⭐ Cross Compile

> [Check supported OS and ARCH](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63)

* Download source code `git clone https://github.com/medzikuser/go-pingbot.git --depth 1`
* Go to folder with source code `cd go-pingbot`
* Build `make GOOS=os GOARCH=arch`
  * e.g. `make GOOS=openbsd GOARCH=arm64	`
* Done your cross compiled binary name is `pingbot.out`

## ⚙️ Configurate

* Complete .env according to schema.env
* And fill in config.toml according to config.schema.toml

## 🔧 Run

* `./pingbot.out` or `./start.sh` (auto restart e.g. if exit on update is enabled)
