# Import modules
from app import app
from flask import Flask, request, jsonify
import pymongo
import os



# Access db containerName:portNum
myclient = pymongo.MongoClient("mongodb://mongo:27017/")
mydb = myclient["mydatabase"]
mycol = mydb["customers"]

mydict = { "name": "John", "address": "Highway 37" }

print("BEFORE INSERT")
x = mycol.insert_one(mydict)
print("AFTER INSERT")

@app.route("/group")
def index():
    return "Hello from Groups Service"

@app.route("/group/get")
def getGroup():
    result = mycol.find({"name": "John"})
    for doc in result:
        print(doc)
    return "Did it work"

# client = pymongo.MongoClient(
#     os.environ['DB_PORT_27017_TCP_ADDR'],
#     27017)
# db = client.tododb


# @app.route('/')
# def todo():

#     _items = db.tododb.find()
#     items = [item for item in _items]

#     return render_template('todo.html', items=items)


# @app.route('/new', methods=['POST'])
# def new():

#     item_doc = {
#         'name': request.form['name'],
#         'description': request.form['description']
#     }
#     db.tododb.insert_one(item_doc)

#     return redirect(url_for('todo'))

