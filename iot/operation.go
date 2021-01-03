/**
** @创建时间: 2020/12/26 8:40 下午
** @作者　　: return
** @描述　　:
 */
package iot

import (
	"encoding/json"
	"github.com/gincmf/alipayEasySdk/data"
	"github.com/gincmf/alipayEasySdk/util"
)

type Operation struct {
}

type operationResponse struct {
	QrCodeUrl string `json:"qr_code_url"`
	BatchNo   string `json:"batch_no"`
	data.AlipayResponse
}

type operationResult struct {
	Response operationResponse `json:"alipay_open_sp_operation_qrcode_query_response"`
	Sign     string            `json:"sign"`
}

func (rest *Operation) Query(bizContent map[string]string) operationResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.sp.operation.qrcode.query", b)
	data := util.GetResult(params)

	result := operationResult{}
	_ = json.Unmarshal(data, &result)

	return result

}
