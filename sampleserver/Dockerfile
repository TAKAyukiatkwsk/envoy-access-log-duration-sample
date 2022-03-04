FROM golang:1.17 AS builder

WORKDIR /go/src/app
COPY . /go/src/app
RUN go build -o /go/bin/app

FROM gcr.io/distroless/base-debian11
COPY --from=builder /go/bin/app /
EXPOSE 50051
CMD ["/app"]
