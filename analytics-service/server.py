from concurrent import futures
import grpc
import analytics_pb2
import analytics_pb2_grpc
import analytics


class AnalyticsServicer(analytics_pb2_grpc.AnlyticsServicer):
    def GradeAccuracy(self, request, context):
        return super().GradeAccuracy(request, context)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    analytics_pb2_grpc.add_AnlyticsServicer_to_server(
        AnalyticsServicer(), server)
    print("Starting Analytics Service at :8050")
    server.add_insecure_port("localhost:8050")
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
