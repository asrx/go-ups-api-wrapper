package ComplexType


type CODType struct {
	CODFundsCode string `xml:"CODFundsCode,omitempty"`

	CODAmount *CurrencyMonetaryType `xml:"CODAmount,omitempty"`
}
