package ComplexType

import (
	"encoding/xml"
	. "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
)

type FreightRateRequest struct {
	XMLName xml.Name `xml:"http://www.ups.com/XMLSchema/XOLTWS/FreightRate/v1.0 FreightRateRequest"`

	Request *RequestType `xml:"Request,omitempty"`

	ShipFrom *Party `xml:"ShipFrom,omitempty"`

	ShipTo *Party `xml:"ShipTo,omitempty"`

	PaymentInformation *PaymentInformationType `xml:"PaymentInformation,omitempty"`

	Service *CodeDescriptionType `xml:"Service,omitempty"`

	HandlingUnitOne *HandlingUnitType `xml:"HandlingUnitOne,omitempty"`

	HandlingUnitTwo *HandlingUnitType `xml:"HandlingUnitTwo,omitempty"`

	Commodity []*FreightCommodityType `xml:"Commodity,omitempty"`

	ShipmentServiceOptions *FreightShipmentServiceOptionsType `xml:"ShipmentServiceOptions,omitempty"`

	PickupRequest *PickupRequestType `xml:"PickupRequest,omitempty"`

	AlternateRateOptions *CodeDescriptionType `xml:"AlternateRateOptions,omitempty"`

	GFPOptions *GFPOptionsType `xml:"GFPOptions,omitempty"`

	AccountType *AccountType `xml:"AccountType,omitempty"`

	ShipmentTotalWeight *ShipmentTotalWeightType `xml:"ShipmentTotalWeight,omitempty"`

	HandlingUnitWeight *HandlingUnitWeightType `xml:"HandlingUnitWeight,omitempty"`

	AdjustedWeightIndicator string `xml:"AdjustedWeightIndicator,omitempty"`

	TimeInTransitIndicator string `xml:"TimeInTransitIndicator,omitempty"`

	HandlingUnits []*HandlingUnitsType `xml:"HandlingUnits,omitempty"`

	AdjustedHeightIndicator string `xml:"AdjustedHeightIndicator,omitempty"`

	DensityEligibleIndicator string `xml:"DensityEligibleIndicator,omitempty"`

	QuoteNumberIndicator string `xml:"QuoteNumberIndicator,omitempty"`
}

type FreightRateResponse struct {
	XMLName xml.Name `xml:"http://www.ups.com/XMLSchema/XOLTWS/FreightRate/v1.0 FreightRateResponse"`

	Response *ResponseType `xml:"Response,omitempty"`

	CustomerServiceCenterPhone *PhoneType `xml:"CustomerServiceCenterPhone,omitempty"`

	Rate []*RateType `xml:"Rate,omitempty"`

	FreightDensityRate *FreightDensityRateType `xml:"FreightDensityRate,omitempty"`

	Commodity []*CommodityWeightType `xml:"Commodity,omitempty"`

	TotalShipmentCharge *TotalShipmentChargeType `xml:"TotalShipmentCharge,omitempty"`

	BillableShipmentWeight *WeightType `xml:"BillableShipmentWeight,omitempty"`

	DimensionalWeight *WeightType `xml:"DimensionalWeight,omitempty"`

	Service *CodeDescriptionType `xml:"Service,omitempty"`

	GuaranteedIndicator string `xml:"GuaranteedIndicator,omitempty"`

	MinimumChargeAppliedIndicator string `xml:"MinimumChargeAppliedIndicator,omitempty"`

	AlternateRatesResponse []*AlternateRatesResponseType `xml:"AlternateRatesResponse,omitempty"`

	MinimumBillableWeightAppliedIndicator string `xml:"MinimumBillableWeightAppliedIndicator,omitempty"`

	RatingSchedule *RatingScheduleType `xml:"RatingSchedule,omitempty"`

	HoldAtAirportPickupDate string `xml:"HoldAtAirportPickupDate,omitempty"`

	NextAvailablePickupDate string `xml:"NextAvailablePickupDate,omitempty"`

	TimeInTransit *FreightTimeInTransitResponseType `xml:"TimeInTransit,omitempty"`

	HandlingUnits *HandlingUnitsInfoType `xml:"HandlingUnits,omitempty"`

	QuoteNumber string `xml:"QuoteNumber,omitempty"`
}

