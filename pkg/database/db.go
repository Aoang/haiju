package database

import (
	"github.com/Aoang/haiju/pkg/model"
	"os"
	"sync"

	"github.com/Aoang/haiju/pkg/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql
)

var log = logger.NewLogger(os.Stdout)

var db *gorm.DB

// ConnectDB connects to the database
func ConnectDB() {
	var err error
	if "" != model.Conf.MySQL {
		db, err = gorm.Open("mysql", model.Conf.MySQL)
	} else {
		log.Fatal("please specify database")
	}
	if nil != err {
		log.Fatalf("opens database failed: " + err.Error())
	}

	db.SingularTable(true)

	if err = db.AutoMigrate(model.Models...).Error; nil != err {
		log.Fatal("auto migrate tables failed: " + err.Error())
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
	db.LogMode(model.Conf.ShowSQL)
}

// DisconnectDB disconnects from the database
func DisconnectDB() {
	if err := db.Close(); nil != err {
		log.Errorf("Disconnect from database failed: " + err.Error())
	}
}

// Use database
func Use() *gorm.DB {
	return db
}

var (
	User = &userService{
		mutex: &sync.Mutex{},
	}
	Building = &buildingService{
		mutex: &sync.Mutex{},
	}
	Community = &communityService{
		mutex: &sync.Mutex{},
	}
	Property = &propertyService{
		mutex: &sync.Mutex{},
	}
	Location = &locationService{
		mutex: &sync.Mutex{},
	}
	File = &fileService{
		mutex: &sync.Mutex{},
	}
	Room = &roomService{
		mutex: &sync.Mutex{},
	}
)

type userService struct {
	mutex *sync.Mutex
}
type buildingService struct {
	mutex *sync.Mutex
}
type communityService struct {
	mutex *sync.Mutex
}
type propertyService struct {
	mutex *sync.Mutex
}
type locationService struct {
	mutex *sync.Mutex
}
type fileService struct {
	mutex *sync.Mutex
}
type roomService struct {
	mutex *sync.Mutex
}
