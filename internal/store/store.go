package store

import (
	"errors"
	"sync"

	"gojobtracker/internal/models"
)

type Store struct {
	mu   sync.RWMutex
	jobs []models.Job
	nextID int
}

func New() *Store {
	return &Store{
		nextID: 4,
		jobs: []models.Job{
			{ID:1, Company:"Acme Technologies", Role:"Backend Developer", Location:"Bengaluru", Status:"Interview", Priority:"High", Salary:"₹8-12 LPA", AppliedDate:"2026-09-02", InterviewDate:"2026-09-24", JobURL:"https://example.com", Notes:"Prepare Go REST API and SQL questions."},
			{ID:2, Company:"Nova Labs", Role:"Software Engineer", Location:"Remote", Status:"Screening", Priority:"High", Salary:"₹7-10 LPA", AppliedDate:"2026-09-06", JobURL:"https://example.com", Notes:"Recruiter contacted me."},
			{ID:3, Company:"Pixel Systems", Role:"Full-Stack Developer", Location:"Hyderabad", Status:"Applied", Priority:"Medium", Salary:"₹6-9 LPA", AppliedDate:"2026-09-10", JobURL:"https://example.com", Notes:"Follow up next week."},
		},
	}
}

func (s *Store) List() []models.Job {
	s.mu.RLock(); defer s.mu.RUnlock()
	out := make([]models.Job, len(s.jobs)); copy(out, s.jobs); return out
}

func (s *Store) Add(j models.Job) models.Job {
	s.mu.Lock(); defer s.mu.Unlock()
	j.ID = s.nextID; s.nextID++
	s.jobs = append(s.jobs, j)
	return j
}

func (s *Store) Update(id int, j models.Job) (models.Job, error) {
	s.mu.Lock(); defer s.mu.Unlock()
	for i := range s.jobs {
		if s.jobs[i].ID == id {
			j.ID = id
			s.jobs[i] = j
			return j, nil
		}
	}
	return models.Job{}, errors.New("job not found")
}

func (s *Store) Delete(id int) error {
	s.mu.Lock(); defer s.mu.Unlock()
	for i := range s.jobs {
		if s.jobs[i].ID == id {
			s.jobs = append(s.jobs[:i], s.jobs[i+1:]...)
			return nil
		}
	}
	return errors.New("job not found")
}
