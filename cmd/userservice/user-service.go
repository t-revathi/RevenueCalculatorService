package main

import (
	"api-traderevenuecalculator/controller"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func main() {
	startService()
}
func LogRequest(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func startService() {
	router := chi.NewRouter()

	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(LogRequest)
	

	userController := controller.NewUserController()
	userController.WireRoutes(router)

	serverAddr := ":3333"

	for _, r := range router.Routes()[0].SubRoutes.Routes() {
		fmt.Printf(r.Pattern)
	}
	if err := http.ListenAndServe(serverAddr, router); err != nil {
		fmt.Println("error")
	}
}
