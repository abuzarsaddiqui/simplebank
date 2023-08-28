package util

// constant for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	PKR = "PKR"
)

// IsSupportedCurrency returns true if supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, PKR:
		return true
	}
	return false
}
