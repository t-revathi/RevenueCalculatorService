# RevenueCalculatorService
RevenueCalculatorService is developed in Go. The service calculates the profit/loss based on the business logic by sorting and processing the stock trading data and show the results grouped based on the financial year. Both the input and output are in JSON.

# Project Structure
## cmd/userservice/userservice.go
Configuration and router intialization.Then the main function continues in controller

## controller
Server handler layer. Server routes are written

## service
Business logic implemented
Datas are sorted and put into appropriate bucket to make a accurate calculation and the results are grouped based on the financial year.
API accepts input and shows results in JSON. 

# Installation

```go
git clone https://github.com/t-revathi/RevenueCalculatorService.git

```

# Sample Input
```
{"Config":{"FinancialYear":"2021",
"StartFinancialMonth":"July",
"EndFinancialMonth":"June",
"SkipCorporateAction":true },
"TransactionData" :[
    {
        "Date" : "2018-11-09T12:42:31Z",
        "Market": "NEXTDC Ltd",
        "Cost":-3133,
        "Direction":"BUY",
        "Price": 6.25,
        "Activity":"TRADE",
        "Quantity": 500
    },
    {
         "Date" : "2019-02-14T12:42:31Z",
        "Market": "NEXTDC Ltd",
        "Cost":-3467.54,
        "Direction":"SELL",
        "Price": 6.965,
        "Activity":"TRADE",
        "Quantity": 499
        
    },
    {
         "Date" : "2019-04-02T12:42:31Z",
        "Market": "NEXTDC Ltd",
        "Cost":-1997.96,
        "Direction":"BUY",
        "Price": 6.18,
        "Activity":"TRADE",
        "Quantity": 322
        
    },
    {
        "Date" : "2019-04-26T12:42:31Z",
        "Market": "NEXTDC Ltd",
        "Cost":-2057.63,
        "Direction":"SELL",
        "Price": 6.415,
        "Activity":"TRADE",
        "Quantity": 322
    },
    {
        "Date" : "2019-06-27T12:42:31Z",
        "Market": "NEXTDC Ltd",
        "Cost":-1288,
        "Direction":"BUY",
        "Price": 6.4,
        "Activity":"TRADE",
        "Quantity": 200

    },
    {
        "Date" : "2019-12-11T12:42:31Z",
        "Market": "NEXTDC Ltd",
        "Cost":-1342,
        "Direction":"SELL",
        "Price": 6.75,
        "Activity":"TRADE",
        "Quantity": 200
    
    }

]
}
```
# Sample Output
```
[
    {
        "Year": "2018-2019",
        "Items": [
            {
                "Date": "2019-02-14T12:42:31Z",
                "Market": "NEXTDC Ltd",
                "Quantity": 499,
                "PandL": 340.8061,
                "SellUnitPrice": 6.948978
            },
            {
                "Date": "2019-04-26T12:42:31Z",
                "Market": "NEXTDC Ltd",
                "Quantity": 322,
                "PandL": 59.60878,
                "SellUnitPrice": 6.390155
            }
        ]
    },
    {
        "Year": "2019-2020",
        "Items": [
            {
                "Date": "2019-12-11T12:42:31Z",
                "Market": "NEXTDC Ltd",
                "Quantity": 200,
                "PandL": 54.235153,
                "SellUnitPrice": 6.71
            }
        ]
    }
]
```



