package test

import (
	"context"

	service "github.com/t-revathi/revenuecalculatorservice/service/userservice"
)

type noDB struct {
}

func NewDBService(dburi string) *noDB {

	return &noDB{}
}

func (db *noDB) Pingdb(context.Context) error {
	return nil
}

func (db *noDB) Insertone(ctx context.Context, dataBase string, col string, doc interface{}) *service.InsertOneResult {
	return &service.InsertOneResult{
		Result: nil,
		Err:    nil,
	}
}
func (db *noDB) FindAll(ctx context.Context, dataBase string, col string, filter interface{}) *string {
	return nil

}
