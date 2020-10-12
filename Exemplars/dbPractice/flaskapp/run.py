from app import create_app
from app import Users, db
from flask import request, jsonify
import os, json, datetime
PORT = os.environ.get('PORT')
POSTGRES_URL = os.environ.get('POSTGRES_URL')
POSTGRES_USER = os.environ.get('POSTGRES_USER')
POSTGRES_PASSWORD = os.environ.get('POSTGRES_PASSWORD')
POSTGRES_DB = os.environ.get("POSTGRES_DB")

class config:
    SQLALCHEMY_DATABASE_URI='postgresql+psycopg2://{user}:{pw}@{url}/{db}'.format(user=POSTGRES_USER,pw=POSTGRES_PASSWORD,url=POSTGRES_URL,db=POSTGRES_DB)
    SQLALCHEMY_TRACK_MODIFICATIONS=False
    
app = create_app(config)

@app.route('/')
def index():
    return "Hello World"

@app.route('/postgres/user', methods=["GET","POST","DELETE", "PATCH"])
def user():
    if request.method == 'GET':
        query = request.args
        print(query)
        data = Users.query.filter_by(email=query["email"]).first()
        return str(data)
    
    data = request.args
    if request.method == 'POST':
        user = Users(email=data['email'], password=data['password'])
        db.session.add(user)
        db.session.commit()
        return 'User has been added'

    if request.method == 'PATCH':
        pass

    if request.method == 'DELETE':
        pass

if __name__ == '__main__':
    with app.app_context():
        db.drop_all()
        db.create_all()
    app.run(debug=True, host='0.0.0.0', port=int(PORT))