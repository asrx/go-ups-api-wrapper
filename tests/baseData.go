package tests

import (
	"github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-ups-api-wrapper/src/Common/SimpleType"
	. "github.com/asrx/go-ups-api-wrapper/src/RateService/ComplexType"
	ComplexType2 "github.com/asrx/go-ups-api-wrapper/src/ShipService/ComplexType"
)

var Shipper = &ComplexType.Party{
	Name:          "ANL",
	AttentionName: "DonovanXu",
	ShipperNumber: _ShipperNumber,
	Address:       &ComplexType.AddressType{
		AddressLine:                []string{"16018 Adelante st Suite D"},
		City:                        "IRWINDALE",
		StateProvinceCode:           "CA",
		PostalCode:                  "91702",
		CountryCode:                 "US",
		//ResidentialAddressIndicator: "",
	},
}

var ShippFrom = &ComplexType.Party{
	Name:          "ANL",
	AttentionName: "DonovanXu",
	Address:       &ComplexType.AddressType{
		AddressLine:                []string{"16018 Adelante st Suite D"},
		City:                        "IRWINDALE",
		StateProvinceCode:           "CA",
		PostalCode:                  "91702",
		CountryCode:                 "US",
		//ResidentialAddressIndicator: "",
	},
}

var Recipient1 = &ComplexType.Party{
	Name:          "ANL",
	AttentionName: "Alex",
	//ShipperNumber: _ShipperNumber,
	Address:       &ComplexType.AddressType{
		AddressLine:                []string{"24300 Nandina Ave"},
		City:                        "Moreno Valley",
		StateProvinceCode:           "CA",
		PostalCode:                  "92551",
		CountryCode:                 "US",
		//ResidentialAddressIndicator: "false",
	},
}

var Recipient = &ComplexType.Party{
	Name:          "ANL",
	AttentionName: "Alex",
	//ShipperNumber: _ShipperNumber,
	Address:       &ComplexType.AddressType{
		AddressLine:                []string{"14306 Flora Ln"},
		City:                        "Wellington",
		StateProvinceCode:           "FL",
		PostalCode:                  "33414",
		CountryCode:                 "US",
		//ResidentialAddressIndicator: "false",
	},
}

var Recipient2 = &ComplexType.Party{
	Name:          "ANL",
	AttentionName: "Alex",
	//ShipperNumber: _ShipperNumber,
	Address:       &ComplexType.AddressType{
		AddressLine:                []string{"4 Gordon Ln"},
		City:                        "Westport",
		StateProvinceCode:           "CT",
		PostalCode:                  "06880",
		CountryCode:                 "US",
		//ResidentialAddressIndicator: "false",
	},
}

func GetRatePackages() []*PackageType {
	// 包装超过最大尺寸总限制165英寸(长+周长，其中周长是2x宽+ 2x高)。
	var packages = []*PackageType{}
	for i:=0;i <10;i++{
		packages = append(packages, &PackageType{
			PackagingType:               &ComplexType.CodeDescriptionType{
				Code:        SimpleType.PT_PACKAGE,
			},
			//Dimensions:                  &ComplexType.DimensionsType{
			//	UnitOfMeasurement: &ComplexType.CodeDescriptionType{
			//		Code:        SimpleType.UOM_IN,
			//	},
			//	Length:            "4",
			//	Width:             "4",
			//	Height:            "4",
			//},
			PackageWeight:               &ComplexType.PackageWeightType{
				UnitOfMeasurement: &ComplexType.CodeDescriptionType{
					Code:        SimpleType.UOM_LBS,
				},
				Weight:            "20",
			},
		})
	}

	return packages
}

