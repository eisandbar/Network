# Network
A simple server that will be continuously upgraded

### About
This project was a way for me to practice the various system design concepts that I went through in the [primer](https://github.com/donnemartin/system-design-primer).
The idea was to start with something basic, a simple client and backend servers. Then, by making the requests from the client more complex, come up with improvements for the server that will solve the issues that arose. The end product should be a scalable service that is resiient to failure.

The steps that I took are explained in more detail in the stages section.
Note that simply cloning this repository and running docker-compose up is not enough to have the service started, more details are in the stages section.


### Resource
https://github.com/donnemartin/system-design-primer

### Stages

### Plan
- ~~Create basic server~~
- ~~Create db, connect to api~~
- ~~Create LB, use several servers~~ (Nginx proxy)
- ~~High intensity queries (better test load)~~
- ~~Use cache~~
- ~~High memory queries (test cache)~~
- ~~Better Load balancer~~
- ~~Limit container resources~~
- Something that checks server load, query response times (can use docker stats for now)
- ~~Message queue (for large tasks that can be done in ||)~~
- Kubernetes
- Create balancer for db, use multiple db
