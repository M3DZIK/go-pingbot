FROM gitpod/workspace-full

RUN brew install goreleaser

RUN npm i -g nodemon
