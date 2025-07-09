# go-rest-api-jwt-starter

A starter template for building secure and scalable REST APIs in Go using [Fiber](https://gofiber.io), JSON Web Tokens (JWT) for authentication, and Microsoft SQL Server as the database. This project includes user registration, login, logout, token refresh, and middleware-protected routes with a clean modular structure.

## ✨ Features

- 🛡️ JWT-based Authentication (Access & Refresh Token)
- 🔐 Secure password hashing with bcrypt
- 🧾 Middleware protection for authenticated routes
- 🗃️ SQL Server integration with prepared statements
- 🧱 MVC structure (model, handler, service, middleware)
- 🧪 Easy to test with Postman or Insomnia
- 🧹 Clean and simple code structure for quick development

## ⚙️ Tech Stack

- [Go](https://golang.org/)
- [Fiber v2](https://gofiber.io/)
- [JWT (golang-jwt v5)](https://github.com/golang-jwt/jwt)
- [Microsoft SQL Server](https://www.microsoft.com/en-us/sql-server)
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- `.env` config support

## 📦 Project Structure

```
.
├── cmd/
│   └── main.go           # App entry point
├── internal/
│   ├── handler/          # Route handlers
│   ├── service/          # Business logic (auth, user)
│   ├── model/            # Data models
│   ├── middleware/       # JWT middleware
│   └── response/         # Reusable JSON response formatter
├── pkg/
│   └── database/         # DB connection
├── go.mod
└── .env
```

## 🚀 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/your-username/go-rest-api-jwt-starter.git
cd go-rest-api-jwt-starter
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Setup environment variables

Buat file `.env` di root project:

```env
PORT=8080
DB_HOST=localhost
DB_PORT=1433
DB_USER=sa
DB_PASS=YourPassword123
DB_NAME=your_db
JWT_SECRET=supersecretkey
```

### 4. Run the app

```bash
go run cmd/main.go
```

## 📬 Example API Request (Login)

```http
POST /api/login
Content-Type: application/json

{
  "email": "admin@example.com",
  "password": "secret"
}
```

## 📌 Available Endpoints

| Method | Endpoint            | Description               |
|--------|---------------------|---------------------------|
| POST   | `/api/register`     | Register new user         |
| POST   | `/api/login`        | Login and get JWT tokens  |
| GET    | `/api/me`           | Get current user (auth)   |
| POST   | `/api/refresh`      | Refresh access token      |
| POST   | `/api/logout`       | Logout and clear token    |

## 👤 Author

**Tauhid Jr**  
GitHub: [@tauhidjr](https://github.com/)

---

> Built with ❤️ and Go to help you bootstrap secure REST APIs faster.
