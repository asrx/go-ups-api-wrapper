package ComplexType

import . "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"

type ShipmentType struct {
	OriginRecordTransactionTimestamp string `xml:"OriginRecordTransactionTimestamp,omitempty"`

	Shipper *Party `xml:"Shipper,omitempty"`

	ShipTo *Party `xml:"ShipTo,omitempty"`

	ShipFrom *Party `xml:"ShipFrom,omitempty"`

	AlternateDeliveryAddress *AlternateDeliveryAddressType `xml:"AlternateDeliveryAddress,omitempty"`

	ShipmentIndicationType []*IndicationType `xml:"ShipmentIndicationType,omitempty"`

	PaymentDetails *PaymentDetailsType `xml:"PaymentDetails,omitempty"`

	FRSPaymentInformation *FRSPaymentInfoType `xml:"FRSPaymentInformation,omitempty"`

	FreightShipmentInformation *FreightShipmentInformationType `xml:"FreightShipmentInformation,omitempty"`

	GoodsNotInFreeCirculationIndicator string `xml:"GoodsNotInFreeCirculationIndicator,omitempty"`

	Service *CodeDescriptionType `xml:"Service,omitempty"`

	NumOfPieces string `xml:"NumOfPieces,omitempty"`

	ShipmentTotalWeight *ShipmentWeightType `xml:"ShipmentTotalWeight,omitempty"`

	DocumentsOnlyIndicator string `xml:"DocumentsOnlyIndicator,omitempty"`

	Package []*PackageType `xml:"Package,omitempty"`

	ShipmentServiceOptions *ShipmentServiceOptionsType `xml:"ShipmentServiceOptions,omitempty"`

	ShipmentRatingOptions *ShipmentRatingOptionsType `xml:"ShipmentRatingOptions,omitempty"`

	InvoiceLineTotal *CurrencyMonetaryType `xml:"InvoiceLineTotal,omitempty"`

	RatingMethodRequestedIndicator string `xml:"RatingMethodRequestedIndicator,omitempty"`

	TaxInformationIndicator string `xml:"TaxInformationIndicator,omitempty"`

	PromotionalDiscountInformation *PromotionalDiscountInformationType `xml:"PromotionalDiscountInformation,omitempty"`

	DeliveryTimeInformation *TimeInTransitRequestType `xml:"DeliveryTimeInformation,omitempty"`

	MasterCartonIndicator string `xml:"MasterCartonIndicator,omitempty"`

	WWEShipmentIndicator string `xml:"WWEShipmentIndicator,omitempty"`
}
