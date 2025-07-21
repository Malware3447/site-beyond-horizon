package requests

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"sbh/internal/requests/new_user"
)

type Api struct {
	router *chi.Mux
	us     *new_user.User
}

func NewApi(us *new_user.User) *Api {
	return &Api{
		router: nil,
		us:     us,
	}
}

func (a *Api) Init() error {
	a.router = chi.NewRouter()

	a.router.Route("/api/v1", func(r chi.Router) {
		r.Post("/addUserData/", a.us.AddUser)
	})

	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%v", 8080), a.router); err != nil {
			fmt.Println(err)
		}
	}()

	return nil
}
