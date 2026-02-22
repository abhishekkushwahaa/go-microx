# Project Structure

When you scaffold a project with **go-microx**, it generates a clean and modular directory structure designed for scalability. Below is a typical structure for a generated project:

```text
my-project/
├── user-service/        # Individual microservice directory
│   ├── main.go          # Entry point for the service
│   ├── Dockerfile       # Container definition
│   ├── config.yaml      # Service-specific configuration
│   └── internal/        # Private application and library code
├── product-service/     # Another microservice
├── api-gateway/         # Central gateway for routing requests
├── configs/             # Shared configuration files (optional)
├── docker-compose.yml   # Orchestration for local development
└── .gitignore           # Standard Go gitignore
```

## Core Directories

### Service Directories (`*-service/`)

Each microservice is contained within its own directory, promoting isolation. This allows you to scale, deploy, and manage services independently.

### `api-gateway/`

The API Gateway acts as the single entry point for all clients. It handles request routing, load balancing, and potentially authentication before forwarding requests to the internal services.

### `configs/`

Contains configuration templates or static configuration files that are shared across services or environment-specific.

### `docker-compose.yml`

Generated out of the box to allow you to spin up the entire microservices ecosystem (including databases) with a single command: `docker-compose up`.

## Design Pattern

The generated code follows a **Modular Monolith-to-Microservices** approach, ensuring that business logic is separated from infrastructure concerns, making the code easy to test and maintain.
