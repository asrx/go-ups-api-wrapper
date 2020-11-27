package ComplexType

type FRSPaymentInfoType struct {
	Type *CodeDescriptionType `xml:"Type,omitempty"`

	AccountNumber string `xml:"AccountNumber,omitempty"`

	Address *PayerAddressType `xml:"Address,omitempty"`
}


type PayerAddressType struct {
	PostalCode string `xml:"PostalCode,omitempty"`

	CountryCode string `xml:"CountryCode,omitempty"`
}