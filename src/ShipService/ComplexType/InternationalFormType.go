package ComplexType

import . "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"

type InternationalFormType struct {
	FormType string `xml:"FormType,omitempty"`

	UserCreatedForm *UserCreatedFormType `xml:"UserCreatedForm,omitempty"`

	CN22Form *CN22FormType `xml:"CN22Form,omitempty"`

	UPSPremiumCareForm *UPSPremiumCareFormType `xml:"UPSPremiumCareForm,omitempty"`

	AdditionalDocumentIndicator string `xml:"AdditionalDocumentIndicator,omitempty"`

	FormGroupIdName string `xml:"FormGroupIdName,omitempty"`

	SEDFilingOption string `xml:"SEDFilingOption,omitempty"`

	EEIFilingOption *EEIFilingOptionType `xml:"EEIFilingOption,omitempty"`

	Contacts *ContactType `xml:"Contacts,omitempty"`

	Product *ProductType `xml:"Product,omitempty"`

	InvoiceNumber string `xml:"InvoiceNumber,omitempty"`

	InvoiceDate string `xml:"InvoiceDate,omitempty"`

	PurchaseOrderNumber string `xml:"PurchaseOrderNumber,omitempty"`

	TermsOfShipment string `xml:"TermsOfShipment,omitempty"`

	ReasonForExport string `xml:"ReasonForExport,omitempty"`

	Comments string `xml:"Comments,omitempty"`

	DeclarationStatement string `xml:"DeclarationStatement,omitempty"`

	Discount *IFChargesType `xml:"Discount,omitempty"`

	FreightCharges *IFChargesType `xml:"FreightCharges,omitempty"`

	InsuranceCharges *IFChargesType `xml:"InsuranceCharges,omitempty"`

	OtherCharges *OtherChargesType `xml:"OtherCharges,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	BlanketPeriod *BlanketPeriodType `xml:"BlanketPeriod,omitempty"`

	ExportDate string `xml:"ExportDate,omitempty"`

	ExportingCarrier string `xml:"ExportingCarrier,omitempty"`

	CarrierID string `xml:"CarrierID,omitempty"`

	InBondCode string `xml:"InBondCode,omitempty"`

	EntryNumber string `xml:"EntryNumber,omitempty"`

	PointOfOrigin string `xml:"PointOfOrigin,omitempty"`

	PointOfOriginType string `xml:"PointOfOriginType,omitempty"`

	ModeOfTransport string `xml:"ModeOfTransport,omitempty"`

	PortOfExport string `xml:"PortOfExport,omitempty"`

	PortOfUnloading string `xml:"PortOfUnloading,omitempty"`

	LoadingPier string `xml:"LoadingPier,omitempty"`

	PartiesToTransaction string `xml:"PartiesToTransaction,omitempty"`

	RoutedExportTransactionIndicator string `xml:"RoutedExportTransactionIndicator,omitempty"`

	ContainerizedIndicator string `xml:"ContainerizedIndicator,omitempty"`

	License *LicenseType `xml:"License,omitempty"`

	ECCNNumber string `xml:"ECCNNumber,omitempty"`

	OverridePaperlessIndicator string `xml:"OverridePaperlessIndicator,omitempty"`

	ShipperMemo string `xml:"ShipperMemo,omitempty"`

	MultiCurrencyInvoiceLineTotal string `xml:"MultiCurrencyInvoiceLineTotal,omitempty"`

	HazardousMaterialsIndicator string `xml:"HazardousMaterialsIndicator,omitempty"`
}

type UserCreatedFormType struct {
	DocumentID []string `xml:"DocumentID,omitempty"`
}

type CN22FormType struct {
	LabelSize string `xml:"LabelSize,omitempty"`

	PrintsPerPage string `xml:"PrintsPerPage,omitempty"`

	LabelPrintType string `xml:"LabelPrintType,omitempty"`

	CN22Type string `xml:"CN22Type,omitempty"`

	CN22OtherDescription string `xml:"CN22OtherDescription,omitempty"`

	FoldHereText string `xml:"FoldHereText,omitempty"`

	CN22Content *CN22ContentType `xml:"CN22Content,omitempty"`
}

type UPSPremiumCareFormType struct {
	ShipmentDate string `xml:"ShipmentDate,omitempty"`

	PageSize string `xml:"PageSize,omitempty"`

	PrintType string `xml:"PrintType,omitempty"`

	NumOfCopies string `xml:"NumOfCopies,omitempty"`

	LanguageForUPSPremiumCare *LanguageForUPSPremiumCareType `xml:"LanguageForUPSPremiumCare,omitempty"`
}

type EEIFilingOptionType struct {
	Code string `xml:"Code,omitempty"`

	EMailAddress string `xml:"EMailAddress,omitempty"`

	Description string `xml:"Description,omitempty"`

	UPSFiled *UPSFiledType `xml:"UPSFiled,omitempty"`

	ShipperFiled *ShipperFiledType `xml:"ShipperFiled,omitempty"`
}

type ContactType struct {
	ForwardAgent *ForwardAgentType `xml:"ForwardAgent,omitempty"`

	UltimateConsignee *UltimateConsigneeType `xml:"UltimateConsignee,omitempty"`

	IntermediateConsignee *IntermediateConsigneeType `xml:"IntermediateConsignee,omitempty"`

	Producer *ProducerType `xml:"Producer,omitempty"`

	SoldTo *SoldToType `xml:"SoldTo,omitempty"`
}

type ProductType struct {
	Description string `xml:"Description,omitempty"`

	Unit *UnitType `xml:"Unit,omitempty"`

	CommodityCode string `xml:"CommodityCode,omitempty"`

	PartNumber string `xml:"PartNumber,omitempty"`

	OriginCountryCode string `xml:"OriginCountryCode,omitempty"`

	JointProductionIndicator string `xml:"JointProductionIndicator,omitempty"`

	NetCostCode string `xml:"NetCostCode,omitempty"`

	NetCostDateRange *NetCostDateType `xml:"NetCostDateRange,omitempty"`

	PreferenceCriteria string `xml:"PreferenceCriteria,omitempty"`

	ProducerInfo string `xml:"ProducerInfo,omitempty"`

	MarksAndNumbers string `xml:"MarksAndNumbers,omitempty"`

	NumberOfPackagesPerCommodity string `xml:"NumberOfPackagesPerCommodity,omitempty"`

	ProductWeight *ProductWeightType `xml:"ProductWeight,omitempty"`

	VehicleID string `xml:"VehicleID,omitempty"`

	ScheduleB *ScheduleBType `xml:"ScheduleB,omitempty"`

	ExportType string `xml:"ExportType,omitempty"`

	SEDTotalValue string `xml:"SEDTotalValue,omitempty"`

	ExcludeFromForm *ExcludeFromFormType `xml:"ExcludeFromForm,omitempty"`

	ProductCurrencyCode string `xml:"ProductCurrencyCode,omitempty"`

	PackingListInfo *PackingListInfoType `xml:"PackingListInfo,omitempty"`

	EEIInformation *EEIInformationType `xml:"EEIInformation,omitempty"`
}

type IFChargesType struct {
	MonetaryValue string `xml:"MonetaryValue,omitempty"`
}

type OtherChargesType struct {
	MonetaryValue string `xml:"MonetaryValue,omitempty"`

	Description string `xml:"Description,omitempty"`
}

type BlanketPeriodType struct {
	BeginDate string `xml:"BeginDate,omitempty"`

	EndDate string `xml:"EndDate,omitempty"`
}

type LicenseType struct {
	Number string `xml:"Number,omitempty"`

	Date string `xml:"Date,omitempty"`

	ExceptionCode string `xml:"ExceptionCode,omitempty"`
}

type CN22ContentType struct {
	CN22ContentQuantity string `xml:"CN22ContentQuantity,omitempty"`

	CN22ContentDescription string `xml:"CN22ContentDescription,omitempty"`

	CN22ContentWeight *ProductWeightType `xml:"CN22ContentWeight,omitempty"`

	CN22ContentTotalValue string `xml:"CN22ContentTotalValue,omitempty"`

	CN22ContentCurrencyCode string `xml:"CN22ContentCurrencyCode,omitempty"`

	CN22ContentCountryOfOrigin string `xml:"CN22ContentCountryOfOrigin,omitempty"`

	CN22ContentTariffNumber string `xml:"CN22ContentTariffNumber,omitempty"`
}

type LanguageForUPSPremiumCareType struct {
	Language string `xml:"Language,omitempty"`
}

type UPSFiledType struct {
	POA *CodeDescriptionType `xml:"POA,omitempty"`
}

type ShipperFiledType struct {
	Code string `xml:"Code,omitempty"`

	Description string `xml:"Description,omitempty"`

	PreDepartureITNNumber string `xml:"PreDepartureITNNumber,omitempty"`

	ExemptionLegend string `xml:"ExemptionLegend,omitempty"`

	EEIShipmentReferenceNumber string `xml:"EEIShipmentReferenceNumber,omitempty"`
}

type ForwardAgentType struct {
	CompanyName string `xml:"CompanyName,omitempty"`

	TaxIdentificationNumber string `xml:"TaxIdentificationNumber,omitempty"`

	Address *AddressType `xml:"Address,omitempty"`
}

type UltimateConsigneeType struct {
	CompanyName string `xml:"CompanyName,omitempty"`

	Address *AddressType `xml:"Address,omitempty"`

	UltimateConsigneeType *CodeDescriptionType `xml:"UltimateConsigneeType,omitempty"`
}


type IntermediateConsigneeType struct {
	CompanyName string `xml:"CompanyName,omitempty"`

	Address *AddressType `xml:"Address,omitempty"`
}

type ProducerType struct {
	Option string `xml:"Option,omitempty"`

	CompanyName string `xml:"CompanyName,omitempty"`

	TaxIdentificationNumber string `xml:"TaxIdentificationNumber,omitempty"`

	Address *AddressType `xml:"Address,omitempty"`

	AttentionName string `xml:"AttentionName,omitempty"`

	Phone *PhoneType `xml:"Phone,omitempty"`

	EMailAddress string `xml:"EMailAddress,omitempty"`
}

type SoldToType struct {
	Name string `xml:"Name,omitempty"`

	AttentionName string `xml:"AttentionName,omitempty"`

	TaxIdentificationNumber string `xml:"TaxIdentificationNumber,omitempty"`

	Phone *PhoneType `xml:"Phone,omitempty"`

	Option string `xml:"Option,omitempty"`

	Address *AddressType `xml:"Address,omitempty"`

	EMailAddress string `xml:"EMailAddress,omitempty"`
}

type UnitType struct {
	Number string `xml:"Number,omitempty"`

	UnitOfMeasurement *CodeDescriptionType `xml:"UnitOfMeasurement,omitempty"`

	Value string `xml:"Value,omitempty"`
}

type NetCostDateType struct {
	BeginDate string `xml:"BeginDate,omitempty"`

	EndDate string `xml:"EndDate,omitempty"`
}

type ProductWeightType struct {
	UnitOfMeasurement *CodeDescriptionType `xml:"UnitOfMeasurement,omitempty"`

	Weight string `xml:"Weight,omitempty"`
}

type ScheduleBType struct {
	Number string `xml:"Number,omitempty"`

	Quantity string `xml:"Quantity,omitempty"`

	UnitOfMeasurement *CodeDescriptionType `xml:"UnitOfMeasurement,omitempty"`
}

type ExcludeFromFormType struct {
	FormType []string `xml:"FormType,omitempty"`
}

type PackingListInfoType struct {
	PackageAssociated []*PackageAssociatedType `xml:"PackageAssociated,omitempty"`
}
type PackageAssociatedType struct {
	PackageNumber string `xml:"PackageNumber,omitempty"`

	ProductAmount string `xml:"ProductAmount,omitempty"`
}

type EEIInformationType struct {
	ExportInformation string `xml:"ExportInformation,omitempty"`

	License *EEILicenseType `xml:"License,omitempty"`

	DDTCInformation *DDTCInformationType `xml:"DDTCInformation,omitempty"`
}
type DDTCInformationType struct {
	ITARExemptionNumber string `xml:"ITARExemptionNumber,omitempty"`

	USMLCategoryCode string `xml:"USMLCategoryCode,omitempty"`

	EligiblePartyIndicator string `xml:"EligiblePartyIndicator,omitempty"`

	RegistrationNumber string `xml:"RegistrationNumber,omitempty"`

	Quantity string `xml:"Quantity,omitempty"`

	UnitOfMeasurement *UnitOfMeasurementType `xml:"UnitOfMeasurement,omitempty"`

	SignificantMilitaryEquipmentIndicator string `xml:"SignificantMilitaryEquipmentIndicator,omitempty"`

	ACMNumber string `xml:"ACMNumber,omitempty"`
}

type EEILicenseType struct {
	Number string `xml:"Number,omitempty"`

	Code string `xml:"Code,omitempty"`

	LicenseLineValue string `xml:"LicenseLineValue,omitempty"`

	ECCNNumber string `xml:"ECCNNumber,omitempty"`
}

