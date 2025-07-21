package new_user

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	model "sbh/internal/models/pg"
	"sbh/internal/services/pg"
)

type User struct {
	repo *pg.PgService
}

func NewUser(repo *pg.PgService) *User {
	return &User{repo: repo}
}

func (u *User) AddUser(w http.ResponseWriter, r *http.Request) {
	log := logrus.New()
	ctx := context.WithValue(r.Context(), "", nil)
	body := model.UserDataParams{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Error("Invalid request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err = u.repo.AddUser(ctx, &body)
	if err != nil {
		log.Error("Ошибка добавления в базу данных")
		http.Error(w, "Ошибка добавления в базу данных", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(true); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
