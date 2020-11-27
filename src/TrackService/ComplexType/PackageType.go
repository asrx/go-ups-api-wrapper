package ComplexType

import "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"

type PackageType struct {
	TrackingNumber string `xml:"TrackingNumber,omitempty"`

	DeliveryIndicator string `xml:"DeliveryIndicator,omitempty"`

	DeliveryDate string `xml:"DeliveryDate,omitempty"`

	Redirect *RedirectType `xml:"Redirect,omitempty"`

	DeliveryDetail []*DeliveryDetailType `xml:"DeliveryDetail,omitempty"`

	PackageAddress []*PackageAddressType `xml:"PackageAddress,omitempty"`

	PackageServiceOption []*ServiceOptionType `xml:"PackageServiceOption,omitempty"`

	COD *CODType `xml:"COD,omitempty"`

	Activity []*ActivityType `xml:"Activity,omitempty"`

	Message []*MessageType `xml:"Message,omitempty"`

	PackageWeight *WeightType `xml:"PackageWeight,omitempty"`

	ReferenceNumber []*ComplexType.ReferenceNumberType `xml:"ReferenceNumber,omitempty"`

	AlternateTrackingNumber []string `xml:"AlternateTrackingNumber,omitempty"`

	AlternateTrackingInfo []*AlternateTrackingInfoType `xml:"AlternateTrackingInfo,omitempty"`

	DimensionalWeightScanIndicator string `xml:"DimensionalWeightScanIndicator,omitempty"`

	PreauthorizedReturnInformation *PreauthorizedReturnInformationType `xml:"PreauthorizedReturnInformation,omitempty"`
}
