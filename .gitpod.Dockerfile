FROM gitpod/workspace-full

# GoReleaser
RUN curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh

RUN yarn global add nodemon
RUN echo "export PATH=\$PATH:`yarn global bin`" >> ~/.bashrc
