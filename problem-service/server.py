from concurrent import futures
import grpc
import problem_pb2
import problem_pb2_grpc
import problem_generation

# use lock / mutex etc


class ProblemServicer(problem_pb2_grpc.ProblemServicer):
    def GenerateProblem(self, request, context):
        print("Generating problem with type: {}".format(request.type))
        resp = problem_pb2.ImageMessage()
        resp.type = request.type

        img = problem_generation.gen(request.type)
        resp.img = img
        return resp


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    problem_pb2_grpc.add_ProblemServicer_to_server(ProblemServicer(), server)
    print("Starting Problem Service at :8040")
    server.add_insecure_port("[::]:8040")
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
