# Import modules
from app import app
from flask import Flask, request, jsonify
import pymongo
import os

# Access db containerName:portNum
# mongo is the container name, and thus our host. It is hosting the mongo instance in the "mongo" container on port 27018, so we access from there.
myClient = pymongo.MongoClient("mongodb://mongoEvents:27017/", connect=False)
myDB = myClient["mydatabase"]
myCol = myDB["event"]

@app.route("/event")
def index():
    return "Welcome to the Events Service! Please access one of our CRUD endpoints. See documentation for formatting"

# CRUD Operations for events:
# CREATE Operation for events
@app.route("/event/create", methods=['POST'])
def createEvent():
    # Retreive data from POST request
    req_data = request.get_json(force=True)
    
    # Assign variables from data
    eventID = req_data['eventID']
    groupID = req_data['groupID']
    eventName = req_data['eventName']
    eventTime = req_data['eventTime']
    eventLocation = req_data['eventLocation']

    # Parameters for DB Query
    PARAMETERS = {
        'eventID': eventID,
        'groupID': groupID,
        'eventName': eventName,
        'eventTime': eventTime,
        'eventLocation': eventLocation
        }

    # Insert entry with all PARAMETERS specified
    # Result is an InsertOneResult
    result = myCol.insert_one(PARAMETERS)
    resultString = "Inserted a DB entry with the following parameters: " + str(PARAMETERS)
    return resultString

# READ Operation for events
@app.route("/event/read", methods=['POST'])
def readEvent():
    # Retreive data from POST request
    req_data = request.get_json(force=True)
    eventID = req_data['eventID']

    # Parameters for DB Query
    PARAMETERS = {'eventID': eventID}

    # Find and return entry corresponding to eventID in PARAMETERS
    result = myCol.find_one(PARAMETERS)
    val = result["eventName"]
    return "eventName for the eventID that was queryed: " + str(val) + "\n"

# UPDATE Operation for events
@app.route("/event/update", methods=['POST'])
def updateEvent():
    # Retreive data from POST request
    req_data = request.get_json(force=True)
    
    # Assign variables from data
    eventID = req_data['eventID']
    eventName = req_data['eventName']
    eventTime = req_data['eventTime']
    eventLocation = req_data['eventLocation']

    # Parameters for DB Query
    # Filter (Updating the Matching eventID)
    filter = { 'eventID': eventID}
    # Values to be updated
    newParams = { "$set": 
                    {
                        "eventName": eventName, 
                        "eventTime": eventTime,
                        "eventLocation": eventLocation
                    }
                }

    # Update eventName and eventMembers based on a eventID. Upsert=true will create a event if no DB events match the params given
    result = myCol.update_one(filter, newParams, upsert=True)
    
    # Check updated entry
    result = myCol.find_one({"eventID": eventID, "groupID": groupID})
    updatedName = result["eventName"]
    updatedTime = result["eventTime"]
    updatedLocation = result["eventLocation"]

    returnStr =  "eventID of the object that was changed: " + str(eventID) + "\n" + "New parameters set for that event: " + "\n" + "eventName: " + str(updatedName) + "\n" + "eventTime: " + str(updatedTime) + "\n" "eventLocation: " + str(updatedLocation) + "\n"
    return returnStr

# DELETE Operation for events
@app.route("/event/delete", methods=['POST'])
def deleteEvent():
    # Retreive data from POST request
    req_data = request.get_json(force=True)
    eventID = req_data['eventID']

    # Parameters for DB Query
    PARAMETERS = {'eventID': eventID}

    # Delete entry corresponding to eventID in PARAMETERS
    result = myCol.delete_one(PARAMETERS)
    return "eventID of the deleted event: " + str(eventID) + "\n"
