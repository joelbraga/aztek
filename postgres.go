package aztek

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"reflect"
	"fmt"
	"time"
	"log"
)

type PostgresCoreRepo struct {
	DB *gorm.DB
}

type PostgresCoreRepoOptions struct {
	Host            string
	User            string
	DB              string
	SSLMode         string
	Password        string
	Log             bool
	MaxIdleConns    int
	MaxOpensConns   int
	CoonMaxLifetime time.Duration
}

func NewPostgresCoreRepoOptions() PostgresCoreRepoOptions {
	options := PostgresCoreRepoOptions{}
	options.Log = false
	options.MaxIdleConns = -1
	options.MaxOpensConns = -1
	options.CoonMaxLifetime = -1
	return options
}

func NewPostgresCoreRepo(options PostgresCoreRepoOptions) *PostgresCoreRepo {

	if options.Host == "" {
		log.Fatalln("options.Host is missing.")
	}

	if options.User == "" {
		log.Fatalln("options.User is missing.")
	}

	if options.DB == "" {
		log.Fatalln("options.DB is missing.")
	}

	if options.SSLMode == "" {
		log.Fatalln("options.SSLMode is missing.")
	}

	if options.Password == "" {
		log.Fatalln("options.Password is missing.")
	}

	connection := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s",
		options.Host,
		options.User,
		options.DB,
		options.SSLMode,
		options.Password,
	)

	db, err := gorm.Open("postgres", connection)

	if err != nil {
		panic("failed to connect database")
	}

	db.LogMode(options.Log)

	if options.MaxIdleConns != -1 {
		db.DB().SetMaxIdleConns(options.MaxIdleConns)
	}

	if options.MaxOpensConns != -1 {
		db.DB().SetMaxOpenConns(options.MaxOpensConns)
	}

	if options.CoonMaxLifetime != -1 {
		db.DB().SetConnMaxLifetime(options.CoonMaxLifetime)
	}

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

func setPreload(db *gorm.DB, preload []string) *gorm.DB {
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
	return rp.DB.Save(model).Error
}

// TODO Remove associations
func (rp *PostgresCoreRepo) Delete(id string, model interface{}) error {
	return rp.DB.Where("id = ?", id).Unscoped().Delete(model).Error
}

func (rp *PostgresCoreRepo) Create(model interface{}) error {
	return rp.DB.Create(model).Error
}
