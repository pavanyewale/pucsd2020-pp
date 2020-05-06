package model

type UserPermission struct {
	Id           int64 `json:"id,omitempty" key:"primary" autoincr:"1" column:"id"`
	UserId       int64 `json:"user_id,omitempty"  column:"user_id"`
	PermissionId int64 `json:"permission_id,omitempty"  column:"permission_id"`
	UpdatedBy    int64 `json:"updated_by" column:"updated_by"`
}

func (up *UserPermission) Table() string {
	return "user_permission"
}

func (up *UserPermission) String() string {
	return Stringify(up)
}
