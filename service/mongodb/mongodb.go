package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DBService struct {
	Client mongo.Client
	Ctx    context.Context
	Cancel context.CancelFunc
	Err    error
}

func NewDBService(dburi string) *DBService {
	client, ctx, cancel, err := connectdb(dburi)

	return &DBService{
		Client: *client,
		Ctx:    ctx,
		Cancel: cancel,
		Err:    err,
	}
}

/*func closedb(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	// CancelFunc to cancel to context
	defer cancel()

	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}*/

// This is a user defined method that returns mongo.Client,
// context.Context, context.CancelFunc and error.
// mongo.Client will be used for further database operation.
// context.Context will be used set deadlines for process.
// context.CancelFunc will be used to cancel context and
// resource associated with it.

func connectdb(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		60*time.Second)

	// mongo.Connect return mongo.Client method

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	if err == nil {
		fmt.Println("Mongodb Connected")
		databases, err := client.ListDatabaseNames(ctx, bson.M{})
		fmt.Println(databases, err)
	}

	return client, ctx, cancel, err
}

// This is a user defined method that accepts
// mongo.Client and context.Context
// This method used to ping the mongoDB, return error if any.
func (db *DBService) Pingdb(client *mongo.Client, ctx context.Context) error {

	// mongo.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// Ping method return error if any occurred, then
	// the error can be handled.
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}

func (db *DBService) Insertone(client *mongo.Client, ctx context.Context, dataBase string, col string, doc interface{}) (*mongo.InsertOneResult, error) {
	collection := client.Database(dataBase).Collection(col)
	result, err := collection.InsertOne(ctx, doc)
	//defer closedb(client,ctx,db.cancel)
	return result, err

}

func (db *DBService) FindOne(client *mongo.Client, ctx context.Context, dataBase string, col string, filter interface{}) *mongo.SingleResult {
	collection := client.Database(dataBase).Collection(col)
	result := collection.FindOne(ctx, filter)
	return result
}

func (db *DBService) FindAll(client *mongo.Client, ctx context.Context, dataBase string, col string, filter interface{}) *mongo.Cursor {
	collection := client.Database(dataBase).Collection(col)
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	return cursor
}
