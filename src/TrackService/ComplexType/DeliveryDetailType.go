package ComplexType

type DeliveryDetailType struct {
	Type *CommonCodeDescriptionType `xml:"Type,omitempty"`

	Date string `xml:"Date,omitempty"`

	Time string `xml:"Time,omitempty"`
}
