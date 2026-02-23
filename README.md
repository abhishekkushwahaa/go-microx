<div align="center">

# go-microx

**The Ultimate Microservices Scaffolding Tool for Go**

[![Go Report Card](https://goreportcard.com/badge/github.com/abhishekkushwahaa/go-microx)](https://goreportcard.com/report/github.com/abhishekkushwahaa/go-microx)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/abhishekkushwahaa/go-microx)](https://github.com/abhishekkushwahaa/go-microx)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=flat-square&logo=linkedin)](https://linkedin.com/in/abhishekkushwahaa)
[![X](https://img.shields.io/badge/X-000000?style=flat-square&logo=x)](https://x.com/AbhishekKushwaa)
[![Docs](https://img.shields.io/badge/Documentation-FF4500?style=flat-square)](https://go-microx.vercel.app/docs/intro)

**go-microx** is a high-performance CLI tool built to eliminate the complexity of bootstrapping microservices architectures in Go.
Scaffold industry-ready project structures in minutes—not hours.

[Core Features](#core-features) • [Installation](#installation) • [Usage](#usage) • [Project Templates](#project-templates) • [Future Roadmap](#future-roadmap)

</div>

## Overview

Building microservices from scratch is repetitive and error-prone. **go-microx** provides a standardized, battle-tested foundation so you can focus on writing business logic instead of boilerplate. Whether you're building a massive e-commerce platform or a high-throughput streaming service, we've got the scaffolding ready.

## Core Features

| Feature                    | Description                                                                          |
| :------------------------- | :----------------------------------------------------------------------------------- |
| **Interactive CLI**        | Seamless experience using `PromptUI` for project configuration.                      |
| **Industry Templates**     | Ready-to-use scaffolding for E-commerce, Streaming, and more.                        |
| **Pluggable Stack**        | Choose your favorite Router (Gin, Fiber, Chi) and Database (Postgres, Mongo, MySQL). |
| **Authentication Options** | Built-in support for JWT, OAuth, and API Keys.                                       |
| **Cloud Ready**            | Pre-configured Docker and containerization out of the box.                           |
| **gRPC Support**           | Built-in scaffolding for high-performance RPC communication.                         |

## Installation

Install the **go-microx** binary directly to your `$GOPATH/bin`:

```bash
go install github.com/abhishekkushwahaa/go-microx/cmd/go-microx@latest
```

## Usage

### Create a New Project

Start the interactive wizard and follow the prompts to build your service.

```bash
go-microx new
```

### Interactive Flow Details

The CLI handles the heavy lifting, prompting you for:

1. **Template Selection:** Industry-specific architecture patterns.
2. **Project Naming:** Your unique service identifier.
3. **HTTP Router:** Gin, Fiber, Chi, Echo, or Standard Mux.
4. **Database:** PostgreSQL, MongoDB, MySQL, SQLite, or None.
5. **Authentication:** Secure your services from the start.

## Project Templates

### E-commerce

_Comprehensive setup for online marketplaces._

- User Management
- Product Catalog & Inventory
- Order Processing & Payments
- Shopping Cart Service
- API Gateway integration

### Video Streaming

_Built for high-traffic media delivery._

- Transcoding Service
- Content Delivery Logic
- Subscription & Analytics
- Auth & Media Uploads

### Food Delivery

_On-demand logistics architecture._

- Restaurant Management
- Order Tracking System
- Delivery Driver Assignment
- Real-time Payment Status

## Generated Structure

Example structure generated for an **E-commerce** project:

```text
myshop/
├── user-service/        # Auth & Profile management
├── product-service/     # Inventory & Catalog
├── order-service/       # Logic for transactions
├── payment-service/     # Gateway integrations
├── api-gateway/         # Unified entry point
├── configs/             # Shared configurations
└── Dockerfile           # Multi-stage build setup
```

## Future Roadmap

- [ ] **Service Mesh:** Native Istio & Linkerd integration.
- [ ] **Observability:** OpenTelemetry, Prometheus, & Grafana dashboarding.
- [ ] **Event-Driven:** Built-in Kafka, NATS, and RabbitMQ scaffolding.
- [ ] **Biometrics:** WebAuthn support for passwordless auth.
- [ ] **K8s:** Helm Chart generation and Kustomize support.
- [ ] **Cloud:** One-click deployment configs for AWS, GCP, and Azure.

## Contributing

We love community contributions! Check out our [Contributing Guidelines](CONTRIBUTING.md) to get started.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

<div align="center">
Built with interest by <a href="https://github.com/abhishekkushwahaa">Abhishek Kushwaha</a>
</div>
