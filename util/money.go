package util

import (
	"github.com/Template7/common/structs"
	"math"
)

func ToPico(money structs.Money) uint {
	switch money.Unit {
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
	}
}
