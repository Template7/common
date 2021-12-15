package structs

type Wallet struct {
	Id      string `json:"id" bson:"id"`
	UserId  string `json:"user_id" bson:"user_id"`
	Balance []struct {
		Amount   int      `json:"amount" bson:"amount"`
		Unit     unit     `json:"unit" bson:"unit"` // should be the min unit
		Currency currency `json:"currency" bson:"currency"`
	} `json:"balance" bson:"balance"`
}
