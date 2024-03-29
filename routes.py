from app import app
from cohere_util import analyze_sentiment, categorize_feedback
from flask import redirect, url_for, render_template, request, jsonify
from db import add_record, id_generator
from nlp_service_server import get_nlp_analysis
import threading


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

    add_record("feedback", feedback)
    async_nlp_analysis(feedback_id, feedback_text)
    return jsonify({'message': 'feedback recieved successfully'})

def async_nlp_analysis(feedback_id, feedback_text):
    def task():
        sentiment, topic = get_nlp_analysis(feedback_text)
        feedback_analysis = {
            'id': feedback_id,
            'topic': topic,
            'sentiment': sentiment 
        }
        add_record("feedback_analysis", feedback_analysis)

    thread = threading.Thread(target=task)
    thread.start()
        
