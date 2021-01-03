/**
** @创建时间: 2020/11/29 9:11 下午
** @作者　　: return
** @描述　　:
 */
package payment

import (
	"encoding/json"
	"github.com/gincmf/alipayEasySdk"
	"github.com/gincmf/alipayEasySdk/data"
	"github.com/gincmf/alipayEasySdk/util"
)

type Common struct {
}

type GoodsDetail struct {
	GoodsId        string  `json:"goods_id"`
	GoodsName      string  `json:"goods_name"`
	Quantity       string  `json:"quantity"`
	Price          float64 `json:"price"`
	GoodsCategory  string  `json:"goods_category,omitempty"`
	CategoriesTree string  `json:"categories_tree,omitempty"`
	Body           string  `json:"body,omitempty"`
	ShowUrl        string  `json:"show_url,omitempty"`
}

type AlipayResponse struct {
	OutTradeNo string `json:"out_trade_no"`
	TradeNo    string `json:"trade_no"`
	data.AlipayResponse
}

type tradeResult struct {
	Response AlipayResponse `json:"alipay_trade_create_response"`
	Sign     string         `json:"sign"`
}

type AlipayTradeRefundResponse struct {
	AlipayResponse
	BuyerLogonId string `json:"buyer_logon_id"`
}

type refundResult struct {
	Response AlipayTradeRefundResponse `json:"alipay_trade_refund_response"`
	Sign     string                    `json:"sign"`
}

// 预创建订单
func (rest *Common) Create(bizContent map[string]interface{}) tradeResult {

	op := alipayEasySdk.Options()
	b, _ := json.Marshal(bizContent)

	paramsMap := make(map[string]string, 0)
	paramsMap["method"] = "alipay.trade.create"
	paramsMap["biz_content"] = string(b)

	if op.AppAuthToken != "" {
		paramsMap["app_auth_token"] = op.AppAuthToken
	}

	data := util.GetResult(paramsMap)
	result := tradeResult{}
	_ = json.Unmarshal(data, &result)

	return result
}

// 预创建退款
func (rest *Common) Refund(bizContent map[string]interface{}) refundResult {
	op := alipayEasySdk.Options()
	b, _ := json.Marshal(bizContent)

	paramsMap := make(map[string]string, 0)
	paramsMap["method"] = "alipay.trade.refund"
	paramsMap["biz_content"] = string(b)

	if op.AppAuthToken != "" {
		paramsMap["app_auth_token"] = op.AppAuthToken
	}

	data := util.GetResult(paramsMap)

	result := refundResult{}
	_ = json.Unmarshal(data, &result)

	return result
}
