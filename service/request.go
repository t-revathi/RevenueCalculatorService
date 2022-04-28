package service

import "time"

type DataCalculateRevenue struct {
	// Transactions []Transaction
	TransactionData   []Transaction
	Config Config
}
type Config struct{
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