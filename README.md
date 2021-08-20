# [Pingbot](https://pingbot.cf/) - Backend

[![Repo Size](https://img.shields.io/github/repo-size/MedzikUser/go-pingbot)](https://github.com/MedzikUser/go-pingbot)
[![Build](https://img.shields.io/github/workflow/status/MedzikUser/go-pingbot/release/main)](https://github.com/MedzikUser/go-pingbot/actions/workflows/release.yml)

This application "pings" websites every few minutes (to be set in config). It can be used to keep the application alive on e.g. [glitch.me](https://glitch.com/) or [repl.it](https://replit.com/).

## Install Pre-Compile binary

* 💻 Linux amd64
  * [Download](https://github.com/MedzikUser/go-pingbot/releases) latest version
  * Unpack file `tar xzf pingbot_{VERSION}_linux_amd64.tar.gz`
  * Create an .env and config.toml file and complete according to .env.schema and config.schema.toml
  * Add permissions `chmod +rwx pingbot.out`
  * Run binary `./pingbot.out`
