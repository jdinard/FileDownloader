# Use the latest golang build image as our build environment
FROM golang:latest AS build-env

# Specify arguments as being required
ARG cmd_args

# Make args available to entrypoint
ENV cmd_args_env=${cmd_args}

# get realize for hotloading (makes dev way nicer)
RUN go get github.com/oxequa/realize

# get an initial copy of our application files over
COPY ./FileDownloader/app /go/src/filedownloader/app

# Download our package files over
COPY ./FileDownloader/app/pkg/downloader /go/src/downloader

# Development entrypoint/workdir
WORKDIR /go/src/filedownloader/app

# Using a shell like this makes it really easy to have the container run unit tests for a CI server (like teamcity) instead of running
ENTRYPOINT ./start.sh

# Only for production, we could skip this in development by just passing an arg to the dockerfile (and making a slightly less simple dockerfile)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o filedownloader .

# Final entrypoint for when we want to target a prod build and dont want to be using a hotloader
FROM alpine:latest AS final
COPY --from=build-env /go/src/filedownloader/app/filedownloader /filedownloader
ENTRYPOINT ["/bin/sh"]