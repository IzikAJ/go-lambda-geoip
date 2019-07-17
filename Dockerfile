FROM golang:latest as builder

# RUN apk add curl make git && \
RUN \
  go get -u github.com/jteeuwen/go-bindata/go-bindata && \
  mkdir -p /app

COPY . /app
WORKDIR /app

RUN make download && make build

# FROM node:lts-slim
FROM node:lts-slim

RUN npm install serverless -g && mkdir -p /usr/local/app/bin
WORKDIR /usr/local/app
COPY --from=builder /app/bin ./bin
COPY serverless.yml ./

CMD ["/usr/local/bin/serverless", "deploy", "--verbose"]
