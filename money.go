package money

import (
	"errors"
	"math/big"
	"strings"
)

type Money struct {
	value    *big.Int
	exp      int32
	currency *Currency
}

// NewMoney creates Money from a value and an exponent
func NewMoney(value int64, exp int32) *Money {
	return &Money{
		value: big.NewInt(value),
		exp:   exp,
	}
}

// NewFromString creates Money from a string
func NewFromString(str string) (*Money, error) {
	parts := strings.Split(str, ".")
	intPart := parts[0]
	exp := int32(0)
	if len(parts) > 1 {
		exp = -int32(len(parts[1]))
		intPart += parts[1]
	}

	val, ok := new(big.Int).SetString(intPart, 10)
	if !ok {
		return nil, errors.New("invalid number format")
	}

	return &Money{
		value: val,
		exp:   exp,
	}, nil
}
