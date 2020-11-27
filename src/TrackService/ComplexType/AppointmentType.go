package ComplexType

type AppointmentType struct {
	Made *DateTimeType `xml:"Made,omitempty"`

	Requested *DateTimeType `xml:"Requested,omitempty"`

	BeginTime string `xml:"BeginTime,omitempty"`

	EndTime string `xml:"EndTime,omitempty"`
}
