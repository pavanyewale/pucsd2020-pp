package model

//IModel prototype of db model
type IModel interface {
	Table() string
	String() string
}
