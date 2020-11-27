package ComplexType

import "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"

type VolumeType struct {
	UnitOfMeasurement *ComplexType.UnitOfMeasurementType `xml:"UnitOfMeasurement,omitempty"`

	Value string `xml:"Value,omitempty"`
}
