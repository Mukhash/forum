package main

import (
	"fmt"
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

	database, err := db.ConnectBD()
	if err != nil {
		log.Fatal(err)
	}
	db.FillDatabase(database)
	fmt.Println("Connected to database...")
	env.SetDB(database)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/css"))))
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./static/js"))))

	mux.Handle("/", env.Middleware(env.MainHandler()))
	mux.Handle("/registration", env.Middleware(env.RegHandler()))
	mux.Handle("/login", env.Middleware(env.LogHandler()))
	mux.Handle("/post/", env.Middleware(env.PostHandler()))
	mux.Handle("/next_posts", env.NextPostsHandler())
	mux.Handle("/test_index", env.TestIndexHandler())

	mux.HandleFunc("/single_sign_on", env.HandleSignOn)
	mux.HandleFunc("/reg_sign_on", env.HandleRegSignOn)
	mux.HandleFunc("/test", env.TestHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
