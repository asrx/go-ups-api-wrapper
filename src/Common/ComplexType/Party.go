package ComplexType

type Party struct {
	Name string `xml:"Name,omitempty"`

	AttentionName string `xml:"AttentionName,omitempty"`

	ShipperNumber string `xml:"ShipperNumber,omitempty"`

	Address *AddressType `xml:"Address,omitempty"`


	CompanyDisplayableName string `xml:"CompanyDisplayableName,omitempty"`

	TaxIdentificationNumber string `xml:"TaxIdentificationNumber,omitempty"`

	TaxIDType *TaxIDCodeDescType `xml:"TaxIDType,omitempty"`

	Phone *PhoneType `xml:"Phone,omitempty"`

	EMailAddress string `xml:"EMailAddress,omitempty"`

	FaxNumber string `xml:"FaxNumber,omitempty"`

	LocationID string `xml:"LocationID,omitempty"`
}

type TaxIDCodeDescType struct {
	Code string `xml:"Code,omitempty"`

	Description string `xml:"Description,omitempty"`
}


type PhoneType struct {
	Number string `xml:"Number,omitempty"`

	Extension string `xml:"Extension,omitempty"`
}