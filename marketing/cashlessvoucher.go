/**
** @创建时间: 2020/12/10 6:13 下午
** @作者　　: return
** @描述　　:
 */
package marketing

import (
	"encoding/json"
	"github.com/gincmf/alipayEasySdk/data"
	"github.com/gincmf/alipayEasySdk/util"
)

type CashlessVoucher struct{}

type alipayResponse struct {
	data.AlipayResponse
	TemplateId string `json:"template_id"`
}

type VoucherCreateResult struct {
	Response alipayResponse `json:"alipay_marketing_cashlessvoucher_template_create_response"`
	Sign     string         `json:"sign"`
}

type VoucherModifyResult struct {
	Response alipayResponse `json:"alipay_marketing_cashlessvoucher_template_modify_response"`
	Sign     string         `json:"sign"`
}

type sendResponse struct {
	data.AlipayResponse
	VoucherId string `json:"voucher_id"`
	UserId    string `json:"user_id"`
}

type VoucherSendResponse struct {
	Response sendResponse `json:"alipay_marketing_voucher_send_response"`
	Sign     string       `json:"sign"`
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 有效期
 * @Date 2020/12/10 18:22:2
 * @Param
 * @return
 **/

type VoucherValidPeriod struct {
	Type     string `json:"type"`
	Start    string `json:"start,omitempty"`
	End      string `json:"end,omitempty"`
	Duration int    `json:"duration,omitempty"`
	Unit     string `json:"unit,omitempty"`
}

type VoucherAvailableTime struct {
	DayRule   string `json:"day_rule,omitempty"`
	TimeBegin string `json:"time_begin,omitempty"`
	TimeEnd   string `json:"time_end,omitempty"`
}

type RuleConf struct {
	Pid   string `json:"PID,omitempty"`
	Store string `json:"STORE,omitempty"`
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 创建优惠券卡券模板
 * @Date 2020/12/10 18:14:55
 * @Param
 * @return
 **/
func (rest *CashlessVoucher) CreateTemplate(bizContent map[string]interface{}, voucherType string) VoucherCreateResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.marketing.cashlessvoucher.template.create", b)
	if voucherType == "item" {
		params = util.GetParams("alipay.marketing.cashlessitemvoucher.template.create", b)
	}

	data := util.GetResult(params)

	vr := VoucherCreateResult{}
	json.Unmarshal(data, &vr)

	return vr

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 修改优惠券
 * @Date 2020/12/11 21:20:27
 * @Param
 * @return
 **/
func (rest *CashlessVoucher) ModifyTemplate(bizContent map[string]interface{}, voucherType string) VoucherModifyResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.marketing.cashlessvoucher.template.modify", b)
	if voucherType == "item" {
		params = util.GetParams("alipay.marketing.cashlessitemvoucher.template.modify", b)
	}

	data := util.GetResult(params)

	vr := VoucherModifyResult{}
	json.Unmarshal(data, &vr)

	return vr

}

func (rest *CashlessVoucher) Send(bizContent map[string]interface{}) VoucherSendResponse {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.marketing.voucher.send", b)
	data := util.GetResult(params)

	vs := VoucherSendResponse{}

	json.Unmarshal([]byte(data), &vs)

	return vs

}
