package database

type SqlHandler interface {
	Create(*interface{}) (Result, error)
}

type Result *interface{}
