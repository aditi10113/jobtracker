package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"gojobtracker/internal/models"
	"gojobtracker/internal/store"
)

type Handler struct{ store *store.Store }

func New(s *store.Store) *Handler { return &Handler{store:s} }

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, map[string]string{"status":"ok","service":"job-tracker-api"})
}

func (h *Handler) Jobs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, 200, h.store.List())
	case http.MethodPost:
		var j models.Job
		if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
			writeJSON(w, 400, map[string]string{"error":"invalid JSON"}); return
		}
		if err := validate(j); err != nil {
			writeJSON(w, 400, map[string]string{"error":err.Error()}); return
		}
		writeJSON(w, 201, h.store.Add(j))
	default:
		writeJSON(w, 405, map[string]string{"error":"method not allowed"})
	}
}

func (h *Handler) JobByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/jobs/"))
	if err != nil { writeJSON(w, 400, map[string]string{"error":"invalid job id"}); return }
	switch r.Method {
	case http.MethodPut:
		var j models.Job
		if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
			writeJSON(w, 400, map[string]string{"error":"invalid JSON"}); return
		}
		if err := validate(j); err != nil {
			writeJSON(w, 400, map[string]string{"error":err.Error()}); return
		}
		updated, err := h.store.Update(id, j)
		if err != nil { writeJSON(w, 404, map[string]string{"error":err.Error()}); return }
		writeJSON(w, 200, updated)
	case http.MethodDelete:
		if err := h.store.Delete(id); err != nil { writeJSON(w,404,map[string]string{"error":err.Error()}); return }
		writeJSON(w, 200, map[string]string{"message":"job deleted"})
	default:
		writeJSON(w, 405, map[string]string{"error":"method not allowed"})
	}
}

func (h *Handler) Stats(w http.ResponseWriter, r *http.Request) {
	var s models.Stats
	for _, j := range h.store.List() {
		s.Total++
		switch strings.ToLower(j.Status) {
		case "applied": s.Applied++
		case "screening": s.Screening++
		case "interview": s.Interview++
		case "offer": s.Offer++
		case "rejected": s.Rejected++
		}
	}
	if s.Total > 0 { s.SuccessRate = float64(s.Interview+s.Offer)/float64(s.Total)*100 }
	writeJSON(w, 200, s)
}

func validate(j models.Job) error {
	if strings.TrimSpace(j.Company)=="" { return errText("company is required") }
	if strings.TrimSpace(j.Role)=="" { return errText("role is required") }
	if strings.TrimSpace(j.Status)=="" { return errText("status is required") }
	allowedStatus := map[string]bool{"Applied":true,"Screening":true,"Interview":true,"Offer":true,"Rejected":true}
	if !allowedStatus[j.Status] { return errText("invalid status") }
	return nil
}

type errText string
func (e errText) Error() string { return string(e) }
