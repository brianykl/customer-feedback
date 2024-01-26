from dotenv import load_dotenv 
import cohere
from cohere.responses.classify import Example
import db
import os

load_dotenv()
co = cohere.Client(os.getenv('COHERE_API_KEY'))
examples = [Example("I absolutely love this product, it's amazing!", "Positive"),
            Example("This is the worst experience I've ever had.", "Negative"),
            Example("The product is okay, nothing special.", "Neutral"),
            Example("Seems whatever", "Neutral"),
            Example("I'm really pleased with the service.", "Positive"),
            Example("I'm very disappointed with the quality.", "Negative")]

def analyze_sentiment(text):
    response = co.classify(
        inputs=[text],
        examples = examples
    )

    return response.classifications[0].prediction

def categorize_feedback(text):
    prompt = f'feedback: {text}\nidentify the main category of this feedback:'
    response = co.generate(
        prompt=prompt,
        max_tokens=5
        )
    generated_text = response.generations[0].text.strip()
    return generated_text
