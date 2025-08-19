package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/AlinaZbk/mini-blog.git/model"
	"github.com/AlinaZbk/mini-blog.git/service"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(service.ListPosts())
	case http.MethodPost:
		var req model.CreatePostRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "invalid JSON"})
			return
		}
		post, err := service.CreatePost(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(post)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func PostByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/posts/")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "invalid id"})
		return
	}

	switch r.Method {
	case http.MethodGet:
		post, err := service.GetPost(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "not found"})
			return
		}
		json.NewEncoder(w).Encode(post)
	case http.MethodDelete:
		if err := service.DeletePost(id); err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "not found"})
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
