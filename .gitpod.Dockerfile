FROM gitpod/workspace-full

# GoReleaser
RUN curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh

RUN npm i -g nodemon
