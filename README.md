# Network
A simple server that will be continuously upgraded

### About
This project was a way for me to practice the various system design concepts that I went through in this [primer](https://github.com/donnemartin/system-design-primer).
The idea was to start with something basic, a simple client and backend servers. Then, by making the requests from the client more complex, come up with improvements for the server that will solve the issues that arose. The end product should be a scalable service that is resiient to failure.

The steps that I took are explained in more detail in the [stages](#stages) section.
Note that simply cloning this repository and running docker-compose up is not enough to have the service started, more details are in the setup section.


### Resource
https://github.com/donnemartin/system-design-primer

### Setup

### Stages

#### 1. Base
The project starts of with a single server with a ping and longQuery(sleep 20s) endpoints. A separate client server sends a request every 2 seconds to either endpoint.
This model was made to test networking and the basic principles. One problem of this model is that since GO is so good at concurrency the longQ doesn't really load the server, so something else needs to be used to test it.

#### 2. Adding DB
Next I added a postgres DB with some dummy data. This way client requests can better mimic real world scenarios. Currently the server only does reads and the edpoints will be cleaned up later on.

#### 3. Docker replicas
Now I added an nginx load balancer between the client and the server. This way I can easily scale the number of server containers if demand increases. To test this I used docker deploy replicas to create another 2 server containers. In the docs it says that it only works in swarm mode, however on a single machine it seems to work fine even in docker-compose up.

#### 4. Load queries
To better test cpu load I created load queries that make the server perform some sort of calculations. However two problems arose. First, the containers didn't have cpu or memory limits set, which made testing their resources more difficult. Second, servers would get bogged down with the loadQ requests and timeout. To solve this I decided that a message queue was the best option, as it would take off cpu load from the servers. However that change was put off until later onn and loadQ was turned off for the time being.

#### 5. Cache
Made the users query endpoints restful. Now requests can Get/Post/Del. Turned up the number of requests the client sends to 200 reads/sec. A lot of these are from a small group of users. The idea here is that by adding a cache we can store the results of popular queries and lessen the load on the db. Added a redis cache that updates 
using the cache-aside strategy. The reasoning is that we want to be able to have a large amount of writes of unpopular content, which we don't really want to store in the cache. After implementing redis the load on the db and servers went from 160/60% to 30/30% with the cache proving greatly effective. After this I added resource limits to the containers to better understand how much each container needed and to better see stress.

#### 6. Haproxy and keepalived
Currently, with only one load balancer, nginx becomes a single point of failure. However if we want the client to be able to communicate with servers we need a single IP. To solve this issue we use multiple nginx servers and put haproxy infront of it. However we don't use haproxy by itself, but together with keeplaived. What this does is that it will use a virtual IP to which all requests are sent, and if the master haproxy ever fails the backup will start taking requests instead, all on the same VIP. 

#### 7. Message queue
Returning to the problem of load queries, I added a RabbitMQ server. This way I can add tasks to the queue and worker containers will do the calculations for the servers. Since in this case I wanted to send data back RPC style I had workers create a respone queue. The result was that load queries no longer overloaded the servers but  instead were handled by worker containers.

#### 8. DB replication
Nect I created more client containers, increasing the load on the servers. The weak point turned out to be the postgres DB as it was having trouble keeping up with all the read/write requests. To combat that I added a replica db and a haproxy load balancer for read requests. This way the reads are split between the two databases. One thing to consider is that if I expect a low read/write ratio I should use a master/master strategy so that both DBs can handle the high write load.


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
- ~~Create balancer for db, use multiple db~~
