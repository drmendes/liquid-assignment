package money

import "testing"

func TestNewFromStringValid(t *testing.T) {
	dec, err := NewFromString("1234.56")
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	expected := "1234.56"
	if dec.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, dec.String())
	}
}

// Test creating Money from an invalid string
func TestNewFromStringInvalid(t *testing.T) {
	_, err := NewFromString("notanumber")
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestAddMoneys(t *testing.T) {
	dec1, _ := NewFromString("1534.00")
	dec2, _ := NewFromString("1.00")
	result := dec1.Add(dec2)
	expected := "1535.00" // Now the test should expect properly formatted output
	if result.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result.String())
	}
}

// Test the subtraction of two Money numbers
func TestSubtractMoneys(t *testing.T) {
	dec1, _ := NewFromString("1000.50")
	dec2, _ := NewFromString("200.25")
	result := dec1.Subtract(dec2)
	expected := "800.25"
	if result.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result.String())
	}
}

// Test the multiplication of two Money numbers
func TestMultiplyMoneys(t *testing.T) {
	dec1, _ := NewFromString("10.5")
	dec2, _ := NewFromString("2.0")
	result := dec1.Multiply(dec2)
	expected := "21.00"
	if result.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result.String())
	}
}

// Test the division of two Money numbers
func TestDivideMoneys(t *testing.T) {
	dec1, _ := NewFromString("21.0")
	dec2, _ := NewFromString("2.0")
	result, err := dec1.Divide(dec2, 1) // Assuming we specify the precision for division
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	expected := "10.5"
	if result.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result.String())
	}
}

// Test getting the absolute value of a Money
func TestAbsMoney(t *testing.T) {
	dec, _ := NewFromString("-123.45")
	result := dec.Abs()
	expected := "123.45"
	if result.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result.String())
	}
}
