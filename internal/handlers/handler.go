package handlers

import (
	"log"
	"net/http"

	"github.com/NhanNT-VNG/go-auth/internal/config"
	"github.com/NhanNT-VNG/go-auth/internal/helpers"
	"github.com/NhanNT-VNG/go-auth/internal/models"
)

var Repo *Repository

type Repository struct {
	Config *config.Config
}

func NewRepository(config *config.Config) *Repository {
	return &Repository{
		Config: config,
	}
}

func NewHandlers(repository *Repository) {
	Repo = repository
}

func (repository *Repository) GetAllUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	err := helpers.ReadJson(w, r, user)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(user)

	response := helpers.JsonResponse{}
	response.Data = user
	helpers.WriteJson(w, http.StatusOK, response)
}
