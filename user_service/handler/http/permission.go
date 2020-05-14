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
	"pucsd2020-pp/user_service/repository/permission"
)

type Permission struct {
	handler.HTTPHandler
	repo repository.IRepository
}

func NewPermissionHandler(conn *sql.DB) *Permission {
	return &Permission{
		repo: permission.NewPermissionRepository(conn),
	}
}

func (perm *Permission) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "permission/{id}", Func: perm.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "permission", Func: perm.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "permission/{id}", Func: perm.Update},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "permission/{id}", Func: perm.Delete},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "permission", Func: perm.GetAll},
	}
}

func (perm *Permission) GetByID(w http.ResponseWriter, r *http.Request) {
	var usr interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		usr, err = perm.repo.GetByID(r.Context(), id)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (perm *Permission) Create(w http.ResponseWriter, r *http.Request) {
	var permission model.Permission
	err := json.NewDecoder(r.Body).Decode(&permission)
	for {
		if nil != err {
			break
		}

		_, err = perm.repo.Create(r.Context(), permission)
		break
	}
	handler.WriteJSONResponse(w, r, permission, http.StatusOK, err)
}

func (perm *Permission) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	usr := model.Permission{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}
		usr.Id = id
		if nil != err {
			break
		}

		// set logged in user id for tracking update
		usr.UpdatedBy = 0

		iUsr, err = perm.repo.Update(r.Context(), usr)
		if nil != err {
			break
		}
		usr = iUsr.(model.Permission)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (perm *Permission) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		err = perm.repo.Delete(r.Context(), id)
		if nil != err {
			break
		}
		payload = "User deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (perm *Permission) GetAll(w http.ResponseWriter, r *http.Request) {
	usrs, err := perm.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}
