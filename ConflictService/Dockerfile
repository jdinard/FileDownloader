# Use the latest golang build image as our build environment
FROM golang:latest AS build-env

# protoc and gRPC setup
RUN apt-get update && apt-get -y install unzip && apt-get clean

# Download and install protoc
ENV PB_VER 3.6.1
ENV PB_URL https://github.com/google/protobuf/releases/download/v${PB_VER}/protoc-${PB_VER}-linux-x86_64.zip
RUN mkdir -p /tmp/protoc && \
    curl -L ${PB_URL} > /tmp/protoc/protoc.zip && \
    cd /tmp/protoc && \
    unzip protoc.zip && \
    cp /tmp/protoc/bin/protoc /usr/local/bin && \
    cp -R /tmp/protoc/include/* /usr/local/include && \
    chmod go+rx /usr/local/bin/protoc && \
    cd /tmp && \
    rm -r /tmp/protoc

# Install go gRPC packages
RUN go get github.com/golang/protobuf/proto
RUN go get github.com/golang/protobuf/protoc-gen-go
RUN go get -u google.golang.org/grpc
RUN go get github.com/golang/protobuf/ptypes

# get realize for hotloading
RUN go get github.com/oxequa/realize

# get an initial copy of our application files over
COPY ./ConflictService/app /go/src/conflictservice/app

# copy our implementation of the golang interface
COPY ./ConflictService/app/pkg/calendar /go/src/calendar

WORKDIR /go/src/conflictservice/app

# protobuf generators, this is where we actually use protoc that we installed earlier to generate the go gRPC code
COPY ./protobufs ./protobufs

# Generate the gRPC interface that we will write an implementation for in our go code
RUN mkdir ../../../src/scheduling
RUN protoc -I=./protobufs calendar.proto --go_out=plugins=grpc:../../../src/scheduling

# Development entrypoint
ENTRYPOINT ./start.sh