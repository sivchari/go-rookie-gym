package handler

import (
	"encoding/json"
	"log"
	"net/http"

	groupuc "github.com/sivchari/go-rookie-gym/usecase/group"
	useruc "github.com/sivchari/go-rookie-gym/usecase/user"
)

type Handler interface {
	GroupHandler() http.HandlerFunc
	GroupsHandler() http.HandlerFunc
	UserHandler() http.HandlerFunc
}

type handle struct {
	groupUsecase groupuc.Usecase
	userUsecase  useruc.Usecase
}

var _ Handler = (*handle)(nil)

func NewHandler(guc groupuc.Usecase, uuc useruc.Usecase) Handler {
	return &handle{
		groupUsecase: guc,
		userUsecase:  uuc,
	}
}

type GroupRequest struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}

type GroupResponse struct {
	ID int64 `json:"id"`
}

func (h *handle) GroupHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var g GroupRequest
		if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
			log.Printf("failed to decode err = %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if g.UserID == 0 {
			log.Println("user_id is required")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		id, err := h.groupUsecase.Group(r.Context(), &groupuc.GroupInput{
			UserID: g.UserID,
			Name:   g.Name,
		})
		if err != nil {
			log.Printf("failed to put group err = %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var gr GroupResponse
		gr.ID = id
		b, err := json.Marshal(gr)
		if err != nil {
			log.Printf("failed to encode err = %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

type UserRequest struct {
	Name string `json:"name"`
}

type UserResponse struct {
	UserID  int64  `json:"user_id"`
	GroupID int64  `json:"group_id"`
	JWT     string `json:"jwt"`
}

func (h *handle) UserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var u UserRequest
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			log.Printf("failed to decode err = %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		puo, err := h.userUsecase.User(r.Context(), &useruc.UserInput{
			Name: u.Name,
		})
		if err != nil {
			log.Printf("failed to put user err = %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var ur UserResponse
		ur.UserID = puo.UserID
		ur.GroupID = puo.GroupID
		ur.JWT = puo.JWT
		b, err := json.Marshal(ur)
		if err != nil {
			log.Printf("failed to encode err = %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (h *handle) GroupsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		id := r.URL.Query().Get("user_id")

		gs, err := h.groupUsecase.Groups(r.Context(), id)
		if err != nil {
			log.Printf("failed to get groups err = %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		groups := make([]*Group, len(gs))
		for i, g := range gs {
			groups[i] = &Group{
				ID:   g.ID,
				Name: g.Name,
			}
		}
		b, err := json.Marshal(groups)
		if err != nil {
			log.Printf("failed to encode err = %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
