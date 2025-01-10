# Local Development Deployment

This document describes how to set up the local development environment for the project using Docker and Nginx.

### Prerequisites for Development

To develop this project locally, ensure you have the following tools and environments installed:

- **Docker** (with **WSL 2.0** for Windows)
- **Node.js** `22.11` or higher
- **Go** `1.23.4` or higher
- **Python** `3.13` or higher
- **.NET 6 Runtime** (with the *ASP.NET and Web Development* workload)

Additionally, use an IDE with support for the mentioned languages and **AWS Toolkit** integration. For .NET development, it's recommended to use **Visual Studio Community 2022** or the latest version.

## Getting Started

Follow these steps to set up the development environment on your local machine.

### 1. Initialize Docker Containers

To start the application and its dependencies locally, use Docker Compose:

```bash
docker-compose up --build -d
```

This will build and run all the necessary services defined in the `docker-compose.yml` file in detached mode.

### 2. Adding More Microservices

If you add new microservices to the `docker-compose.yml` file, don't forget to also add them to the `nginx.conf` file, specifying their respective domains.

### 3. Initialize Front-End Project

To run the front-end project, navigate to the `/FrontEnd/global-tune` directory and run it in development mode:

```bash
npm install
npm run dev
```

This will start the front-end application on port `3000`.

### 4. Ports in Use

By default, the following ports are used for local development. Ensure that they are not occupied by other processes on your machine. If necessary, change the ports in the `docker-compose.yml` file:

- `3000` – Front-end project
- `27017` – MongoDB
- `8080-8089` – Microservices
- `6379` – Redis

### Notes

- If you encounter any conflicts with the ports listed above, modify the port numbers in the `docker-compose.yml` file to use other available ports.
- Ensure that Docker and Nginx are properly configured for communication between the microservices and the front-end project.