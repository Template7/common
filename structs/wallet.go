package structs

import "time"

type Wallet struct {
	Id     string `gorm:"type:VARCHAR(36) NOT NULL primary_key;"`
	UserId string `gorm:"type:VARCHAR(36) NOT NULL index;"`
	//Status string
	CreatedAt time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

type Balance struct {
	WalletId string   `gorm:"type:VARCHAR(36) NOT NULL index"`
	Currency currency `gorm:"type:VARCHAR(36) NOT NULL"`
	Amount   int      `gorm:"type:bigint(20);"`
	Unit     unit     `gorm:"type:VARCHAR(10) NOT NULL"` // should be the min unit
	UpdatedAt time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
