package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"schoolManagementAPI/config"
	"schoolManagementAPI/model"
)

func main() {

	router := mux.NewRouter()


	config.GetDB().Model(model.RegisteredCourse{}).AddForeignKey("student_id","student(id)","CASCADE","CASCADE")


	config.GetDB().Debug().AutoMigrate(&model.Student{}, &model.Faculty{},&model.Department{}, &model.Course{},model.RegisteredCourse{},model.DepartmentCourse{})



	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}

	log.Fatal(err)
}
