# 🚀 Job Tracker

> A modern full-stack job application tracking system built to help job seekers organize applications, track hiring progress, manage interviews, and stay on top of their job search.

---

## 📌 About

Job Tracker is a lightweight and responsive web application for managing job applications in one place.

Instead of maintaining applications across spreadsheets, notes, and emails, users can use Job Tracker to record important information about each opportunity and track its progress through the hiring process.

The application provides a simple dashboard for managing:

- Company information
- Job roles
- Locations
- Application status
- Priority
- Interview dates
- Personal notes
- Application statistics

---

## 🛠️ Tech Stack

### 🎨 Frontend

| Technology | Purpose |
|---|---|
| **HTML5** | Provides the structure and semantic markup |
| **CSS3** | Handles styling, layouts, animations, and responsive design |
| **JavaScript** | Handles frontend logic and user interactions |
| **Vanilla JavaScript** | Keeps the frontend lightweight without a framework |
| **Fetch API** | Communicates with the backend REST API |
| **JSON** | Data exchange between frontend and backend |

### ⚙️ Backend

| Technology | Purpose |
|---|---|
| **Go (Golang)** | Backend application and server |
| **net/http** | HTTP server and API handling |
| **REST API** | CRUD operations for job applications |
| **JSON** | API request and response format |
| **sync.RWMutex** | Safe concurrent access to in-memory data |

### 💾 Data Storage

| Technology | Purpose |
|---|---|
| **In-Memory Storage** | Stores job applications during runtime |
| **Go Structs** | Represents job application data |

> ⚠️ The current version uses in-memory storage, so data may be lost when the server restarts. A persistent database such as PostgreSQL can be added in a future version.

### ☁️ Deployment & DevOps

| Technology | Purpose |
|---|---|
| **Vercel** | Application deployment and hosting |
| **Docker** | Containerization |
| **Dockerfile.vercel** | Vercel container deployment configuration |
| **Git** | Version control |
| **GitHub** | Source code hosting and collaboration |

---

## ✨ Features

### 📊 Dashboard

View an overview of your job search activity.

The dashboard provides statistics such as:

- Total applications
- Applied applications
- Screening applications
- Interviews
- Offers
- Rejected applications

---

### ➕ Add Applications

Create a new job application with details such as:

- Company
- Job title
- Location
- Status
- Priority
- Interview date
- Notes

---

### ✏️ Edit Applications

Update application information whenever there is a change in the hiring process.

Example:

```text
Applied
   ↓
Screening
   ↓
Interview
   ↓
Offer
