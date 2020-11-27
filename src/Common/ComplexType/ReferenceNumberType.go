package ComplexType

type ReferenceNumberType struct {
	BarCodeIndicator string `xml:"BarCodeIndicator,omitempty"`

	Code string `xml:"Code,omitempty"`

	Value string `xml:"Value,omitempty"`
}

