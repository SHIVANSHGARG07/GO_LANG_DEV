# CRUD Application in Go

A simple REST API CRUD (Create, Read, Update, Delete) application built with Go and Gorilla Mux for learning purposes.

## 🚀 Features

- **GET** all users
- **GET** single user by ID
- **POST** create new user
- **PUT** update existing user
- In-memory data storage
- Clean architecture with separated handlers and models

## 📋 Prerequisites

- Go 1.16 or higher
- Basic understanding of REST APIs

## 🛠️ Installation

1. Clone the repository:
```bash
git clone <your-repo-url>
cd crud-app
```

2. Install dependencies:
```bash
go mod download
```

## 🏃 Running the Application

Start the server:
```bash
go run main.go
```

The server will start on `http://localhost:7777`

You should see:
```
Server running on port 7777....
```

## 📚 API Endpoints

### 1. Get All Users
```bash
GET /users
```

**Example:**
```bash
curl http://localhost:7777/users
```

**Response:**
```json
[
  {
    "id": "1",
    "name": "Shivansh Garg",
    "email": "shivansh@example.com",
    "password": ""
  }
]
```

---

### 2. Get Single User
```bash
GET /users/{id}
```

**Example:**
```bash
curl http://localhost:7777/users/1
```

**Response:**
```json
{
  "id": "1",
  "name": "Shivansh Garg",
  "email": "shivansh@example.com",
  "password": ""
}
```

---

### 3. Create User
```bash
POST /users
Content-Type: application/json
```

**Example:**
```bash
curl -X POST http://localhost:7777/users \
  -H "Content-Type: application/json" \
  -d '{
    "id": "1",
    "name": "Shivansh Garg",
    "email": "shivansh@example.com"
  }'
```

**Response:**
```json
{
  "id": "1",
  "name": "Shivansh Garg",
  "email": "shivansh@example.com",
  "password": ""
}
```

---

### 4. Update User
```bash
PUT /users/{id}
Content-Type: application/json
```

**Example:**
```bash
curl -X PUT http://localhost:7777/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Shivansh Garg Updated",
    "email": "shivansh.updated@example.com"
  }'
```

**Response:**
```json
{
  "id": "1",
  "name": "Shivansh Garg Updated",
  "email": "shivansh.updated@example.com",
  "password": ""
}
```

## 🧪 Testing with cURL

Run these commands in order to test all endpoints:

```bash
# 1. Get all users (should return empty or null initially)
curl http://localhost:7777/users

# 2. Create first user
curl -X POST http://localhost:7777/users \
  -H "Content-Type: application/json" \
  -d '{"id":"1","name":"Shivansh Garg","email":"shivansh@example.com"}'

# 3. Create second user
curl -X POST http://localhost:7777/users \
  -H "Content-Type: application/json" \
  -d '{"id":"2","name":"John Doe","email":"john@example.com"}'

# 4. Get all users (should return both users)
curl http://localhost:7777/users

# 5. Get single user by ID
curl http://localhost:7777/users/1

# 6. Update user
curl -X PUT http://localhost:7777/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Shivansh Updated","email":"updated@example.com"}'

# 7. Verify update
curl http://localhost:7777/users
```

## 📁 Project Structure

```
crud-app/
├── main.go              # Application entry point and route definitions
├── handlers/
│   └── user.go         # HTTP handlers for user operations
├── models/
│   └── user.go         # User data model
├── go.mod              # Go module file
├── go.sum              # Dependency checksums
└── README.md           # This file
```

## 🔧 Technologies Used

- **Go** - Programming language
- **Gorilla Mux** - HTTP router and URL matcher

## ⚠️ Important Notes

- This application uses **in-memory storage**. All data will be lost when the server restarts.
- No database is connected - this is a learning project.
- For production use, consider adding:
  - Database integration (PostgreSQL, MongoDB, etc.)
  - Authentication and authorization
  - Input validation
  - Error handling improvements
  - Logging
  - DELETE endpoint

## 📝 License

This project is for educational purposes.

## 👤 Author

**Shivansh Garg**

## 🤝 Contributing

Feel free to fork this project and submit pull requests for improvements!