# API exposing inbound/outbound IP

[![Build Status](https://github.com/bertrandmartel/external-ip/workflows/build%20and%20deploy/badge.svg)](https://github.com/bertrandmartel/external-ip/actions?workflow=build%20and%20deploy)
[![Go Report Card](https://goreportcard.com/badge/github.com/bertrandmartel/external-ip)](https://goreportcard.com/report/github.com/bertrandmartel/external-ip)
[![](https://img.shields.io/docker/pulls/bertrandmartel/external-ip.svg)](https://hub.docker.com/r/bertrandmartel/external-ip)
[![License](http://img.shields.io/:license-mit-blue.svg)](LICENSE.md)

Small server written in go exposing a single API displaying inbound & outbound IP

* outbound IP(s) is(are) retrieved using [ipify API](https://www.ipify.org/)
* inbound IP(s) is(are) retrieved using [Google DNS API](https://dns.google.com/)

Environment variables :

|   name   | description |
|----------|-------------|
| PORT     | server port |
| HOSTNAME | hostname to check inbound ip |

## Using Docker

* DockerHub

```
docker run -p 4242:4242 -e PORT=4242 -e HOSTNAME=example.com -it bertrandmartel/external-ip
```

* locally

```
docker build . -t external-ip
docker run -p 4242:4242 -e PORT=4242 -e HOSTNAME=example.com -it external-ip
```

## Using Go

```
go install
go run ./main.go
```

or 

```
go install
go build
./external-ip
```

## Dependencies

* [echo](https://echo.labstack.com/)