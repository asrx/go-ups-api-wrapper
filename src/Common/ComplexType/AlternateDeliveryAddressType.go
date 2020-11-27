package ComplexType

type AlternateDeliveryAddressType struct {
	Name string `xml:"Name,omitempty"`

	Address *ADRType `xml:"Address,omitempty"`
}

type ADRType struct {
	AddressLine string `xml:"AddressLine,omitempty"`

	City string `xml:"City,omitempty"`

	StateProvinceCode string `xml:"StateProvinceCode,omitempty"`

	PostalCode string `xml:"PostalCode,omitempty"`

	CountryCode string `xml:"CountryCode,omitempty"`

	ResidentialAddressIndicator string `xml:"ResidentialAddressIndicator,omitempty"`

	POBoxIndicator string `xml:"POBoxIndicator,omitempty"`
}