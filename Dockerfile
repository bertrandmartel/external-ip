FROM golang:1.13-alpine
WORKDIR /go/src/app
COPY . .
RUN go install 
RUN go build
RUN ls
CMD ["./external-ip"]