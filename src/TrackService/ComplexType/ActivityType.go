package ComplexType

type ActivityType struct {
	AlternateTrackingInfo []*AlternateTrackingInfoType `xml:"AlternateTrackingInfo,omitempty"`

	ActivityLocation *ActivityLocationType `xml:"ActivityLocation,omitempty"`

	Status *StatusType `xml:"Status,omitempty"`

	Date string `xml:"Date,omitempty"`

	Time string `xml:"Time,omitempty"`

	GMTDate string `xml:"GMTDate,omitempty"`

	GMTTime string `xml:"GMTTime,omitempty"`

	GMTOffset string `xml:"GMTOffset,omitempty"`

	DeliveryDateFromManifestIndicator string `xml:"DeliveryDateFromManifestIndicator,omitempty"`

	NextScheduleActivity *NextScheduleActivityType `xml:"NextScheduleActivity,omitempty"`

	Document []*DocumentType `xml:"Document,omitempty"`

	AdditionalAttribute []*AdditionalCodeDescriptionValueType `xml:"AdditionalAttribute,omitempty"`
}
