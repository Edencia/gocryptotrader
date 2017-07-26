package pair

import "strings"

// CurrencyItem is an exported string with methods to manipulate the data instead
// of using array/slice access modifiers
type CurrencyItem string

// Lower converts the CurrencyItem object c to lowercase
func (c CurrencyItem) Lower() CurrencyItem {
	return CurrencyItem(strings.ToLower(string(c)))
}

// Upper converts the CurrencyItem object c to uppercase
func (c CurrencyItem) Upper() CurrencyItem {
	return CurrencyItem(strings.ToUpper(string(c)))
}

// String converts the CurrencyItem object c to string
func (c CurrencyItem) String() string {
	return string(c)
}

// CurrencyPair holds currency pair information
type CurrencyPair struct {
	Delimiter      string       `json:"delimiter"`
	FirstCurrency  CurrencyItem `json:"first_currency"`
	SecondCurrency CurrencyItem `json:"second_currency"`
}

// GetFirstCurrency returns the first currency item
func (c CurrencyPair) GetFirstCurrency() CurrencyItem {
	return c.FirstCurrency
}

// GetSecondCurrency returns the second currency item
func (c CurrencyPair) GetSecondCurrency() CurrencyItem {
	return c.SecondCurrency
}

// Pair returns a currency pair string
func (c CurrencyPair) Pair() CurrencyItem {
	return c.FirstCurrency + CurrencyItem(c.Delimiter) + c.SecondCurrency
}

// NewCurrencyPairDelimiter splits the desired currency string at delimeter,
// the returns a CurrencyPair struct
func NewCurrencyPairDelimiter(currency, delimiter string) CurrencyPair {
	result := strings.Split(currency, delimiter)
	return CurrencyPair{
		Delimiter:      delimiter,
		FirstCurrency:  CurrencyItem(result[0]),
		SecondCurrency: CurrencyItem(result[1]),
	}
}

// NewCurrencyPair returns a CurrencyPair without a delimiter
func NewCurrencyPair(firstCurrency, secondCurrency string) CurrencyPair {
	return CurrencyPair{
		FirstCurrency:  CurrencyItem(firstCurrency),
		SecondCurrency: CurrencyItem(secondCurrency),
	}
}

// NewCurrencyPairFromString converts currency string into a new CurrencyPair
// with or without delimeter
func NewCurrencyPairFromString(currency string) CurrencyPair {
	delimiters := []string{"_", "-"}
	var delimiter string
	for _, x := range delimiters {
		if strings.Contains(currency, x) {
			delimiter = x
			return NewCurrencyPairDelimiter(currency, delimiter)
		}
	}
	return NewCurrencyPair(currency[0:3], currency[3:])
}