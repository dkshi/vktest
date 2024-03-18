package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dkshi/vktest"
	"github.com/dkshi/vktest/internal/handler"
	"github.com/stretchr/testify/assert"
)

func TestHandler_EncodePretty(t *testing.T) {
	h := &handler.Handler{}
	obj := map[string]interface{}{
		"key": "value",
	}
	expectedJSON := `{
  "key": "value"
}`

	w := httptest.NewRecorder()

	err := h.EncodePretty(w, obj)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedJSON, strings.TrimSpace(w.Body.String()))
}

func TestHandler_GetIntFromUrl(t *testing.T) {
	h := &handler.Handler{}
	req := httptest.NewRequest("GET", "/path/123", nil)

	num, err := h.GetIntFromUrl(req)

	assert.NoError(t, err)
	assert.Equal(t, 123, num)
}

func TestHandler_WriteError(t *testing.T) {
	h := &handler.Handler{}
	w := httptest.NewRecorder()
	err := fmt.Errorf("Test Error")

	h.WriteError(w, err, http.StatusBadRequest)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Test Error", strings.TrimSpace(w.Body.String()))
}

func TestHandler_GetStringFromUrl(t *testing.T) {
	h := &handler.Handler{}
	req := httptest.NewRequest("GET", "/path/string", nil)

	str := h.GetStringFromUrl(req)

	assert.Equal(t, "string", str)
}

func TestHandler_DecodeActor(t *testing.T) {
	h := &handler.Handler{}
	actor := vktest.Actor{Name: "John"}
	body, _ := json.Marshal(actor)
	req := httptest.NewRequest("POST", "/path", bytes.NewReader(body))

	decodedActor, err := h.DecodeActor(req)

	assert.NoError(t, err)
	assert.Equal(t, actor, *decodedActor)
}

func TestHandler_DecodeToMap(t *testing.T) {
	h := &handler.Handler{}
	input := map[string]interface{}{
		"key": "value",
	}
	body, _ := json.Marshal(input)
	req := httptest.NewRequest("POST", "/path", bytes.NewReader(body))

	decodedMap, err := h.DecodeToMap(req)

	assert.NoError(t, err)
	assert.Equal(t, input, *decodedMap)
}

func TestHandler_DecodeFilm(t *testing.T) {
	h := &handler.Handler{}
	film := vktest.Film{Title: "Interstellar"}
	body, _ := json.Marshal(film)
	req := httptest.NewRequest("POST", "/path", bytes.NewReader(body))

	decodedFilm, err := h.DecodeFilm(req)

	assert.NoError(t, err)
	assert.Equal(t, film, *decodedFilm)
}

func TestHandler_DecodeAdmin(t *testing.T) {
	h := &handler.Handler{}
	admin := vktest.Admin{Adminname: "admin", Password: "password"}
	body, _ := json.Marshal(admin)
	req := httptest.NewRequest("POST", "/path", bytes.NewReader(body))

	decodedAdmin, err := h.DecodeAdmin(req)

	assert.NoError(t, err)
	assert.Equal(t, admin, *decodedAdmin)
}
