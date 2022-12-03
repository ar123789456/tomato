package http

import (
	"encoding/json"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"strconv"
	"tomato/models"
	"tomato/servis"
)

type Handler struct {
	useCase servis.UseCase
}

//NewHandler Handler constructor
func NewHandler(useCase servis.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

//----------------------------------
// Api
//----------------------------------

//----------------------------------
//------------ User
//----------------------------------

// UserSignUp struct for user sign up
type UserSignUp struct {
	Name       string  `json:"name"`
	SecondName *string `json:"secondName"`
	Nick       string  `json:"nick"`
	Email      *string `json:"email"`
	Password   string  `json:"password"`
}

// SignUp create new user handler
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	//check request method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	//read request body
	body, err := ioutil.ReadAll(r.Body) // response body is []byte
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//unmarshal request body
	var result UserSignUp
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Set to user model
	var user models.User
	user.Name = result.Name
	user.SecondName = result.SecondName
	user.Nick = result.Nick
	user.Email = result.Email
	user.Password = result.Password

	//Create user
	err = h.useCase.CreateUser(&user, r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// UserSignIn User signIn model
type UserSignIn struct {
	Nick     *string `json:"nick"`
	Email    *string `json:"email"`
	Password string  `json:"password"`
}

// SignIn User signIn handler
func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	//check request method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	//read request body
	body, err := ioutil.ReadAll(r.Body) // response body is []byte
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//unmarshal request body
	var result UserSignIn
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//check input data
	if result.Nick == nil && result.Email == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Set to user model
	var user models.User
	if result.Nick != nil {
		user.Nick = *result.Nick
	}
	if result.Email != nil {
		user.Email = result.Email
	}
	user.Password = result.Password

	//Create user
	token, err := h.useCase.SignIn(&user, r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Set cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: token,
	})
	w.WriteHeader(http.StatusOK)
}

// SignOut User signOut handler
func (h *Handler) SignOut(w http.ResponseWriter, r *http.Request) {
	//check request method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	//Delete session
	err := h.useCase.SignOut(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Delete cookie
	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
	})
}

//----------------------------------
//------------ Task
//----------------------------------

// TaskCreate struct for task create
type TaskCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// CreateTask create new task handler
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	//check request method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	//read request body
	body, err := ioutil.ReadAll(r.Body) // response body is []byte
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//unmarshal request body
	var result TaskCreate
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Set to task model
	var task models.Task
	task.Title = result.Title
	task.Description = result.Description

	//Create task
	err = h.useCase.CreateTask(&task, r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetTasks get all tasks handler
func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	//check request method
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	//get filter on request params
	filter := r.URL.Query().Get("filter")

	//parse filter
	time, err := strconv.ParseInt(filter, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Get session
	cookie, err := r.Cookie("token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	//Get tasks
	tasks, err := h.useCase.GetTasks(cookie.Value, time, r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Marshal tasks
	jsonTasks, err := json.Marshal(tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Write response
	_, err = w.Write(jsonTasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// CompleteTask complete task handler
func (h *Handler) CompleteTask(w http.ResponseWriter, r *http.Request) {
	//check request method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	//get task id on request params
	id := r.URL.Query().Get("id")
	//Get session
	cookie, err := r.Cookie("token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	//Complete task
	err = h.useCase.CompletedTask(cookie.Value, id, r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

//----------------------------------
//------------ Habit
//----------------------------------

// HabitCreate struct for habit create
type HabitCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// CreateHabit create new habit handler
func (h *Handler) CreateHabit(w http.ResponseWriter, r *http.Request) {
	//check request method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	//read request body
	body, err := ioutil.ReadAll(r.Body) // response body is []byte
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//unmarshal request body
	var result HabitCreate
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Set to habit model
	var habit models.Habit
	habit.Title = result.Title
	habit.Description = result.Description

	//Create habit
	err = h.useCase.CreateHabit(&habit, r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetHabits get all habits handler
func (h *Handler) GetHabits(w http.ResponseWriter, r *http.Request) {
	//check request method
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	//get filter on request params
	filter := r.URL.Query().Get("filter")

	//parse filter
	time, err := strconv.ParseInt(filter, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Get habits
	habits, err := h.useCase.GetHabits(time, r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Marshal habits
	jsonHabits, err := json.Marshal(habits)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Write response
	_, err = w.Write(jsonHabits)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// CompleteHabit complete habit handler
func (h *Handler) CompleteHabit(w http.ResponseWriter, r *http.Request) {
	//check request method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	//get habit id on request params
	id := r.URL.Query().Get("id")
	//Complete habit
	err := h.useCase.CompletedHabit(id, r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

//----------------------------------
//------------ Tomatoes
//----------------------------------

// TomatoCreate struct for tomato create
type TomatoCreate struct {
	TimerId uuid.UUID `json:"categoryTomato"` //Timer Id
	TaskId  uuid.UUID `json:"taskId"`         //Task Id
}

// CreateTomato create new tomato handler
func (h *Handler) CreateTomato(w http.ResponseWriter, r *http.Request) {
	//check request method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	//read request body
	body, err := ioutil.ReadAll(r.Body) // response body is []byte
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//unmarshal request body
	var result TomatoCreate
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Set to tomato model
	var tomato models.Tomato
	tomato.TimerId = result.TimerId
	tomato.TaskId = result.TaskId

	//Create tomato
	err = h.useCase.CreateTomato(&tomato, r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetTomatoes get all tomatoes handler
func (h *Handler) GetTomatoes(w http.ResponseWriter, r *http.Request) {
	//check request method
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	//get filter on request params
	filter := r.URL.Query().Get("filter")

	//parse filter
	time, err := strconv.ParseInt(filter, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Get tomatoes
	tomatoes, err := h.useCase.GetTomatoes(time, r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Marshal tomatoes
	jsonTomatoes, err := json.Marshal(tomatoes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Write response
	_, err = w.Write(jsonTomatoes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// StartTomato start tomato handler
func (h *Handler) StartTomato(w http.ResponseWriter, r *http.Request) {
	//check request method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	//get tomato id on request params
	id := r.URL.Query().Get("id")
	//Start tomato
	err := h.useCase.StartTomato(id, r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
