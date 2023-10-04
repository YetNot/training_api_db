package api

import (
	"dev/projects/ServerAndDB/internal/app/middleware"
	"dev/projects/ServerAndDB/internal/app/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
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
	initHeaders(writer)
	api.logger.Info("Get Article by ID /api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing {id} params", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unapropriate id value. don`t use ID as uncasting to int value",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	article, ok, err := api.storage.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (articles) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Can not find article with that ID in database")
		msg := Message{
			StatusCode: 404,
			Message:    "Article with that ID does not exists in database",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(article)
}

func (api *API) DeleteArticleById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete Article by id DELETE /api/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing {id} params", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unapropriate id value. don`t use ID as uncasting to int value",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, ok, err := api.storage.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (articles) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Can not find article with that ID in database")
		msg := Message{
			StatusCode: 404,
			Message:    "Article with that ID does not exists in database",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, err = api.storage.Article().DeleteById(id)
	if err != nil {
		api.logger.Info("Troubles while deleting element from database table (articles) with id. err:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	msg := Message{
		StatusCode: 202,
		Message:    fmt.Sprintf("Article with ID %d successfully deleted", id),
		IsError:    false,
	}
	writer.WriteHeader(202)
	json.NewEncoder(writer).Encode(msg)
}

func (api *API) PostArticle(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post Article POST /api/v1/articles")
	var article models.Article
	err := json.NewDecoder(request.Body).Decode(&article)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	a, err := api.storage.Article().Create(&article)
	if err != nil {
		api.logger.Info("Troubles while creating new articles", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try againg",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(a)
}

func (api *API) PostUserRegister(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post User Register /api/v1/user/register")
	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, ok, err := api.storage.User().FindByLogin(user.Login)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (users) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if ok {
		api.logger.Info("User with that ID already exists")
		msg := Message{
			StatusCode: 400,
			Message:    "User with that login already exists in database",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	userAdded, err := api.storage.User().Create(&user)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (users) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	msg := Message{
		StatusCode: 201,
		Message:    fmt.Sprintf("User {login:%s} successfully registered!", userAdded.Login),
		IsError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}

func (api *API) PostToAuth(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post User Register /api/v1/user/register")
	var userFromJson models.User
	err := json.NewDecoder(request.Body).Decode(&userFromJson)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	userInDB, ok, err := api.storage.User().FindByLogin(userFromJson.Login)
	if err != nil {
		api.logger.Info("Can not make user search in database", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles while accessing database",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("User with that login does not exists")
		msg := Message{
			StatusCode: 400,
			Message:    "User with that login does not exists in database. Try register first",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if userInDB.Password != userFromJson.Password {
		api.logger.Info("Invalid credetials to auth")
		msg := Message{
			StatusCode: 404,
			Message:    "Your password is invalid",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	claims["admin"] = true
	claims["name"] = userInDB.Login
	tokenString, err := token.SignedString(middleware.SecretKey)
	if err != nil {
		api.logger.Info("Can not claim jwt-token")
		msg := Message {
			StatusCode: 500,
			Message: "We have some troubles. Try again",
			IsError: true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	msg := Message {
		StatusCode: 201,
		Message: tokenString,
		IsError: false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}
