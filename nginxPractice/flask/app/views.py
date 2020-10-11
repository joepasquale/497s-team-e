# Import modules
from app import app
import sys

print(sys.path)

from flask_cors import CORS
import requests
from flask import Flask, request, jsonify

# Define the API Key, Define the Endpoint, Define the Header
API_KEY = '9DtY40amt8NwFuCTtVPweR9w10EJ2XYXi47mTBC4fXeIR2wBvJkLFKi5n20PO-OdgcGXoolgBqITpxyd8qh194tI7xM5mxjJYg3gURkgvAGzLq4GN_aIjcSGdLdzX3Yx'
ENDPOINT = 'https://api.yelp.com/v3/businesses/search'
HEADERS = {'Authorization': 'bearer %s' % API_KEY}

CORS(app)

@app.route("/")
def index():
    return "This is a test"

@app.route("/yelp")
def yelp_resp():

    terms = request.args.get('terms')
    location = request.args.get('location')

    PARAMETERS = {'terms': terms,
    'limit': 10,
    'location': location}

    # Make a request to the yelp api
    response = requests.get(url = ENDPOINT, params= PARAMETERS, headers= HEADERS)

    # Convert JSON string to dictionary
    business_data = response.json()
    businessArr = []

    for biz in business_data['businesses']:
        businessArr.append(biz['name'])

    businessStr = ' '.join(map(str, businessArr))
    return(businessStr)