# Event Service Documentation

## Service Author
Ron Arbo

## Description
This is a service that manages our "Event" objects for our application. It is written with the Flask Python framework. As of now, the service just maintains the object definition for event, and allows database access to our event mongoDB. This service is routed to by our nginx service, and connects to one of our mongoDB services through docker-compose networking. 

## Object Definition
```
Event {  
    eventID: int  
    groupID: int
    eventName: string
    eventTime: string
    eventLocation: string
}
```

## Request Formatting
Note that this service will only process POST requests in JSON formatting. Listed below are the attributes you must include in the request for each operation.  
### Create an event:  
| Attribute     | Description                                                 |
|---------------|-------------------------------------------------------------|
| eventID       | ID of the new event you'd like to create                    |
| groupID       | ID of the existing group you'd like to assign this event to |
| eventName     | Name of the new event you're creating                       |
| eventTime     | Time that this new event will occur at                      |
| eventLocation | Location of this new event                                  |

### Read an event:
| Attribute | Description                                  |
|-----------|----------------------------------------------|
| eventID   | ID of the event that you'd like to read from |

### Update an event:
| Attribute     | Description                                        |
|---------------|----------------------------------------------------|
| eventID       | ID of the event you'd like to update               |
| eventName     | New name of the event corresponding to eventID     |
| eventTime     | New time of the event corresponding to eventID     |
| eventLocation | New location of the event corresponding to eventID |

### Delete an event
| Attribute | Description                          |
|-----------|--------------------------------------|
| eventID   | ID of the event you'd like to delete |

<br> 

## Response Formatting
As of now, reponses will simply be a string explaining what happened (typically involving the DB)
### Create a event
reponse: "Inserted a DB entry with the following parameters: (event Object inserted)"

### Read a event
response: "eventName for the eventID that was queryed: (eventName)"

### Update a event


### Delete a event
response: "eventID of the deleted event: (eventID)"
