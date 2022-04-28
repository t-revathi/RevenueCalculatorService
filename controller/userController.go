package controller

import (
	"api-traderevenuecalculator/service"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/go-chi/render"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (u *UserController) WireRoutes(r chi.Router) {
	r.Route("/", func(r chi.Router) {
		r.Post("/calculateRevenue", u.PerformCalculateProfit)
		r.Get("/healthCheck", u.healthCheck)
	})
}

func (u *UserController) PerformCalculateProfit(w http.ResponseWriter, r *http.Request) {
	var dataCalculateRevenue service.DataCalculateRevenue

	if err := render.DecodeJSON(r.Body, &dataCalculateRevenue); err != nil {
		return
	}
	u.userService.PerformCalculateProfit(r.Context(), w, r, &dataCalculateRevenue)
}

func (u *UserController) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Healthcheck good"))
}
