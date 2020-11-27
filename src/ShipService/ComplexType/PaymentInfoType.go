package ComplexType

type PaymentInfoType struct {
	ShipmentCharge *ShipmentChargeType `xml:"ShipmentCharge,omitempty"`

	SplitDutyVATIndicator string `xml:"SplitDutyVATIndicator,omitempty"`
}

type ShipmentChargeType struct {
	Type string `xml:"Type,omitempty"`

	BillShipper *BillShipperType `xml:"BillShipper,omitempty"`

	BillReceiver *BillReceiverType `xml:"BillReceiver,omitempty"`

	BillThirdParty *BillThirdPartyChargeType `xml:"BillThirdParty,omitempty"`

	ConsigneeBilledIndicator string `xml:"ConsigneeBilledIndicator,omitempty"`
}


type BillShipperType struct {
	AccountNumber string `xml:"AccountNumber,omitempty"`

	CreditCard *CreditCardType `xml:"CreditCard,omitempty"`

	AlternatePaymentMethod string `xml:"AlternatePaymentMethod,omitempty"`
}


type BillReceiverType struct {
	AccountNumber string `xml:"AccountNumber,omitempty"`

	Address *BillReceiverAddressType `xml:"Address,omitempty"`
}


type BillThirdPartyChargeType struct {
	AccountNumber string `xml:"AccountNumber,omitempty"`

	Address *AccountAddressType `xml:"Address,omitempty"`
}


type CreditCardType struct {
	Type string `xml:"Type,omitempty"`

	Number string `xml:"Number,omitempty"`

	ExpirationDate string `xml:"ExpirationDate,omitempty"`

	SecurityCode string `xml:"SecurityCode,omitempty"`

	Address *CreditCardAddressType `xml:"Address,omitempty"`
}

type BillReceiverAddressType struct {
	PostalCode string `xml:"PostalCode,omitempty"`
}

type AccountAddressType struct {
	PostalCode string `xml:"PostalCode,omitempty"`

	CountryCode string `xml:"CountryCode,omitempty"`
}


type CreditCardAddressType struct {
	AddressLine string `xml:"AddressLine,omitempty"`

	City string `xml:"City,omitempty"`

	StateProvinceCode string `xml:"StateProvinceCode,omitempty"`

	PostalCode string `xml:"PostalCode,omitempty"`

	CountryCode string `xml:"CountryCode,omitempty"`
}
