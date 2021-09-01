FROM gitpod/workspace-full

# Tools
RUN curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh

RUN sudo mv bin/* /usr/bin/ && \
  rm -rf bin/

RUN yarn global add nodemon
RUN echo "export PATH=\$PATH:`yarn global bin`" >> ~/.bashrc
