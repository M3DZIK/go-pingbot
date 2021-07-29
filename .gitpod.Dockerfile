FROM gitpod/workspace-full

RUN brew install goreleaser

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.41.1

RUN golangci-lint --version

RUN npm i -g nodemon
