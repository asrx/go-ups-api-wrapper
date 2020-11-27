package ComplexType

type AddressType struct {
	AddressLine []string `xml:"AddressLine,omitempty"`

	City string `xml:"City,omitempty"`

	StateProvinceCode string `xml:"StateProvinceCode,omitempty"`

	//Town string `xml:"Town,omitempty"`

	PostalCode string `xml:"PostalCode,omitempty"`

	CountryCode string `xml:"CountryCode,omitempty"`

	ResidentialAddressIndicator string `xml:"ResidentialAddressIndicator,omitempty"`
}
