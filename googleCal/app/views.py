# Import modules
from app import app
from flask_cors import CORS
import requests
from flask import Flask, request, jsonify

# Define the API Key, Define the Endpoint, Define the Header
API_KEY = 
ENDPOINT = 
HEADERS = {'Authorization': 'bearer %s' % API_KEY}

CORS(app)


@app.route("/gcal")
def index():
    return "This is a test for the gCal API"

