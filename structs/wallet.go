package structs

type Wallet struct {
	Id     string `gorm:"primaryKey;column:id;type:VARCHAR(36) NOT NULL"`
	UserId string `gorm:"uniqueIndex:user_id;column:userId;type:VARCHAR(36) NOT NULL"`
	//Status string
	CreatedAt int64 `gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt int64 `gorm:"column:updatedAt;autoUpdateTime:milli"`
}

func (w Wallet) TableName() string {
	return "wallet"
}

type Balance struct {
	WalletId  string `gorm:"column:walletId;type:VARCHAR(36) NOT NULL;uniqueIndex:walletId_currency,unique"`
	Money     `gorm:"embedded"`
	UpdatedAt int64 `gorm:"column:updatedAt;autoUpdateTime:milli"`
}

func (b Balance) TableName() string {
	return "balance"
}

type WalletData struct {
	Id      string  `json:"id"`
	UserId  string  `json:"user_id"`
	Balance []Money `json:"balance" validate:"required,dive"`
}

type Money struct {
	Currency Currency `json:"currency" bson:"currency" gorm:"column:currency;type:VARCHAR(36) NOT NULL;uniqueIndex:walletId_currency,unique" validate:"oneof=NTD USD"`
	Amount   uint     `json:"amount" bson:"amount" gorm:"column:amount;type:bigint(20);default:0"`
	Unit     Unit     `json:"unit" bson:"unit" gorm:"column:unit;type:VARCHAR(10) NOT NULL;default:pico" validate:"oneof=cent milli micro nano pico"` // should be the min Unit in db
}
