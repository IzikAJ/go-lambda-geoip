FROM golang:latest as builder

# RUN apk add curl make git && \
RUN \
  go install github.com/a-urth/go-bindata/go-bindata@latest && \
  mkdir -p /app

COPY . /app
WORKDIR /app

# ENV MAXMIND_LICENSE_KEY=$MAXMIND_LICENSE_KEY
ARG MAXMIND_LICENSE_KEY
RUN make download && make build

FROM node:lts-slim

RUN mkdir -p /usr/local/app/bin
WORKDIR /usr/local/app
COPY --from=builder /app/bin ./bin
COPY serverless.yml package.json package-lock.json ./
RUN npm install

CMD ["npm", "run", "deploy"]
