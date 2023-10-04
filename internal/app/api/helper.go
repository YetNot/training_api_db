package api

import (
	"dev/projects/ServerAndDB/internal/app/middleware"
	"dev/projects/ServerAndDB/storage"
	"net/http"

	_ "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	prefix string = "/api/v1"
)

func (a *API) configreLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

func (a *API) configreRouterField() {
	a.router.HandleFunc(prefix + "/articles", a.GetAllArticles).Methods("GET")
	a.router.Handle(prefix + "/articles/{id}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.GetArticleById),
	)).Methods("GET")
	a.router.HandleFunc(prefix + "/articles/{id}", a.DeleteArticleById).Methods("DELETE")
	a.router.HandleFunc(prefix + "/articles", a.PostArticle).Methods("POST")
	a.router.HandleFunc(prefix + "/user/register", a.PostUserRegister).Methods("POST")
	a.router.HandleFunc(prefix + "/user/auth", a.PostToAuth).Methods("POST")
}

func (a *API) configreStorageField() error {
	storage := storage.New(a.config.Storage)
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
