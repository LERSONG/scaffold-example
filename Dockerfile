
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/github.com\LERSONG\scaffold-example
COPY . .
RUN apk add --no-cache git
RUN go get ./...
RUN go build -o scaffold-example ./cmd/main.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache bash
COPY --from=builder /go/src/github.com\LERSONG\scaffold-example/scaffold-example /scaffold-example
COPY --from=builder /go/src/github.com\LERSONG\scaffold-example/config/config.yml /config.yml

ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh ./wait-for-it.sh
RUN ["chmod", "+x", "./wait-for-it.sh"]
LABEL Name=scaffold-example
EXPOSE 8080
