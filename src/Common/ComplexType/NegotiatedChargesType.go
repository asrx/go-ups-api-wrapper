package ComplexType


type NegotiatedChargesType struct {
	ItemizedCharges []*ChargesType `xml:"ItemizedCharges,omitempty"`
}