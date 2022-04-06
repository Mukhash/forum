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
		fmt.Println("connect")
		log.Fatal(err)
	}
	_ = db.FillDatabase(database)
	fmt.Println("Connected to database...")
	env.SetDB(database)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/css"))))
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./static/js"))))

	mux.Handle("/", env.Middleware(env.MainHandler()))
	mux.Handle("/registration", env.Middleware(env.RegHandler()))
	mux.Handle("/login", env.Middleware(env.LogHandler()))
	mux.Handle("/logout", env.Middleware(env.LogoutHandler()))
	mux.Handle("/post/", env.Middleware(env.PostHandler()))
	mux.Handle("/comment", env.Middleware(env.CommentHandler()))
	mux.Handle("/like_post", env.Middleware(env.PostLikeHandler()))
	mux.Handle("/like_comment", env.Middleware(env.CommentLikeHandler()))
	mux.Handle("/search", env.Middleware(env.SearchHandler()))

	mux.Handle("/next_posts", env.NextPostsHandler())
	mux.Handle("/next_comments", env.NextComments())

	mux.HandleFunc("/single_sign_on", env.HandleSignOn)
	mux.HandleFunc("/reg_sign_on", env.HandleRegSignOn)

	fmt.Println("listening at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
