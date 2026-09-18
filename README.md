## 🛠️ Tech Stack

### Frontend

- **HTML5** — Structure and semantic markup
- **CSS3** — Styling, responsive layout, and UI design
- **JavaScript (Vanilla JS)** — Client-side logic and API communication
- **Fetch API** — Communication between the frontend and Go REST API

### Backend

- **Go (Golang)** — Backend application and HTTP server
- **Go `net/http`** — HTTP server and REST API routing
- **REST API** — CRUD operations for job applications
- **JSON** — Data exchange between frontend and backend

### Data Storage

- **In-Memory Store** — Lightweight storage for the current version
- **Mutex (`sync.RWMutex`)** — Safe concurrent access to application data

> ⚠️ The current version uses in-memory storage. Data may be lost when the server restarts. A persistent database such as PostgreSQL can be added for production use.

### Deployment & DevOps

- **Vercel** — Application deployment and hosting
- **Docker** — Containerized deployment
- **Dockerfile.vercel** — Vercel container configuration
- **Git & GitHub** — Source control and repository management

### Development Tools

- **Go Modules** — Dependency and package management
- **Git** — Version control
- **GitHub** — Source code hosting
- **Vercel** — Continuous deployment

---

## 🏗️ Architecture

```text
┌─────────────────────────────┐
│          Browser            │
│                             │
│  HTML + CSS + JavaScript    │
└──────────────┬──────────────┘
               │
               │ HTTP / JSON
               ▼
┌─────────────────────────────┐
│        Go Backend           │
│                             │
│       net/http              │
│       REST API              │
└──────────────┬──────────────┘
               │
               ▼
┌─────────────────────────────┐
│      In-Memory Store        │
│                             │
│      Go structs + Mutex     │
└─────────────────────────────┘
