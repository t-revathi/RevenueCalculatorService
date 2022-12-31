package main

import (
	"encoding/json"
	"fmt"

	service "github.com/t-revathi/RevenueCalculatorService/service/userservice"
)

func Main(args map[string]interface{}) map[string]interface{} {
	var dataCalculateRevenue service.DataCalculateRevenue

	blob, _ := json.Marshal(args)

	json.Unmarshal(blob, &dataCalculateRevenue)

	name, ok := args["name"].(string)
	if !ok {
		name = "stranger"
	}
	msg := make(map[string]interface{})
	msg["body"] = "Hello From Revenue Service - " + name
	msg["args"] = fmt.Sprintf("%v+", args)
	msg["obj"] = fmt.Sprintf("%v+", dataCalculateRevenue)
	return msg
}
