# THE FOLLOWING COMMANDS NEED TO BE EXECUTED ONE BY ONE
#!/usr/bin/env bash
set -x
trap read debug

# initiate config server
docker exec -it mongocfg1 bash -c "echo 'rs.initiate(\
{   _id: \"mongocfg\", \
    configsvr: true,    \
    members:[ \
        { _id : 0, host : \"mongocfg1\" }, \
        { _id : 1, host : \"mongocfg2\" }, \
        { _id : 2, host : \"mongocfg3\" }] \
})' | mongo"           
docker exec -it mongocfg1 bash -c "echo 'rs.status()' | mongo"

# build shard replica sets
docker exec -it mongors1n1 bash -c "echo 'rs.initiate(\
{   _id: \"mongors1\", \
    members:[ \
        { _id : 0, host : \"mongors1n1\" }, \
        { _id : 1, host : \"mongors1n2\" }, \
        { _id : 2, host : \"mongors1n3\" }] \
})' | mongo"

docker exec -it mongors1n1 bash -c "echo 'rs.status()' | mongo"

# Finally, we will introduce our shard to the routers:
docker exec -it mongos1 bash -c "echo 'sh.addShard(\"mongors1/mongors1n1\")' | mongo --port 27019"
docker exec -it mongos1 bash -c "echo 'sh.status()' | mongo --port 27019"

# Now, our sharded cluster configuration is complete. 
# We don’t have any databases yet. 
# We will create a database and will enable sharding.
docker exec -it mongors1n1 bash -c "echo 'use uplink-test' | mongo"
docker exec -it mongos1 bash -c "echo 'sh.enableSharding(\"uplink-test\")' | mongo --port 27019"
# Create collection in the database
docker exec -it mongors1n1 bash -c "echo 'db.createCollection(\"uplink-test.users\")' | mongo"
# The collection is not shareded yet.
# We will shard it on a field named email
docker exec -it mongos1 bash -c "echo 'sh.shardCollection(\"uplink-test.users\", {\"_id\" : \"hashed\"})' | mongo --port 27019"
docker exec -it mongos1 bash -c "echo 'sh.getBalancerState()' | mongo --port 27019"
