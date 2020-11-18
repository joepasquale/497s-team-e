# Import modules
from app import app, calSetup
from flask_cors import CORS
import requests
from google.oauth2 import id_token
from google.auth.transport import requests
from googleapiclient.discovery import build
from flask import Flask, request, jsonify
from datetime import datetime, timedelta

CORS(app)

service = None


@app.route("/gcal/auth", methods=['POST'])
def index():
    global service
    print("This is a test for the Google API")
    req_data = request.get_json(force=True)
    creds = req_data['tokenId']
    try:
        # Specify the CLIENT_ID of the app that accesses the backend:
        idinfo = id_token.verify_oauth2_token(creds, requests.Request(
        ), '350429252210-hq617ss9idkeat0h66hbop59ul53mpnf.apps.googleusercontent.com')

        # Or, if multiple clients access the backend server:
        # idinfo = id_token.verify_oauth2_token(token, requests.Request())
        # if idinfo['aud'] not in [CLIENT_ID_1, CLIENT_ID_2, CLIENT_ID_3]:
        #     raise ValueError('Could not verify audience.')

        # If auth request is from a G Suite domain:
        # if idinfo['hd'] != GSUITE_DOMAIN_NAME:
        #     raise ValueError('Wrong hosted domain.')

        # ID token is valid. Get the user's Google Account ID from the decoded token.
        userid = idinfo['sub']
        service = build("calendar", "v3", credentials=userid)
        print('calendar service created successfully')
    except Exception as e:
        print(e)
        print(req_data['tokenId'])
        return None


@app.route("/gcal/add", methods=['POST'])
def create_event():
   # creates one hour event tomorrow 10 AM IST
   d = datetime.now().date()
   print(d)
   tomorrow = datetime(d.year, d.month, d.day, 10)+timedelta(days=1)
   start = tomorrow.isoformat()
   end = (tomorrow + timedelta(hours=1)).isoformat()

   event_result = service.events().insert(calendarId='primary',
                                          body={
                                              "summary": 'Automating calendar',
                                              "description": 'This is a tutorial example of automating google calendar with python',
                                              "start": {"dateTime": start, "timeZone": 'Asia/Kolkata'},
                                              "end": {"dateTime": end, "timeZone": 'Asia/Kolkata'},
                                          }
                                          ).execute()

   print("created event")
   print("id: ", event_result['id'])
   print("summary: ", event_result['summary'])
   print("starts at: ", event_result['start']['dateTime'])
   print("ends at: ", event_result['end']['dateTime'])


@app.route("/gcal/update")
def update_event():
    # update the event to tomorrow 9 AM IST
    d = datetime.now().date()
    tomorrow = datetime(d.year, d.month, d.day, 9)+timedelta(days=1)
    start = tomorrow.isoformat()
    end = (tomorrow + timedelta(hours=2)).isoformat()

    event_result = service.events().update(
        calendarId='primary',
        eventId='<place your event ID here>',
        body={
            "summary": 'Updated Automating calendar',
            "description": 'This is a tutorial example of automating google calendar with python, updated time.',
            "start": {"dateTime": start, "timeZone": 'Asia/Kolkata'},
                "end": {"dateTime": end, "timeZone": 'Asia/Kolkata'},
        },
    ).execute()

    print("updated event")
    print("id: ", event_result['id'])
    print("summary: ", event_result['summary'])
    print("starts at: ", event_result['start']['dateTime'])
    print("ends at: ", event_result['end']['dateTime'])
