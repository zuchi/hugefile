# Compile image
FROM golang:1.20-alpine3.18 as BUILDER
RUN apk update && apk add --no-cache curl ca-certificates git
RUN  mkdir /go/project /go/project/src
COPY go.mod /go/project
COPY go.sum /go/project
COPY ./src/. /go/project/src
WORKDIR /go/project
RUN env GOOS=linux GOARCH=amd64 go build -o rest ./src/cmd/rest.go

#Running image
FROM alpine:3.18.2 as RUNNER
RUN apk update && apk add --update bash ca-certificates
RUN apk add -U tzdata && cp /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime && echo "America/Sao_Paulo" > /etc/timezone
COPY --from=builder /go/project/rest /rest
RUN ["chmod", "+x", "rest"]
ENTRYPOINT ["./rest"]