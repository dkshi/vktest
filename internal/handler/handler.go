package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/dkshi/vktest"
	_ "github.com/dkshi/vktest/docs"
	"github.com/dkshi/vktest/internal/service"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Handler struct {
	port    string
	service *service.Service
}

func NewHandler(port string, service *service.Service) *Handler {
	return &Handler{
		port:    port,
		service: service,
	}
}

func (h *Handler) InitRoutes() error {
	http.Handle("/actors/get", http.HandlerFunc(h.getActors))
	http.Handle("/actors/add", h.userIdentity(http.HandlerFunc(h.addActor)))
	http.Handle("/actors/update/", h.userIdentity(http.HandlerFunc(h.updateActor)))
	http.Handle("/actors/delete/", h.userIdentity(http.HandlerFunc(h.deleteActor)))

	http.Handle("/films/get", http.HandlerFunc(h.getFilms))
	http.Handle("/films/add", h.userIdentity(http.HandlerFunc(h.addFilm)))
	http.Handle("/films/update/", h.userIdentity(http.HandlerFunc(h.updateFilm)))
	http.Handle("/films/delete/", h.userIdentity(http.HandlerFunc(h.deleteFilm)))

	http.Handle("/auth/signUp", http.HandlerFunc(h.signUp))
	http.Handle("/auth/signIn", http.HandlerFunc(h.signIn))

	http.Handle("/swagger/", httpSwagger.Handler())

	return http.ListenAndServe(":"+h.port, nil)
}

func (h *Handler) EncodePretty(w http.ResponseWriter, object any) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(object)
}

func (h *Handler) GetIntFromUrl(r *http.Request) (int, error) {
	parts := strings.Split(r.URL.Path, "/")
	num, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return 0, err
	}
	return num, nil
}

func (h *Handler) WriteError(w http.ResponseWriter, err error, httpStatus int) {
	w.WriteHeader(httpStatus)
	w.Write([]byte(err.Error()))
	logrus.Println(err)
}

func (h *Handler) GetStringFromUrl(r *http.Request) string {
	parts := strings.Split(r.URL.Path, "/")
	return parts[len(parts)-1]
}

func (h *Handler) DecodeActor(r *http.Request) (*vktest.Actor, error) {
	var actor vktest.Actor
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&actor)
	return &actor, err
}

func (h *Handler) DecodeToMap(r *http.Request) (*map[string]any, error) {
	var object map[string]any
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&object)
	return &object, err
}

func (h *Handler) DecodeFilm(r *http.Request) (*vktest.Film, error) {
	var film vktest.Film
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&film)
	return &film, err
}

func (h *Handler) DecodeAdmin(r *http.Request) (*vktest.Admin, error) {
	var admin vktest.Admin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&admin)
	return &admin, err
}
