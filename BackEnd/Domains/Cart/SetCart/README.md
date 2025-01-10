# Set Cart Service

The Set Cart Service is a RESTful API developed in Go (Golang) that facilitates the set management of shopping carts items within the Cart Domain of the Global Tune eCommerce platform. The service uses Redis as a temporary database, with plans to migrate to DynamoDB.

## Architecture
<p align="center">
    <img alt="Set Cart Service architecture diagram" src="/assets/set-cart-arch.webp"/>
</p>


## Features
- Retrieve a shopping cart by ID.
- Create or update a shopping cart.

## Getting Started for Local testing

### Prerequisites
- [Go](https://golang.org/dl/) 1.23.4 installed.
- Docker (optional for containerized deployment).

### Install Go Dependencies
Before running the application locally, install the required Go modules:
```bash
go mod tidy
```

### Environment Variables
Create a `.env` file in the root directory with the following variables:
```
REDIS_ADDR=endpoint:port
REDIS_PASSWORD=password
```
Replace `endpoint:port` and `password` with your Redis server details.

## Running the Application

### Run Locally
Run the application directly using:
```bash
go run main.go
```
Ensure your Redis server is running and accessible.

### Run with Docker
#### Build the Docker Image
```bash
docker build -t cart-rest .
```

#### Using Docker Network
If your Redis server is also running in a Docker container, use a Docker network for communication:
1. Create a Docker network:
   ```bash
   docker network create my_network
   ```
2. Start the Redis container:
   ```bash
   docker run --network my_network --name redis -p 6379:6379 -d redis
   ```
3. Start the API container:
   ```bash
   docker run --network my_network -e REDIS_ADDR="redis:6379" -e REDIS_PASSWORD="" -p 8080:8080 cart-rest
   ```

#### Using an External Redis Endpoint
To connect to an external Redis server:
```bash
docker run -e REDIS_ADDR="endpoint:port" -e REDIS_PASSWORD="password" -p 8080:8080 cart-rest
```
Replace `endpoint:port` and `password` with your Redis server details.

## API Endpoint

### POST /cart/
Create or update a shopping cart.
- **Request Body:**
```json
{
  "id": "123",
  "items": [
    { "product_id": "product_1", "quantity": 2 },
    { "product_id": "product_2", "quantity": 3 }
  ]
}
```
- **Response:**
  - 201: Cart created successfully.
  - 400: Invalid request payload.
  - 500: Internal server error.

**Example Request:**
```bash
curl -X POST http://localhost:8080/cart/ \
-H "Content-Type: application/json" \
-d '{
  "id": "123",
  "items": [
    { "product_id": "product_1", "quantity": 2 },
    { "product_id": "product_2", "quantity": 3 }
  ]
}'
```
**Example Response:**
```json
{
  "message": "Cart created successfully"
}
```

## Notes
- Ensure Redis is running and accessible before starting the application.
- In production, use `gin.SetMode(gin.ReleaseMode)` to enable release mode for better performance.
- For security, avoid trusting all proxies; configure trusted proxies explicitly.
