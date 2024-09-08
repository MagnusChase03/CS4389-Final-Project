FROM docker.io/library/golang:alpine as builder

# Install deps and build ssl cert
RUN apk update && \
    apk add --no-cache openssl binutils && \
    mkdir -p /code && \
    mkdir -p /certs && \
    openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
    -keyout /certs/server.key -out /certs/server.crt \
    -subj "/C=/ST=/L=/O=/CN="

# Copy over code and build the application
WORKDIR /code
COPY . .
RUN go build main.go && strip main

FROM docker.io/library/alpine

RUN mkdir -p /code && \
    mkdir -p /certs

# Copy over needed files from builder
WORKDIR /code
COPY --from=builder /code/main apiServer
COPY --from=builder /certs/* /certs

CMD ["/code/apiServer"]
