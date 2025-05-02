from concurrent.futures import ThreadPoolExecutor
import grpc
import mllm_pb2
import mllm_pb2_grpc

import torch
from transformers import AutoProcessor, LlavaForConditionalGeneration


MODEL_NAME = "llava-hf/llava-1.5-7b-hf"
MODEL_REVISION = "6ceb2ed33cb8f107a781c431fe2e61574da69369"

model_dtype = torch.bfloat16 if torch.cuda.is_bf16_supported() else torch.float16
class MLLMServicer(mllm_pb2_grpc.MllmService):
    def __init__(self):
        self.model = LlavaForConditionalGeneration.from_pretrained(
            MODEL_NAME,
            trust_remote_code=True,
            revision=MODEL_REVISION,
            torch_dtype=model_dtype,
            device_map="auto"
        )
        self.processor = AutoProcessor.from_pretrained(
            MODEL_NAME,
            trust_remote_code=True,
            revision=MODEL_REVISION,
            use_fast=True,
        )
    def Chat(self, request, context):
        print(request)
        conversation = [
            {
                "role": "user",
                "content": [
                    {"type": "image", "url": "https://www.ilankelman.org/stopsigns/australia.jpg"},
                    {"type": "text", "text": "What is shown in this image?"},
                ],
            },
        ]

        inputs = self.processor.apply_chat_template(
            conversation,
            add_generation_prompt=True,
            tokenize=True,
            return_dict=True,
            return_tensors="pt"
        ).to(self.model.device, model_dtype)

        # Generate
        generate_ids = self.model.generate(**inputs, max_new_tokens=128, do_sample=True, min_p=.02, repetition_penalty=1.02)
        outputs = self.processor.batch_decode(generate_ids, skip_special_tokens=True)
        print(outputs)
        return mllm_pb2.MllmResponse(text=outputs)
    
import os
import sys

if __name__ == "__main__":
    server = grpc.server(ThreadPoolExecutor(max_workers=2))
    mllm_pb2_grpc.add_MllmServiceServicer_to_server(
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