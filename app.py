from flask import Flask, redirect, url_for, session, request, render_template
from flask_cors import CORS
import requests

app = Flask(__name__)
app.debug = True
app.secret_key = 'development'
CORS(app)

from routes import *

if __name__ == '__main__':
    app.run()

