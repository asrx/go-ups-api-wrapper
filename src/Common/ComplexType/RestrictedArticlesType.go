package ComplexType


type RestrictedArticlesType struct {
	DiagnosticSpecimensIndicator string `xml:"DiagnosticSpecimensIndicator,omitempty"`

	AlcoholicBeveragesIndicator string `xml:"AlcoholicBeveragesIndicator,omitempty"`

	PerishablesIndicator string `xml:"PerishablesIndicator,omitempty"`

	PlantsIndicator string `xml:"PlantsIndicator,omitempty"`

	SeedsIndicator string `xml:"SeedsIndicator,omitempty"`

	SpecialExceptionsIndicator string `xml:"SpecialExceptionsIndicator,omitempty"`

	TobaccoIndicator string `xml:"TobaccoIndicator,omitempty"`
}