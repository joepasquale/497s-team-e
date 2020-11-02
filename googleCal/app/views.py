# Import modules
from app import app
from flask_cors import CORS
import requests
from flask import Flask, request, jsonify

# Define the API Key, Define the Endpoint, Define the Header
API_KEY = 'AIzaSyDdVUMKArJ1GpFar9QbxuZWvFLzROVqI0g'
CLIENT_ID = '350429252210-abjunlthml9hfblfhv9rb949s1vra1u0.apps.googleusercontent.com'
SCOPE = "https://www.googleapis.com/auth/calendar.readonly""
HEADERS = {'Authorization': 'bearer %s' % API_KEY}

CORS(app)


@app.route("/gcal")
def index():
    return "This is a test for the gCal API"

