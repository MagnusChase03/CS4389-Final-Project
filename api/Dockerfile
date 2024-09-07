FROM docker.io/library/go

RUN mkdir -p /code
WORKDIR /code

COPY . .
RUN go build main.go && strip main

CMD ["main"]
