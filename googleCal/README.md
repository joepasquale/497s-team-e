# Google Calendar Exporter Service Documentation

## Service Author
Joe Pasquale

## Description
This is a Python Flask service that allows the user to export events to their Google calendar. The service connects to the Google API through a client ID and a user login using a React component. The verified OAuth credentials are then passed to the back end, which allows the user to create events in their Google Calendar that they organize using LinkUp via the Google Calendar API.  

This microservice is not fully working, as I cannot figure out how to manipulate the OAuth credentials that are returned from react-google-login so that the backend flask server can execute the event actions. The credentials are logged and an access token is generated, but I don't know where to interpret that in order to build the API service for the Python app. It could be done straight in the React app, but that defeats the point of making the microservice in the first place. 
