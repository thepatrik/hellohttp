# hellohttp

[![Build Status](https://travis-ci.org/thepatrik/hellohttp.svg?branch=master)](https://travis-ci.org/thepatrik/hellohttp) [![Go Report Card](https://goreportcard.com/badge/github.com/thepatrik/hellohttp)](https://goreportcard.com/report/github.com/thepatrik/hellohttp) [![GoDoc](https://godoc.org/github.com/thepatrik/hellohttp?status.svg)](https://godoc.org/github.com/thepatrik/hellohttp)

A simple barebone HTTP server implementation in golang, that uses [chi](https://github.com/go-chi/chi) as a router with cors, compression, and logging middlewares enabled.

```bash
$ go run .
Running HTTP server on :8080
```

Check the health endpoint.

```bash
curl http://localhost:8080/health
OK
```

#### Docker

The Dockerfile is built from scratch and produces a small image (~8.1 MB).

Build with:

```bash
$ docker build -t hellohttp .
```

Run with:

```bash
$ docker run -p 8080:8080 hellohttp
```

#### Environment variables

The following environment variables are available (use a .env file for convenience).

| Name | Default | Description |
| ---- | ------- | ----------- |
| PORT | 8080    | Port number |
