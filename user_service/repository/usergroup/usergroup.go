package usergroup

import (
	"context"
	"database/sql"

	"pucsd2020-pp/user_service/driver"
	"pucsd2020-pp/user_service/model"
)

type userGroupRepository struct {
	conn *sql.DB
}

//NewUserGroupRepository ...
func NewUserGroupRepository(conn *sql.DB) *userGroupRepository {
	return &userGroupRepository{conn: conn}
}

func (ugr *userGroupRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.UserGroup)
	return driver.GetById(ugr.conn, obj, id)
}

func (ugr *userGroupRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	permission := obj.(model.UserGroup)
	result, err := driver.Create(ugr.conn, &permission)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	permission.Id = id
	return id, nil
}

func (ugr *userGroupRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	permission := obj.(model.UserGroup)
	err := driver.UpdateById(ugr.conn, &permission)
	return obj, err
}

func (ugr *userGroupRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.UserGroup{Id: id}
	return driver.SoftDeleteById(ugr.conn, obj, id)
}

func (ugr *userGroupRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.UserGroup{}
	return driver.GetAll(ugr.conn, obj, 0, 0)
}
