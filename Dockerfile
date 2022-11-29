FROM golang:1.18.0-bullseye as base

RUN apt update && \
    apt-get install -y \
        build-essential \
        ca-certificates \
        curl

# enable faster module downloading.
ENV GOPROXY https://proxy.golang.org

# Install Ignite
RUN curl -L https://get.ignite.com/cli@v0.25.2! | bash

WORKDIR /app

# cache dependencies.
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

COPY app app
COPY cmd cmd
COPY docs docs
COPY proto proto
COPY config.yml config.yml
COPY ts-client ts-client
COPY vue vue
COPY testutil testutil
COPY x x

# Install Node
#RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash -
#RUN apt-get install -y nodejs

RUN ignite chain build

ENTRYPOINT ["bash"]

# CMD ["ignite", "network", "net" ]
