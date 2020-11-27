package ComplexType

import . "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"

type PaymentType struct {
	Code string `xml:"Code,omitempty"`

	Description string `xml:"Description,omitempty"`
}

type PreAlertNotificationType struct {
	EMailMessage *PreAlertEMailMessageType `xml:"EMailMessage,omitempty"`

	VoiceMessage *PreAlertVoiceMessageType `xml:"VoiceMessage,omitempty"`

	TextMessage *PreAlertTextMessageType `xml:"TextMessage,omitempty"`

	Locale *LocaleType `xml:"Locale,omitempty"`
}

type PreAlertEMailMessageType struct {
	EMailAddress string `xml:"EMailAddress,omitempty"`

	UndeliverableEMailAddress string `xml:"UndeliverableEMailAddress,omitempty"`
}

type LocaleType struct {
	Language string `xml:"Language,omitempty"`

	Dialect string `xml:"Dialect,omitempty"`
}

type PreAlertVoiceMessageType struct {
	PhoneNumber string `xml:"PhoneNumber,omitempty"`
}

type PreAlertTextMessageType struct {
	PhoneNumber string `xml:"PhoneNumber,omitempty"`
}

type ContactInfoType struct {
	Name string `xml:"Name,omitempty"`

	Phone *PhoneType `xml:"Phone,omitempty"`
}

type NotificationType struct {
	NotificationCode string `xml:"NotificationCode,omitempty"`

	EMail *EmailDetailsType `xml:"EMail,omitempty"`

	VoiceMessage *ShipmentServiceOptionsNotificationVoiceMessageType `xml:"VoiceMessage,omitempty"`

	TextMessage *ShipmentServiceOptionsNotificationTextMessageType `xml:"TextMessage,omitempty"`

	Locale *LocaleType `xml:"Locale,omitempty"`
}

type PackageType struct {
	Description string `xml:"Description,omitempty"`

	PackagingType *CodeDescriptionType `xml:"Packaging,omitempty"`

	Dimensions *DimensionsType `xml:"Dimensions,omitempty"`

	DimWeight *PackageWeightType `xml:"DimWeight,omitempty"`

	PackageWeight *PackageWeightType `xml:"PackageWeight,omitempty"`

	LargePackageIndicator string `xml:"LargePackageIndicator,omitempty"`

	ReferenceNumber *ReferenceNumberType `xml:"ReferenceNumber,omitempty"`

	AdditionalHandlingIndicator string `xml:"AdditionalHandlingIndicator,omitempty"`

	PackageServiceOptions *PackageServiceOptionsType `xml:"PackageServiceOptions,omitempty"`

	Commodity *CommodityType `xml:"Commodity,omitempty"`

	HazMatPackageInformation *HazMatPackageInformationType `xml:"HazMatPackageInformation,omitempty"`
}


type PackageServiceOptionsType struct {
	DeliveryConfirmation *DeliveryConfirmationType `xml:"DeliveryConfirmation,omitempty"`

	AccessPointCOD *PackageServiceOptionsAccessPointCODType `xml:"AccessPointCOD,omitempty"`

	COD *CODType `xml:"COD,omitempty"`

	DeclaredValue *PackageDeclaredValueType `xml:"DeclaredValue,omitempty"`

	VerbalConfirmation *VerbalConfirmationType `xml:"VerbalConfirmation,omitempty"`

	ShipperReleaseIndicator string `xml:"ShipperReleaseIndicator,omitempty"`

	Notification *PSONotificationType `xml:"Notification,omitempty"`

	HazMat *HazMatType `xml:"HazMat,omitempty"`

	DryIce *DryIceType `xml:"DryIce,omitempty"`

	UPSPremiumCareIndicator string `xml:"UPSPremiumCareIndicator,omitempty"`

	ProactiveIndicator string `xml:"ProactiveIndicator,omitempty"`

	PackageIdentifier string `xml:"PackageIdentifier,omitempty"`

	ClinicalTrialsID string `xml:"ClinicalTrialsID,omitempty"`

	RefrigerationIndicator string `xml:"RefrigerationIndicator,omitempty"`
}

