# 📦 Go REST API - Task Manager

A simple RESTful API built with Go, designed to manage tasks efficiently. This API allows you to create, read, update, and delete (CRUD) tasks with persistent storage using a JSON file.

## 🚀 Features

- ✅ RESTful architecture using `net/http`
- ✅ CRUD operations for tasks
- ✅ JSON file-based persistence
- ✅ Modular and easy to extend
- ✅ Security best practices in place

---

## 🔐 Security Implementations

This API is built with several security measures in place to protect data and ensure safe usage:

### 1. Input Validation and Sanitization
All incoming request data is validated and sanitized to prevent:
- **SQL Injection** (N/A in this project, but good habits apply)
- **Cross-Site Scripting (XSS)**
- **Command Injection** via untrusted input

### 2. Secure HTTP Headers
Several HTTP headers are added to protect against common web vulnerabilities:
- `X-Content-Type-Options: nosniff`
- `X-Frame-Options: DENY`
- `Content-Security-Policy` (to be added in future updates)
- `Strict-Transport-Security` (when behind HTTPS)

### 3. CORS Configuration
CORS policies are configured to restrict access to trusted domains:
```go
w.Header().Set("Access-Control-Allow-Origin", "https://yourdomain.com")
```
### 4. HTTPS Enforcement

While HTTPS is not handled directly in Go, the API is expected to run behind a secure proxy (e.g., NGINX or Caddy) that enforces HTTPS on all endpoints.

### 5. Rate Limiting (Planned)

Rate limiting will be added to prevent abuse and denial-of-service (DoS) attacks.

---

## 🚀 Features

- List all tasks
- Retrieve a task by ID
- Create a new task
- Delete a task by ID

---

## 🔧 Technologies

- [Go](https://golang.org/)
- `net/http` (standard library)

---

## ▶️ How to Run

1. **Clone the repository:**

```bash
git clone https://github.com/yourusername/task-api-go.git
cd task-api-go

**Run the server:**
go run main.go
