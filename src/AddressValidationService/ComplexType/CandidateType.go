package ComplexType



type CandidateType struct {
	AddressClassification *AddressClassificationType `xml:"AddressClassification,omitempty"`

	AddressKeyFormat *AddressKeyFormatType `xml:"AddressKeyFormat,omitempty"`
}
