# API

*This is the HTTP API that the frontend client interacts with to create the
functionallity of an end-to-end encrypted group messenger.*

## Table of Contents

**Auth**

- [Login](#login)
- [Logout](#logout)

**Misc.**

- [Healthcheck](#healthcheck)

## Build

To build the application:

```
$ sudo podman build -t localhost/cs4389-api -f ./Dockerfile ../
```

## Usage

To run the API:

```
$ sudo podman run --env-file ../.env --name cs4389-api -p 8081:8080 \
  -d localhost/cs4389-api
```

To stop the API:

```
$ sudo podman stop cs4389-api
```

## Documentation

### /healthcheck

*Route to test the current status of the API.*

**Method**: `GET`

**Parameters**: `N/A`

**Example**: `https://api.application.com/healthcheck`

**Returns**: `200`

```JSON
{
    "StatusCode": 200,
    "Data": "Ok"
}
```

### /login

*Route to login a user.*

**Method**: `POST`

**Body**: `username`, `password`

**Example**: `https://api.application.com/login`

**Returns**: `200`, `401`, `400`, `500`

```JSON
{
    "StatusCode": 200,
    "Data": "Ok"
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```

```JSON
{
    "StatusCode": 401,
    "Data": "Invalid username or password."
}
```

```JSON
{
    "StatusCode": 500,
    "Data": "Internal Server Error"
}
```

### /logout

*Route to logout a user.*

**Method**: `POST`

**Body**: N/A

**Example**: `https://api.application.com/logout`

**Returns**: `200`, `401`, `400`

```JSON
{
    "StatusCode": 200,
    "Data": "Ok"
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```

```JSON
{
    "StatusCode": 401,
    "Data": "Unauthorized"
}
```
