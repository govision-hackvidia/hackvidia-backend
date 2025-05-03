# GoVision Backend

## File Structure
This project organized into the following directories:
- ```docker-compose/```:  Contains the Docker Compose files (docker-compose.yml and potentially others) to define and run the multi-container application.
- ```go/```: Contains the source code for the Go service(s).
- ```python/```: Contains the source code for the Python service(s).

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. Make sure the server have CUDA driver installed.

### Prerequisites
- Docker 27.5
- Go 1.24
- Python 3.10

### Installation
1. Clone the repository
    ```bash
    git clone https://github.com/govision-hackvidia/hackvidia-backend.git
    ```

2. Install protoc for Go
    ```bash
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    ```

3. Install grpcio-tools for Python
    ```bash
    python3 -m pip install grpcio-tools
    ```

4. Update PATH for protoc
    ```bash
    export PATH="$PATH:$(go env GOPATH)/bin"
    ```

5. Compile protobuf files:
    ```bash
    make generate_mllm
    make generate_hazalert
    ```

6. Install python libraries:
    ```bash
    cd python/mllm/
    python -m pip install -r requirements.txt
    cd ../hazalert/
    python -m pip install -r requirements.txt

## Deployment (Non-Docker)

1. Run main_service in go/ directory:
    ```bash
    cd go/
    export SERVICE_HOST=0.0.0.0
    export SERVICE_PORT=8080
    export MLLM_HOST=0.0.0.0
    export MLLM_PORT=50051
    export HAZALERT_HOST=0.0.0.0
    export HAZALERT_PORT=50052
    go run .
    ```

2. Run mllm_service in python/ directory:
    ```bash
    cd python/
    export MLLM_PORT=50051
    python3 main.py
    ```

3. Run hazalert_service in python/ directory:
    ```bash
    cd python/
    export HAZALERT_PORT=50052
    python3 main.py
    ```

## Deployment

Before deploying, ensure you have the following installed:

* [Docker](https://docs.docker.com/get-docker/)
* [Docker Compose](https://docs.docker.com/compose/install/)

Follow these steps to deploy the project:
1. Build Dockerfile in go/ directory:

    ```bash
    cd go/
    sudo docker build -t govision_main_service:v1.0 .
    ```

2. Navigate to the docker-compose/main_service/ directory and start the core application services.

    ```bash
    cd docker-compose/main_service/
    sudo docker compose up -d
    ```

3. Navigate to the Caddy service's docker-compose/caddy/ directory and start the services defined there.

    ```bash
    cd docker-compose/caddy/
    sudo docker compose up -d
    ```

For mLLM and HazAlert Service, use non-docker deployment instruction.