type LabelMethodType struct {
	Code string `xml:"Code,omitempty"`

	Description string `xml:"Description,omitempty"`
}

type VerbalConfirmationType struct {
	ContactInfo *ContactInfoType `xml:"ContactInfo,omitempty"`
}



type PSONotificationType struct {
	NotificationCode string `xml:"NotificationCode,omitempty"`

	EMail *EmailDetailsType `xml:"EMail,omitempty"`
}


type ShipmentChargesType struct {
	RateChart string `xml:"RateChart,omitempty"`

	BaseServiceCharge *ChargesType `xml:"BaseServiceCharge,omitempty"`

	TransportationCharges *ChargesType `xml:"TransportationCharges,omitempty"`

	ItemizedCharges []*ChargesType `xml:"ItemizedCharges,omitempty"`

	ServiceOptionsCharges *ChargesType `xml:"ServiceOptionsCharges,omitempty"`

	TaxCharges []*TaxChargeType `xml:"TaxCharges,omitempty"`

	TotalCharges *ChargesType `xml:"TotalCharges,omitempty"`

	TotalChargesWithTaxes *ChargesType `xml:"TotalChargesWithTaxes,omitempty"`
}

type NegotiatedRateChargesType struct {
	ItemizedCharges []*ChargesType `xml:"ItemizedCharges,omitempty"`

	TaxCharges []*TaxChargeType `xml:"TaxCharges,omitempty"`

	TotalCharge *ChargesType `xml:"TotalCharge,omitempty"`

	TotalChargesWithTaxes *ChargesType `xml:"TotalChargesWithTaxes,omitempty"`
}

type FRSShipmentDataType struct {
	TransportationCharges *TransportationChargeType `xml:"TransportationCharges,omitempty"`

	FreightDensityRate *FreightDensityRateType `xml:"FreightDensityRate,omitempty"`

	HandlingUnits []*HandlingUnitsResponseType `xml:"HandlingUnits,omitempty"`
}

type TransportationChargeType struct {
	GrossCharge *ChargesType `xml:"GrossCharge,omitempty"`

	DiscountAmount *ChargesType `xml:"DiscountAmount,omitempty"`

	DiscountPercentage string `xml:"DiscountPercentage,omitempty"`

	NetCharge *ChargesType `xml:"NetCharge,omitempty"`
}

type PackageResultsType struct {
	TrackingNumber string `xml:"TrackingNumber,omitempty"`

	BaseServiceCharge *ChargesType `xml:"BaseServiceCharge,omitempty"`

	ServiceOptionsCharges *ChargesType `xml:"ServiceOptionsCharges,omitempty"`

	ShippingLabel *LabelType `xml:"ShippingLabel,omitempty"`

	ShippingReceipt *ReceiptType `xml:"ShippingReceipt,omitempty"`

	USPSPICNumber string `xml:"USPSPICNumber,omitempty"`

	CN22Number string `xml:"CN22Number,omitempty"`

	Accessorial []*AccessorialType `xml:"Accessorial,omitempty"`

	Form *FormType `xml:"Form,omitempty"`

	ItemizedCharges []*ChargesType `xml:"ItemizedCharges,omitempty"`

	NegotiatedCharges *NegotiatedChargesType `xml:"NegotiatedCharges,omitempty"`
}


type LabelType struct {
	*ImageType

	GraphicImagePart []string `xml:"GraphicImagePart,omitempty"`

	InternationalSignatureGraphicImage string `xml:"InternationalSignatureGraphicImage,omitempty"`

	HTMLImage string `xml:"HTMLImage,omitempty"`

	PDF417 string `xml:"PDF417,omitempty"`
}

type SCReportType struct {
	Image *ImageType `xml:"Image,omitempty"`
}

