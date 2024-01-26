from dotenv import load_dotenv 
import cohere
import db
import os

load_dotenv()
co = cohere.Client(os.getenv('COHERE_API_KEY'))

def analyze_sentiment(text):
    response = co.classify(
        model="large",
        inputs=[text],
        examples = [
            {"text": "I absolutely love this product, it's amazing!", "label": "Positive"},
            {"text": "This is the worst experience I've ever had.", "label": "Negative"},
            {"text": "The product is okay, nothing special.", "label": "Neutral"},
            {"text": "I'm really pleased with the service.", "label": "Positive"},
            {"text": "I'm very disappointed with the quality.", "label": "Negative"},]
    )

    return response.classifications[0].prediction

# response = co.classify(
#     model='large',
#     inputs=["Your feedback text"],
#     examples=[{"text": "I love this product", "label": "Positive"},
#               {"text": "I hate this product", "label": "Negative"},
#               {"text": "This product is okay", "label": "Neutral"}]
# )
# sentiment = response.classifications[0].prediction
# print(f"Predicted sentiment: {sentiment}")