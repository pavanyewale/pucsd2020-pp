package grouppermission

import (
	"context"
	"database/sql"

	"pucsd2020-pp/user_service/driver"
	"pucsd2020-pp/user_service/model"
)

type groupPermissionRepository struct {
	conn *sql.DB
}

//NewGroupPermissionRepository ...
func NewGroupPermissionRepository(conn *sql.DB) *groupPermissionRepository {
	return &groupPermissionRepository{conn: conn}
}

func (gpr *groupPermissionRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.GroupPermission)
	return driver.GetById(gpr.conn, obj, id)
}

func (gpr *groupPermissionRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	permission := obj.(model.GroupPermission)
	result, err := driver.Create(gpr.conn, &permission)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	permission.Id = id
	return id, nil
}

func (gpr *groupPermissionRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	permission := obj.(model.GroupPermission)
	err := driver.UpdateById(gpr.conn, &permission)
	return obj, err
}

func (gpr *groupPermissionRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.GroupPermission{Id: id}
	return driver.SoftDeleteById(gpr.conn, obj, id)
}

func (gpr *groupPermissionRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.GroupPermission{}
	return driver.GetAll(gpr.conn, obj, 0, 0)
}
