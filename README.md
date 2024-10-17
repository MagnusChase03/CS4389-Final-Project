# Silent Circle

*A real time end-to-end encrypted scalable group messenger.*

## Build

To build the entire application stack:

```
$ ./build.sh [docker]
```

## Usage

To run the entire application stack:

```
sudo podman-compose up -d
```

To stop the entire application stack:

```
$ sudo podman-compose down
```

The application will be available at https://localhost:8080.

The docker equivalent is:

```
$ docker compose up -d
$ docker compose down
```

## Tech Stack

- Go
- Vue
- MySQL
- Redis
- Nginx
- Docker

## Contributers

- MagnusChase03
