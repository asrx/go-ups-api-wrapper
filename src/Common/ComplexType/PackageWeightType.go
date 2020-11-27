package ComplexType

type PackageWeightType struct {
	UnitOfMeasurement *CodeDescriptionType `xml:"UnitOfMeasurement,omitempty"`

	Weight string `xml:"Weight,omitempty"`
}
