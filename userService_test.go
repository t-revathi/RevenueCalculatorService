package RevenueCalculatorService_test

import (
	"api-traderevenuecalculator/service"
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"time"
)

var _ = Describe("ProcessTransactions", func() {
	var transactions []service.Transaction
	var config service.Config

	BeforeEach(func() {
		transactions = make([]service.Transaction, 0)
		transactions = append(transactions, service.Transaction{
			Market:    "NEXTDC Ltd",
			Direction: "BUY",
			Cost:      -3133,
			Price:     6.25,
			Quantity:  500,
			Date:      time.Date(2018, time.November, 9, 12, 42, 31, 0, &time.Location{}),
			Activity:  "TRADE",
		})
		transactions = append(transactions, service.Transaction{
			Market:    "NEXTDC Ltd",
			Direction: "SELL",
			Cost:      -3467.54,
			Price:     6.965,
			Quantity:  499,
			Date:      time.Date(2019, time.February, 14, 12, 42, 31, 0, &time.Location{}),
			Activity:  "TRADE",
		})
		transactions = append(transactions, service.Transaction{
			Market:    "NEXTDC Ltd",
			Direction: "BUY",
			Cost:      -1997.96,
			Price:     6.18,
			Quantity:  322,
			Date:      time.Date(2019, time.April, 2, 12, 42, 31, 0, &time.Location{}),
			Activity:  "TRADE",
		})
		transactions = append(transactions, service.Transaction{
			Market:    "NEXTDC Ltd",
			Direction: "SELL",
			Cost:      -2057.63,
			Price:     6.415,
			Quantity:  322,
			Date:      time.Date(2019, time.April, 26, 12, 42, 31, 0, &time.Location{}),
			Activity:  "TRADE",
		})
		transactions = append(transactions, service.Transaction{
			Market:    "NEXTDC Ltd",
			Direction: "BUY",
			Cost:      -1288,
			Price:     6.4,
			Quantity:  200,
			Date:      time.Date(2019, time.June, 27, 12, 42, 31, 0, &time.Location{}),
			Activity:  "TRADE",
		})
		transactions = append(transactions, service.Transaction{
			Market:    "NEXTDC Ltd",
			Direction: "SELL",
			Cost:      -1342,
			Price:     6.75,
			Quantity:  200,
			Date:      time.Date(2019, time.December, 11, 12, 42, 31, 0, &time.Location{}),
			Activity:  "TRADE",
		})
		config = service.Config{true, "2021", "June", "July"}

	})
	Context("Checking the Profit/Loss", func() {
		It("Checking result after calculating profit/loss", func() {
			result := service.ProcessTransactions(transactions, config)
			Expect(len(result.Items)).To(Equal(2))
			Expect(result.Items[0].Year).To(Equal("2018-2019"))
			Expect(result.Items[1].Year).To(Equal("2019-2020"))
			Expect(math.Round(float64(result.Items[0].Items[0].PandL))).To(Equal(math.Round(340.81)))
			Expect(math.Round(float64(result.Items[0].Items[1].PandL))).To(Equal(math.Round(59.60878)))
			Expect(math.Round(float64(result.Items[1].Items[0].PandL))).To(Equal(math.Round(54.235153)))
		})
	})
})
