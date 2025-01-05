### Products REST API (CRUD) - Golang

A RESTful API for managing products, written in Golang. The service uses MongoDB for data storage.

---

## Features
- Retrieve all products.
- Retrieve a product by SKU.
- Add a new product.
- Delete a product by SKU.

---

## Getting Started for Local Testing

### Prerequisites
- [Go](https://golang.org/dl/) 1.23.4 installed.
- Docker (for MongoDB containerized deployment, optional).

### Install Go Dependencies
Before running the application locally, install the required Go modules:
```bash
go mod tidy
```

### Environment Variables
Create a `.env` file in the root directory with the following variable:
```
MONGO_URI=mongodb://user:password@endpoint:port
```
Replace the `MONGO_URI` values for connecting to the MongoDB instance.

---

## Running the Application

### Run Locally
Run the application directly using:
```bash
go run main.go
```
Ensure your MongoDB server is running and accessible.

### Run with Docker
#### Build the Docker Image
```bash
docker build -t products-rest-api .
```

#### Using Docker Network
If your MongoDB server is also running in a Docker container, use a Docker network for communication:
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
   docker run --network my_network -e MONGO_URI="mongodb://user:password@products-db:port" -p 8000:8000 products-rest-api
   ```

#### Using an External MongoDB Endpoint
To connect to an external MongoDB server:
```bash
docker run -e MONGO_URI="mongodb://user:password@external-host:port" -p 8000:8000 products-rest-api
```
Replace `user:password@external-host:port` with the MongoDB server details.

---

## API Endpoints

### GET /products/
Retrieve all products.
- **Response:**
  - `200 OK`: Returns all products.
  - `404 Not Found`: No products found.
  - `500 Internal Server Error`: Server error.

**Example Request:**
```bash
curl -X GET http://localhost:8000/products/
```

---

### GET /products/:sku
Retrieve a product by its SKU.
- **URL Parameters:**
  - `sku` (string): The SKU of the product.
- **Response:**
  - `200 OK`: Returns the product details.
  - `404 Not Found`: Product not found.
  - `500 Internal Server Error`: Server error.

**Example Request:**
```bash
curl -X GET http://localhost:8000/products/GTR-GIB-RD-001
```

---

### POST /products/
Add a new product.
- **Request Body:**
```json
{
  "sku": "GTR-GIB-RD-002",
  "title": "New Product",
  "description": "A description",
  "price": 200,
  "stock": 10,
  "brand": "Brand",
  "category": "Category",
  "tags": ["tag1", "tag2"],
  "thumbnail": "url-to-image",
  "images": ["url1", "url2"],
  "warranty": "2 years",
  "weight": 1.5
}
```
- **Response:**
  - `201 Created`: Product created successfully.
  - `400 Bad Request`: Invalid request payload or duplicate SKU.
  - `500 Internal Server Error`: Server error.

**Example Request:**
```bash
curl -X POST http://localhost:8000/products/ \
-H "Content-Type: application/json" \
-d '{
  "sku": "GTR-GIB-RD-002",
  "title": "New Product",
  "description": "A description",
  "price": 200,
  "stock": 10,
  "brand": "Brand",
  "category": "Category",
  "tags": ["tag1", "tag2"],
  "thumbnail": "url-to-image",
  "images": ["url1", "url2"],
  "warranty": "2 years",
  "weight": 1.5
}'
```

---

### DELETE /products/:sku
Delete a product by SKU.
- **URL Parameters:**
  - `sku` (string): The SKU of the product.
- **Response:**
  - `200 OK`: Product deleted successfully.
  - `404 Not Found`: Product not found.
  - `500 Internal Server Error`: Server error.

**Example Request:**
```bash
curl -X DELETE http://localhost:8000/products/GTR-GIB-RD-001
```

---

## Notes
- Ensure MongoDB is running and accessible before starting the application.
- Use `gin.SetMode(gin.ReleaseMode)` for production to enable release mode for better performance.
- For security, avoid exposing sensitive MongoDB credentials in plain text; use a secret management tool or environment variables.

---

## Future Improvements
- Add authentication and authorization mechanisms.
- Implement advanced search and filtering options for products.
- Implement other needed requests.