package ComplexType

import "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"

type WeightType struct {
	UnitOfMeasurement *ComplexType.UnitOfMeasurementType `xml:"UnitOfMeasurement,omitempty"`

	Weight string `xml:"Weight,omitempty"`
}
