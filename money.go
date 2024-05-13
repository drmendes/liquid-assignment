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

// Add adds two Money
func (d *Money) Add(other *Money) *Money {
	if other == nil {
		return d
	}

	dVal := new(big.Int).Set(d.value)
	otherVal := new(big.Int).Set(other.value)
	var finalExp int32

	// Determine the smallest exponent to use as the base
	if d.exp < other.exp {
		finalExp = d.exp
	} else {
		finalExp = other.exp
	}

	// Scale dVal and otherVal to have the same exponent
	scaleDiffD := finalExp - d.exp
	scaleDiffOther := finalExp - other.exp

	if scaleDiffD != 0 {
		scaleFactor := big.NewInt(10).Exp(big.NewInt(10), big.NewInt(int64(-scaleDiffD)), nil)
		dVal.Mul(dVal, scaleFactor)
	}
	if scaleDiffOther != 0 {
		scaleFactor := big.NewInt(10).Exp(big.NewInt(10), big.NewInt(int64(-scaleDiffOther)), nil)
		otherVal.Mul(otherVal, scaleFactor)
	}

	// Add the values now that they are aligned
	newVal := new(big.Int).Add(dVal, otherVal)

	return &Money{
		value:    newVal,
		exp:      finalExp,
		currency: d.currency,
	}
}

// String converts Money into string representation
func (d *Money) String() string {
	// Convert the entire number to a string
	absoluteValue := d.value.Abs(d.value).String()

	if d.exp >= 0 {
		// If exp is non-negative, just append zeros
		return absoluteValue + strings.Repeat("0", int(d.exp))
	}

	// Handle negative exponent (Money places)
	exp := int(-d.exp)
	if exp >= len(absoluteValue) {
		return "0." + strings.Repeat("0", exp-len(absoluteValue)) + absoluteValue
	}

	// Insert Money point at the correct position
	MoneyPointIndex := len(absoluteValue) - exp
	result := absoluteValue[:MoneyPointIndex] + "." + absoluteValue[MoneyPointIndex:]

	// Ensure no trailing Money points
	result = strings.TrimRight(result, ".")
	return result
}

// Multiply multiplies two Moneys
func (d *Money) Multiply(other *Money) *Money {
	if other == nil {
		return nil
	}
	newValue := new(big.Int).Mul(d.value, other.value)
	newExp := d.exp + other.exp
	return &Money{
		value: newValue,
		exp:   newExp,
	}
}

// Subtract subtracts another Money from this Money, taking into account different exponents.
func (d *Money) Subtract(other *Money) *Money {
	negOther := &Money{value: new(big.Int).Neg(other.value), exp: other.exp}
	return d.Add(negOther)
}

// Divide divides this Money by another, adjusting precision as necessary.
func (d *Money) Divide(other *Money, precision int32) (*Money, error) {
	if other.value.Sign() == 0 {
		return nil, errors.New("division by zero")
	}
	factor := big.NewInt(10)
	factor.Exp(factor, big.NewInt(int64(precision)), nil)
	numerator := new(big.Int).Mul(d.value, factor)
	result := numerator.Div(numerator, other.value)
	return &Money{value: result, exp: d.exp - other.exp - precision}, nil
}

// Abs returns the absolute value of this Money.
func (d *Money) Abs() *Money {
	absValue := new(big.Int).Abs(d.value)
	return &Money{value: absValue, exp: d.exp}
}

// ConvertCurrency converts this Money to another currency using a conversion rate and the target currency code.
func (d *Money) ConvertCurrency(rate *Money, targetCurrencyCode string) (*Money, error) {
	if d == nil || rate == nil {
		return nil, errors.New("nil Money value")
	}
	if d.currency.Code == rate.currency.Code {
		convertedValue := new(big.Int).Mul(d.value, rate.value)
		// Adjust exponent based on the rate's exponent
		finalExp := d.exp + rate.exp
		return &Money{
			value:    convertedValue,
			exp:      finalExp,
			currency: &Currency{Code: targetCurrencyCode},
		}, nil
	}
	return nil, errors.New("conversion rate currency mismatch with current currency")
}

func (d *Money) Round(places int32) *Money {
	deltaExp := int64(d.exp - places) // Difference in exponent from current to target.
	factor := big.NewInt(10)
	factor = factor.Exp(factor, big.NewInt(abs(deltaExp)), nil) // 10^|deltaExp| for scaling.

	rounded := new(big.Int).Set(d.value)
	if deltaExp < 0 {
		// Positive deltaExp means we need to multiply to scale down.
		rounded.Div(rounded, factor) // Divide to scale down.
	} else {
		// Negative deltaExp means we need to multiply to scale up.
		rounded.Mul(rounded, factor) // Multiply to scale up.
	}

	// No need to adjust back since we're setting the exponent directly.
	return &Money{value: rounded, exp: places}
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

// Compare compares this Money to another. Returns -1, 0, or 1.
func (d *Money) Compare(other *Money) int {
	// Normalize to the same exponent
	if d.exp > other.exp {
		temp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(d.exp-other.exp)), nil)
		otherValue := new(big.Int).Mul(other.value, temp)
		return d.value.Cmp(otherValue)
	} else if d.exp < other.exp {
		temp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(other.exp-d.exp)), nil)
		dValue := new(big.Int).Mul(d.value, temp)
		return dValue.Cmp(other.value)
	}
	return d.value.Cmp(other.value)
}

// Equals checks if two Moneys are exactly equal.
func (d *Money) Equals(other *Money) bool {
	return d.Compare(other) == 0
}
