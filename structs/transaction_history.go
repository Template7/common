package structs

const (
	UnitOne   Unit = "unit"  // 1
	UnitCent  Unit = "cent"  // 10^(-2)
	UnitMilli Unit = "milli" // 10^(-3)
	UnitMicro Unit = "micro" // 10^(-6)
	UnitNano  Unit = "nano"  // 10^(-9)
	UnitPico  Unit = "pico"  // 10^(-12)

	CurrencyNTD Currency = "NTD"
	CurrencyUSD Currency = "USD"
)

type Unit string
type Currency string

type TransactionHistory struct {
	TransactionId string   `json:"transaction_id" bson:"transaction_id"`
	From          string   `json:"from" bson:"from"` // from wallet ID
	To            string   `json:"to" bson:"to"`     // to wallet ID
	Amount        int      `json:"amount" bson:"amount"`
	Unit          Unit     `json:"unit" bson:"unit"`
	Currency      Currency `json:"currency" bson:"currency"`
	Note          string   `json:"note" bson:"note"`
}
