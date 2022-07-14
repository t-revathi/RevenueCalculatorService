package service

import "time"

type DataCalculateRevenue struct {
	// Transactions []Transaction
	TransactionData []Transaction
	Config          Config
}
type Config struct {
	SkipCorporateAction bool
	FinancialYear       string
	StartFinancialMonth string
	EndFinancialMonth   string
}

type Transaction struct {
	Market    string
	Direction string
	Cost      float32
	Price     float32
	Quantity  int
	Date      time.Time
	Activity  string
	UnitPrice float32
}

type Income struct {
	Date          time.Time
	Market        string
	Quantity      int
	PandL         float32
	SellUnitPrice float32
}

type Revenue struct {
	Year  string
	Items []Income
}

type RevenueCollection struct {
	Items []*Revenue
}

func (r *RevenueCollection) Add(year string, inc Income) {
	revenue := r.Get(year)
	if revenue == nil {
		revenue = &Revenue{
			Year:  year,
			Items: make([]Income, 0),
		}
		r.Items = append(r.Items, revenue)
	}
	
	revenue.Items = append(revenue.Items, inc)
}

func (r *RevenueCollection) Get(year string) *Revenue {
	for i := 0; i < len(r.Items); i++ {
		if r.Items[i].Year == year {
			return r.Items[i]
		}
	}

	return nil
}
