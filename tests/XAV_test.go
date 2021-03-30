package tests

import (
	"encoding/json"
	"fmt"
	"github.com/asrx/go-ups-api-wrapper/src/AddressValidationService"
	"github.com/asrx/go-ups-api-wrapper/src/AddressValidationService/ComplexType"
	. "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	"testing"
)

func addr(AddressLine, city, state, postCode string) *ComplexType.AddressKeyFormatType {
	return &ComplexType.AddressKeyFormatType{
		// 收货人名字
		ConsigneeName: "Donovan.xu",
		// 建筑物的名称
		//AttentionName: "",
		AddressLine:   []string{
			AddressLine,
		},
		//Region:        "",
		PoliticalDivision2: city,// 城市/城镇
		PoliticalDivision1: state,// 州/省
		PostcodePrimaryLow: postCode,// 邮编
		//Urbanization:  "",
		CountryCode:   "US",
	}
}

func Test_XAV(t *testing.T) {

	//addr := addr("12380 Morris Road","Alpharetta","GA","30009")
	//addr := addr("BILLY MITCHELL","Addison","TX","75001")
	//addr := addr("109 Iron Bridge Rd","Cliff","NM","88028")
	//addr := addr("790 Rock Lake Rd","Jamestown","CO","80455")
	addr := addr("790 Rock Lake Rd","Jamestown","CO","80455")

	request := &ComplexType.XAVRequest{
		Request:                  &RequestType{
			RequestOption:        []string{"3"},
		},
		// 如果存在此标记，将不会进行US和PR街道级别的地址验证。默认值是提供街道级别的地址验证
		//RegionalRequestIndicator: "",
		// 最大候选值，默认15
		//MaximumCandidateListSize: "",
		AddressKeyFormat:         addr,
	}

	c := AddressValidationService.NewXAVPortType(GetSoapHeaderSecurity(), false)

	resp, fault := c.ProcessXAV(request)
	if fault != nil {
		fmt.Println("==================================================")
		if fault.Errors == nil {
			fmt.Println("not have Error Message")
			return
		}
		for _, exp := range fault.Errors.ErrorDetail{
			fmt.Printf("%s: %s\n", exp.PrimaryErrorCode.Code, exp.PrimaryErrorCode.Description)
		}
		fmt.Println("**************************************************")
		return
	}
	fmt.Println("============== SUCCESS ==============")

	byte, err := json.Marshal(resp)

	if err != nil {
		fmt.Println("Marshal Error: ")
		fmt.Printf("%+v\n",resp)
	}else{
		fmt.Println("json: ")
		fmt.Println(string(byte))
	}

	// 查询找到一个有效的匹配项。
	//resp.ValidAddressIndicator
	// 无法匹配到地址信息，候选名单,模糊的地址指示器
	//resp.AmbiguousAddressIndicator
	//for addr := range resp.AmbiguousAddressIndicator {
	//	//addr
	//}
	// 无候选项
	// 地址类型 0 -不保密的，1 -商业，2 -住宅
	//resp.AddressClassification
}