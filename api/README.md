# API

*This is the HTTP API that the frontend client interacts with to create the
functionallity of an end-to-end encrypted group messenger.*

## Table of Contents

**Auth**

- [Login](#login)
- [Logout](#logout)

**User**
- [Get User](#userget)
- [Update User](#userupdate)
- [Create User](#usercreate)
- [Delete User](#userdelete)
- [Send friend request](#userfriendinvite)

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
    "Data": "Invalid username or password."
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

### /user/get

*Route to get a users public key.*

**Method**: `POST`

**Body**: `username`

**Example**: `https://api.application.com/user/get`

**Returns**: `200`, `400`

```JSON
{
    "StatusCode": 200,
    "Data": {
        "PublicKey": "supersecretpublickey"
    }
}
```

```JSON
{
    "StatusCode": 400,
    "Data": "Bad Request"
}
```

### /user/update

*Route to update user information.*

**Method**: `POST`

**Body**: `password` (optional), `publicKey` (optional)

**Example**: `https://api.application.com/user/update`

**Returns**: `200`, `400`, `401`

```JSON
{
    "StatusCode": 200,
    "Data": {
        "PublicKey": "supersecretpublickey"
    }
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

### /user/create

*Route to create a new user.*

**Method**: `POST`

**Body**: `username`, `password`, `publicKey`

**Example**: `https://api.application.com/user/create`

**Returns**: `200`, `400`

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

### /user/delete

*Route to create a delete a user.*

**Method**: `POST`

**Body**: N/A

**Example**: `https://api.application.com/user/delete`

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

### /user/friend/invite

*Route to send a friend request to a user.*

**Method**: `POST`

**Body**: `username`

**Example**: `https://api.application.com/user/friend/invite`

**Returns**: `200`, `400`

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
