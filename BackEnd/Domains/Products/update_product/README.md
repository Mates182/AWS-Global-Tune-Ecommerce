# UpdateProduct

A Gin-based microservice created with ginshot.

## Getting Started

These instructions will help you run the project on your local machine.

### Prerequisites

- Go 1.16 or higher

### Running the service

1. Start the server:
	Run the project locally
	```bash
	go run main.go
	```
	Or run the project in a Docker container
	```bash
	docker-compose up --build -d
	```
	```bash
	curl http://localhost:80/ping
	```

## API Endpoints

- GET /ping - Health check endpoint that returns "pong"

## Built With

- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [ginshot](https://github.com/yourusername/ginshot) - Project scaffolding tool
	
	