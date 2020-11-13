# Import modules
from app import app
from flask import Flask, request, jsonify
import pymongo
import os

# Access db containerName:portNum
# mongo is the container name, and thus our host. It is hosting the mongo instance in the "mongo" container on port 27017, so we access from there.
myClient = pymongo.MongoClient("mongodb://mongoGroups:27017/", connect=False)
myDB = myClient["mydatabase"]
myCol = myDB["group"]


@app.route("/group")
def index():
    return "Welcome to the Groups Service! Please access one of our CRUD endpoints. See documentation for formatting"

# CRUD Operations for Groups:
# CREATE Operation for Groups


@app.route("/group/create", methods=['POST'])
def createGroup():
    # Retreive data from POST request
    req_data = request.get_json(force=True)

    # Assign variables from data
    groupID = req_data['groupID']
    groupName = req_data['groupName']
    groupMembers = req_data["groupMembers"]

    # Parameters for DB Query
    PARAMETERS = {
        'groupID': groupID,
        'groupName': groupName,
        'groupMembers': groupMembers
    }

    # Insert entry with all PARAMETERS specified
    # Result is an InsertOneResult
    result = myCol.insert_one(PARAMETERS)
    resultString = "Inserted a DB entry with the following parameters: " + \
        str(PARAMETERS)
    return resultString

# READ Operation for Groups


@app.route("/group/read", methods=['POST'])
def readGroup():
    # Retreive data from POST request
    req_data = request.get_json(force=True)
    groupID = req_data['groupID']

    # Parameters for DB Query
    PARAMETERS = {'groupID': groupID}

    # Find and return entry corresponding to groupID in PARAMETERS
    result = myCol.find_one(PARAMETERS)
    val = result["groupName"]
    return "groupName for the groupID that was queryed: " + str(val) + "\n"

# UPDATE Operation for Groups


@app.route("/group/update", methods=['POST'])
def updateGroup():
    # Retreive data from POST request
    req_data = request.get_json(force=True)

    # Assign variables from data
    groupID = req_data['groupID']
    groupName = req_data['groupName']
    groupMembers = req_data["groupMembers"]

    # Parameters for DB Query
    # Filter (Updating the Matching groupID)
    filter = {'groupID': groupID}
    # Values to be updated
    newParams = {"$set":
                 {
                     "groupName": groupName,
                     "groupMembers": groupMembers
                 }
                 }

    # Update groupName and groupMembers based on a groupID. Upsert=true will create a group if no DB groups match the params given
    result = myCol.update_one(filter, newParams, upsert=True)

    # Check updated entry
    result = myCol.find_one({"groupID": groupID})
    updatedName = result["groupName"]
    updatedMembers = result["groupMembers"]

    returnStr = "groupID of the object that was changed: " + \
        str(groupID) + "\n" + "New parameters set for that group: " + "\n" + "groupName: " + \
        str(updatedName) + "\n" + "groupMembers: " + \
        ' '.join(groupMembers) + "\n"
    return returnStr

# DELETE Operation for Groups


@app.route("/group/delete", methods=['POST'])
def deleteGroup():
    # Retreive data from POST request
    req_data = request.get_json(force=True)
    groupID = req_data['groupID']

    # Parameters for DB Query
    PARAMETERS = {'groupID': groupID}

    # Delete entry corresponding to groupID in PARAMETERS
    result = myCol.delete_one(PARAMETERS)
    return "groupID of the deleted group: " + str(groupID) + "\n"
