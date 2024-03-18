package handler

import (
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			w.WriteHeader(http.StatusUnauthorized)
			logrus.Println("empty auth header")
			return
		}

		adminID, err := h.service.Authorization.ParseToken(header)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			logrus.Println(err)
			return
		}
		w.Write([]byte(strconv.Itoa(adminID)))
		next(w, r)
	}

}
