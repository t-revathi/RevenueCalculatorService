package main

import (
	service "api-traderevenuecalculator/service/userservice"
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, dataCalculateRevenue service.DataCalculateRevenue) (service.RevenueCollection, error) {

	return service.NewUserService().PerformCalculateProfit(ctx, &dataCalculateRevenue), nil
	//return fmt.Sprintf("Hello %+v %+v", dataCalculateRevenue.Config, dataCalculateRevenue.TransactionData), nil
}

func main() {
	lambda.Start(HandleRequest)
}
