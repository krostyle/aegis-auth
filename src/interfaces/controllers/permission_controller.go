package controller

import (
	"encoding/json"
	"net/http"

	"github.com/krostyle/auth-systme-go/src/application/usecase"
)

type PermissionController struct {
	permissionUseCase usecase.PermissionUseCase
}

func NewPermissionController(permissionUseCase usecase.PermissionUseCase) *PermissionController {
	return &PermissionController{
		permissionUseCase: permissionUseCase,
	}
}

func (p *PermissionController) CreatePermission(w http.ResponseWriter, r *http.Request) {
	var permission usecase.PermissionUseCase
	if err := json.NewDecoder(r.Body).Decode(&permission); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (p *PermissionController) GetPermissionByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	permission, err := p.permissionUseCase.GetPermissionByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(permission)
}

func (p *PermissionController) GetPermissionByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	permission, err := p.permissionUseCase.GetPermissionByName(r.Context(), name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(permission)
}

func (p *PermissionController) UpdatePermission(w http.ResponseWriter, r *http.Request) {
	var permission usecase.PermissionUseCase
	if err := json.NewDecoder(r.Body).Decode(&permission); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (p *PermissionController) DeletePermission(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := p.permissionUseCase.DeletePermission(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
