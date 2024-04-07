package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!\n"))
	}).Methods(http.MethodGet)

	// r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	// r.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	// r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	// r.HandleFunc("/article/nice", aCon.PostNiceHandler).Methods(http.MethodPost)

	// r.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)

	return r
}
