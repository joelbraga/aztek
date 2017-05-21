package aztek

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"reflect"
)

type PostgresCoreRepo struct {
	DB *gorm.DB
}

func NewPostgesCoreRepo(connectionString string) *PostgresCoreRepo {
	db, err := gorm.Open("postgres", connectionString)

	if err != nil {
		panic("failed to connect database")
	}

	db.LogMode(true)

	return &PostgresCoreRepo{
		DB: db,
	}
}

func (rp *PostgresCoreRepo) Migrations(models []interface{}) {
	for _, elm := range models {
		if !rp.DB.HasTable(elm) {
			rp.DB.CreateTable(elm)
		}
	}
}

func setPreload (db *gorm.DB, preload []string) *gorm.DB {
	if preload != nil {
		for _, elm := range preload {
			db = db.Preload(elm)
		}
	}
	return db
}

func (rp *PostgresCoreRepo) GetAll(model interface{}) (interface{}, error) {
	sliceOfModel := reflect.SliceOf(reflect.TypeOf(model))
	ptr := reflect.New(sliceOfModel)
	ptr.Elem().Set(reflect.MakeSlice(sliceOfModel, 0, 0))
	modelPtr := ptr.Interface()

	if err := rp.DB.Find(modelPtr).Error; err != nil {
		return nil, err
	} else {
		return modelPtr, nil
	}
}

func (rp *PostgresCoreRepo) GetById(id string, model interface{}, preload []string) (interface{}, error) {
	ptr := reflect.New(reflect.TypeOf(model))
	modelPtr := ptr.Interface()

	db := setPreload(rp.DB, preload)
	if err := db.First(modelPtr, "id = ?", id).Error; err != nil {
		return modelPtr, err
	} else {
		return modelPtr, nil
	}
}

func (rp *PostgresCoreRepo) GetByCode(code string, model interface{}, preload []string) (interface{}, error) {
	ptr := reflect.New(reflect.TypeOf(model))
	modelPtr := ptr.Interface()

	db := setPreload(rp.DB, preload)
	if err := db.First(modelPtr, "code = ?", code).Error; err != nil {
		return modelPtr, err
	} else {
		return modelPtr, nil
	}
}

func (rp *PostgresCoreRepo) GetWhere(model interface{}, preload []string) (interface{}, error) {
	ptr := reflect.New(reflect.TypeOf(model))
	modelPtr := ptr.Interface()

	db := setPreload(rp.DB, preload)
	if err := db.First(modelPtr, model).Error; err != nil {
		return modelPtr, err
	} else {
		return modelPtr, nil
	}
}

func (rp *PostgresCoreRepo) GetWhereMultiple(model interface{}, preload []string) (interface{}, error) {
	sliceOfModel := reflect.SliceOf(reflect.TypeOf(model))
	ptr := reflect.New(sliceOfModel)
	ptr.Elem().Set(reflect.MakeSlice(sliceOfModel, 0, 0))
	modelPtr := ptr.Interface()

	db := setPreload(rp.DB, preload)
	if err := db.Where(model).Find(modelPtr).Error; err != nil {
		return modelPtr, err
	} else {
		return modelPtr, nil
	}
}

func (rp *PostgresCoreRepo) Update(id string, model interface{}) error {
	if err := rp.DB.Save(model).Error; err != nil {
		return err
	}

	return nil
}

// TODO Remove associations
func (rp *PostgresCoreRepo) Delete(id string, model interface{}) error {
	if err := rp.DB.Where("id = ?", id).Unscoped().Delete(model).Error; err != nil {
		return err
	}

	return nil
}

func (rp *PostgresCoreRepo) Create(model interface{}) error {
	tx := rp.DB.Begin()

	if err := rp.DB.Create(model).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
