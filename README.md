# User Age Management System

A full-stack User Management application built with **Go Fiber**, **MySQL**, **SQLC**, **Docker**, and **React**. The application allows users to perform CRUD operations, dynamically calculate age from date of birth, and manage records through a modern responsive dashboard.

## Features

### Backend

* RESTful API using Go Fiber
* MySQL database integration
* SQLC for type-safe database queries
* Repository-Service-Handler architecture
* Input validation using Validator
* Structured logging using Uber Zap
* Pagination support
* Dockerized backend
* Environment variable configuration
* Health check endpoint

### Frontend

* React + Vite
* Responsive dashboard UI
* Dark Mode / Light Mode
* Search users by name
* User statistics dashboard
* Create User
* Update User
* Delete User
* Pagination support
* Mobile-friendly design

---

## Tech Stack

### Backend

* Go
* Fiber
* MySQL
* SQLC
* Validator
* Uber Zap
* Docker

### Frontend

* React
* Vite
* Axios
* CSS3

---

## Project Structure

```text
go-user-age-api/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── config/
│
├── db/
│   ├── query/
│   ├── schema/
│   └── sqlc/
│
├── internal/
│   ├── handler/
│   ├── logger/
│   ├── middleware/
│   ├── models/
│   ├── repository/
│   ├── routes/
│   └── service/
│
├── frontend/
│   ├── src/
│   ├── public/
│   └── package.json
│
├── Dockerfile
├── go.mod
├── go.sum
├── sqlc.yaml
└── README.md
```

---

## API Endpoints

### Health Check

```http
GET /health
```

### Create User

```http
POST /users
```

Request Body:

```json
{
  "name": "John Doe",
  "dob": "2000-05-15"
}
```

### Get User By ID

```http
GET /users/:id
```

### Get All Users

```http
GET /users?page=1&limit=5
```

### Update User

```http
PUT /users/:id
```

Request Body:

```json
{
  "name": "Updated Name",
  "dob": "1999-12-10"
}
```

### Delete User

```http
DELETE /users/:id
```

---

## Dynamic Age Calculation

The application calculates age dynamically based on the user's Date of Birth. Age is computed during retrieval and does not need to be stored separately in the database.

---

## Database Schema

```sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    dob DATE NOT NULL
);
```

---

## Local Setup

### Clone Repository

```bash
git clone https://github.com/Mehwish4610/user-age-management-system.git
cd user-age-management-system
```

### Backend Setup

Install dependencies:

```bash
go mod tidy
```

Create `.env`

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=user_age_db
SERVER_PORT=8081
```

Run backend:

```bash
go run cmd/server/main.go
```

---

### Frontend Setup

Navigate to frontend:

```bash
cd frontend
```

Install dependencies:

```bash
npm install
```

Run frontend:

```bash
npm run dev
```

Frontend:

```text
http://localhost:5173
```

Backend:

```text
http://localhost:8081
```

---

## Docker

Build image:

```bash
docker build -t go-user-age-api .
```

Run container:

```bash
docker run --env-file .env -p 8081:8081 go-user-age-api
```

---

## Future Enhancements

* JWT Authentication
* User Profile Images
* Advanced Filtering
* Sorting by Age and Name
* Export to CSV
* Deployment on Vercel and Render
* CI/CD using GitHub Actions

---

## Author

**Mehwish**

Computer Science Engineer | Data Science & AI Enthusiast

GitHub:
https://github.com/Mehwish4610
