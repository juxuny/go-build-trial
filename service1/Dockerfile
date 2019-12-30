FROM golang:1.13 as builder
ENV GO111MODULE=auto
WORKDIR /go/src/app
COPY ./tools ./tools
COPY ./service1 ./service1

WORKDIR /go/src/app/tools
RUN go mod download
WORKDIR /go/src/app/service1
RUN go mod download

WORKDIR /go/src/app

RUN go build -o server /go/src/app/service1/main.go

FROM centos:7
RUN yum -y update && yum clean all
COPY --from=builder /go/src/app/server .
EXPOSE 3000
CMD ["./server"]