package tests

import (
	. "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
)

const (
	_AccessKey = ""
	_UserId = ""
	_Password = ""
	_ShipperNumber = ""

	_URL = "https://onlinetools.ups.com/webservices/"
	//_URL = "https://wwwcie.ups.com/webservices/"
)

func GetRequestUrl(addr string) string {
	return _URL + addr
}


var SoapHeader = &UPSSecurity{
	UsernameToken:      &UsernameToken{
		Username: _UserId,
		Password: _Password,
	},
	ServiceAccessToken: &ServiceAccessToken{
		AccessLicenseNumber: _AccessKey,
	},
}

func GetSoapHeader() *Certification {
	return &Certification{
		UPSSecurity: &UPSSecurity{
			UsernameToken:      &UsernameToken{
				Username: _UserId,
				Password: _Password,
			},
			ServiceAccessToken: &ServiceAccessToken{
				AccessLicenseNumber: _AccessKey,
			},
		},
	}
}

func GetSoapHeaderSecurity() *UPSSecurity {
	return &UPSSecurity{
		UsernameToken:      &UsernameToken{
			Username: _UserId,
			Password: _Password,
		},
		ServiceAccessToken: &ServiceAccessToken{
			AccessLicenseNumber: _AccessKey,
		},
	}
}
