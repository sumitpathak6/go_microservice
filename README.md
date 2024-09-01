# 🚀 Orders Microservice in GoLang

## Overview

Welcome to the Orders Microservice! This project is a RESTful API built with GoLang, designed to handle CRUD operations for orders and items. The service is engineered for high performance, leveraging Go-Chi for efficient routing and Redis for fast data storage. Deployment is containerized using Docker, ensuring seamless cross-platform compatibility.

## Features

- **🔗 RESTful API:** Perform Create, Read, Update, and Delete operations on orders and items with clean, intuitive endpoints.
- **⚙️ Go-Chi Router:** Optimized routing for handling HTTP requests efficiently.
- **📦 Redis Integration:** Utilized as the primary data store with custom environment configurations for flexibility.
- **🐳 Dockerized:** The entire service runs smoothly in Docker containers, supporting cross-platform deployment.

## Project Structure

```plaintext
├── cmd/
│   └── main.go        # Entry point of the service
├── pkg/
│   ├── orders/        # Order-related operations
│   ├── items/         # Item-related operations
│   └── router/        # API routing logic using Go-Chi
├── config/
│   └── config.go      # Configuration settings for Redis and other environments
├── Dockerfile         # Docker setup for containerizing the service
├── go.mod             # Go module dependencies
└── README.md          # Project documentation
```
## Getting Started

### Prerequisites

- [GoLang](https://golang.org/doc/install) 1.16+
- [Docker](https://docs.docker.com/get-docker/)
- [Redis](https://redis.io/download)

### Installation

1. **Clone the repository:**
    ```bash
    git clone https://github.com/yourusername/orders-microservice.git
    cd orders-microservice
    ```

2. **Build the Docker image:**
    ```bash
    docker build -t orders-microservice .
    ```

3. **Run the Docker container:**
    ```bash
    docker run -d -p 8080:8080 --name orders-service orders-microservice
    ```

4. **Access the service:**
    Open your browser or use Postman to access the API at `http://localhost:8080`.

### Configuration

Configure the Redis server address and port in the `config/config.go` file:

```go
RedisAddress = "localhost:6379"
RedisPassword = ""
```

## API Endpoints

- **`GET /orders`** - Retrieve all orders
- **`POST /orders`** - Create a new order
- **`GET /orders/{id}`** - Retrieve a specific order
- **`PUT /orders/{id}`** - Update an existing order
- **`DELETE /orders/{id}`** - Delete an order

## Deployment

Deploying this service is as easy as running it in a Docker container. The provided Dockerfile ensures that all dependencies are installed, and the service is ready to run on any platform that supports Docker.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests for improvements and new features.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.


