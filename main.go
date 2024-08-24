package main

import (
	"log"
	"net/http"

	"github.com/YamaguchiKoki/myapi/handlers"
)

func main() {
	
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	//log.Fatal: 実行時異常としてプログラム終了
	log.Fatal(http.ListenAndServe(":8080", nil))
}