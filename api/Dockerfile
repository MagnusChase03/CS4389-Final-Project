FROM docker.io/library/golang:alpine

RUN apk update && \
    apk add --no-cache openssl binutils && \
    mkdir -p /code && \
    mkdir -p /certs && \
    openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
    -keyout /certs/server.key -out /certs/server.crt \
    -subj "/C=/ST=/L=/O=/CN="

WORKDIR /code

COPY . .
RUN go build main.go && strip main

CMD ["/code/main"]
