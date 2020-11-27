package ComplexType

import (
	. "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	"encoding/xml"
)

type LabelRecoveryRequest struct {
	XMLName xml.Name `xml:"http://www.ups.com/XMLSchema/XOLTWS/LBRecovery/v1.0 LabelRecoveryRequest"`

	Request *RequestType `xml:"Request,omitempty"`

	LabelSpecification *LabelSpecificationType `xml:"LabelSpecification,omitempty"`

	Translate *TranslateType `xml:"Translate,omitempty"`

	LabelDelivery *LabelDeliveryType `xml:"LabelDelivery,omitempty"`

	TrackingNumber string `xml:"TrackingNumber,omitempty"`

	MailInnovationsTrackingNumber string `xml:"MailInnovationsTrackingNumber,omitempty"`

	ReferenceValues *ReferenceValuesType `xml:"ReferenceValues,omitempty"`

	UPSPremiumCareForm *LRUPSPremiumCareFormType `xml:"UPSPremiumCareForm,omitempty"`
}

type LabelRecoveryResponse struct {
	XMLName xml.Name `xml:"http://www.ups.com/XMLSchema/XOLTWS/LBRecovery/v1.0 LabelRecoveryResponse"`

	Response *ResponseType `xml:"Response,omitempty"`

	ShipmentIdentificationNumber string `xml:"ShipmentIdentificationNumber,omitempty"`

	CODTurnInPage *ResponseImageType `xml:"CODTurnInPage,omitempty"`

	Form *FormType `xml:"Form,omitempty"`

	HighValueReport *HighValueReportType `xml:"HighValueReport,omitempty"`

	LabelResults *LabelResultsType `xml:"LabelResults,omitempty"`

	TrackingCandidate *TrackingCandidateType `xml:"TrackingCandidate,omitempty"`
}

type ReferenceValuesType struct {
	ReferenceNumber *ReferenceNumberType `xml:"ReferenceNumber,omitempty"`

	ShipperNumber string `xml:"ShipperNumber,omitempty"`
}

type TranslateType struct {
	LanguageCode string `xml:"LanguageCode,omitempty"`

	DialectCode string `xml:"DialectCode,omitempty"`

	Code string `xml:"Code,omitempty"`
}

type LabelResultsType struct {
	TrackingNumber string `xml:"TrackingNumber,omitempty"`

	LabelImage *LabelImageType `xml:"LabelImage,omitempty"`

	Receipt *ReceiptType `xml:"Receipt,omitempty"`

	Form *FormType `xml:"Form,omitempty"`

	MailInnovationsTrackingNumber string `xml:"MailInnovationsTrackingNumber,omitempty"`

	MailInnovationsLabelImage *LabelImageType `xml:"MailInnovationsLabelImage,omitempty"`
}

type LabelImageType struct {
	LabelImageFormat *LabelImageFormatType `xml:"LabelImageFormat,omitempty"`

	GraphicImage string `xml:"GraphicImage,omitempty"`

	HTMLImage string `xml:"HTMLImage,omitempty"`

	PDF417 string `xml:"PDF417,omitempty"`

	InternationalSignatureGraphicImage string `xml:"InternationalSignatureGraphicImage,omitempty"`

	URL string `xml:"URL,omitempty"`
}

type TrackingCandidateType struct {
	TrackingNumber string `xml:"TrackingNumber,omitempty"`

	DestinationPostalCode string `xml:"DestinationPostalCode,omitempty"`

	DestinationCountryCode string `xml:"DestinationCountryCode,omitempty"`

	PickupDateRange *PickupDateRangeType `xml:"PickupDateRange,omitempty"`
}

type ResponseImageType struct {
	Image *ResponseImageDetailType `xml:"Image,omitempty"`
}

type ResponseImageDetailType struct {
	ImageFormat *ImageFormatType `xml:"ImageFormat,omitempty"`

	GraphicImage string `xml:"GraphicImage,omitempty"`
}

type LRUPSPremiumCareFormType struct {
	PageSize string `xml:"PageSize,omitempty"`

	PrintType string `xml:"PrintType,omitempty"`
}
