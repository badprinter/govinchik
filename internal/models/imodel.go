package models

type IModel interface {
	Create()
	Get(id int64)
	Update()
	Delete()
}
