package ComplexType

type PackageDeclaredValueType struct {
	Type *CodeDescriptionType `xml:"Type,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	MonetaryValue string `xml:"MonetaryValue,omitempty"`
}

