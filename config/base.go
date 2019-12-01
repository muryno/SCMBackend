package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

var db *gorm.DB

func init() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	conn.SingularTable(true)

	//config.GetDB().Model(model.RegisteredCourse{}).AddForeignKey("student_id","student(id)","CASCADE","CASCADE")
	//
	//
	//config.GetDB().Debug().AutoMigrate(&model.Student{}, &model.Faculty{},&model.Department{}, &model.Course{},model.RegisteredCourse{},model.DepartmentCourse{})

}
func GetDB() *gorm.DB {
	return db
}

