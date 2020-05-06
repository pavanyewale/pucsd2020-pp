package model

type UserGroup struct {
	Id        int64 `json:"id,omitempty" key:"primary" autoincr:"1" column:"id"`
	UserId    int64 `json:"user_id" column:"user_id"`
	GroupId   int64 `json:"group_id" column:"group_id"`
	UpdatedBy int64 `json:"updated_by" column:"updated_by"`
}

func (ug *UserGroup) Table() string {
	return "user_group"
}

func (ug *UserGroup) String() string {
	return Stringify(ug)
}