func GetRatePackagesGfp() []*PackageType {
	// 包装超过最大尺寸总限制165英寸(长+周长，其中周长是2x宽+ 2x高)。
	var packages = []*PackageType{}
	for i:=0;i<10;i++{
		packages = append(packages, &PackageType{
			PackagingType:               &ComplexType.CodeDescriptionType{
				Code:        SimpleType.PT_PACKAGE,
			},
			//Dimensions:                  &ComplexType.DimensionsType{
			//	UnitOfMeasurement: &ComplexType.CodeDescriptionType{
			//		Code:        SimpleType.UOM_IN,
			//	},
			//	Length:            "4",
			//	Width:             "4",
			//	Height:            "4",
			//},
			PackageWeight:               &ComplexType.PackageWeightType{
				UnitOfMeasurement: &ComplexType.CodeDescriptionType{
					Code:        SimpleType.UOM_LBS,
				},
				Weight:            "20",
			},
			// gfp must
			Commodity:                   &ComplexType.CommodityType{
				FreightClass: SimpleType.CLASS_50,
			},
		})
	}
	return packages
}

func GetRatePackagesHundredweight() []*PackageType {
	// 包装超过最大尺寸总限制165英寸(长+周长，其中周长是2x宽+ 2x高)。
	var packages = []*PackageType{}
	for i:=0;i<10;i++{
		packages = append(packages, &PackageType{
			PackagingType:               &ComplexType.CodeDescriptionType{
				Code:        SimpleType.PT_PACKAGE,
			},
			//Dimensions:                  &ComplexType.DimensionsType{
			//	UnitOfMeasurement: &ComplexType.CodeDescriptionType{
			//		Code:        SimpleType.UOM_IN,
			//	},
			//	Length:            "4",
			//	Width:             "4",
			//	Height:            "4",
			//},
			PackageWeight:               &ComplexType.PackageWeightType{
				UnitOfMeasurement: &ComplexType.CodeDescriptionType{
					Code:        SimpleType.UOM_LBS,
				},
				Weight:            "20",
			},
			Commodity:                   &ComplexType.CommodityType{
				FreightClass: SimpleType.CLASS_50,
			},
		})
	}
	return packages
}

func GetShipPackages() []*ComplexType2.PackageType {
	// 包装超过最大尺寸总限制165英寸(长+周长，其中周长是2x宽+ 2x高)。
	var packages = []*ComplexType2.PackageType{}
	packages = append(packages, &ComplexType2.PackageType{
		PackagingType:               &ComplexType.CodeDescriptionType{
			Code:        SimpleType.PT_PACKAGE,
		},
		//Dimensions:                  &DimensionsType{
		//	UnitOfMeasurement: &ComplexType.CodeDescriptionType{
		//		Code:        SimpleType.UOM_IN,
		//	},
		//	Length:            "10",
		//	Width:             "10",
		//	Height:            "10",
		//},
		PackageWeight:               &ComplexType.PackageWeightType{
			UnitOfMeasurement: &ComplexType.CodeDescriptionType{
				Code:        SimpleType.UOM_LBS,
			},
			Weight:            "22.05",
		},
		ReferenceNumber:                    &ComplexType.ReferenceNumberType{
			Code:  SimpleType.CODE_PRODUCTION_CODE,
			//Code:  SimpleType.CODE_INVOICE_NUMBER,
			Value: "TEST001",
		},
		//Commodity:                   &CommodityType{
		//	FreightClass: SimpleType.CLASS_50,
		//},
	})

	//packages = append(packages, &PackageType{
	//	PackagingType:               &ComplexType.CodeDescriptionType{
	//		Code:        SimpleType.PT_PACKAGE,
	//	},
	//	Dimensions:                  &DimensionsType{
	//		UnitOfMeasurement: &ComplexType.CodeDescriptionType{
	//			Code:        SimpleType.UOM_IN,
	//		},
	//		Length:            "10",
	//		Width:             "20",
	//		Height:            "15",
	//	},
	//	Commodity:                   &CommodityType{
	//		FreightClass: SimpleType.CLASS_50,
	//	},
	//	PackageWeight:               &PackageWeightType{
	//		UnitOfMeasurement: &ComplexType.CodeDescriptionType{
	//			Code:        SimpleType.UOM_LBS,
	//		},
	//		Weight:            "22.5",
	//	},
	//})
	return packages
}