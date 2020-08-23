FROM golang:latest

WORKDIR /go/src/go-hls-hosting

COPY ./ /go/src/go-hls-hosting

EXPOSE 1323

CMD ["go", "run", "main.go"]