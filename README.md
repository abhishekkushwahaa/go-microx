# go-microx

<a href="https://linkedin.com/in/abhishekkushwahaa">![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=flat-square&logo=linkedin)</a>
<a href="https://x.com/AbhishekKushwaa">![X](https://img.shields.io/badge/X-000000?style=flat-square&logo=x)</a>
<a href="https://abhishekkushwaha.tech">![Website](https://img.shields.io/badge/Website-FF4500?style=flat-square)</a>
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

**go-microx** is a powerful CLI tool designed to streamline the creation of scalable microservices architectures in Golang. It provides a well-structured setup tailored for various industries, including e-commerce, video streaming, food delivery, and more.

## Features

- **Automated Boilerplate Generation**: Quickly scaffold a complete microservices project.
- **Multiple Project Templates**: Choose from predefined templates like E-commerce, Video Streaming, Food Delivery, and more.
- **Database Support**: Choose from PostgreSQL, MongoDB, MySQL, and more.
- **Authentication Options**: Supports JWT, OAuth, and API Key authentication mechanisms.
- **Pre-configured Docker Setup**: Seamless containerization with Docker.
- **Customizable**: Extend and modify configurations based on your project needs.

## Installation

To install go-microx, use the following command:

```sh
go get github.com/abhishekkushwahaa/gomicrox@latest
```

## Usage

### Create a New Microservices Project

To generate a new microservices project, run:

```sh
gomicrox new <project-type> <project-name> --db=postgres --auth=jwt
```

#### Example:

```sh
gomicrox new ecommerce myshop --db=postgres --auth=jwt
```

```sh
gomicrox new videostream myvideo --db=mongodb --auth=oauth
```

### Available Project Templates

- **E-commerce**: Includes user management, product catalog, order processing, payments, and cart service.
- **Video Streaming**: Includes user authentication, video upload service, transcoding service, content delivery, and analytics.
- **Food Delivery**: Includes restaurant management, order tracking, delivery assignment, and payment service.

### Project Structure

Upon execution, go-microx generates the following structure based on the selected template:

```
myshop/ (E-commerce)
├── user/
│   ├── main.go
│   ├── Dockerfile
│   ├── config.yaml
├── product/
├── order/
├── payment/
├── cart/
├── api-gateway/
├── configs/
```

```
myvideo/ (Video Streaming)
├── user/
├── video-upload/
├── transcoding/
├── content-delivery/
├── analytics/
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
