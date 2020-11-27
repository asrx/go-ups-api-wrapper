package ComplexType

type ShipmentRatingOptionsType struct {
	NegotiatedRatesIndicator string `xml:"NegotiatedRatesIndicator,omitempty"`

	// The account number must be enabled for GFP.
	// 必须为GFP启用帐号
	FRSShipmentIndicator string `xml:"FRSShipmentIndicator,omitempty"`

	RateChartIndicator string `xml:"RateChartIndicator,omitempty"`

	UserLevelDiscountIndicator string `xml:"UserLevelDiscountIndicator,omitempty"`
}
