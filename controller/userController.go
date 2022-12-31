package controller

import (
	service "api-traderevenuecalculator/service/userservice"
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/go-chi/render"
)

type DBInterface interface {
	Pingdb(context.Context) error
	Insertone(ctx context.Context, dataBase string, col string, doc interface{}) *service.InsertOneResult
	FindAll(ctx context.Context, dataBase string, col string, filter interface{}) *string
}

type UserController struct {
	userService *service.UserService
	dbService   DBInterface
}

func NewUserController(savetodb string, dburi string, dbinterface DBInterface) *UserController {
	return &UserController{
		userService: service.NewUserService(),
		dbService:   dbinterface,
	}
}
func Health(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (u *UserController) WireRoutes(r chi.Router) {
	r.Route("/", func(r chi.Router) {
		r = r.With(Health)
		r.Post("/calculateRevenue", u.PerformCalculateProfit)

		r.Get("/healthCheck", u.healthCheck)
		r.Get("/allRevenue", u.ShowAllRevenue)
	})
}

func (u *UserController) PerformCalculateProfit(w http.ResponseWriter, r *http.Request) {
	var dataCalculateRevenue service.DataCalculateRevenue

	if err := render.DecodeJSON(r.Body, &dataCalculateRevenue); err != nil {
		return
	}
	result := u.userService.PerformCalculateProfit(r.Context(), &dataCalculateRevenue)

	// Insert and Listing opertaions
	dbname := "stockprofitcalculator"
	collection := "plResults"

	res := u.dbService.Insertone(r.Context(), dbname, collection, result.Items)

	if res.Err != nil {
		fmt.Println("Error Occured during insertion" + res.Err.Error())
	}
	fmt.Println(res, res.Err)
	//Listing the last inserted Record
	// id, _ := res.Result.(primitive.ObjectID)
	// //primitive.ObjectIDFromHex(string("62ccdf87b79b0e2fc4ea67f0"))
	// filter := bson.D{{Key: "_id", Value: id}}

	// var record bson.M
	// u.dbService.FindOne(&u.dbService.Client, u.dbService.Ctx, dbname, collection, filter).Decode(&record)

	// fmt.Println(record)

	//Return results to client
	render.JSON(w, r,
		result.Items)
}

func (u *UserController) ShowAllRevenue(w http.ResponseWriter, r *http.Request) {

	collection := "plResults"

	// // Get all records
	filter := ""
	results := u.dbService.FindAll(r.Context(), "stockprofitcalculator", collection, filter)

	fmt.Println(results)
	// //Return results to client
	render.JSON(w, r, results)

}

func (u *UserController) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Healthcheck good"))
}
