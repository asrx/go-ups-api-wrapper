package ComplexType

import (
	. "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
)

type ShipmentType struct {
	Description string `xml:"Description,omitempty"`

	ReturnService *ReturnServiceType `xml:"ReturnService,omitempty"`

	DocumentsOnlyIndicator string `xml:"DocumentsOnlyIndicator,omitempty"`

	Shipper *Party `xml:"Shipper,omitempty"`

	ShipTo *Party `xml:"ShipTo,omitempty"`

	AlternateDeliveryAddress *AlternateDeliveryAddressType `xml:"AlternateDeliveryAddress,omitempty"`

	ShipFrom *Party `xml:"ShipFrom,omitempty"`

	PaymentInformation *PaymentInfoType `xml:"PaymentInformation,omitempty"`

	FRSPaymentInformation *FRSPaymentInfoType `xml:"FRSPaymentInformation,omitempty"`

	FreightShipmentInformation *FreightShipmentInformationType `xml:"FreightShipmentInformation,omitempty"`

	GoodsNotInFreeCirculationIndicator string `xml:"GoodsNotInFreeCirculationIndicator,omitempty"`

	ShipmentRatingOptions *ShipmentRatingOptionsType `xml:"ShipmentRatingOptions,omitempty"`

	MovementReferenceNumber string `xml:"MovementReferenceNumber,omitempty"`

	ReferenceNumber *ReferenceNumberType `xml:"ReferenceNumber,omitempty"`

	Service *CodeDescriptionType `xml:"Service,omitempty"`

	InvoiceLineTotal *CurrencyMonetaryType `xml:"InvoiceLineTotal,omitempty"`

	NumOfPiecesInShipment string `xml:"NumOfPiecesInShipment,omitempty"`

	USPSEndorsement string `xml:"USPSEndorsement,omitempty"`

	MILabelCN22Indicator string `xml:"MILabelCN22Indicator,omitempty"`

	SubClassification string `xml:"SubClassification,omitempty"`

	CostCenter string `xml:"CostCenter,omitempty"`

	CostCenterBarcodeIndicator string `xml:"CostCenterBarcodeIndicator,omitempty"`

	PackageID string `xml:"PackageID,omitempty"`

	PackageIDBarcodeIndicator string `xml:"PackageIDBarcodeIndicator,omitempty"`

	IrregularIndicator string `xml:"IrregularIndicator,omitempty"`

	ShipmentIndicationType []*IndicationType `xml:"ShipmentIndicationType,omitempty"`

	MIDualReturnShipmentKey string `xml:"MIDualReturnShipmentKey,omitempty"`

	MIDualReturnShipmentIndicator string `xml:"MIDualReturnShipmentIndicator,omitempty"`

	RatingMethodRequestedIndicator string `xml:"RatingMethodRequestedIndicator,omitempty"`

	TaxInformationIndicator string `xml:"TaxInformationIndicator,omitempty"`

	PromotionalDiscountInformation *PromotionalDiscountInformationType `xml:"PromotionalDiscountInformation,omitempty"`

	DGSignatoryInfo *DGSignatoryInfoType `xml:"DGSignatoryInfo,omitempty"`

	MasterCartonID string `xml:"MasterCartonID,omitempty"`

	MasterCartonIndicator string `xml:"MasterCartonIndicator,omitempty"`

	ShipmentServiceOptions *ShipmentServiceOptionsType `xml:"ShipmentServiceOptions,omitempty"`

	Package []*PackageType `xml:"Package,omitempty"`
}
