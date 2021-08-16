/**
** @创建时间: 2021/1/8 10:59 上午
** @作者　　: return
** @描述　　:
 */
package merchant

import (
	"encoding/json"
	"fmt"
	"github.com/gincmf/alipayEasySdk/data"
	"github.com/gincmf/alipayEasySdk/util"
)

type Qrcode struct{}

type QrcodeCreateResult struct {
	Response QrcodeCreateResponse `json:"alipay_open_app_qrcode_create_response"`
	Sign string `json:"sign"`
}

type QrcodeCreateResponse struct {
	data.AlipayResponse
	QrCodeUrl string              `json:"qr_code_url"`
}

/**
 * @Author return <1140444693@qq.com>
 * @Description // 创建二维码
 * @Date 2021/1/8 11:7:30
 * @Param
 * @return
 **/
func (rest *Qrcode) Create(bizContent map[string]interface{}) QrcodeCreateResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.app.qrcode.create", b)

	data := util.GetResult(params)

	result := QrcodeCreateResult{}

	json.Unmarshal(data,&result)

	return result

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 绑定二维码
 * @Date 2020/12/30 12:49:6
 * @Param
 * @return
 **/
func (rest *Qrcode) bind(bizContent map[string]interface{}) {
	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.mini.qrcode.bind", b)

	data := util.GetResult(params)

	fmt.Println(string(data))
}
