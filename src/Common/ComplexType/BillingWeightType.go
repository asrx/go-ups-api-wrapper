package ComplexType


type BillingWeightType struct {
	UnitOfMeasurement *CodeDescriptionType `xml:"UnitOfMeasurement,omitempty"`

	Weight string `xml:"Weight,omitempty"`
}
