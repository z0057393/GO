FROM  --platform=linux/amd64 debian:stable-slim

SHELL ["/bin/bash", "-c"]

RUN apt-get -y update
RUN apt-get -y install curl

WORKDIR /dl

RUN curl -OL https://go.dev/dl/go1.20.5.linux-amd64.tar.gz

RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"

RUN go version

WORKDIR /app

COPY . .