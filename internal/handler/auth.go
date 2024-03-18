package handler

import (
	"fmt"
	"net/http"
)

// @Summary "Registration for admins"
// @Description "Registers new admin"
// @Tags auth
// @Accept  json
// @Produce  json
// @Param input body vktest.Admin true "Adminname and password"
// @Success 200 {string} string "Registration success"
// @Router /auth/signUp [post]
func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		admin, err := h.DecodeAdmin(r)
		if err != nil {
			h.WriteError(w, err, http.StatusBadRequest)
			return
		}

		_, err = h.service.CreateAdmin(admin)
		if err != nil {
			h.WriteError(w, err, http.StatusBadRequest)
			return
		}
	default:
		h.WriteError(w, fmt.Errorf("incorrect method: %s", r.Method), http.StatusMethodNotAllowed)
		return
	}
}

// @Summary "Authentificate admin"
// @Description "Authentificates admin and returns JWT-token"
// @Tags auth
// @Accept  json
// @Produce  json
// @Param input body vktest.Admin true "Adminname and password"
// @Success 200 {string} string "JWT-token"
// @Router /auth/signIn [post]
func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		admin, err := h.DecodeAdmin(r)
		if err != nil {
			h.WriteError(w, err, http.StatusBadRequest)
			return
		}

		token, err := h.service.Authorization.GenerateToken(admin.Adminname, admin.Password)
		if err != nil {
			h.WriteError(w, err, http.StatusBadRequest)
			return
		}
		w.Write([]byte(token))
	default:
		h.WriteError(w, fmt.Errorf("incorrect method: %s", r.Method), http.StatusMethodNotAllowed)
		return
	}
}
