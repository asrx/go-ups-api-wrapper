package ComplexType

import . "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"

type RatedPackageType struct {
	TransportationCharges *ChargesType `xml:"TransportationCharges,omitempty"`

	BaseServiceCharge *ChargesType `xml:"BaseServiceCharge,omitempty"`

	ServiceOptionsCharges *ChargesType `xml:"ServiceOptionsCharges,omitempty"`

	TotalCharges *ChargesType `xml:"TotalCharges,omitempty"`

	Weight string `xml:"Weight,omitempty"`

	BillingWeight *BillingWeightType `xml:"BillingWeight,omitempty"`

	Accessorial []*AccessorialType `xml:"Accessorial,omitempty"`

	ItemizedCharges []*ChargesType `xml:"ItemizedCharges,omitempty"`

	NegotiatedCharges *NegotiatedChargesType `xml:"NegotiatedCharges,omitempty"`
}

type ReturnContractServicesType struct {
	Code string `xml:"Code,omitempty"`

	Description string `xml:"Description,omitempty"`
}

type TimeInTransitResponseType struct {
	PickupDate string `xml:"PickupDate,omitempty"`

	DocumentsOnlyIndicator string `xml:"DocumentsOnlyIndicator,omitempty"`

	PackageBillType string `xml:"PackageBillType,omitempty"`

	AutoDutyCode string `xml:"AutoDutyCode,omitempty"`

	Disclaimer string `xml:"Disclaimer,omitempty"`

	ServiceSummary *ServiceSummaryType `xml:"ServiceSummary,omitempty"`
}

type ServiceSummaryType struct {
	Service *CodeDescriptionType `xml:"Service,omitempty"`

	EstimatedArrival *EstimatedArrivalType `xml:"EstimatedArrival,omitempty"`

	GuaranteedIndicator string `xml:"GuaranteedIndicator,omitempty"`

	Disclaimer string `xml:"Disclaimer,omitempty"`

	SaturdayDelivery string `xml:"SaturdayDelivery,omitempty"`

	SaturdayDeliveryDisclaimer string `xml:"SaturdayDeliveryDisclaimer,omitempty"`
}

type EstimatedArrivalType struct {
	Arrival *PickupType `xml:"Arrival,omitempty"`

	BusinessDaysInTransit string `xml:"BusinessDaysInTransit,omitempty"`

	Pickup *PickupType `xml:"Pickup,omitempty"`

	DayOfWeek string `xml:"DayOfWeek,omitempty"`

	CustomerCenterCutoff string `xml:"CustomerCenterCutoff,omitempty"`

	DelayCount string `xml:"DelayCount,omitempty"`

	HolidayCount string `xml:"HolidayCount,omitempty"`

	RestDays string `xml:"RestDays,omitempty"`

	TotalTransitDays string `xml:"TotalTransitDays,omitempty"`
}

type TotalChargeType struct {
	ItemizedCharges []*ChargesType `xml:"ItemizedCharges,omitempty"`

	TaxCharges []*TaxChargeType `xml:"TaxCharges,omitempty"`

	TotalCharge *ChargesType `xml:"TotalCharge,omitempty"`

	TotalChargesWithTaxes *ChargesType `xml:"TotalChargesWithTaxes,omitempty"`
}

type RatedShipmentInfoType struct {
	Code string `xml:"Code,omitempty"`

	Description string `xml:"Description,omitempty"`
}

type TransportationChargesType struct {
	GrossCharge *ChargesType `xml:"GrossCharge,omitempty"`

	DiscountAmount *ChargesType `xml:"DiscountAmount,omitempty"`

	DiscountPercentage string `xml:"DiscountPercentage,omitempty"`

	NetCharge *ChargesType `xml:"NetCharge,omitempty"`
}

type FRSShipmentType struct {
	TransportationCharges *TransportationChargesType `xml:"TransportationCharges,omitempty"`

	FreightDensityRate *FreightDensityRateType `xml:"FreightDensityRate,omitempty"`

	HandlingUnits []*HandlingUnitsResponseType `xml:"HandlingUnits,omitempty"`
}

type PackageType struct {
	PackagingType *CodeDescriptionType `xml:"PackagingType,omitempty"`

	Dimensions *DimensionsType `xml:"Dimensions,omitempty"`

	DimWeight *PackageWeightType `xml:"DimWeight,omitempty"`

	PackageWeight *PackageWeightType `xml:"PackageWeight,omitempty"`

	Commodity *CommodityType `xml:"Commodity,omitempty"`

	LargePackageIndicator string `xml:"LargePackageIndicator,omitempty"`

	PackageServiceOptions *PackageServiceOptionsType `xml:"PackageServiceOptions,omitempty"`

	AdditionalHandlingIndicator string `xml:"AdditionalHandlingIndicator,omitempty"`
}

