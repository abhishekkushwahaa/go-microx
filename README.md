# go-microx

<a href="https://linkedin.com/in/abhishekkushwahaa">![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=flat-square&logo=linkedin)</a>
<a href="https://x.com/AbhishekKushwaa">![X](https://img.shields.io/badge/X-000000?style=flat-square&logo=x)</a>
<a href="https://abhishekkushwaha.tech">![Website](https://img.shields.io/badge/Website-FF4500?style=flat-square)</a>
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

**go-microx** is a powerful CLI tool designed to streamline the creation of scalable microservices architectures in Golang. It provides a well-structured setup tailored for various industries, including e-commerce, video streaming, food delivery, and more.

## Features

- **Interactive CLI with PromptUI**: User-friendly selection for project templates, database choices, and configurations.
- **Automated Boilerplate Generation**: Quickly scaffold a complete microservices project.
- **Multiple Project Templates**: Choose from predefined templates like E-commerce, Video Streaming, Food Delivery, and more.
- **Database Support**: Select from PostgreSQL, MongoDB, MySQL, SQLite, or no database.
- **Authentication Options**: Supports JWT, OAuth, and API Key authentication mechanisms.
- **Pre-configured Docker Setup**: Seamless containerization with Docker.
- **Customizable**: Extend and modify configurations based on your project needs.

## Installation

To install go-microx, use the following command:

```sh
go install github.com/abhishekkushwahaa/gomicrox@latest
```

## Usage

### Create a New Microservices Project

Run the following command to start an interactive project setup:

```sh
gomicrox new
```

### Interactive CLI Flow

The CLI will guide you through project creation using **promptui**, displaying interactive options for selecting a template, naming your project, and choosing a database.

#### Example Interaction:

```
? Select a project template:
  ▸ E-commerce
    Video Streaming
    Food Delivery
    Custom

? Enter project name: myshop

? Select a database:
  ▸ PostgreSQL
    MongoDB
    MySQL
    SQLite
    None

🚀 Creating project: myshop
📦 Type: E-commerce
🫙 Database: PostgreSQL
✅ Project myshop has been successfully created!
```

### Available Project Templates

- **E-commerce**: Includes user management, product catalog, order processing, payments, and cart service.
- **Video Streaming**: Includes user authentication, video upload service, transcoding service, content delivery, and analytics.
- **Food Delivery**: Includes restaurant management, order tracking, delivery assignment, and payment service.

### Project Structure

Upon execution, go-microx generates the following structure based on the selected template:

```
myshop/ (E-commerce)
├── user-service/
│   ├── main.go
│   ├── Dockerfile
│   ├── config.yaml
├── product-service/
├── order-service/
├── payment-service/
├── cart-service/
├── api-gateway/
├── configs/
```

```
myvideo/ (Video Streaming)
├── user-service/
├── video-upload-service/
├── transcoding-service/
├── content-delivery-service/
├── analytics-service/
├── api-gateway/
```

## Roadmap

- **More Industry Templates**: Add support for FinTech, SaaS, and more.
- **gRPC Integration**: Support for high-performance RPC communication.
- **CI/CD Templates**: Pre-configured pipelines for seamless deployment.
- **Kubernetes Helm Charts**: Simplified Kubernetes deployments.
- **Cloud Deployment Support**: AWS, GCP, and Azure integrations.

## Contributing

We welcome contributions from the community! Feel free to **submit issues, feature requests, or pull requests** to help improve **go-microx**.

### How to Contribute:

1. Fork the repository and create a new branch.
2. Implement your changes with clear commit messages.
3. Submit a pull request with a detailed description.

## License

**go-microx** is released under the **MIT License**. See [LICENSE](LICENSE) for more details.
