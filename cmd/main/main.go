package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_"github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yash_death/go-bookstore/pkg/routes"

)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9010",r))
}