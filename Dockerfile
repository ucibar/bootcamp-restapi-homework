FROM golang:1.17-alpine AS builder
RUN mkdir /build
COPY . /build/
WORKDIR /build
RUN go build ./cmd/http

FROM alpine
WORKDIR /app
COPY --from=builder /build/http /app/
WORKDIR /app
CMD ["./http"]