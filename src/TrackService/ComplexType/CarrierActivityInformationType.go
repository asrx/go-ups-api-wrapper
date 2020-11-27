package ComplexType

type CarrierActivityInformationType struct {
	CarrierId string `xml:"CarrierId,omitempty"`

	Description string `xml:"Description,omitempty"`

	Status string `xml:"Status,omitempty"`

	Arrival *DateTimeType `xml:"Arrival,omitempty"`

	Departure *DateTimeType `xml:"Departure,omitempty"`

	OriginPort string `xml:"OriginPort,omitempty"`

	DestinationPort string `xml:"DestinationPort,omitempty"`
}
