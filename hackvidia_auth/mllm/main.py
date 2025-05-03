from concurrent.futures import ThreadPoolExecutor
import grpc
import mllm_pb2
import mllm_pb2_grpc

import torch
from transformers import AutoModel, AutoProcessor, LlavaForConditionalGeneration

from PIL import Image
import io

import os
import sys

MODEL_ENVIROSCAN_NAME = "llava-hf/llava-1.5-7b-hf"
MODEL_ENVIROSCAN_REVISION = "6ceb2ed33cb8f107a781c431fe2e61574da69369"

MODEL_LUMEN_NAME = "Qwen/Qwen2.5-7B-Instruct"
MODEL_LUMEN_REVISION = "a09a35458c702b33eeacc393d103063234e8bc28"

model_dtype = torch.bfloat16 if torch.cuda.is_bf16_supported() else torch.float16
class MLLMServicer(mllm_pb2_grpc.MllmService):
    def __init__(self):
        self.model = LlavaForConditionalGeneration.from_pretrained(
            MODEL_ENVIROSCAN_NAME,
            trust_remote_code=True,
            revision=MODEL_ENVIROSCAN_REVISION,
            torch_dtype=model_dtype,
            device_map="auto"
        )
        self.processor = AutoProcessor.from_pretrained(
            MODEL_ENVIROSCAN_NAME,
            trust_remote_code=True,
            revision=MODEL_ENVIROSCAN_REVISION,
            use_fast=True,
        )

    def ChatLumen(self, request, context):
        if self.model.config.name_or_path != MODEL_ENVIROSCAN_NAME:
            self.model = LlavaForConditionalGeneration.from_pretrained(
                MODEL_ENVIROSCAN_NAME,
                trust_remote_code=True,
                revision=MODEL_ENVIROSCAN_REVISION,
                torch_dtype=model_dtype,
                device_map="auto"
            )
            self.processor = AutoProcessor.from_pretrained(
                MODEL_ENVIROSCAN_NAME,
                trust_remote_code=True,
                revision=MODEL_ENVIROSCAN_REVISION,
                use_fast=True,
            )
        text = request.text
        content = []
        if request.text != "":
            content.append(
                {"type": "text", "text": f"""You are a Chat Assistant. Respond to the user's questions. The user is blind, so please be especially kind, clear, and descriptive in your language. 
                
                User: {text}"""},
            )
        
        conversation = [
            {
                "role": "user",
                "content": content
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
        generate_ids = generate_ids[:, inputs["input_ids"].shape[1]:]
        output = self.processor.batch_decode(generate_ids, skip_special_tokens=True)
        return mllm_pb2.MllmResponse(text=output[0])
    def Chat(self, request, context):
        if self.model.config.name_or_path != MODEL_ENVIROSCAN_NAME:
            self.model = LlavaForConditionalGeneration.from_pretrained(
                MODEL_ENVIROSCAN_NAME,
                trust_remote_code=True,
                revision=MODEL_ENVIROSCAN_REVISION,
                torch_dtype=model_dtype,
                device_map="auto"
            )
            self.processor = AutoProcessor.from_pretrained(
                MODEL_ENVIROSCAN_NAME,
                trust_remote_code=True,
                revision=MODEL_ENVIROSCAN_REVISION,
                use_fast=True,
            )
        text = request.text
        content = []
        if request.text != "":
            content.append({"type": "text", "text": text})
        
        print(len(request.image))
        if request.image != None and len(request.image) > 0:
            img_buffer = io.BytesIO(request.image)
            img = Image.open(img_buffer)
            content.append({"type": "image", "image": img})
        conversation = [
            {
                "role": "user",
                "content": content
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
        generate_ids = generate_ids[:, inputs["input_ids"].shape[1]:]
        output = self.processor.batch_decode(generate_ids, skip_special_tokens=True)
        return mllm_pb2.MllmResponse(text=output[0])

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