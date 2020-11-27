package ComplexType

type DestinationPortDetailType struct {
	DestinationPort string `xml:"DestinationPort,omitempty"`

	EstimatedArrival *DateTimeType `xml:"EstimatedArrival,omitempty"`
}
