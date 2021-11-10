FROM golang:1.16 AS builder

########
# Prep
########


# add the source
COPY . /go/src/go-microservice
WORKDIR /go/src/go-microservice/

########
# Build Go Wrapper
########

#build the go app
RUN go mod init
RUN go get -d
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-microservice main.go

# final stage
FROM alpine:3.10
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/go-microservice/certs ./certs
COPY --from=builder /go/src/go-microservice/.env .
COPY --from=builder /go/src/go-microservice/go-microservice .

EXPOSE 8081

CMD ["./go-microservice"]