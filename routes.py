from app import app
from flask import redirect, url_for, render_template, request, jsonify
from db import add_record, get_records, update_record, delete_record, id_generator


@app.route('/')
def index():
    return render_template('index.html')

@app.route('/api/feedback', methods=['POST'])
def recieve_feedback():
    response = request.json
    feedback_id = id_generator()
    feedback = {
        'id': feedback_id, 
        'email': response.get('email'), 
        'category': response.get('category'), 
        'feedback_text': response.get('feedback')}
    add_record("feedback", feedback)
    return jsonify({'message': 'feedback recieved successfully'})

