# FileDownloader

A command line utility written in Go for download the first N chunks of a remote file in parallel. 

In order to avoid having to deal with issues like versions of Go and GO_PATH/GO_ROUTE setups, I've dockerized the utility.

## Using this application

### Building and running the application

Building the application is all taken care of by docker, just run

`docker-compose -f docker-compose.dev.yml -f docker-compose.prod.yml build`

This will build a production container named filedownloader, which you can then run (it will only be about 10.6mb).

After building the container, open an interactive shell with the container:
`docker run -it filedownloader`

Then you can immediately start using the filedownloader application, for example by running:
`./filedownloader --source_url=http://f39bf6aa.bwtest-aws.pravala.com/384MB.jar`

or

`./filedownloader --source_url=http://f39bf6aa.bwtest-aws.pravala.com/384MB.jar --output_file=anotherfile.name`

Note: The container is based on alpine

### Building the application for development

Building the container for development will boot up a docker container that contains realize for hotloading. This container will be much larger, about 1GB.

You can boot the container into a development environment with
`docker-compose -f docker-compose.dev.yml up --build`

Please note that when in development/hotloading mode, arguments that get passed to the go program are defined in the .realize.yaml file. For hotloading to work, make sure you've enabled shared drives with docker.

### Running unit tests

Running unit tests for a CI server are also quite straightforward. 
- Uncomment the environment variable section (in the docker-compose.dev.yml) so that a TEST environment variable is injected into the container when its built
- Then boot the container with `docker-compose -f docker-compose.dev.yml up --build`

You should now see the results of unit tests, in teamcity report format, outputted to your console. In an actual CI system I would recommend you configure it to inject the environment variable, instead of using the docker-compose file to inject it.


## Assumptions
- Passing around a few dockerfiles is much easier than passing around/sharing uncompiled go code
- The binary that the docker container produces is specifically compiled for linux
