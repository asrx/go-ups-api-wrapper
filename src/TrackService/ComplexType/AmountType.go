package ComplexType

type AmountType struct {
	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	MonetaryValue string `xml:"MonetaryValue,omitempty"`
}