type AccountType struct {
	Code string `xml:"Code,omitempty"`

	Description string `xml:"Description,omitempty"`
}

type ShipmentTotalWeightType struct {
	Value string `xml:"Value,omitempty"`

	UnitOfMeasurement *UnitOfMeasurementType `xml:"UnitOfMeasurement,omitempty"`
}

type HandlingUnitWeightType struct {
	Value string `xml:"Value,omitempty"`

	UnitOfMeasurement *UnitOfMeasurementType `xml:"UnitOfMeasurement,omitempty"`
}

type AlternateRatesResponseType struct {
	AlternateRateType *CodeDescriptionType `xml:"AlternateRateType,omitempty"`

	Rate []*RateType `xml:"Rate,omitempty"`

	FreightDensityRate *FreightDensityRateType `xml:"FreightDensityRate,omitempty"`

	BillableShipmentWeight *WeightType `xml:"BillableShipmentWeight,omitempty"`

	TimeInTransit *FreightTimeInTransitResponseType `xml:"TimeInTransit,omitempty"`
}

type PaymentInformationType struct {
	Payer *Party `xml:"Payer,omitempty"`

	ShipmentBillingOption *CodeDescriptionType `xml:"ShipmentBillingOption,omitempty"`
}

type HandlingUnitType struct {
	Quantity string `xml:"Quantity,omitempty"`

	Type *CodeDescriptionType `xml:"Type,omitempty"`
}

type FreightCommodityType struct {
	CommodityID string `xml:"CommodityID,omitempty"`

	Description string `xml:"Description,omitempty"`

	Weight *WeightType `xml:"Weight,omitempty"`

	AdjustedWeight *AdjustedWeightType `xml:"AdjustedWeight,omitempty"`

	Dimensions *DimensionsType `xml:"Dimensions,omitempty"`

	NumberOfPieces string `xml:"NumberOfPieces,omitempty"`

	PackagingType *CodeDescriptionType `xml:"PackagingType,omitempty"`

	DangerousGoodsIndicator string `xml:"DangerousGoodsIndicator,omitempty"`

	CommodityValue *CommodityValueType `xml:"CommodityValue,omitempty"`

	FreightClass string `xml:"FreightClass,omitempty"`

	NMFCCommodityCode string `xml:"NMFCCommodityCode,omitempty"`

	NMFCCommodity *NMFCCommodityType `xml:"NMFCCommodity,omitempty"`
}



type WeightType struct {
	Value string `xml:"Value,omitempty"`

	UnitOfMeasurement *UnitOfMeasurementType `xml:"UnitOfMeasurement,omitempty"`
}

type AdjustedWeightType struct {
	Value string `xml:"Value,omitempty"`

	UnitOfMeasurement *UnitOfMeasurementType `xml:"UnitOfMeasurement,omitempty"`
}

type CommodityValueType struct {
	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	MonetaryValue string `xml:"MonetaryValue,omitempty"`
}

