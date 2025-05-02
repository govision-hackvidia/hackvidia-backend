FROM nvidia/cuda:12.8.1-cudnn-devel-ubuntu24.04

RUN apt-get update && \
    apt-get install -y python3-venv python3-pip python3-dev && \
    rm -rf /var/lib/apt/lists/* && \
    python3 -m venv ~/.pyvenv/

WORKDIR /app
COPY ./ ./
RUN ~/.pyvenv/bin/pip install -r ./requirements.txt
RUN ~/.pyvenv/bin/python3 main.py