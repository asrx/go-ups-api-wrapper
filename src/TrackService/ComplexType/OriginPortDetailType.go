package ComplexType

type OriginPortDetailType struct {
	OriginPort string `xml:"OriginPort,omitempty"`

	EstimatedDeparture *DateTimeType `xml:"EstimatedDeparture,omitempty"`
}
