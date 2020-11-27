package ComplexType

type ServiceOptionType struct {
	Type *CommonCodeDescriptionType `xml:"Type,omitempty"`

	Value string `xml:"Value,omitempty"`
}
