/**
** @创建时间: 2021/5/9 10:47 上午
** @作者　　: return
** @描述　　:
 */
package mini

import (
	"encoding/json"
	"github.com/gincmf/alipayEasySdk/data"
	"github.com/gincmf/alipayEasySdk/util"
)

type TemplateMessage struct {}

type sendResponse struct {
	Response data.AlipayResponse `json:"alipay_open_app_mini_templatemessage_send_response"`
	Sign     string              `json:"sign"`
}

func (rest *TemplateMessage) Send(bizContent map[string]interface{}) sendResponse {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.app.mini.templatemessage.send", b)

	data := util.GetResult(params)

	result := sendResponse{}

	json.Unmarshal(data, &result)

	return result

}
