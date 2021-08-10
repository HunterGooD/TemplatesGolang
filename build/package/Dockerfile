FROM golang:1.16-alpine3.14 AS builder

RUN go version
RUN apk --no-cache add gcc g++ make git

WORKDIR /go/src/app
COPY ./ ./


RUN go mod download && go get -u ./...
RUN GOOS=linux go build -o ./.bin/app ./cmd/template/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder ./.bin/app .
COPY --from=builder ./config/ ./config/

EXPOSE 8000

CMD [ "./app"]