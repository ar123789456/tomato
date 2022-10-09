package http

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"
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

//----------------------------------
// Page
//----------------------------------

func (h *Handler) MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/main.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, 4)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

//----------------------------------
// Api
//----------------------------------

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

	var result models.EditUser
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.useCase.EditUser(context.TODO(), result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	id := path.Base(r.URL.Path)

	uuId, err := uuid.Parse(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res, err := h.useCase.GetUser(context.TODO(), uuId)
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

func (h *Handler) CreateTomato(w http.ResponseWriter, r *http.Request) {
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

	var result models.CreateTomatoIn
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := h.useCase.CreateTomato(context.TODO(), result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	out, err := json.Marshal(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(out)
}

func (h *Handler) StartTomato(w http.ResponseWriter, r *http.Request) {
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

	var result models.TomatoIn
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.useCase.StartTomato(context.TODO(), result.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetTomato(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body) // response body is []byte

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var result models.TomatoIn
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tomato, err := h.useCase.GetTomato(context.TODO(), result.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	out, err := json.Marshal(tomato)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(out)
}

func (h *Handler) DeleteTomato(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body) // response body is []byte

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var result models.TomatoIn
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.useCase.DeleteTomato(context.TODO(), result.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetTomatoNltx(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body) // response body is []byte

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var result models.TomatoIn
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tomato, err := h.useCase.GetTomatoNltx(context.TODO(), result.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	out, err := json.Marshal(tomato)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(out)
}
