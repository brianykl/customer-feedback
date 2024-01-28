from concurrent import futures
import grpc
import nlp_service_pb2
import nlp_service_pb2_grpc

from cohere_util import analyze_sentiment, categorize_feedback

class NLPService(nlp_service_pb2_grpc.NLPServiceServicer):

    def AnalyzeSentiment(self, request, context):
        sentiment = analyze_sentiment(request.feedback_text)
        return nlp_service_pb2.SentimentResponse(sentiment=sentiment)

    def CategorizeFeedback(self, request, context):
        category = categorize_feedback(request.feedback_text)
        return nlp_service_pb2.CategoryResponse(category=category)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    nlp_service_pb2_grpc.add_NLPServiceServicer_to_server(NLPService(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print("running!")
    server.wait_for_termination()

if __name__ == "__main__":
    serve()


def get_nlp_analysis(feedback_text):
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = nlp_service_pb2_grpc.NLPServiceStub(channel)
        sentiment_response = stub.AnalyzeSentiment(nlp_service_pb2.FeedbackRequest(feedback_text=feedback_text))
        category_response = stub.CategorizeFeedback(nlp_service_pb2.FeedbackRequest(feedback_text=feedback_text))
    return sentiment_response.sentiment, category_response.category