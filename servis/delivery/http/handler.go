package http

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"tomato/models"
	"tomato/servis"
)

type Handler struct {
	useCase servis.UseCase
}

func NewHandler(useCase servis.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body) // response body is []byte

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var result models.CreateUser
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.useCase.CreateUser(context.TODO(), result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	out, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(out)
}

func (h *Handler) EditUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("editUser"))
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) CreateTomato(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) CompleteTomato() {

}

func (h *Handler) GetTomato() {

}

func (h *Handler) GetTomatoNltx() {

}
