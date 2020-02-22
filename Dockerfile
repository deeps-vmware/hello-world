FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /build

COPY main.go .

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" main.go

FROM scratch

COPY --from=builder /build/main .

USER 65534

ENTRYPOINT [ "./main" ]
