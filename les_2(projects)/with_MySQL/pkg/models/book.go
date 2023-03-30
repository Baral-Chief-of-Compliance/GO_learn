package models


import(
	"github.com/jinzhu/gorm"
	"bookstore/pkg/config"
)


var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `gotm: ""json:"name`
	Author string `json: "author"`
	Publication string `json: "publication"`
}

func init(){
	config.Connetc()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

