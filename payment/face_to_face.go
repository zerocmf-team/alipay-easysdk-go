/**
** @创建时间: 2020/10/15 9:19 下午
** @作者　　: return
** @描述　　:
 */
package payment

import (
	"encoding/json"
	"fmt"
	"github.com/gincmf/alipayEasySdk"
	"github.com/gincmf/alipayEasySdk/data"
	"github.com/gincmf/alipayEasySdk/util"
)

type FaceToFace struct{}

type payResult struct {
	Response payResponse `json:"alipay_trade_pay_response"`
	Sign     string      `json:"sign"`
}

type voucherDetailList struct {
	Id                         string  `json:"id"`
	Name                       string  `json:"name"`
	Type                       string  `json:"type"`
	Amount                     float64 `json:"amount"`
	MerchantContribute         float64 `json:"merchant_contribute,omitempty"`
	OtherContribute            float64 `json:"other_contribute,omitempty"`
	Memo                       string  `json:"memo,omitempty"`
	TemplateId                 string  `json:"template_id,omitempty"`
	PurchaseBuyerContribute    float64 `json:"purchase_buyer_contribute,omitempty"`
	PurchaseMerchantContribute float64 `json:"purchase_merchant_contribute,omitempty"`
	PurchaseAntContribute      float64 `json:"purchase_ant_contribute,omitempty"`
}

type payResponse struct {
	data.AlipayResponse
	TradeNo             string            `json:"trade_no"`
	OutTradeNo          string            `json:"out_trade_no"`
	BuyerLogonId        string            `json:"buyer_logon_id"`
	TotalAmount         string           `json:"total_amount"`
	ReceiptAmount       string           `json:"receipt_amount"`
	BuyerPayAmount      string           `json:"buyer_pay_amount,omitempty"`
	PointAmount         float64           `json:"point_amount,omitempty"`
	InvoiceAmount       float64           `json:"invoice_amount,omitempty"`
	GmtPayment          string            `json:"gmt_payment,omitempty"`
	FundBillList        fundBillList      `json:"fund_bill_list"`
	StoreName           string            `json:"store_name,omitempty"`
	BuyerUserId         string            `json:"buyer_user_id"`
	DiscountGoodsDetail string            `json:"discount_goods_detail,omitempty"`
	VoucherDetailList   voucherDetailList `json:"voucher_detail_list,omitempty"`
	BuyerUserType       string            `json:"buyer_user_type,omitempty"`
	MdiscountAmount     float64           `json:"mdiscount_amount,omitempty"`
	DiscountAmount      float64           `json:"discount_amount"`
	BuyerUserName       string            `json:"buyer_user_name"`
}

type fundBillList struct {
	FundChannel string  `json:"fund_channel"`
	Amount      float64 `json:"amount"`
	RealAmount  float64 `json:"real_amount"`
}

// 扫用户出示的付款码，完成付款
func (rest *FaceToFace) Pay(bizContent map[string]interface{}) payResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.trade.pay", b)
	data := util.GetResult(params)

	var result payResult
	json.Unmarshal(data, &result)

	return result

}

// 生成二维码付款
func (rest *FaceToFace) PreCreate(bizContent map[string]interface{}) {

	op := alipayEasySdk.Options()

	b, _ := json.Marshal(bizContent)

	// 参数集合
	paramsMap := make(map[string]string, 0)

	paramsMap["method"] = "alipay.trade.precreate"
	paramsMap["biz_content"] = string(b)

	if op.AppAuthToken != "" {
		paramsMap["app_auth_token"] = op.AppAuthToken
	}

	data := util.GetResult(paramsMap)
	var result map[string]interface{}
	json.Unmarshal(data, &result)

	fmt.Println("data", result)

}
