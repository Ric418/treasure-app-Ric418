package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/voyagegroup/treasure-app/httputil"
	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
	"github.com/voyagegroup/treasure-app/service"
)

type Articlecomment struct {
	dbx *sqlx.DB
}

func NewArticlecomment(dbx *sqlx.DB) *Articlecomment {
	return &Articlecomment{dbx: dbx}
}

func (a *Articlecomment) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	articles, err := repository.AllArticlecomment(a.dbx)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, articles, nil
}

func (a *Articlecomment) Show(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	article, err := repository.FindArticlecomment(a.dbx, aid)
	if err != nil && err == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, article, nil
}

func (a *Articlecomment) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	newArticlecomment := &model.Articlecomment{}
	if err := json.NewDecoder(r.Body).Decode(&newArticlecomment); err != nil {
		return http.StatusBadRequest, nil, err
	}

	contextUser, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		log.Print(err)
	}
	newArticlecomment.UserID = contextUser.ID

	articleService := service.NewArticlecommentService(a.dbx)
	id, err := articleService.Create(newArticlecomment)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	newArticlecomment.ID = id

	return http.StatusCreated, newArticlecomment, nil
}

func (a *Articlecomment) Update(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	reqArticlecomment := &model.Articlecomment{}
	if err := json.NewDecoder(r.Body).Decode(&reqArticlecomment); err != nil {
		return http.StatusBadRequest, nil, err
	}

	articleService := service.NewArticlecommentService(a.dbx)
	err = articleService.Update(aid, reqArticlecomment)
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusNoContent, nil, nil
}

func (a *Articlecomment) Destroy(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	articleService := service.NewArticleService(a.dbx)
	err = articleService.Destroy(aid)
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusNoContent, nil, nil
}