type PackageServiceOptionsType struct {
	DeliveryConfirmation *DeliveryConfirmationType `xml:"DeliveryConfirmation,omitempty"`

	AccessPointCOD *PackageServiceOptionsAccessPointCODType `xml:"AccessPointCOD,omitempty"`

	COD *CODType `xml:"COD,omitempty"`

	DeclaredValue *PackageDeclaredValueType `xml:"DeclaredValue,omitempty"`

	ShipperDeclaredValue *CurrencyMonetaryType `xml:"ShipperDeclaredValue,omitempty"`

	ProactiveIndicator string `xml:"ProactiveIndicator,omitempty"`

	RefrigerationIndicator string `xml:"RefrigerationIndicator,omitempty"`

	Insurance *InsuranceType `xml:"Insurance,omitempty"`

	VerbalConfirmationIndicator string `xml:"VerbalConfirmationIndicator,omitempty"`

	UPSPremiumCareIndicator string `xml:"UPSPremiumCareIndicator,omitempty"`

	HazMat *HazMatType `xml:"HazMat,omitempty"`

	DryIce *DryIceType `xml:"DryIce,omitempty"`
}

type HazMatType struct {
	PackageIdentifier string `xml:"PackageIdentifier,omitempty"`

	QValue string `xml:"QValue,omitempty"`

	OverPackedIndicator string `xml:"OverPackedIndicator,omitempty"`

	AllPackedInOneIndicator string `xml:"AllPackedInOneIndicator,omitempty"`

	HazMatChemicalRecord *HazMatChemicalRecordType `xml:"HazMatChemicalRecord,omitempty"`
}

type HazMatChemicalRecordType struct {
	ChemicalRecordIdentifier string `xml:"ChemicalRecordIdentifier,omitempty"`

	ClassDivisionNumber string `xml:"ClassDivisionNumber,omitempty"`

	IDNumber string `xml:"IDNumber,omitempty"`

	TransportationMode string `xml:"TransportationMode,omitempty"`

	RegulationSet string `xml:"RegulationSet,omitempty"`

	EmergencyPhone string `xml:"EmergencyPhone,omitempty"`

	EmergencyContact string `xml:"EmergencyContact,omitempty"`

	ReportableQuantity string `xml:"ReportableQuantity,omitempty"`

	SubRiskClass string `xml:"SubRiskClass,omitempty"`

	PackagingGroupType string `xml:"PackagingGroupType,omitempty"`

	Quantity string `xml:"Quantity,omitempty"`

	UOM string `xml:"UOM,omitempty"`

	PackagingInstructionCode string `xml:"PackagingInstructionCode,omitempty"`

	ProperShippingName string `xml:"ProperShippingName,omitempty"`

	TechnicalName string `xml:"TechnicalName,omitempty"`

	AdditionalDescription string `xml:"AdditionalDescription,omitempty"`

	PackagingType string `xml:"PackagingType,omitempty"`

	HazardLabelRequired string `xml:"HazardLabelRequired,omitempty"`

	PackagingTypeQuantity string `xml:"PackagingTypeQuantity,omitempty"`

	CommodityRegulatedLevelCode string `xml:"CommodityRegulatedLevelCode,omitempty"`

	TransportCategory string `xml:"TransportCategory,omitempty"`

	TunnelRestrictionCode string `xml:"TunnelRestrictionCode,omitempty"`
}

type DryIceType struct {
	RegulationSet string `xml:"RegulationSet,omitempty"`

	DryIceWeight *DryIceWeightType `xml:"DryIceWeight,omitempty"`

	MedicalUseIndicator string `xml:"MedicalUseIndicator,omitempty"`

	AuditRequired string `xml:"AuditRequired,omitempty"`
}

type InsuranceType struct {
	BasicFlexibleParcelIndicator *CurrencyMonetaryType `xml:"BasicFlexibleParcelIndicator,omitempty"`

	ExtendedFlexibleParcelIndicator *CurrencyMonetaryType `xml:"ExtendedFlexibleParcelIndicator,omitempty"`

	TimeInTransitFlexibleParcelIndicator *CurrencyMonetaryType `xml:"TimeInTransitFlexibleParcelIndicator,omitempty"`
}

type UOMCodeDescriptionType struct {
	Code string `xml:"Code,omitempty"`

	Description string `xml:"Description,omitempty"`
}

type TimeInTransitRequestType struct {
	PackageBillType string `xml:"PackageBillType,omitempty"`

	Pickup *PickupType `xml:"Pickup,omitempty"`

	ReturnContractServices []*ReturnContractServicesType `xml:"ReturnContractServices,omitempty"`
}

