package database

type SqlHandler interface {
	Create(*interface{}) (Result, error)
	Update(interface{}) (Result, error)
	Query(interface{}) (Result, error)
	Destroy(interface{}) error
}

type Result *interface{}
