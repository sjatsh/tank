package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"tank/rest"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	rest.DbInit()

	//将运行时参数装填到config中去。
	rest.PrepareConfigs()
	context := rest.NewContext()
	defer context.Destroy()

	r := mux.NewRouter().StrictSlash(false)
	r.PathPrefix("/").Handler(context.Router)

	//http.Handle("/", context.Router)

	dotPort := fmt.Sprintf(":%v", rest.CONFIG.ServerPort)

	info := fmt.Sprintf("App started at http://localhost%v", dotPort)
	rest.LogInfo(info)

	err := http.ListenAndServe(dotPort, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
