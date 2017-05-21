package aztek

type Repository interface {
	Migrations(models []interface{})
	GetAll(model interface{}) (interface{}, error)
	GetById(id string, model interface{}, preload []string) (interface{}, error)
	GetByCode(code string, model interface{}, preload []string) (interface{}, error)
	GetWhere(model interface{}, preload []string) (interface{}, error)
	GetWhereMultiple(model interface{}, preload []string) (interface{}, error)
	Update(id string, model interface{}) error
	Delete(id string, model interface{}) error
	Create(model interface{}) error
}
