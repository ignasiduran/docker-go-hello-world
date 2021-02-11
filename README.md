# Sample Go application with Docker

Go webapp with Docker.

## Build container image

```
docker build -t go-hello-world .
```

## Start container

```
docker run -p 8080:8080 go-hello-world
```

## Test

```
http://127.0.0.1/8080
```