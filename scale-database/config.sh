# initiate config server
docker exec -it mongocfg1 bash -c "echo 'rs.initiate(\
{   _id: \"mongocfg\", \
    configsvr: true,    \
    members:[ \
        { _id : 0, host : \"mongocfg1:27019\" }, \
        { _id : 1, host : \"mongocfg2:27019\" }, \
        { _id : 2, host : \"mongocfg3:27019\" }] \
})' | mongo --port 27019"           
docker exec -it mongocfg1 bash -c "echo 'rs.status()' | mongo --port 27019"

# build shard replica sets
docker exec -it mongors1n1 bash -c "echo 'rs.initiate(\
{   _id: \"mongors1\", \
    members:[ \
        { _id : 0, host : \"mongors1n1:27018\" }, \
        { _id : 1, host : \"mongors1n2:27018\" }, \
        { _id : 2, host : \"mongors1n3:27018\" }] \
})' | mongo --port 27018"

docker exec -it mongors1n1 bash -c "echo 'rs.status()' | mongo --port 27018"

docker exec -it mongors2n1 bash -c "echo 'rs.initiate(\
{   _id: \"mongors2\", \
    members:[ \
        { _id : 0, host : \"mongors2n1:27018\" }, \
        { _id : 1, host : \"mongors2n2:27018\" }, \
        { _id : 2, host : \"mongors2n3:27018\" }] \
})' | mongo --port 27018"
docker exec -it mongors2n1 bash -c "echo 'rs.status()' | mongo --port 27018"

# Finally, we will introduce our shard to the routers:
docker exec -it mongos1 bash -c "echo 'sh.addShard(\"mongors1/mongors1n1:27018\")' | mongo"
docker exec -it mongos1 bash -c "echo 'sh.addShard(\"mongors2/mongors2n1:27018\")' | mongo"
docker exec -it mongos1 bash -c "echo 'sh.status()' | mongo "

# Now, our sharded cluster configuration is complete. 
# We don’t have any databases yet. 
# We will create a database and will enable sharding.
docker exec -it mongors1n1 bash -c "echo 'use uplink-test' | mongo"
docker exec -it mongos1 bash -c "echo 'sh.enableSharding(\"uplink-test\")' | mongo "
# Create collection in the database
docker exec -it mongors1n1 bash -c "echo 'db.createCollection(\"uplink-test.users\")' | mongo "
# The collection is not shareded yet.
# We will shard it on a field named email
docker exec -it mongos1 bash -c "echo 'sh.shardCollection(\"uplink-test.users\", {\"_id\" : \"hashed\"})' | mongo "