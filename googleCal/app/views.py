# Import modules
from app import app, calSetup
from flask_cors import CORS
import requests
from google.oauth2 import client
from google.auth.transport import requests
from googleapiclient.discovery import build
from flask import Flask, request, jsonify
from datetime import datetime, timedelta

CORS(app)

service = None


@app.route("/gcal/auth", methods=['POST'])
def index():
    print("This is a test for the Google API")


@app.route("/gcal/add", methods=['POST'])
def create_event():
    """
    creds = client.AccessTokenCredentials(
        'ACCESS_TOKEN',
        'my-calendar-bot/1.0'
    )
    service = build('calendar', 'v3', credentials=creds)
    settings = service.settings().list().execute()
    print(settings)
    """
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