type PickupType struct {
	Date string `xml:"Date,omitempty"`

	Time string `xml:"Time,omitempty"`
}

type ShipmentWeightType struct {
	UnitOfMeasurement *CodeDescriptionType `xml:"UnitOfMeasurement,omitempty"`

	Weight string `xml:"Weight,omitempty"`
}

type PaymentDetailsType struct {
	ShipmentCharge *ShipmentChargeType `xml:"ShipmentCharge,omitempty"`

	SplitDutyVATIndicator string `xml:"SplitDutyVATIndicator,omitempty"`
}

type ShipmentChargeType struct {
	Type string `xml:"Type,omitempty"`

	BillShipper *BillShipperChargeType `xml:"BillShipper,omitempty"`

	BillReceiver *BillReceiverChargeType `xml:"BillReceiver,omitempty"`

	BillThirdParty *BillThirdPartyChargeType `xml:"BillThirdParty,omitempty"`

	ConsigneeBilledIndicator string `xml:"ConsigneeBilledIndicator,omitempty"`
}

type BillShipperChargeType struct {
	AccountNumber string `xml:"AccountNumber,omitempty"`
}

type BillReceiverChargeType struct {
	AccountNumber string `xml:"AccountNumber,omitempty"`

	Address *BillReceiverAddressType `xml:"Address,omitempty"`
}

type BillThirdPartyChargeType struct {
	AccountNumber string `xml:"AccountNumber,omitempty"`

	Address *AddressType `xml:"Address,omitempty"`
}

type BillReceiverAddressType struct {
	PostalCode string `xml:"PostalCode,omitempty"`
}

type ShipmentServiceOptionsType struct {
	SaturdayPickupIndicator string `xml:"SaturdayPickupIndicator,omitempty"`

	SaturdayDeliveryIndicator string `xml:"SaturdayDeliveryIndicator,omitempty"`

	AccessPointCOD *PackageServiceOptionsAccessPointCODType `xml:"AccessPointCOD,omitempty"`

	DeliverToAddresseeOnlyIndicator string `xml:"DeliverToAddresseeOnlyIndicator,omitempty"`

	DirectDeliveryOnlyIndicator string `xml:"DirectDeliveryOnlyIndicator,omitempty"`

	COD *CODType `xml:"COD,omitempty"`

	DeliveryConfirmation *DeliveryConfirmationType `xml:"DeliveryConfirmation,omitempty"`

	ReturnOfDocumentIndicator string `xml:"ReturnOfDocumentIndicator,omitempty"`

	UPScarbonneutralIndicator string `xml:"UPScarbonneutralIndicator,omitempty"`

	CertificateOfOriginIndicator string `xml:"CertificateOfOriginIndicator,omitempty"`

	PickupOptions *PickupOptionsType `xml:"PickupOptions,omitempty"`

	DeliveryOptions *DeliveryOptionsType `xml:"DeliveryOptions,omitempty"`

	RestrictedArticles *RestrictedArticlesType `xml:"RestrictedArticles,omitempty"`

	ShipperExportDeclarationIndicator string `xml:"ShipperExportDeclarationIndicator,omitempty"`

	CommercialInvoiceRemovalIndicator string `xml:"CommercialInvoiceRemovalIndicator,omitempty"`

	ImportControl *ImportControlType `xml:"ImportControl,omitempty"`

	ReturnService *ReturnServiceType `xml:"ReturnService,omitempty"`

	SDLShipmentIndicator string `xml:"SDLShipmentIndicator,omitempty"`

	EPRAIndicator string `xml:"EPRAIndicator,omitempty"`

	InsideDelivery string `xml:"InsideDelivery,omitempty"`

	ItemDisposalIndicator string `xml:"ItemDisposalIndicator,omitempty"`
}

type ImportControlType struct {
	Code string `xml:"Code,omitempty"`

	Description string `xml:"Description,omitempty"`
}

type PickupOptionsType struct {
	LiftGateAtPickupIndicator string `xml:"LiftGateAtPickupIndicator,omitempty"`

	HoldForPickupIndicator string `xml:"HoldForPickupIndicator,omitempty"`
}

type DeliveryOptionsType struct {
	LiftGateAtDeliveryIndicator string `xml:"LiftGateAtDeliveryIndicator,omitempty"`

	DropOffAtUPSFacilityIndicator string `xml:"DropOffAtUPSFacilityIndicator,omitempty"`
}

type GuaranteedDeliveryType struct {
	BusinessDaysInTransit string `xml:"BusinessDaysInTransit,omitempty"`

	DeliveryByTime string `xml:"DeliveryByTime,omitempty"`
}
