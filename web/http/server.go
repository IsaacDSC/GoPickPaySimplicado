package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func StartServerHttp() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	StartRoutes(router)
	fmt.Println("[ * ] start application 3000")
	http.ListenAndServe(":3000", router)
}
