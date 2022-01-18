package util

import (
	"github.com/Template7/common/logger"
	"github.com/Template7/common/structs"
	"math"
)

func ToPico(money structs.Money) uint {
	switch money.Unit {
	case structs.UnitOne:
		return money.Amount * uint(math.Pow10(12))
	case structs.UnitCent:
		return money.Amount * uint(math.Pow10(10))
	case structs.UnitMilli:
		return money.Amount * uint(math.Pow10(9))
	case structs.UnitMicro:
		return money.Amount * uint(math.Pow10(6))
	case structs.UnitNano:
		return money.Amount * uint(math.Pow10(3))
	case structs.UnitPico:
		return money.Amount
	default:
		logger.GetLogger().Warn("invalid unit: ", money.Unit)
		return 0
	}
}
