package model

type GroupPermission struct {
	Id           int64 `json:"id,omitempty" key:"primary" autoincr:"1" column:"id"`
	GroupId      int64 `json:"group_id" column:"group_id"`
	PermissionId int64 `json:"permission_id" column:"permission_id"`
	UpdatedBy    int64 `json:"updated_by" column:"updated_by"`
}

func (gp *GroupPermission) Table() string {
	return "group_permission"
}

func (gp *GroupPermission) String() string {
	return Stringify(gp)
}
