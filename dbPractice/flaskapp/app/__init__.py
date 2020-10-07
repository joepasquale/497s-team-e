import os, json
from flask import Flask
from flask_sqlalchemy import SQLAlchemy

db = SQLAlchemy()

class Users(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    email = db.Column(db.String(100), unique=True, nullable=False)
    password = db.Column(db.String(100), nullable=False)
    
    def __repr__(self):
        return f"Users('{self.email}, '{self.password}')"

def create_app(config):
    app = Flask(__name__)
    app.config.from_object(config)
    db.init_app(app)
    return app