package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"students-api/internal/services/student"

	"github.com/gorilla/mux"
)

func NewHandler(service *student.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

type Handler struct {
	Router  *mux.Router
	Service *student.Service
}

func (h *Handler) InitRoutes() {
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/students", h.GetAllStudents).Methods("GET")
	h.Router.HandleFunc("/api/students", h.PostStudent).Methods("POST")
	h.Router.HandleFunc("/api/students/{school}", h.GetStudentsBySchool).Methods("GET")
	h.Router.HandleFunc("/api/students/{id}", h.GetStudentByID).Methods("GET")
	h.Router.HandleFunc("/api/students/{id}", h.UpdateStudent).Methods("PUT")
	h.Router.HandleFunc("/api/students/{id}", h.DeleteStudent).Methods("DELETE")
	h.Router.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "Status Up!"}); err != nil {
			panic(err)
		}
	})
}

func (h *Handler) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	students, err := h.Service.GetAllStudents()
	if err != nil {
		respondWithError(w, "Failed to retrieve all students", err)
	}
	if err := json.NewEncoder(w).Encode(students); err != nil {
		panic(err)
	}
}
func (h *Handler) PostStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var student student.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		respondWithError(w, "Could not decode json body", err)
	}

	student, err := h.Service.PostStudent(student)
	if err != nil {
		respondWithError(w, "Failed to post new student", err)
	}

	if err = json.NewEncoder(w).Encode(student); err != nil {
		panic(err)
	}
}

func (h *Handler) GetStudentsBySchool(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	school := vars["school"]
	students, err := h.Service.GetStudentsBySchool(school)
	if err != nil {
		respondWithError(w, "Error Retrieving Students by School", err)
	}
	if err := json.NewEncoder(w).Encode(students); err != nil {
		panic(err)
	}
}

func (h *Handler) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id := vars["id"]
	studentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		respondWithError(w, "Error Parsing ID to UINT", err)
	}
	student, err := h.Service.GetStudentByID(uint(studentID))
	if err != nil {
		respondWithError(w, "Error Retrieving Student by ID", err)
	}
	if err := json.NewEncoder(w).Encode(student); err != nil {
		panic(err)
	}
}
func (h *Handler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id := vars["id"]
	var student student.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		respondWithError(w, "Could not decode json body", err)
	}
	studentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		respondWithError(w, "Error Parsing ID to UINT", err)
	}
	student, err = h.Service.UpdateStudent(uint(studentID), student)
	if err != nil {
		respondWithError(w, "Failed to update student", err)
	}
	if err = json.NewEncoder(w).Encode(student); err != nil {
		panic(err)
	}
}
func (h *Handler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id := vars["id"]

	studentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		respondWithError(w, "Error Parsing ID to UINT", err)
	}

	err = h.Service.DeleteStudent(uint(studentID))
	if err != nil {
		respondWithError(w, "Failed to delete student by student ID", err)
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Student successfully deleted"}); err != nil {
		panic(err)
	}
}
