from app import app
from cohere_util import analyze_sentiment, categorize_feedback
from flask import redirect, url_for, render_template, request, jsonify
from db import add_record, get_records, update_record, delete_record, id_generator


@app.route('/')
def index():
    return render_template('index.html')

@app.route('/api/feedback', methods=['POST'])
def recieve_feedback():
    response = request.json
    feedback_id = id_generator()
    feedback_text = response.get('feedback')
    
    feedback = {
        'id': feedback_id, 
        'email': response.get('email'), 
        'category': response.get('category'), 
        'feedback_text': feedback_text}
    
    feedback_analysis = {
        'id': feedback_id,
        'topic': categorize_feedback(feedback_text),
        'sentiment': analyze_sentiment(feedback_text)
    }

    add_record("feedback", feedback)
    add_record("feedback_analysis", feedback_analysis)
    return jsonify({'message': 'feedback recieved successfully'})

