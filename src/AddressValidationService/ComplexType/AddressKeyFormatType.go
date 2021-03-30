package ComplexType


type AddressKeyFormatType struct {
	ConsigneeName string `xml:"ConsigneeName,omitempty"`

	AttentionName string `xml:"AttentionName,omitempty"`

	AddressLine []string `xml:"AddressLine,omitempty"`

	Region string `xml:"Region,omitempty"`

	PoliticalDivision2 string `xml:"PoliticalDivision2,omitempty"`

	PoliticalDivision1 string `xml:"PoliticalDivision1,omitempty"`

	PostcodePrimaryLow string `xml:"PostcodePrimaryLow,omitempty"`

	PostcodeExtendedLow string `xml:"PostcodeExtendedLow,omitempty"`

	Urbanization string `xml:"Urbanization,omitempty"`

	CountryCode string `xml:"CountryCode,omitempty"`
}