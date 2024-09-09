FROM docker.io/library/golang:alpine as builder

# Install deps
RUN apk update && \
    apk add --no-cache binutils && \
    mkdir -p /code

# Copy over code and build the application
WORKDIR /code
COPY ./api/ .
RUN go build main.go && strip main

FROM docker.io/library/alpine

RUN mkdir -p /code && \
    mkdir -p /certs

# Copy over needed files from builder
WORKDIR /code
COPY ./certs/server.* /certs
COPY --from=builder /code/main apiServer

CMD ["/code/apiServer"]
