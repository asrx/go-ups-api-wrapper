package ComplexType


//type PackageDeclaredValueType struct {
type CurrencyMonetaryType struct {
	Type *CodeDescriptionType `xml:"Type,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	MonetaryValue string `xml:"MonetaryValue,omitempty"`
}
