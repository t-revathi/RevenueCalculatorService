package main

import (
	"api-traderevenuecalculator/controller"
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func main() {
	var dburi string
	flag.StringVar(&dburi, "db-uri", "mongodb://0.0.0.0:27017/stockprofitcalculator", "data base name to connect")
	savetodb := "mongodb"
	startService(savetodb, dburi)
}
func LogRequest(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func startService(savetodb string, dburi string) {
	router := chi.NewRouter()

	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(LogRequest)

	userController := controller.NewUserController(savetodb, dburi)
	userController.WireRoutes(router)

	serverAddr := ":3333"

	for _, r := range router.Routes()[0].SubRoutes.Routes() {
		fmt.Printf(r.Pattern)
	}
	if err := http.ListenAndServe(serverAddr, router); err != nil {
		fmt.Println("error")
	}
}
