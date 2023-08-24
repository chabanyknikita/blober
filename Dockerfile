FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/ gitlab.com/nikchabanyk/blober
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/blober /go/src/ gitlab.com/nikchabanyk/blober


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/blober /usr/local/bin/blober
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["blob-svc"]
