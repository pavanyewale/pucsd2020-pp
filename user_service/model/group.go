package model

type Group struct {
	Id        int64  `json:"id,omitempty" key:"primary" autoincr:"1" column:"id"`
	GroupName string `json:"group_name" column:"group_name"`
	UpdatedBy int64  `json:"updated_by" column:"updated_by"`
}

func (group *Group) Table() string {
	return "group"
}

func (group *Group) String() string {
	return Stringify(group)
}
