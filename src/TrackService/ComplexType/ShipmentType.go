package ComplexType

type ShipmentType struct {
	InquiryNumber *CodeDescriptionValueType `xml:"InquiryNumber,omitempty"`

	ShipmentType *RefShipmentType `xml:"ShipmentType,omitempty"`

	CandidateBookmark string `xml:"CandidateBookmark,omitempty"`

	ShipperNumber string `xml:"ShipperNumber,omitempty"`

	ShipmentAddress []*ShipmentAddressType `xml:"ShipmentAddress,omitempty"`

	ShipmentWeight *WeightType `xml:"ShipmentWeight,omitempty"`

	Service *CommonCodeDescriptionType `xml:"Service,omitempty"`

	ReferenceNumber []*ShipmentReferenceNumberType `xml:"ReferenceNumber,omitempty"`

	CurrentStatus *CommonCodeDescriptionType `xml:"CurrentStatus,omitempty"`

	PickupDate string `xml:"PickupDate,omitempty"`

	ServiceCenter []*ServiceCenterType `xml:"ServiceCenter,omitempty"`

	DeliveryDateUnavailable *DeliveryDateUnavailableType `xml:"DeliveryDateUnavailable,omitempty"`

	DeliveryDetail []*DeliveryDetailType `xml:"DeliveryDetail,omitempty"`

	Volume *VolumeType `xml:"Volume,omitempty"`

	BillToName string `xml:"BillToName,omitempty"`

	NumberOfPackagingUnit []*NumberOfPackagingUnitType `xml:"NumberOfPackagingUnit,omitempty"`

	COD *CODType `xml:"COD,omitempty"`

	SignedForByName string `xml:"SignedForByName,omitempty"`

	Activity []*ShipmentActivityType `xml:"Activity,omitempty"`

	OriginPortDetail *OriginPortDetailType `xml:"OriginPortDetail,omitempty"`

	DestinationPortDetail *DestinationPortDetailType `xml:"DestinationPortDetail,omitempty"`

	DescriptionOfGoods string `xml:"DescriptionOfGoods,omitempty"`

	CargoReady *DateTimeType `xml:"CargoReady,omitempty"`

	Manifest *DateTimeType `xml:"Manifest,omitempty"`

	CarrierActivityInformation []*CarrierActivityInformationType `xml:"CarrierActivityInformation,omitempty"`

	Document []*DocumentType `xml:"Document,omitempty"`

	FileNumber string `xml:"FileNumber,omitempty"`

	Appointment *AppointmentType `xml:"Appointment,omitempty"`

	Package []*PackageType `xml:"Package,omitempty"`

	AdditionalAttribute []*AdditionalCodeDescriptionValueType `xml:"AdditionalAttribute,omitempty"`
}
