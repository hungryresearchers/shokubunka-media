package database

type SqlHandler interface {
	Create(*interface{}) (*interface{}, error)
}
