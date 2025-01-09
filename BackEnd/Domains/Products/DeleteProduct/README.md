# Delete Product Service

The Delete Product Service is a RESTful API developed in Go (Golang) that facilitates the delete management of products within the Products Domain of the Global Tune eCommerce platform.

## Architecture
<p align="center">
    <img alt="Edit Product architecture diagram" src="/assets/rest-admin-documentdb.webp"/>
</p>

## Getting Started for Local Testing

### Prerequisites
- Docker

## Running the Application

### Run Locally with Docker
#### Build the Docker Image
```bash
docker build -t delete-product-service .
```

#### Using Docker Network
If the MongoDB server is also running in a Docker container, use a Docker network for communication:
1. Create a Docker network:
   ```bash
   docker network create my_network
   ```
2. Start the MongoDB container:
   ```bash
   docker run --network my_network --name products-db -p PORT:PORT -e env_credentials -d products-db
   ```
3. Start the API container:
   ```bash
   docker run --network my_network -e MONGO_URI="mongodb://user:password@products-db:port" -p LOCAL_PORT:80 delete-product-service
   ```

#### Using an External MongoDB Endpoint
To connect to an external MongoDB server:
```bash
docker run -e MONGO_URI="mongodb://user:password@external-host:port" -p LOCAL_PORT:80 delete-product-service
```
Replace `user:password@external-host:port` with the MongoDB server details.
Replace `LOCAL_PORT` to a free port.

---

## API Endpoint

### DELETE /products/:id
Delete a product by ID.
- **URL Parameters:**
  - `id` (string): The SKU of the product.
- **Response:**
  - `200 OK`: Product deleted successfully.
  - `404 Not Found`: Product not found.
  - `500 Internal Server Error`: Server error.

**Example Request:**
```bash
curl -X DELETE http://localhost:8000/products/GTR-GIB-RD-001
```

## Notes
- Ensure MongoDB is running and accessible before starting the application.
- Use `gin.SetMode(gin.ReleaseMode)` for production to enable release mode for better performance.
- For security, avoid exposing sensitive MongoDB credentials in plain text; use a secret management tool or environment variables.

---