type HazMatPackageInformationType struct {
	AllPackedInOneIndicator string `xml:"AllPackedInOneIndicator,omitempty"`

	OverPackedIndicator string `xml:"OverPackedIndicator,omitempty"`

	QValue string `xml:"QValue,omitempty"`

	OuterPackagingType string `xml:"OuterPackagingType,omitempty"`
}

type HazMatType struct {
	PackagingTypeQuantity string `xml:"PackagingTypeQuantity,omitempty"`

	RecordIdentifier1 string `xml:"RecordIdentifier1,omitempty"`

	RecordIdentifier2 string `xml:"RecordIdentifier2,omitempty"`

	RecordIdentifier3 string `xml:"RecordIdentifier3,omitempty"`

	SubRiskClass string `xml:"SubRiskClass,omitempty"`

	ADRItemNumber string `xml:"aDRItemNumber,omitempty"`

	ADRPackingGroupLetter string `xml:"aDRPackingGroupLetter,omitempty"`

	TechnicalName string `xml:"TechnicalName,omitempty"`

	HazardLabelRequired string `xml:"HazardLabelRequired,omitempty"`

	ClassDivisionNumber string `xml:"ClassDivisionNumber,omitempty"`

	ReferenceNumber string `xml:"ReferenceNumber,omitempty"`

	Quantity string `xml:"Quantity,omitempty"`

	UOM string `xml:"UOM,omitempty"`

	PackagingType string `xml:"PackagingType,omitempty"`

	IDNumber string `xml:"IDNumber,omitempty"`

	ProperShippingName string `xml:"ProperShippingName,omitempty"`

	AdditionalDescription string `xml:"AdditionalDescription,omitempty"`

	PackagingGroupType string `xml:"PackagingGroupType,omitempty"`

	PackagingInstructionCode string `xml:"PackagingInstructionCode,omitempty"`

	EmergencyPhone string `xml:"EmergencyPhone,omitempty"`

	EmergencyContact string `xml:"EmergencyContact,omitempty"`

	ReportableQuantity string `xml:"ReportableQuantity,omitempty"`

	RegulationSet string `xml:"RegulationSet,omitempty"`

	TransportationMode string `xml:"TransportationMode,omitempty"`

	CommodityRegulatedLevelCode string `xml:"CommodityRegulatedLevelCode,omitempty"`

	TransportCategory string `xml:"TransportCategory,omitempty"`

	TunnelRestrictionCode string `xml:"TunnelRestrictionCode,omitempty"`

	ChemicalRecordIdentifier string `xml:"ChemicalRecordIdentifier,omitempty"`

	LocalTechnicalName string `xml:"LocalTechnicalName,omitempty"`

	LocalProperShippingName string `xml:"LocalProperShippingName,omitempty"`
}

type DryIceType struct {
	RegulationSet string `xml:"RegulationSet,omitempty"`

	DryIceWeight *DryIceWeightType `xml:"DryIceWeight,omitempty"`

	MedicalUseIndicator string `xml:"MedicalUseIndicator,omitempty"`
}



type ShipmentServiceOptionsNotificationVoiceMessageType struct {
	PhoneNumber string `xml:"PhoneNumber,omitempty"`
}

type ShipmentServiceOptionsNotificationTextMessageType struct {
	PhoneNumber string `xml:"PhoneNumber,omitempty"`
}

type ADLAddressType struct {
	AddressLine string `xml:"AddressLine,omitempty"`

	City string `xml:"City,omitempty"`

	StateProvinceCode string `xml:"StateProvinceCode,omitempty"`

	PostalCode string `xml:"PostalCode,omitempty"`

	CountryCode string `xml:"CountryCode,omitempty"`
}

type DGSignatoryInfoType struct {
	Name string `xml:"Name,omitempty"`

	Title string `xml:"Title,omitempty"`

	Place string `xml:"Place,omitempty"`

	Date string `xml:"Date,omitempty"`

	ShipperDeclaration string `xml:"ShipperDeclaration,omitempty"`

	UploadOnlyIndicator string `xml:"UploadOnlyIndicator,omitempty"`
}
