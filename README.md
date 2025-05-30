## Golang Dagger

A Go server built with the Gin web framework. This project use [Dagger](https://dagger.io) for CI/CD workflows and utilizes GitHub Actions as the runtime.

### 🚀 Project Overview

This API provides basic CRUD operations for managing items.

#### 📘 Available API Routes

- `POST /items` – Create a new item  
- `GET /items` – Retrieve all items  
- `GET /items/:id` – Retrieve an item by ID  
- `PUT /items/:id` – Update an existing item by ID  
- `DELETE /items/:id` – Delete an item by ID  

### ⚙️ CI/CD

The CI/CD pipeline is defined using Dagger, with GitHub Actions serving as the runtime environment. This setup enables reproducible builds, testing, and deployment with minimal configuration overhead. You can find the Dagger configuration in the [.dagger](.dagger) directory.

### 📜 License

This project is licensed under the Apache-2.0 License. See the [LICENSE](LICENSE) file for full details.

### 🛡 Security

If you discover a security vulnerability in this project, please refer to our [SECURITY](SECURITY.md) policy for guidelines on responsible disclosure.