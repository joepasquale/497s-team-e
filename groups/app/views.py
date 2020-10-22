# Import modules
from app import app
from flask import Flask, request, jsonify
import pymongo
import os

# Access db containerName:portNum
myclient = pymongo.MongoClient("mongodb://mongo:27017/")
mydb = myclient["mydatabase"]
mycol = mydb["customers"]

@app.route("/group")
def index():
    return "Hello from Groups Service"

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
    result = mycol.insert_one(PARAMETERS)
    return str(result)

# READ Operation for Groups
@app.route("/group/read", methods=['POST'])
def readGroup():
    # Retreive data from POST request
    req_data = request.get_json(force=True)
    groupID = req_data['groupID']

    # Parameters for DB Query
    PARAMETERS = {'groupID': groupID}

    # Find and return entry corresponding to groupID in PARAMETERS
    result = mycol.find_one(PARAMETERS)
    return str(result)

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
    filter = { 'groupID': groupID }
    # Values to be updated
    newParams = { "$set": 
                    {
                        "groupName": groupName, 
                        "groupMembers": groupMembers
                    }
                }

    # Update groupName and groupMembers based on a groupID. Upsert=true will create a group if no DB groups match the params given
    result = mycol.update_one(filter, newParams, { 'upsert': True })
    return str(result)

# DELETE Operation for Groups
@app.route("/group/delete", methods=['POST'])
def deleteGroup():
    # Retreive data from POST request
    req_data = request.get_json(force=True)
    groupID = req_data['groupID']

    # Parameters for DB Query
    PARAMETERS = {'groupID': groupID}

    # Delete entry corresponding to groupID in PARAMETERS
    result = mycol.delete_one(PARAMETERS)
    return str(result)
