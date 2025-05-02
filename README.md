# GoVision Backend

## File Structure
This project organized into the following directories:
- ```docker-compose/```:  Contains the Docker Compose files (docker-compose.yml and potentially others) to define and run the multi-container application.
- ```go/```: Contains the source code for the Go service(s).
- ```python/```: Contains the source code for the Python service(s).

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.   

### Prerequisites
- Docker
- Docker Compose

### Installation
1. Clone the repository
```$ git clone https://github.com/govision-hackvidia/hackvidia-backend.git
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

## Deployment

Before deploying, ensure you have the following installed:

* [Docker](https://docs.docker.com/get-docker/)
* [Docker Compose](https://docs.docker.com/compose/install/)

Follow these steps to deploy the project:
1. Navigate to go/ directory to build app:

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