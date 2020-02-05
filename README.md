# API exposing inbound/outbound IP

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

## Depedencies

* [echo](https://echo.labstack.com/)