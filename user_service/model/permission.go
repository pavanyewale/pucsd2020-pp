package model

type Permission struct {
	Id             int64  `json:"id,omitempty" key:"primary" autoincr:"1" column:"id"`
	PermissionName string `json:"permission_name" column:"permission_name"`
	UpdatedBy      int64  `json:"updated_by" column:"updated_by"`
}

func (permission *Permission) Table() string {
	return "permission"
}

func (permission *Permission) String() string {
	return Stringify(permission)
}
