# Usage Guide

Learn how to use **go-microx** to scaffold your next microservices project.

## Creating a New Project

The core command for **go-microx** is `new`, which launches an interactive CLI wizard.

```bash
go-microx new
```

### 1. Select a Project Template

Choose a template that best fit your industry needs:

- **E-commerce**: Includes services for products, orders, payments, and users.
- **Video Streaming**: Scaffolds services for transcoding, content delivery, and user subscriptions.
- **Food Delivery**: Set up a restaurant management and delivery tracking system.
- **Custom**: Start with a blank slate for your own custom architecture.

### 2. Enter Project Name

Provide a name for your project folder (e.g., `my-shop`).

### 3. Select an HTTP Router

Choose the web framework you prefer:

- **Gin**: High performance, feature-rich.
- **Fiber**: Express-inspired, extremely fast.
- **Echo**: Minimalist, high performance.
- **Chi**: Lightweight, idiomatic Go.
- **Mux**: Standard library compatible router.

### 4. Select a Database

Integrate a database out of the box:

- **PostgreSQL / MySQL**: Reliable SQL databases.
- **MongoDB**: Popular NoSQL document store.
- **SQLite**: Great for prototyping or light applications.
- **None**: For stateless services.

### 5. Authentication Method

Secure your services from day one:

- **JWT**: Stateless JSON Web Token authentication.
- **OAuth**: Integrate with external providers.
- **API Key**: Simple header-based security.

## Command Reference

| Command               | Description                          |
| :-------------------- | :----------------------------------- |
| `go-microx new`       | Launch the interactive project setup |
| `go-microx help`      | Display help for any command         |
| `go-microx --version` | Show current version                 |

## What's Next?

Once your project is created, navigate into the directory and check the individual service `README.md` files for specific run instructions!
