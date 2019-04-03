# double-booked
A simple gRPC microservice designed to let you know if a list of events conflict with eachother

## Using the service
Using the service is pretty straightforward, it implements the gRPC service described in the protobuf located in the /protobufs folder of this project. This makes it easy to consume in a variety of languages.

The service implements a single gRPC method `GetConflicts` that accepts an EventList message and returns a ConflictList message.

## Running it in development (with hotloading)
To run the service locally with hotloading, just run
```
docker-compose -f docker-compose.dev.yml up --build
```

## Running unit tests
### Locally
Running unit tests locally just requires adding the following line to the environement/conflictService.env file:
```
TEST=1
```

Then just run
```
docker-compose -f docker-compose.dev.yml up --build
```

### In a CI environment
Have team city or another CI server build the container an environment variable of TEST=1 into the docker container prior to starting it. The service already outputs unit test data in a way the is acceptable to teamcity.

## Creating a production container build
Creating a production container is a must for deploying this guy! In its development mode the container takes an entire gigabyte. Thanks to the docker multi-step builds though, a production container that is less than 3 megabytes can be built by using the docker-compose.prod file, like this:
```
docker-compose -f docker-compose.dev.yml -f docker-compose.prod.yml up --build
```

## Deployment on Kubernetes
### Health checks
The service answers standard HTTP health checks on the /healthz endpoint, by default it will listen on port 8080, the port can be overridden by setting the HEALTH_CHECK_HOST environment variable. Note, this environment variable is also responsible for setting the full host address.

### Handling HTTP/2 traffic
Because this is a gRPC service, it needs kubernetes services/ingresses that are able to properly load balance HTTP/2 traffic. When deploying a service like this to kubernetes, make sure you're exposing it through an l7 load balancer under the hood. 

As an example, on GKE you can use the following service with a special HTTP2 annotation
```
apiVersion: v1
kind: Service
metadata:
  annotations:
    cloud.google.com/app-protocols: '{"double-booked-port":"HTTP2"}'
  name: double-booked
  labels:
    app: double-booked
spec:
  type: NodePort
  ports:
  - port: 443
    targetPort: 7070
    protocol: TCP
    name: double-booked-port
  selector:
    app: double-booked
```
By default the application listens on port 7070, but this can be overriden with the SERVICE_HOST environment variable, again note that this overrides the entire host address.

## Areas for improvement
- Improve unit tests so that they provide better coverage of cases
- Provide a GetConflicts method that returns a stream of conflicts, because of the nature of the algorithm implemented for detecting conflicts, its possible for this service to stream conflict groups back as it sees them, using gRPC streams, instead of waiting until it has processed the entire list to return

As an example, adding:
```
rpc GetConflicts (EventList) returns (stream ConflictList) {}
```