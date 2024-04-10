FROM golang:1.22.2 as builder
WORKDIR /go/src/app/
COPY . /go/src/app/
WORKDIR /go/src/app/init/app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

FROM gcr.io/distroless/static:latest-amd64
COPY --from=builder /go/src/app /go/src/app
WORKDIR /go/src/app/init/app
ENTRYPOINT ["./app"]
EXPOSE 80
