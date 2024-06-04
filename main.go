package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{"John", 30},
	{"Alice", 25},
	{"Bob", 40},
}

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/users":
		h.GetUsers(w, r)
	case "/user":
		h.GetUser(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Here are the users\n")
	json.NewEncoder(w).Encode(users)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	for _, u := range users {
		if u.Name == name {
			json.NewEncoder(w).Encode(u)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`{"error": "User not found"}`))
}

func main() {
	h := &Handler{}
	http.ListenAndServe(":8000", h)
}
