package main

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/heyuan110/go-mega-code/controller"
	"github.com/heyuan110/go-mega-code/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//setup db
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	//start up application
	controller.Startup()
	//start up http server, listen 8888 port
	log.Println("Start up http server, listen 8888 port")
	log.Println("Http server running...")
	log.Println("Open http://127.0.0.1:8888")
	http.ListenAndServe(":8888", context.ClearHandler(http.DefaultServeMux))
}
