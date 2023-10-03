package api

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func (api *API) GetAllArticles(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get all Articles GET /api/v1/articles")
	articles, err := api.storage.Article().SelectAll()
	if err != nil {
		api.logger.Info("Error while Articles.SelectAll :", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again later",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(articles)
}

func (api *API) GetArticleById(writer http.ResponseWriter, request *http.Request) {

}

func (api *API) DeleteArticleById(writer http.ResponseWriter, request *http.Request) {

}

func (api *API) PostArticle(writer http.ResponseWriter, request *http.Request) {

}

func (api *API) PostUserRegister(writer http.ResponseWriter, request *http.Request) {

}
