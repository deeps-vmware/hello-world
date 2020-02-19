FROM golang:alpine as builder

WORKDIR /build

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" main.go

FROM scratch

COPY --from=builder /build/main .

ENTRYPOINT [ "./main" ]
