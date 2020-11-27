package ComplexType

import "encoding/xml"

type Certification struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	UPSSecurity *UPSSecurity
}

type UPSSecurity struct {
	XMLName xml.Name `xml:"http://www.ups.com/XMLSchema/XOLTWS/UPSS/v1.0 UPSSecurity"`

	UsernameToken *UsernameToken `xml:"UsernameToken,omitempty"`

	ServiceAccessToken *ServiceAccessToken `xml:"ServiceAccessToken,omitempty"`
}
type UsernameToken struct {
	Username string `xml:"Username,omitempty"`

	Password string `xml:"Password,omitempty"`
}

type ServiceAccessToken struct {
	AccessLicenseNumber string `xml:"AccessLicenseNumber,omitempty"`
}
