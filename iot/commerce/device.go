/**
** @创建时间: 2020/12/30 3:33 下午
** @作者　　: return
** @描述　　:
 */
package commerce

import (
	"encoding/json"
	"github.com/gincmf/alipayEasySdk/data"
	"github.com/gincmf/alipayEasySdk/util"
)

type Device struct {}

type bindResult struct {
	Response data.AlipayResponse `json:"alipay_commerce_iot_device_bind_response"`
	Sign     string       `json:"sign"`
}

type unBindResult struct {
	Response data.AlipayResponse `json:"alipay_commerce_iot_device_unbind_response"`
	Sign     string       `json:"sign"`
}

// 设备绑定
func (rest *Device) Bind(bizContent map[string]string) bindResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.commerce.iot.device.bind", b)
	data := util.GetResult(params)

	result := bindResult{}
	_ = json.Unmarshal(data, &result)

	return result

}

// 设备解绑
func (rest *Device) UnBind(bizContent map[string]string) unBindResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.commerce.iot.device.unbind", b)
	data := util.GetResult(params)

	result := unBindResult{}
	_ = json.Unmarshal(data, &result)

	return result

}


