# Redis for Shopping Cart

A lightweight Redis image configured for the Shopping Cart service. This image uses the official Redis Alpine version for minimal size and efficient performance.

## Features
- Based on Redis 7.0 (Alpine).
- Configured to expose port 6379.
- Ready to run in Docker containers.

## Building the Redis Image
To build the Redis Docker image:
```bash
docker build -t redis-shopping-cart .
```

## Running the Redis Container

### Using a Docker Network
If the Redis container is part of a multi-container setup (e.g., with the Shopping Cart service), ensure that it runs on the same Docker network.

1. **Create a Docker Network:**
   ```bash
   docker network create my_network
   ```

2. **Run the Redis Container:**
   ```bash
   docker run --network my_network --name redis-cart -d -p 6379:6379 redis-shopping-cart
   ```

This command:
- Connects the Redis container to the `my_network` network.
- Names the container `redis-cart` for easier reference.
- Maps the host port `6379` to the container port `6379`.

## Exposing Redis Configuration
For advanced setups, you can mount a custom Redis configuration file. Update the `Dockerfile` or use volume mounting during runtime as needed.

## Future Improvements
- Migrate to a DynamoDB service

