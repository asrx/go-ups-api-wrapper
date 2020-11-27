package ComplexType


type CommodityType struct {
	FreightClass string `xml:"FreightClass,omitempty"`

	NMFC *NMFCType `xml:"NMFC,omitempty"`
}


type NMFCType struct {
	PrimeCode string `xml:"PrimeCode,omitempty"`

	SubCode string `xml:"SubCode,omitempty"`
}