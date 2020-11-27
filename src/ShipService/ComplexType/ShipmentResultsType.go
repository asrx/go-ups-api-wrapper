package ComplexType

import (
	. "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
)

type ShipmentResultsType struct {
	Disclaimer []*DisclaimerType `xml:"Disclaimer,omitempty"`

	ShipmentCharges *ShipmentChargesType `xml:"ShipmentCharges,omitempty"`

	NegotiatedRateCharges *NegotiatedRateChargesType `xml:"NegotiatedRateCharges,omitempty"`

	FRSShipmentData *FRSShipmentDataType `xml:"FRSShipmentData,omitempty"`

	RatingMethod string `xml:"RatingMethod,omitempty"`

	BillableWeightCalculationMethod string `xml:"BillableWeightCalculationMethod,omitempty"`

	BillingWeight *BillingWeightType `xml:"BillingWeight,omitempty"`

	ShipmentIdentificationNumber string `xml:"ShipmentIdentificationNumber,omitempty"`

	MIDualReturnShipmentKey string `xml:"MIDualReturnShipmentKey,omitempty"`

	ShipmentDigest string `xml:"ShipmentDigest,omitempty"`

	PackageResults []*PackageResultsType `xml:"PackageResults,omitempty"`

	ControlLogReceipt []*ImageType `xml:"ControlLogReceipt,omitempty"`

	Form *FormType `xml:"Form,omitempty"`

	CODTurnInPage *SCReportType `xml:"CODTurnInPage,omitempty"`

	HighValueReport *HighValueReportType `xml:"HighValueReport,omitempty"`

	LabelURL string `xml:"LabelURL,omitempty"`

	LocalLanguageLabelURL string `xml:"LocalLanguageLabelURL,omitempty"`

	ReceiptURL string `xml:"ReceiptURL,omitempty"`

	LocalLanguageReceiptURL string `xml:"LocalLanguageReceiptURL,omitempty"`

	DGPaperImage []string `xml:"DGPaperImage,omitempty"`

	MasterCartonID string `xml:"MasterCartonID,omitempty"`
}
