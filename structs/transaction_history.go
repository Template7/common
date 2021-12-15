package structs

const (
	UnitCent  unit = "cent"  // 10^(-2)
	UnitMilli unit = "milli" // 10^(-3)
	UnitMicro unit = "micro" // 10^(-6)
	UnitNano  unit = "nano"  // 10^(-9)
	UnitPico  unit = "pico"  // 10^(-12)

	CurrencyNTD currency = "NTD"
	CurrencyUSD currency = "USD"
)

type unit string
type currency string

type TransactionHistory struct {
	TransactionId string   `json:"transaction_id" bson:"transaction_id"`
	From          string   `json:"from" bson:"from"` // from wallet ID
	To            string   `json:"to" bson:"to"`     // to wallet ID
	Amount        int      `json:"amount" bson:"amount"`
	Unit          unit     `json:"unit" bson:"unit"`
	Currency      currency `json:"currency" bson:"currency"`
	Note          string   `json:"note" bson:"note"`
}
