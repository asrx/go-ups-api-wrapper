package ComplexType


type DimensionsType struct {
	UnitOfMeasurement *CodeDescriptionType `xml:"UnitOfMeasurement,omitempty"`

	Length string `xml:"Length,omitempty"`

	Width string `xml:"Width,omitempty"`

	Height string `xml:"Height,omitempty"`
}