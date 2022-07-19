package controller

import (
	mongo "api-traderevenuecalculator/service/mongodb"
	service "api-traderevenuecalculator/service/userservice"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/go-chi/render"
)

type UserController struct {
	userService *service.UserService
	dbService   *mongo.DBService
}

func NewUserController(savetodb string, dburi string) *UserController {
	return &UserController{
		userService: service.NewUserService(),
		dbService:   mongo.NewDBService(dburi),
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
	result := u.userService.PerformCalculateProfit(r.Context(), w, r, &dataCalculateRevenue)

	if u.dbService.Err != nil {
		panic(u.dbService.Err)
	}
	///err = u.dbService.Pingdb(client, ctx)

	//defer u.dbService.Closedb(client, ctx, cancel)

	// Insert and Listing opertaions
	dbname := "stockprofitcalculator"
	collection := "plResults"

	doc := bson.D{{Key: "data", Value: result.Items}}

	res, err := u.dbService.Insertone(&u.dbService.Client, u.dbService.Ctx, dbname, collection, doc)
	if err != nil {
		fmt.Println("Error Occured during insertion" + err.Error())
	}
	fmt.Println(res, err)
	//Listing the last inserted Record
	id, _ := res.InsertedID.(primitive.ObjectID)
	//primitive.ObjectIDFromHex(string("62ccdf87b79b0e2fc4ea67f0"))
	filter := bson.D{{Key: "_id", Value: id}}

	var record bson.M
	u.dbService.FindOne(&u.dbService.Client, u.dbService.Ctx, dbname, collection, filter).Decode(&record)

	fmt.Println(record)

	//Return results to client
	render.JSON(w, r,
		result.Items)
}

func (u *UserController) ShowAllRevenue(w http.ResponseWriter, r *http.Request) {

	collection := "plResults"
	filter := bson.M{}

	// Get all records
	cursor := u.dbService.FindAll(&u.dbService.Client, u.dbService.Ctx, "stockprofitcalculator", collection, filter)
	var results []bson.M
	if err := cursor.All(u.dbService.Ctx, &results); err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
	//Return results to client
	render.JSON(w, r, results)
}

func (u *UserController) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Healthcheck good"))
}
