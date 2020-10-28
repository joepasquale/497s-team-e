# Groups Service Documentation

## Description
This is a service that manages our "Group" objects for our application. It is written with the Flask Python framework. As of now, the service just maintains the object definition for Groups, and allows database access to our Groups mongoDB. This service is routed to by our nginx service, and connects to one of our mongoDB services through docker-compose netowrking. 

## Object Definition
```
Group {  
    groupID: int  
    groupName: string  
    groupMembers: [string]  
}
```

## Request Formatting
Note that this service will only process POST requests in JSON formatting. Listed below are the attributes you must include in the request for each operation.  
### Create a Group:  
| Attribute    | Description                                          |
|--------------|------------------------------------------------------|
| groupID      | ID of the new group you'd like to create             |
| groupName    | Name of the new group you'd like to create           |
| groupMembers | Members array for the new group you'd like to create |

### Read a Group:
| Attribute | Description                                      |
|-----------|--------------------------------------------------|
| groupID   | ID of the existing group you'd like to read from |

### Update a Group:
| Attribute    | Description                                            |
|--------------|--------------------------------------------------------|
| groupID      | ID of the existing group you'd like to update          |
| groupName    | New name of the existing group given by the groupID    |
| groupMembers | New members of the existing group given by the groupID |

### Delete a Group:
| Attribute | Description                                   |
|-----------|-----------------------------------------------|
| groupID   | ID of the existing group you'd like to delete |

<br> 

## Response Formatting
As of now, reponses will simply be a string explaining what happened (typically involving the DB)
### Create a Group
reponse: "Inserted a DB entry with the following parameters: (group Object inserted)"

### Read a Group
response: "groupName for the groupID that was queryed: (groupName)"

### Update a Group
response: "groupID of the object that was changed: (groupID) <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;New parameters set for that group:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;groupName: (groupName>) <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;groupMembers: (groupMembers)<br>

### Delete a Group
response: "groupID of the deleted group: (groupID)"
