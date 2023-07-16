FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
# WORKDIR $GOPATH/src/github.com/ctfrancia/go-code-review
WORKDIR /app

COPY . .
# Fetch dependencies.
# RUN go mod init github.com/ctfrancia/go-code-review
# RUN go mod download

# Build the binary.
RUN go mod init github.com/ctfrancia/go-code-review
RUN go mod tidy
RUN go mod download
# RUN go get -d -v ./...
RUN go build -o /go/bin/coupon_service ./cmd/

FROM scratch

COPY --from=builder /go/bin/coupon_service /go/bin/coupon_service

ENTRYPOINT ["/go/bin/coupon_service"]
EXPOSE 8080
