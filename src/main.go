package main

import (
	"forum/db"
	"forum/handlers"
	"forum/utils"
	"log"
	"net/http"
)

func main() {
	env := handlers.InitEnv()

	tmpl, err := utils.GetTmpl()
	if err != nil {
		log.Fatal(err)
	}
	env.SetTmpl(tmpl)

	db, err := db.ConnectBD()
	if err != nil {
		log.Fatal(err)
	}
	env.SetDB(db)

	mux := http.NewServeMux()
	mux.Handle("../templates/", http.StripPrefix("../templates/", http.FileServer(http.Dir("../templates/css"))))

	mux.Handle("/", handlers.Middleware(env.MainHandler()))
	mux.Handle("/registration", handlers.Middleware(env.RegHandler()))
	log.Fatal(http.ListenAndServe(":8080", mux))
}