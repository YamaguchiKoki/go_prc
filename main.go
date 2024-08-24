package main

import (
	"log"
	"net/http"

	"github.com/YamaguchiKoki/go_prc/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler)
	r.HandleFunc("/article/list", handlers.ArticleListHandler)

	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler)
	r.HandleFunc("/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")

	//log.Fatal: 実行時異常としてプログラム終了
	log.Fatal(http.ListenAndServe(":8080", r))
}