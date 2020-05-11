package http

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"pucsd2020-pp/user_service/handler"
	"pucsd2020-pp/user_service/model"
	"pucsd2020-pp/user_service/repository"
)

type Group struct {
	handler.HTTPHandler
	repo repository.IRepository
}

func NewGroupHandler(conn *sql.DB) *User {
	return &User{
		repo: group.NewGroupRepository(),
	}
}

func (group *Group) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "user/{id}", Func: user.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "user", Func: user.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "user/{id}", Func: user.Update},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "user/{id}", Func: user.Delete},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "user", Func: user.GetAll},
	}
}

func (group *Group) GetByID(w http.ResponseWriter, r *http.Request) {
	var grp interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		grp, err := group.repo.GetByID(r.Context(), id)
		break
	}

	handler.WriteJSONResponse(w, r, grp, http.StatusOK, err)
}

func (group *Group) Create(w http.ResponseWriter, r *http.Request) {
	var grp model.Group
	err := json.NewDecoder(r.Body).Decode(&grp)
	for {
		if nil != err {
			break
		}

		_, err = user.repo.Create(r.Context(), grp)
		break
	}
	handler.WriteJSONResponse(w, r, grp, http.StatusOK, err)
}

func (group *Group) Update(w http.ResponseWriter, r *http.Request) {
	var iGrp interface{}
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	grp := model.Group{}
	err := json.NewDecoder(r.Body).Decode(&grp)
	for {
		if nil != err {
			break
		}
		grp.Id = id
		if nil != err {
			break
		}

		// set logged in user id for tracking update
		grp.UpdatedBy = 0

		iGrp, err = group.repo.Update(r.Context(), grp)
		if nil != err {
			break
		}
		grp = iGrp.(model.Group)
		break
	}

	handler.WriteJSONResponse(w, r, grp, http.StatusOK, err)
}

func (group *Group) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		err = group.repo.Delete(r.Context(), id)
		if nil != err {
			break
		}
		payload = "User deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (group *Group) GetAll(w http.ResponseWriter, r *http.Request) {
	usrs, err := group.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}
