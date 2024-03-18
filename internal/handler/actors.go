package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

// @Summary Get all actors
// @Description Get a list of all actors
// @Produce json
// @Tags actors
// @Success 200 {array} vktest.Actor
// @Router /actors/get [get]
func (h *Handler) getActors(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		actors, err := h.service.GetActors()
		if err != nil {
			h.WriteError(w, err, http.StatusBadRequest)
			return
		}
		err = h.EncodePretty(w, actors)
		if err != nil {
			h.WriteError(w, err, http.StatusBadRequest)
			return
		}
	default:
		h.WriteError(w, fmt.Errorf("incorrect method: %s", r.Method), http.StatusMethodNotAllowed)
		return
	}
}

// @Summary Add a new actor
// @Security ApiKeyAuth
// @Description Add a new actor to the database
// @Accept json
// @Tags actors
// @Produce json
// @Param input body vktest.Actor true "Actor object to be added (ignore actor_id and films)"
// @Success 200 {string} string "id of the added actor"
// @Router /actors/add [post]
func (h *Handler) addActor(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		actor, err := h.DecodeActor(r)
		if err != nil {
			h.WriteError(w, err, http.StatusBadRequest)
			return
		}

		id, err := h.service.InsertActor(actor)
		if err != nil {
			h.WriteError(w, err, http.StatusBadRequest)
			return
		}
		w.Write([]byte(strconv.Itoa(id)))
	default:
		h.WriteError(w, fmt.Errorf("incorrect method: %s", r.Method), http.StatusMethodNotAllowed)
		return
	}
}

// @Summary Update an existing actor
// @Security ApiKeyAuth
// @Description Update an existing actor in the database
// @Accept json
// @Tags actors
// @Produce json
// @Param id path int true "ID of the actor to be updated"
// @Param input body vktest.Actor true "Actor object with updated information (ignore actor_id and films)"
// @Router /actors/update/{id} [patch]
func (h *Handler) updateActor(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PATCH":
		actorID, err := h.GetIntFromUrl(r)
		if err != nil {
			h.WriteError(w, err, http.StatusBadRequest)
			return
		}
		actor, err := h.DecodeToMap(r)
		if err != nil {
			h.WriteError(w, err, http.StatusBadRequest)
			return
		}
		(*actor)["actor_id"] = actorID

		err = h.service.UpdateActor(actor)
		if err != nil {
			h.WriteError(w, err, http.StatusBadRequest)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		logrus.Printf("incorrect method: %s", r.Method)
		return
	}
}

// @Summary Delete an actor
// @Security ApiKeyAuth
// @Description Delete an actor from the database
// @Produce json
// @Tags actors
// @Param id path int true "ID of the actor to be deleted"
// @Router /actors/delete/{id} [delete]
func (h *Handler) deleteActor(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		actorID, err := h.GetIntFromUrl(r)
		if err != nil {
			h.WriteError(w, err, http.StatusBadRequest)
			return
		}

		err = h.service.DeleteActor(actorID)
		if err != nil {
			h.WriteError(w, err, http.StatusBadRequest)
			return
		}
	default:
		h.WriteError(w, fmt.Errorf("incorrect method: %s", r.Method), http.StatusMethodNotAllowed)
		return
	}
}
