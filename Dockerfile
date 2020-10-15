FROM golang:1.14 AS builder
ENV CGO_ENABLED 0
WORKDIR /go/src/app
ADD . .
RUN go build -o /auto-remove-evicted

FROM alpine:3.12
COPY --from=builder /auto-remove-evicted /auto-remove-evicted
CMD ["/auto-remove-evicted"]