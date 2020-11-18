# THE FOLLOWING COMMANDS NEED TO BE EXECUTED ONE BY ONE
#!/usr/bin/env bash
set -x
trap read debug
# To scale the database simply add another shard into the cluster, we will introduce our shard to the routers:
docker exec -it mongors2n1 bash -c "echo 'rs.initiate(\
{   _id: \"mongors2\", \
    members:[ \
        { _id : 0, host : \"mongors2n1\" }, \
        { _id : 1, host : \"mongors2n2\" }, \
        { _id : 2, host : \"mongors2n3\" }] \
})' | mongo"

docker exec -it mongors2n1 bash -c "echo 'rs.status()' | mongo"

docker exec -it mongos1 bash -c "echo 'sh.addShard(\"mongors2/mongors2n1:27017\")' | mongo --port 27019"

docker exec -it mongos1 bash -c "echo 'sh.status()' | mongo --port 27019"
