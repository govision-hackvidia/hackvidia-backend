FROM python:3.12.10-slim-bookworm

WORKDIR /app
COPY ./ ./
RUN pip3 install -r ./requirements.txt
RUN python3 /app/main.py