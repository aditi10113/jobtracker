package models

type Job struct {
	ID           int    `json:"id"`
	Company      string `json:"company"`
	Role         string `json:"role"`
	Location     string `json:"location"`
	Status       string `json:"status"`
	Priority     string `json:"priority"`
	Salary       string `json:"salary"`
	AppliedDate  string `json:"appliedDate"`
	InterviewDate string `json:"interviewDate"`
	JobURL       string `json:"jobUrl"`
	Notes        string `json:"notes"`
}

type Stats struct {
	Total      int     `json:"total"`
	Applied    int     `json:"applied"`
	Screening  int     `json:"screening"`
	Interview  int     `json:"interview"`
	Offer      int     `json:"offer"`
	Rejected   int     `json:"rejected"`
	SuccessRate float64 `json:"successRate"`
}