type FreightShipmentServiceOptionsType struct {
	PickupOptions *FreightPickupOptionsType `xml:"PickupOptions,omitempty"`

	DeliveryOptions *FreightDeliveryOptionsType `xml:"DeliveryOptions,omitempty"`

	OverSeasLeg *OverSeasLegType `xml:"OverSeasLeg,omitempty"`

	COD *FreightCODType `xml:"COD,omitempty"`

	DangerousGoods *DangerousGoodsType `xml:"DangerousGoods,omitempty"`

	SortingAndSegregating *SortingAndSegregatingType `xml:"SortingAndSegregating,omitempty"`

	DeclaredValue *DeclaredValueType `xml:"DeclaredValue,omitempty"`

	ExcessDeclaredValue *DeclaredValueType `xml:"ExcessDeclaredValue,omitempty"`

	CustomsValue *CustomsValueType `xml:"CustomsValue,omitempty"`

	DeliveryDutiesPaidIndicator string `xml:"DeliveryDutiesPaidIndicator,omitempty"`

	DeliveryDutiesUnpaidIndicator string `xml:"DeliveryDutiesUnpaidIndicator,omitempty"`

	HandlingCharge *HandlingChargeType `xml:"HandlingCharge,omitempty"`

	CustomsClearanceIndicator string `xml:"CustomsClearanceIndicator,omitempty"`

	FreezableProtectionIndicator string `xml:"FreezableProtectionIndicator,omitempty"`

	ExtremeLengthIndicator string `xml:"ExtremeLengthIndicator,omitempty"`

	LinearFeet string `xml:"LinearFeet,omitempty"`

	AdjustedHeight *AdjustedHeightType `xml:"AdjustedHeight,omitempty"`
}

type EmailInformationType struct {
	EMailType *CodeDescriptionType `xml:"EMailType,omitempty"`

	EMail *EMailType `xml:"EMail,omitempty"`
}

type EMailType struct {
	EMailAddress []string `xml:"EMailAddress,omitempty"`

	EMailText string `xml:"EMailText,omitempty"`

	UndeliverableEMailAddress string `xml:"UndeliverableEMailAddress,omitempty"`

	Subject string `xml:"Subject,omitempty"`
}

type FreightPickupOptionsType struct {
	HolidayPickupIndicator string `xml:"HolidayPickupIndicator,omitempty"`

	InsidePickupIndicator string `xml:"InsidePickupIndicator,omitempty"`

	ResidentialPickupIndicator string `xml:"ResidentialPickupIndicator,omitempty"`

	WeekendPickupIndicator string `xml:"WeekendPickupIndicator,omitempty"`

	LiftGateRequiredIndicator string `xml:"LiftGateRequiredIndicator,omitempty"`

	HoldAtAirportForPickupIndicator string `xml:"HoldAtAirportForPickupIndicator,omitempty"`

	PickupFromDoorIndicator string `xml:"PickupFromDoorIndicator,omitempty"`

	LimitedAccessPickupIndicator string `xml:"LimitedAccessPickupIndicator,omitempty"`
}

type FreightDeliveryOptionsType struct {
	CallBeforeDeliveryIndicator string `xml:"CallBeforeDeliveryIndicator,omitempty"`

	HolidayDeliveryIndicator string `xml:"HolidayDeliveryIndicator,omitempty"`

	InsideDeliveryIndicator string `xml:"InsideDeliveryIndicator,omitempty"`

	ResidentialDeliveryIndicator string `xml:"ResidentialDeliveryIndicator,omitempty"`

	WeekendDeliveryIndicator string `xml:"WeekendDeliveryIndicator,omitempty"`

	LiftGateRequiredIndicator string `xml:"LiftGateRequiredIndicator,omitempty"`

	SaturdayDeliveryIndicator string `xml:"SaturdayDeliveryIndicator,omitempty"`

	DeliveryToDoorIndicator string `xml:"DeliveryToDoorIndicator,omitempty"`

	LimitedAccessDeliveryIndicator string `xml:"LimitedAccessDeliveryIndicator,omitempty"`
}

type OverSeasLegType struct {
	Dimensions *DimensionType `xml:"Dimensions,omitempty"`

	Value *DimensionValueType `xml:"Value,omitempty"`
}

type DimensionType struct {
	Volume string `xml:"Volume,omitempty"`

	Height string `xml:"Height,omitempty"`

	Length string `xml:"Length,omitempty"`

	Width string `xml:"Width,omitempty"`

	UnitOfMeasurement *UnitOfMeasurementType `xml:"UnitOfMeasurement,omitempty"`
}

type DimensionValueType struct {
	Cube *CubeType `xml:"Cube,omitempty"`

	CWT *CWTType `xml:"CWT,omitempty"`
}

