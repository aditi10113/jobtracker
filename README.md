# Job Tracker — Go Full-Stack Project

A complete job application tracking system built with a Go REST API backend and a responsive HTML/CSS/JavaScript frontend.

## Features
- Dashboard with application statistics
- Add, edit and delete job applications
- Search jobs by company, role or location
- Filter by application status
- Status workflow: Applied → Screening → Interview → Offer/Rejected
- Priority tracking
- Interview date tracking
- Notes
- REST API
- In-memory storage for zero setup
- Responsive UI

## Run

```bash
go run ./cmd/server
```

Open: http://localhost:8080

## API
- GET /api/jobs
- POST /api/jobs
- PUT /api/jobs/{id}
- DELETE /api/jobs/{id}
- GET /api/stats
- GET /api/health

For production, replace the in-memory store with PostgreSQL/MySQL and add authentication.
