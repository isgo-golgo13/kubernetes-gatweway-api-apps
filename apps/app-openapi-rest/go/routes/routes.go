package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"go-chi-rest-app/repository"
	"github.com/go-chi/chi/v5"
)

type RouteHandler struct {
	OwnerRepo *repository.OwnerRepository
}

func NewRouteHandler(ownerRepo *repository.OwnerRepository) *RouteHandler {
	return &RouteHandler{
		OwnerRepo: ownerRepo,
	}
}

func (h *RouteHandler) RegisterRoutes(r chi.Router) {
	r.Get("/owners/{ownerID}", h.GetOwnerByID)
	r.Get("/owners", h.GetAllOwners)
}

// GetOwnerByID is the HTTP handler for retrieving a specific owner
func (h *RouteHandler) GetOwnerByID(w http.ResponseWriter, r *http.Request) {
	ownerID := chi.URLParam(r, "ownerID")

	owner, err := h.OwnerRepo.GetOwnerByID(context.Background(), ownerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(owner)
}

// GetAllOwners is the HTTP handler for retrieving all owners
func (h *RouteHandler) GetAllOwners(w http.ResponseWriter, r *http.Request) {
	owners, err := h.OwnerRepo.GetAllOwners(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(owners)
}