type CubeType struct {
	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	MonetaryValue string `xml:"MonetaryValue,omitempty"`
}

type CWTType struct {
	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	MonetaryValue string `xml:"MonetaryValue,omitempty"`
}

type FreightCODType struct {
	CODValue *CODValueType `xml:"CODValue,omitempty"`

	CODPaymentMethod *CodeDescriptionType `xml:"CODPaymentMethod,omitempty"`

	CODBillingOption *CodeDescriptionType `xml:"CODBillingOption,omitempty"`

	RemitTo *RemitToType `xml:"RemitTo,omitempty"`
}

type CODValueType struct {
	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	MonetaryValue string `xml:"MonetaryValue,omitempty"`
}

type RemitToType struct {
	Name string `xml:"Name,omitempty"`

	Address *AddressType `xml:"Address,omitempty"`

	AttentionName string `xml:"AttentionName,omitempty"`
}

type DangerousGoodsType struct {
	Name string `xml:"Name,omitempty"`

	Phone *PhoneType `xml:"Phone,omitempty"`

	TransportationMode *CodeDescriptionType `xml:"TransportationMode,omitempty"`
}

type SortingAndSegregatingType struct {
	Quantity string `xml:"Quantity,omitempty"`
}

type DeclaredValueType struct {
	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	MonetaryValue string `xml:"MonetaryValue,omitempty"`
}

type CustomsValueType struct {
	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	MonetaryValue string `xml:"MonetaryValue,omitempty"`
}

type HandlingChargeType struct {
	Percentage string `xml:"Percentage,omitempty"`

	Amount *HandlingChargeAmountType `xml:"Amount,omitempty"`
}

type HandlingChargeAmountType struct {
	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	MonetaryValue string `xml:"MonetaryValue,omitempty"`
}

type PickupRequestType struct {
	PickupDate string `xml:"PickupDate,omitempty"`

	AdditionalComments string `xml:"AdditionalComments,omitempty"`
}

type RateType struct {
	Type *CodeDescriptionType `xml:"Type,omitempty"`

	SubTypeCode string `xml:"SubTypeCode,omitempty"`

	Factor *FactorType `xml:"Factor,omitempty"`
}

type CommodityWeightType struct {
	CommodityID string `xml:"CommodityID,omitempty"`

	Description string `xml:"Description,omitempty"`

	Weight *WeightType `xml:"Weight,omitempty"`

	AdjustedWeight *AdjustedWeightType `xml:"AdjustedWeight,omitempty"`
}

type FactorType struct {
	Value string `xml:"Value,omitempty"`

	UnitOfMeasurement *UnitOfMeasurementType `xml:"UnitOfMeasurement,omitempty"`
}

type AmountType struct {
	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	MonetaryValue string `xml:"MonetaryValue,omitempty"`
}

type TotalShipmentChargeType struct {
	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	MonetaryValue string `xml:"MonetaryValue,omitempty"`
}

type RatingScheduleType struct {
	Code string `xml:"Code,omitempty"`

	Description string `xml:"Description,omitempty"`
}

type FreightTimeInTransitResponseType struct {
	DaysInTransit string `xml:"DaysInTransit,omitempty"`
}

type GFPOptionsType struct {
	GPFAccesorialRateIndicator string `xml:"GPFAccesorialRateIndicator,omitempty"`

	OnCallInformation *OnCallInformationType `xml:"OnCallInformation,omitempty"`
}

type OnCallInformationType struct {
	OnCallPickupIndicator string `xml:"OnCallPickupIndicator,omitempty"`
}

type HandlingUnitsInfoType struct {
	Quantity string `xml:"Quantity,omitempty"`

	Type *CodeDescriptionType `xml:"Type,omitempty"`

	Dimensions *HandlingUnitsDimensionsType `xml:"Dimensions,omitempty"`

	AdjustedHeight *AdjustedHeightType `xml:"AdjustedHeight,omitempty"`
}
