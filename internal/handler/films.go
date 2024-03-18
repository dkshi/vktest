package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

// @Summary Get all films
// @Description Get a list of all films
// @Produce json
// @Param order query []string false "Order of sorting"
// @Param ascending query bool false "Sorting order"
// @Param title query string false "Title to search"
// @Param actor query string false "Actor to search"
// @Tags films
// @Success 200 {array} vktest.Film
// @Router /films/get [get]
func (h *Handler) getFilms(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		r.ParseForm()

		attrs := make(map[string]any)

		for k, v := range r.Form {
			switch k {
			case "order":
				attrs[k] = v
			case "ascending":
				asc, err  := strconv.ParseBool(v[0])
				if err != nil {
					h.WriteError(w, fmt.Errorf("incorrect format of ascending attribute, needs to be: bool"), http.StatusBadRequest)
					return
				}
				attrs[k] = asc
			default:
				attrs[k] = v[0]
			}
		}

		films, err := h.service.GetFilms(&attrs)
		if err != nil {
			h.WriteError(w, err, http.StatusBadRequest)
			return
		}
		err = h.EncodePretty(w, films)
		if err != nil {
			h.WriteError(w, err, http.StatusBadRequest)
			return
		}
	default:
		h.WriteError(w, fmt.Errorf("incorrect method: %s", r.Method), http.StatusMethodNotAllowed)
		return
	}
}

// @Summary Add a new film
// @Security ApiKeyAuth
// @Description Add a new film to the database
// @Accept json
// @Tags films
// @Produce json
// @Param input body vktest.Film true "Film object to be added (ignore film_id)"
// @Success 200 {string} string "id of the added film"
// @Router /films/add [post]
func (h *Handler) addFilm(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		film, err := h.DecodeFilm(r)
		if err != nil {
			h.WriteError(w, err, http.StatusBadRequest)
			return
		}

		_, err = h.service.InsertFilm(film)
		if err != nil {
			h.WriteError(w, err, http.StatusBadRequest)
			return
		}
	default:
		h.WriteError(w, fmt.Errorf("incorrect method: %s", r.Method), http.StatusMethodNotAllowed)
		return
	}
}

// @Summary Update an existing film
// @Security ApiKeyAuth
// @Description Update an existing film in the database
// @Accept json
// @Tags films
// @Produce json
// @Param id path int true "ID of the film to be updated"
// @Param input body vktest.Film true "Film object with updated information (ignore film_id)"
// @Router /films/update/{id} [patch]
func (h *Handler) updateFilm(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PATCH":
		filmID, err := h.GetIntFromUrl(r)
		if err != nil {
			h.WriteError(w, err, http.StatusBadRequest)
			return
		}
		film, err := h.DecodeToMap(r)
		if err != nil {
			h.WriteError(w, err, http.StatusBadRequest)
			return
		}
		(*film)["film_id"] = filmID

		err = h.service.UpdateFilm(film)
		if err != nil {
			h.WriteError(w, err, http.StatusBadRequest)
			return
		}
	default:
		h.WriteError(w, fmt.Errorf("incorrect method: %s", r.Method), http.StatusMethodNotAllowed)
		return
	}
}

// @Summary Delete a film
// @Security ApiKeyAuth
// @Description Delete a film from the database
// @Produce json
// @Tags films
// @Param id path int true "ID of the film to be deleted"
// @Router /films/delete/{id} [delete]
func (h *Handler) deleteFilm(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		filmID, err := h.GetIntFromUrl(r)
		if err != nil {
			h.WriteError(w, err, http.StatusBadRequest)
			return
		}

		err = h.service.DeleteFilm(filmID)
		if err != nil {
			h.WriteError(w, err, http.StatusBadRequest)
			return
		}
	default:
		h.WriteError(w, fmt.Errorf("incorrect method: %s", r.Method), http.StatusMethodNotAllowed)
		return
	}
}
