# Edit Product Service

The Edit Product Service is a RESTful API developed in Go (Golang) that facilitates the creation and management of products within the Products Domain of the Global Tune eCommerce platform. This service handles product data entry, including SKU, pricing, stock levels, and other key attributes.

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
docker build -t edit-product-service .
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
   docker run --network my_network -e MONGO_URI="mongodb://user:password@products-db:port" -p LOCAL_PORT:80 edit-product-service
   ```

#### Using an External MongoDB Endpoint
To connect to an external MongoDB server:
```bash
docker run -e MONGO_URI="mongodb://user:password@external-host:port" -p LOCAL_PORT:80 edit-product-service
```
Replace `user:password@external-host:port` with the MongoDB server details.
Replace `LOCAL_PORT` to a free port.

---

## API Endpoint

### PATCH /products/:id
Add a new product.
- **Request Body:**
```json
{
  "atribute": value,
}

```
- **Response:**
  - `201 Created`: "Product updated successfully"
  - `400 Bad Request`: Failed to update product.
  - `404 Not Found`: Product not found.
  - `500 Internal Server Error`: Server error.

**Example Request:**
```bash
curl -X POST http://localhost:8000/products/GTR-GIB-RD-001 \
-H "Content-Type: application/json" \
-d '{
  "price": 1500,
  "stock": 10
}
'
```

## Notes
- Ensure MongoDB is running and accessible before starting the application.
- Use `gin.SetMode(gin.ReleaseMode)` for production to enable release mode for better performance.
- For security, avoid exposing sensitive MongoDB credentials in plain text; use a secret management tool or environment variables.

---