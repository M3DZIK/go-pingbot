FROM gitpod/workspace-full

# GoReleaser
RUN curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh

# GoLangCI Lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.41.1

RUN npm i -g nodemon
