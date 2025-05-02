from concurrent.futures import ThreadPoolExecutor
from transformers import pipeline

import grpc
import hazalert_pb2
import hazalert_pb2_grpc

MODEL_NAME = "depth-anything/Depth-Anything-V2-Large-hf"
MODEL_REVISION = "7581137eff8d4e94f6e796d3baea0e9fa79b22d2"

class MLLMServicer(mllm_pb2_grpc.MllmService):
    def __init__(self):
        self.pipe = pipeline(task="depth-estimation",
                model=MODEL_NAME,
                revision=MODEL_REVISION,
                use_fast=True)

    def Predict(self, request, context):
        print(request)
        return hazalert_pb2.HazalertResponse(nearest_distance=0.0)
    

import os
import sys

if __name__ == "__main__":
    server = grpc.server(ThreadPoolExecutor(max_workers=2))
    hazalert_pb2_grpc.add_MllmServiceServicer_to_server(
        MLLMServicer(), server
    )
    port = os.environ.get("SERVICE_PORT")
    if port == None:
        print("SERVICE_PORT not set")
        sys.exit(1)
    print(f"Listening on :{port}")
    server.add_insecure_port(f"[::]:{port}")
    server.start()
    server.wait_for_termination()