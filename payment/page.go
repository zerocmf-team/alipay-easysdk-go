/**
** @创建时间: 2020/12/24 11:54 下午
** @作者　　: return
** @描述　　:
 */
package payment

import (
	"encoding/json"
	"fmt"
	"github.com/gincmf/alipayEasySdk/data"
	"github.com/gincmf/alipayEasySdk/util"
)

type Page struct {
}

type pageResponse struct {
	OutTradeNo      string `json:"out_trade_no"`
	TradeNo         string `json:"trade_no"`
	TotalAmount     string `json:"total_amount"`
	MerchantOrderNo string `json:"merchant_order_no"`
	data.AlipayResponse
}

type pageResult struct {
	Response pageResponse `json:"alipay_trade_page_pay_response"`
	Sign     string       `json:"sign"`
}

func (rest *Page) Create(bizContent map[string]interface{}) pageResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.trade.page.pay", b)
	data := util.GetResult(params)

	fmt.Println(string(data))

	result := pageResult{}
	_ = json.Unmarshal(data, &result)

	return result
}
