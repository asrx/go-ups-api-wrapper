package ComplexType

type AddressRequestType struct {
	PostalCode string `xml:"PostalCode,omitempty"`

	CountryCode string `xml:"CountryCode,omitempty"`
}
