/**
** @创建时间: 2022/6/21 16:12
** @作者　　: return
** @描述　　:
 */

package payment

import (
	"encoding/json"
	"github.com/zerocmf/alipayEasySdkGo/data"
)

// 面对面类

type FaceToFace struct {
	data.PublicParams
}

type PreCreateResult struct {
	Response preCreateResponse `json:"alipay_trade_precreate_response"`
	data.Sign
}

type preCreateResponse struct {
	data.AlipayResponse
	OutTradeNo string `json:"out_trade_no"`
	QrCode     string `json:"qr_code"`
}

/**
 * @Author return <1140444693@qq.com>
 * @Description [alipay.trade.precreate(统一收单线下交易预创建)]https://opendocs.alipay.com/open/02ekfg
 * @Date 2022/6/22 11:44:4
 * @Param
 * @return
 **/

func (rest *FaceToFace) PreCrete(bizContent map[string]interface{}) (result PreCreateResult,err error) {
	options := data.GetOptions()
	options.AppAuthToken = rest.AppAuthToken
	options.Method = "alipay.trade.precreate"
	bizBytes, _ := json.Marshal(bizContent)
	options.BizContent = string(bizBytes)
	encode := options.EncodeAndSign()
	data, err := options.Post(encode)
	json.Unmarshal(data,&result)
	return
}
