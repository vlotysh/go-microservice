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

# Install go dependencies
RUN go get github.com/joho/godotenv

#build the go app
#RUN go env -w GO111MODULE=auto
RUN go mod init
RUN go get -d
RUN GOOS=linux GOARCH=amd64 go build -o ./go-microservice ./main.go


# final stage
#FROM alpine:3.8

#WORKDIR /
#COPY --from=builder /go/src/go-microservice/certs/docker.localhost.* /
#COPY --from=builder /go/src/go-microservice/go-microservice /

#EXPOSE 8081

#CMD ["/bin/sh"]

EXPOSE 8081

CMD ["./go-microservice"]