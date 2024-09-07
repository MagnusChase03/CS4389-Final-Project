# API

*This is the HTTP API that the frontend client interacts with to create the
functionallity of an end-to-end encrypted group messenger.*

## Table of Contents

**Misc.**

- [Healthcheck](#healthcheck)

## Build

To build the application:

```
$ sudo podman build -t localhost/cs4389-project-api .
```

## Usage

To run the API:

```
$ sudo podman run --env-file ../.env --name cs4389-project-api -p 8080:8080 \
  -d localhost/cs4389-project-api
```

To stop the API:

```
$ sudo podman stop cs4389-project-api
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
    "status": "ok"
}
```
