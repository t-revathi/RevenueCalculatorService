package service

import (
	"context"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) PerformCalculateProfit(ctx context.Context, req *DataCalculateRevenue) RevenueCollection {

	fmt.Printf("Received financial year:%s on %s", req.Config.FinancialYear, time.Now().String())
	transactionData := req.TransactionData

	/*render.JSON(w, r,
	DataCalculateRevenue{
		transactionData,
		req.Config,
	})*/
	income := ProcessTransactions(transactionData, req.Config)

	/*render.JSON(w, r,
	income.Items)*/
	return income

}

func ProcessTransactions(transactions []Transaction, config Config) RevenueCollection {
	CalculateUnitPrice(&transactions)
	fmt.Printf("\n %v+", transactions)
	buyShares := GetbuyShares(transactions, config)
	//fmt.Println(buyShares)
	sellShares := GetsellShares(transactions, config)
	//fmt.Println(sellShares)
	//return
	income := calculatePandL(buyShares, sellShares, config)
	//fmt.Println(income)
	return income
}

func CalculateUnitPrice(transactions *[]Transaction) {

	for i := 0; i < len(*transactions); i++ {

		unitprice := (*transactions)[i].Cost / float32((*transactions)[i].Quantity)

		(*transactions)[i].UnitPrice = float32(math.Abs(float64(unitprice)))

	}

}

func GetbuyShares(transactions []Transaction, config Config) []Transaction {

	buytransactions := make([]Transaction, 0)

	for _, t := range transactions {
		if config.SkipCorporateAction {
			if strings.ToLower(t.Activity) != "trade" {
				continue
			}
		}
		if strings.ToLower(t.Direction) == "buy" {
			buytransactions = append(buytransactions, t)

		}

	}
	return buytransactions
}

func GetsellShares(transaction []Transaction, config Config) []Transaction {
	selltransactions := make([]Transaction, 0)
	for _, t := range transaction {
		if config.SkipCorporateAction {
			if strings.ToLower(t.Activity) != "trade" {
				continue
			}
		}
		if strings.ToLower(t.Direction) == "sell" {
			selltransactions = append(selltransactions, t)

		}
	}
	sort.Slice(selltransactions, func(i, j int) bool {
		return selltransactions[i].Date.Before(selltransactions[j].Date)
	})
	return selltransactions
}

func calculatePandL(buyshares []Transaction, sellShares []Transaction, config Config) RevenueCollection {

	income := RevenueCollection{}

	for idx := range sellShares {
		pq := 0
		var pl float32 = 0.0
		var inc Income
		var currentRecordSellYear string
		currentSellRecord := sellShares[idx]
		if currentSellRecord.Date.Month() < 7 {
			currentRecordSellYear = strconv.Itoa((currentSellRecord.Date.Year() - 1)) + "-" + strconv.Itoa(currentSellRecord.Date.Year())
		} else {
			currentRecordSellYear = strconv.Itoa(currentSellRecord.Date.Year()) + "-" + strconv.Itoa((currentSellRecord.Date.Year() + 1))
		}

		inc.Date = currentSellRecord.Date
		inc.Market = currentSellRecord.Market
		inc.Quantity = currentSellRecord.Quantity
		inc.SellUnitPrice = currentSellRecord.UnitPrice
		fmt.Printf("Sell: %v \n", currentSellRecord)
		for currentSellRecord.Quantity > 0 {
			buyt := GetearlierbuyShare(buyshares, currentSellRecord.Market)
			if buyt.Quantity >= currentSellRecord.Quantity {
				pq = currentSellRecord.Quantity
			} else {
				pq = buyt.Quantity
			}
			//if currentSellRecord.Market == "Pilbara Minerals Limited" {
			fmt.Printf("buy records:%v  \n\n", buyt)
			//fmt.Printf("%d\n%d\n%d\n\n", buyt.Quantity, sellShares[idx].Quantity, currentSellRecord.Quantity)
			//}
			pl += (currentSellRecord.UnitPrice - buyt.UnitPrice) * float32(pq)

			buyt.Quantity, sellShares[idx].Quantity, currentSellRecord.Quantity = buyt.Quantity-pq, currentSellRecord.Quantity-pq, currentSellRecord.Quantity-pq
			/*if _, ok := income[currentSellRecord.Market]; !ok {
				income[currentSellRecord.Market] = make([]Income, 0)
			}*/

			// if _, ok := income[currentRecordSellYear]; !ok {
			// 	income[currentRecordSellYear] = make([]Income, 0)
			// }

			//income = append(income,income[sellShares[idx].Market])
		}

		inc.PandL = pl

		//income[currentSellRecord.Market] = append(income[currentSellRecord.Market], inc)
		income.Add(currentRecordSellYear, inc)

	}
	//fmt.Printf("%v+ \n", income)
	return income
}

func GetearlierbuyShare(buyshares []Transaction, market string) *Transaction {
	mindate := time.Now()
	//var earliershare Transaction
	earlierShareIdx := 0
	for idx := range buyshares {
		bshares := buyshares[idx]
		if bshares.Market == market {
			if bshares.Date.Before(mindate) && bshares.Quantity > 0 {
				mindate = bshares.Date
				earlierShareIdx = idx
			}
		}
	}
	return &buyshares[earlierShareIdx]
}
