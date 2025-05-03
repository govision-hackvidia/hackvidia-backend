from concurrent.futures import ThreadPoolExecutor
from transformers import pipeline, AutoProcessor

import grpc
import hazalert_pb2
import hazalert_pb2_grpc

import torch
from PIL import Image
import io

import os
import sys

MODEL_NAME = "depth-anything/Depth-Anything-V2-Metric-Indoor-Large-hf"
MODEL_REVISION = "d2fc6a93601aabb1139a3bf0ebfcb4e89c67817f"

# model_dtype = torch.float16
# model_dtype = torch.bfloat16 if torch.cuda.is_bf16_supported() else torch.float16
class HazalertService(hazalert_pb2_grpc.HazalertService):
    def __init__(self):
        processor = AutoProcessor.from_pretrained(
            MODEL_NAME,
            trust_remote_code=True,
            # torch_dtype=model_dtype,
            revision=MODEL_REVISION,
            device_map="auto",
            use_fast=True,
        )
        self.pipe = pipeline(task="depth-estimation",
            model=MODEL_NAME,
            image_processor=processor,
            revision=MODEL_REVISION,
            # torch_dtype=model_dtype,
        )

    def Predict(self, request, context):
        img_buffer = io.BytesIO(request.image)
        img = Image.open(img_buffer)
        # img_torch = pil_to_tensor(img)
        # img_torch = torch.astype(model_dtype)
        depth = self.pipe(img)["predicted_depth"]
        nearest_depth = torch.min(depth)
        print(nearest_depth.item())
        return hazalert_pb2.HazalertResponse(nearest_distance=nearest_depth.item())

import os
import sys

if __name__ == "__main__":
    server = grpc.server(ThreadPoolExecutor(max_workers=2))
    hazalert_pb2_grpc.add_HazalertServiceServicer_to_server(
        HazalertService(), server
    )
    port = os.environ.get("SERVICE_PORT")
    if port == None:
        print("SERVICE_PORT not set")
        sys.exit(1)
    print(f"Listening on :{port}")
    server.add_insecure_port(f"[::]:{port}")
    server.start()
    server.wait_for_termination()