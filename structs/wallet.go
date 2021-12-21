package structs

type Wallet struct {
	Id     string `gorm:"primaryKey;column:id;type:VARCHAR(36) NOT NULL"`
	UserId string `gorm:"index:user_id;column:userId;type:VARCHAR(36) NOT NULL"`
	//Status string
	CreatedAt int64 `gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt int64 `gorm:"column:updatedAt;autoUpdateTime:milli"`
}

type Balance struct {
	WalletId  string   `gorm:"column:walletId;type:VARCHAR(36) NOT NULL;index:walletId_currency,unique"`
	Currency  Currency `gorm:"column:currency;type:VARCHAR(36) NOT NULL;index:walletId_currency,unique"`
	Amount    uint     `gorm:"column:amount;type:bigint(20);default:0"`
	Unit      Unit     `gorm:"column:unit;type:VARCHAR(10) NOT NULL;default:pico"` // should be the min Unit
	UpdatedAt int64    `gorm:"column:updatedAt;autoUpdateTime:milli"`
}
