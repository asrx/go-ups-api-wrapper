package ComplexType

import . "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"

type RatedShipmentType struct {
	Disclaimer []*DisclaimerType `xml:"Disclaimer,omitempty"`

	Service *CodeDescriptionType `xml:"Service,omitempty"`

	RateChart string `xml:"RateChart,omitempty"`

	RatedShipmentAlert []*RatedShipmentInfoType `xml:"RatedShipmentAlert,omitempty"`

	BillableWeightCalculationMethod string `xml:"BillableWeightCalculationMethod,omitempty"`

	RatingMethod string `xml:"RatingMethod,omitempty"`

	BillingWeight *BillingWeightType `xml:"BillingWeight,omitempty"`

	TransportationCharges *ChargesType `xml:"TransportationCharges,omitempty"`

	BaseServiceCharge *ChargesType `xml:"BaseServiceCharge,omitempty"`

	ItemizedCharges []*ChargesType `xml:"ItemizedCharges,omitempty"`

	FRSShipmentData *FRSShipmentType `xml:"FRSShipmentData,omitempty"`

	ServiceOptionsCharges *ChargesType `xml:"ServiceOptionsCharges,omitempty"`

	TaxCharges []*TaxChargeType `xml:"TaxCharges,omitempty"`

	TotalCharges *ChargesType `xml:"TotalCharges,omitempty"`

	TotalChargesWithTaxes *ChargesType `xml:"TotalChargesWithTaxes,omitempty"`

	NegotiatedRateCharges *TotalChargeType `xml:"NegotiatedRateCharges,omitempty"`

	GuaranteedDelivery *GuaranteedDeliveryType `xml:"GuaranteedDelivery,omitempty"`

	RatedPackage []*RatedPackageType `xml:"RatedPackage,omitempty"`

	TimeInTransit *TimeInTransitResponseType `xml:"TimeInTransit,omitempty"`
}
