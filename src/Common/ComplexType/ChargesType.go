package ComplexType

type ChargesType struct {
	Code string `xml:"Code,omitempty"`

	Description string `xml:"Description,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	MonetaryValue string `xml:"MonetaryValue,omitempty"`

	SubType string `xml:"SubType,omitempty"`
}
