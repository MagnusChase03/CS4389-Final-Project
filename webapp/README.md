# Webapp

*This is the webapp that the client interacts with for an end-to-end encrypted 
group messenger.*

## Build

To build the application:

```
$ sudo podman build -t localhost/cs4389-webapp -f ./Dockerfile ../
```

## Usage

To run the API:

```
$ sudo podman run --env-file ../.env --name cs4389-webapp -p 8080:443 \
  -d localhost/cs4389-webapp
```

To stop the application:

```
$ sudo podman stop cs4389-webapp
```
