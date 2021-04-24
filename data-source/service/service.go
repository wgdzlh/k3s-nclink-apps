package service

type Service interface {
	New() interface{}
	IdOf(interface{}) string
	Dup(string, interface{}) interface{}
	Slice() interface{}
	LenOf(interface{}) int64
	Create(interface{}) error
	Save(interface{}) error
	FindById(string, interface{}) error
	FindAll(interface{}) error
	FindWithFilter(map[string]string, interface{}) (int64, error)
	Delete(interface{}) error
	DeleteById(string) error
	Update(interface{}) error
	UpdateById(string, interface{}) (bool, error)
	Rename(string, string) error
}
