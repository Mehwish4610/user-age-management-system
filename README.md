# User Age Management System

A full-stack User Age Management System built using **Go (Fiber)**, **MySQL**, **React**, and **Docker**.

The application allows users to create, update, delete, search, and manage user records while automatically calculating age from the stored date of birth. It includes a modern responsive dashboard with light/dark mode support and real-time statistics.

## Live Demo

**Frontend:** https://user-age-management-system.vercel.app

**Backend API:** https://user-age-management-system.onrender.com

---

## Features

### User Management

* Add new users
* Update existing users
* Delete users
* View all users
* Search users by name

### Age Calculation

* Automatic age calculation from Date of Birth
* Average age calculation
* Youngest user identification
* Oldest user identification

### Dashboard Analytics

* Total users
* Average age
* Youngest user
* Oldest user

### User Experience

* Fully responsive design
* Mobile-friendly layout
* Light Mode / Dark Mode
* Interactive cards and hover effects
* Modern dashboard UI

### Backend Features

* RESTful API architecture
* Layered architecture (Handler → Service → Repository)
* Request logging middleware
* CORS support
* MySQL database integration
* Environment variable configuration

### DevOps

* Dockerized backend
* GitHub integration
* Render deployment
* Vercel deployment

---

## Tech Stack

### Frontend

* React
* Vite
* CSS3
* JavaScript

### Backend

* Go
* Fiber Framework

### Database

* MySQL

### DevOps & Deployment

* Docker
* Render
* Vercel
* GitHub

---

## Project Structure

```text
go-user-age-api/
│
├── cmd/
├── config/
├── db/
├── internal/
│   ├── handler/
│   ├── middleware/
│   ├── repository/
│   ├── routes/
│   └── service/
│
├── frontend/
│   ├── src/
│   └── public/
│
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

---

## API Endpoints

### Get All Users

```http
GET /users
```

### Get User By ID

```http
GET /users/:id
```

### Create User

```http
POST /users
```

Request Body:

```json
{
  "name": "John Doe",
  "dob": "2002-05-20"
}
```

### Update User

```http
PUT /users/:id
```

### Delete User

```http
DELETE /users/:id
```

### Health Check

```http
GET /health
```

---

## Installation

### Clone Repository

```bash
git clone https://github.com/Mehwish4610/user-age-management-system.git
```

```bash
cd user-age-management-system
```

### Backend Setup

```bash
go mod download
```

Create a `.env` file:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=user_age_db

SERVER_PORT=8081
```

Run backend:

```bash
go run cmd/main.go
```

### Frontend Setup

```bash
cd frontend
```

```bash
npm install
```

```bash
npm run dev
```

---

## Docker Setup

Build Docker Image:

```bash
docker build -t user-age-management-system .
```

Run Container:

```bash
docker run -p 8081:8081 user-age-management-system
```


---

## Future Enhancements

* JWT Authentication
* User Roles & Permissions
* Profile Images
* Export User Data to CSV
* Pagination Improvements
* Advanced Analytics Dashboard

---

## Author

**Mehwish**

Computer Science Engineer | Data Science & AI Enthusiast

GitHub: https://github.com/Mehwish4610

LinkedIn: www.linkedin.com/in/mehwish-18476a266

---

## License

This project is developed for learning, portfolio, and demonstration purposes.
