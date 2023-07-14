# build stage
FROM golang:latest AS builder

RUN apt-get install git gcc libc-dev

WORKDIR /go/src/coupon-service

COPY . .

RUN go mod init coupon-service
RUN go mod tidy
RUN go mod download

WORKDIR /go/src/coupon-service/cmd/export
# RUN go build -o /go/src/coupon-service/main /go/src/coupon-service
RUN go build -o main .


ENTRYPOINT ./main
EXPOSE 8080
##### git needed for fetching dependencies
####RUN apt-get install git gcc libc-dev

####WORKDIR ./go

####COPY . /go/src/coupon_service

####RUN go mod init coupon_service
####RUN go mod tidy
####RUN go mod download

##### Build the binary.
##### WORKDIR /go/src/coupon-service/cmd
##### RUN mkdir -p /go/bin/coupon_service
####RUN go build -o /go/bin/coupon_service /go/bin/coupon_service

##### final stage
####FROM golang:latest

##### Copy the binary from build to the /go/bin directory.
####COPY --from=builder /go/bin/coupon_service /go/bin/coupon_service

##### Run the binary.
####ENTRYPOINT ["/go/bin/go-code-review"]
####EXPOSE 8080


#RUN apk add git gcc libc-dev

# WORKDIR ./

# COPY . .

# WORKDIR /go/src/coupon-service/cmd
# RUN go build -o main .
# RUN ./main

# ENTRYPOINT ./main
# EXPOSE 8080
