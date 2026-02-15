# 🚀 GoNews

<div align="center">

[![GitHub stars](https://img.shields.io/github/stars/muhammadidrusalawi/gonews?style=for-the-badge)](https://github.com/muhammadidrusalawi/gonews/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/muhammadidrusalawi/gonews?style=for-the-badge)](https://github.com/muhammadidrusalawi/gonews/network)
[![GitHub issues](https://img.shields.io/github/issues/muhammadidrusalawi/gonews?style=for-the-badge)](https://github.com/muhammadidrusalawi/gonews/issues)
[![GitHub license](https://img.shields.io/github/license/muhammadidrusalawi/gonews?style=for-the-badge)](LICENSE)

**A robust and scalable RESTful API for news management, built with Go and Fiber.**

</div>

## 📖 Overview

GoNews is a high-performance backend API designed for managing news articles and user authentication. Leveraging the power of Go and the speed of the Fiber web framework, it provides a solid foundation for applications requiring efficient data handling and secure user access. The project integrates PostgreSQL for data persistence using GORM, JWT for authentication, and robust input validation to ensure data integrity.

## ✨ Features

-   🎯 **RESTful API**: Clean and efficient API endpoints for managing news articles.
-   🔐 **JWT Authentication**: Secure user registration, login, and protected routes using JSON Web Tokens.
-   📰 **News Management**: Functionality for creating, reading, updating, and deleting news articles.
-   💾 **PostgreSQL Database**: Persistent storage for news data and user accounts, managed with GORM ORM.
-   ⚙️ **Environment Configuration**: Easy setup and management of application settings using `.env` files.
-   ✅ **Input Validation**: Ensures data integrity and security for incoming API requests.
-   📝 **Structured Logging**: Integrated `logrus` for clear and insightful application logging.

## 🛠️ Tech Stack

**Backend:**
[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![Fiber](https://img.shields.io/badge/Fiber-000000?style=for-the-badge&logo=fiber&logoColor=white)](https://gofiber.io/)
[![GORM](https://img.shields.io/badge/GORM-000000?style=for-the-badge&logo=go&logoColor=white)](https://gorm.io/)
[![JWT](https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=json-web-tokens&logoColor=white)](https://jwt.io/)
[![Bcrypt](https://img.shields.io/badge/Bcrypt-Password%20Hashing-62BEE5?style=for-the-badge)](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
[![Godotenv](https://img.shields.io/badge/Godotenv-Env%20Mgmt-orange?style=for-the-badge)](https://github.com/joho/godotenv)
[![Logrus](https://img.shields.io/badge/Logrus-Logging-red?style=for-the-badge)](https://github.com/sirupsen/logrus)
[![Validator](https://img.shields.io/badge/Go--Playground%20Validator-Validation-blue?style=for-the-badge)](https://github.com/go-playground/validator)

**Database:**
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)](https://www.postgresql.org/)

## 🚀 Quick Start

Follow these steps to get GoNews up and running on your local machine.

### Prerequisites
Before you begin, ensure you have the following installed:
-   **Go**: Version 1.16 or higher
-   **PostgreSQL**: Database server (e.g., v13+)

### Installation

1.  **Clone the repository**
    ```bash
    git clone https://github.com/muhammadidrusalawi/gonews.git
    cd gonews
    ```

2.  **Install Go modules**
    ```bash
    go mod tidy
    ```

3.  **Environment setup**
    Create a `.env` file by copying the example:
    ```bash
    cp .env.example .env
    ```
    Configure your environment variables in the `.env` file:

    ```ini
    # Application settings
    APP_PORT=8080

    # Database settings
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=root
    DB_PASS=root
    DB_NAME=gonews_db

    # JWT settings
    JWT_SECRET=supersecretjwtkey
    ```
    **Note**: Ensure your `JWT_SECRET` is strong and kept confidential in production environments.

4.  **Database setup**
    Ensure your PostgreSQL server is running. The application will attempt to connect to the database specified in your `.env` file. GORM will handle table creation/migration on startup.

### Start development server

Run the application:
```bash
go run ./cmd/
```
*(Assumes `main.go` is located directly within the `cmd/` directory. Adjust path if your main package is in a subdirectory like `cmd/api`.)*

The API server will start on `http://localhost:[APP_PORT]` (default `http://localhost:8080`).

## 📁 Project Structure

```
gonews/
├── .env.example       # Example environment variables
├── .gitignore         # Git ignored files
├── LICENSE            # Project license
├── cmd/               # Main application entry points
│   └── ...            # (Expected: main.go for the API server)
├── go.mod             # Go module dependencies
├── go.sum             # Go module checksums
├── internal/          # Internal application logic (handlers, models, services, etc.)
│   └── ...
├── public/            # Static assets (if any, typically for backend serving)
└── README.md          # This README file
```

## ⚙️ Configuration

### Environment Variables
The application's behavior can be configured using environment variables. A `.env.example` file is provided for reference.

| Variable   | Description                              | Default   | Required |
|------------|------------------------------------------|-----------|----------|
| `APP_PORT` | Port on which the API server will run    | `8080`    | Yes      |
| `DB_HOST`  | PostgreSQL database host                 | `localhost` | Yes      |
| `DB_PORT`  | PostgreSQL database port                 | `5432`    | Yes      |
| `DB_USER`  | PostgreSQL database username             | `root`    | Yes      |
| `DB_PASS`  | PostgreSQL database password             | `root`    | Yes      |
| `DB_NAME`  | PostgreSQL database name                 | `gonews_db` | Yes      |
| `JWT_SECRET` | Secret key for signing and verifying JWTs | `supersecretjwtkey` | Yes |

### Configuration Files
-   `.env`: Used to load environment variables for local development.

## 🔧 Development

### Building the Executable
To build a production-ready executable:
```bash
go build -o gonews ./cmd/
```
The executable `gonews` will be created in the project root.

## 🧪 Testing

GoNews utilizes Go's built-in testing framework. You can run tests for your packages as follows:

```bash
# Run all tests in the project
go test ./...

# Run tests in a specific package (e.g., internal/handlers)
go test ./internal/handlers
```

## 🚀 Deployment

For production deployment, it is recommended to build a static executable and run it on your server, or containerize the application using Docker.

### Production Build
```bash
go build -o gonews ./cmd/
```
The `gonews` executable can then be deployed to your target environment.

### Deployment Options
-   **Traditional Hosting**: Copy the compiled `gonews` executable and the `.env` file to your server.
-   **Docker**: Create a `Dockerfile` to containerize your application for easier deployment and scaling. (TODO: Add Dockerfile)

## 📚 API Reference

The GoNews API provides endpoints for user authentication and news article management.

### Authentication

-   **Register User**
    -   `POST /auth/register`
    -   Body: `{ "email": "user@example.com", "password": "password123" }`
    -   Returns: JWT Token
-   **Login User**
    -   `POST /auth/login`
    -   Body: `{ "email": "user@example.com", "password": "password123" }`
    -   Returns: JWT Token

### News Endpoints (Requires Authentication)

-   **Get All News**
    -   `GET /news`
    -   Requires: `Authorization: Bearer <JWT_TOKEN>`
-   **Get News by ID**
    -   `GET /news/:id`
    -   Requires: `Authorization: Bearer <JWT_TOKEN>`
-   **Create New News Article**
    -   `POST /news`
    -   Body: `{ "title": "New Article", "content": "This is the content..." }`
    -   Requires: `Authorization: Bearer <JWT_TOKEN>`
-   **Update News Article**
    -   `PUT /news/:id`
    -   Body: `{ "title": "Updated Title", "content": "Updated content..." }`
    -   Requires: `Authorization: Bearer <JWT_TOKEN>`
-   **Delete News Article**
    -   `DELETE /news/:id`
    -   Requires: `Authorization: Bearer <JWT_TOKEN>`

## 🤝 Contributing

We welcome contributions! Please feel free to open issues or submit pull requests.

### Development Setup for Contributors
1.  Fork the repository.
2.  Clone your forked repository:
    ```bash
    git clone https://github.com/YOUR_USERNAME/gonews.git
    cd gonews
    ```
3.  Set up your environment variables (`.env` file) and database as described in the [Quick Start](#quick-start) section.
4.  Make your changes and ensure tests pass.
5.  Submit a pull request.

## 📄 License

This project is licensed under the [MIT License](LICENSE) - see the LICENSE file for details.

## 🙏 Acknowledgments

-   **Go**: The efficient and powerful programming language.
-   **Fiber**: For providing a fast and unopinionated web framework.
-   **GORM**: An excellent ORM library for Go.
-   **PostgreSQL**: The robust open-source relational database.
-   **`joho/godotenv`**: For seamless environment variable management.
-   **`sirupsen/logrus`**: For flexible and structured logging.

---

<div align="center">

**⭐ Star this repo if you find it helpful!**

Made with ❤️ by Muhammad Idrus Alawi

</div>