# Import modules
from app import app, calSetup
from flask_cors import CORS
import requests
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
    creds = req_data['code']
    try:
        service = build("calendar", "v3", credentials=creds)
        print('calendar service created successfully')
    except Exception as e:
        print(e)
        print(req_data['code'])
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
