package grouppermission

import (
	"context"
	"database/sql"

	"pucsd2020-pp/user_service/driver"
	"pucsd2020-pp/user_service/model"
)

type permissionRepository struct {
	conn *sql.DB
}

//NewUserRepository ...
func NewUserRepository(conn *sql.DB) *permissionRepository {
	return &permissionRepository{conn: conn}
}

func (perm *permissionRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.Permission)
	return driver.GetById(perm.conn, obj, id)
}

func (perm *permissionRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	permission := obj.(model.Permission)
	result, err := driver.Create(perm.conn, &permission)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	permission.Id = id
	return id, nil
}

func (perm *permissionRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	permission := obj.(model.Permission)
	err := driver.UpdateById(perm.conn, &permission)
	return obj, err
}

func (perm *permissionRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.Permission{Id: id}
	return driver.SoftDeleteById(perm.conn, obj, id)
}

func (perm *permissionRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.Permission{}
	return driver.GetAll(perm.conn, obj, 0, 0)
}
