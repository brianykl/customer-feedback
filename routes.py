from app import app
from flask import redirect, url_for, render_template, request, jsonify


@app.route('/')
def index():
    return render_template('index.html')

@app.route('/api/feedback', methods=['POST'])
def recieve_feedback():
    response = request.json
    feedback = response.get('feedback')
    # process feedback, save to database
    print("hoorah")
    return jsonify({'message': 'feedback recieved successfully'})

# @app.route('/submit', methods = ['POST'])
# def submit():

#     return redirect(url_for('/'))