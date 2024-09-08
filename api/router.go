package api

import (
	"database/sql"
	"net/http"

	"github.com/YamaguchiKoki/go_prc/api/middlewares"
	"github.com/YamaguchiKoki/go_prc/controllers"
	"github.com/YamaguchiKoki/go_prc/services"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	ser := services.NewMyAppService(db)
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCpmmentController(ser)

	r := mux.NewRouter()

	r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)

	r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", aCon.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)

	r.Use(middlewares.LoggingMiddleware)

	return r
}