package ComplexType

type PreauthorizedReturnInformationType struct {
	ReturnEligibilityIndicator string `xml:"ReturnEligibilityIndicator,omitempty"`

	ReturnExpirationDate string `xml:"ReturnExpirationDate,omitempty"`

	ReturnRequestURL string `xml:"ReturnRequestURL,omitempty"`

	OriginalTrackingNumber string `xml:"OriginalTrackingNumber,omitempty"`

	ReturnTrackingNumber string `xml:"ReturnTrackingNumber,omitempty"`
}
