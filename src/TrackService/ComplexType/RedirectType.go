package ComplexType

type RedirectType struct {
	CompanyName string `xml:"CompanyName,omitempty"`

	LocationID string `xml:"LocationID,omitempty"`

	PickupDate string `xml:"PickupDate,omitempty"`
}
