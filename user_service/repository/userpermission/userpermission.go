package userpermission

import (
	"context"
	"database/sql"

	"pucsd2020-pp/user_service/driver"
	"pucsd2020-pp/user_service/model"
)

type userPermissionRepository struct {
	conn *sql.DB
}

//NewUserPermissionRepository ...
func NewUserPermissionRepository(conn *sql.DB) *userPermissionRepository {
	return &userPermissionRepository{conn: conn}
}

func (upr *userPermissionRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.UserPermission)
	return driver.GetById(upr.conn, obj, id)
}

func (upr *userPermissionRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	permission := obj.(model.UserPermission)
	result, err := driver.Create(upr.conn, &permission)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	permission.Id = id
	return id, nil
}

func (upr *userPermissionRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	permission := obj.(model.UserPermission)
	err := driver.UpdateById(upr.conn, &permission)
	return obj, err
}

func (upr *userPermissionRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.UserPermission{Id: id}
	return driver.SoftDeleteById(upr.conn, obj, id)
}

func (upr *userPermissionRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.UserPermission{}
	return driver.GetAll(upr.conn, obj, 0, 0)
}
