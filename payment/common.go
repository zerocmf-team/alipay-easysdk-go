/**
** @创建时间: 2022/6/27 08:40
** @作者　　: return
** @描述　　: 通用接口 Common
 */

package payment

import (
	"encoding/json"
	"github.com/daifuyang/alipayEasySdkGo/data"
	"github.com/daifuyang/alipayEasySdkGo/util"
	"github.com/jinzhu/copier"
	"strings"
)

type Common struct {
	data.Options
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 创建交易
 * @Date 2022/6/27 8:45:8
 * @Param
 * @return 可前往alipay.trade.create查看更加详细的参数说明。
 **/

type createResult struct {
	Response createResponse `json:"alipay_trade_create_response"`
	data.Sign
}

type createResponse struct {
	data.AlipayResponse
	OutTradeNo string `json:"out_trade_no"`
	TradeNo    string `json:"trade_no"`
}

func (rest *Common) Create(bizContent map[string]string) (res *createResult, err error) {
	config := data.GetOptions()
	options := new(Common)
	err = copier.Copy(&options, &config)
	if err != nil {
		return nil, err
	}
	options.AppAuthToken = rest.AppAuthToken
	options.Method = "alipay.trade.create"
	bizBytes, _ := json.Marshal(bizContent)
	options.BizContent = string(bizBytes)
	resp, err := util.Post(options)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}
	return

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 查询交易
 * @Date 2022/6/27 8:46:55
 * @Param
 * @return 可前往alipay.trade.query查看更加详细的参数说明。
 **/

type queryResult struct {
	Response queryResponse `json:"alipay_trade_query_response"`
	data.Sign
}

type queryResponse struct {
	data.AlipayResponse
	TradeNo         string  `json:"trade_no"`
	OutTradeNo      string  `json:"out_trade_no"`
	BuyerLogonID    string  `json:"buyer_logon_id"`
	TradeStatus     string  `json:"trade_status"`
	TotalAmount     float64 `json:"total_amount"`
	TransCurrency   string  `json:"trans_currency"`
	SettleCurrency  string  `json:"settle_currency"`
	SettleAmount    float64 `json:"settle_amount"`
	PayCurrency     int     `json:"pay_currency"`
	PayAmount       string  `json:"pay_amount"`
	SettleTransRate string  `json:"settle_trans_rate"`
	TransPayRate    string  `json:"trans_pay_rate"`
	BuyerPayAmount  float64 `json:"buyer_pay_amount"`
	PointAmount     int     `json:"point_amount"`
	InvoiceAmount   float64 `json:"invoice_amount"`
	SendPayDate     string  `json:"send_pay_date"`
	ReceiptAmount   string  `json:"receipt_amount"`
	StoreID         string  `json:"store_id"`
	TerminalID      string  `json:"terminal_id"`
	FundBillList    []struct {
		FundChannel string  `json:"fund_channel"`
		Amount      int     `json:"amount"`
		RealAmount  float64 `json:"real_amount"`
	} `json:"fund_bill_list"`
	StoreName             string `json:"store_name"`
	BuyerUserID           string `json:"buyer_user_id"`
	IndustrySepcDetailGov string `json:"industry_sepc_detail_gov"`
	IndustrySepcDetailAcc string `json:"industry_sepc_detail_acc"`
	ChargeAmount          string `json:"charge_amount"`
	ChargeFlags           string `json:"charge_flags"`
	SettlementID          string `json:"settlement_id"`
	TradeSettleInfo       struct {
		TradeSettleDetailList []struct {
			OperationType     string `json:"operation_type"`
			OperationSerialNo string `json:"operation_serial_no"`
			OperationDt       string `json:"operation_dt"`
			TransOut          string `json:"trans_out"`
			TransIn           string `json:"trans_in"`
			Amount            int    `json:"amount"`
			OriTransOut       string `json:"ori_trans_out"`
			OriTransIn        string `json:"ori_trans_in"`
		} `json:"trade_settle_detail_list"`
	} `json:"trade_settle_info"`
	AuthTradePayMode    string `json:"auth_trade_pay_mode"`
	BuyerUserType       string `json:"buyer_user_type"`
	MdiscountAmount     string `json:"mdiscount_amount"`
	DiscountAmount      string `json:"discount_amount"`
	Subject             string `json:"subject"`
	Body                string `json:"body"`
	AlipaySubMerchantID string `json:"alipay_sub_merchant_id"`
	ExtInfos            string `json:"ext_infos"`
	PassbackParams      string `json:"passback_params"`
	HbFqPayInfo         struct {
		UserInstallNum string `json:"user_install_num"`
	} `json:"hb_fq_pay_info"`
	CreditPayMode    string `json:"credit_pay_mode"`
	CreditBizOrderID string `json:"credit_biz_order_id"`
}

func (rest *Common) Query(bizContent map[string]string) (res *queryResult, err error) {
	config := data.GetOptions()
	options := new(Common)
	err = copier.Copy(&options, &config)
	if err != nil {
		return
	}
	options.AppAuthToken = rest.AppAuthToken
	options.Method = "alipay.trade.query"
	bizBytes, _ := json.Marshal(bizContent)
	options.BizContent = string(bizBytes)
	resp, err := util.Post(options)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}
	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 交易退款
 * @Date 2022/6/27 8:49:26
 * @Param
 * @return 可前往alipay.trade.refund查看更加详细的参数说明。
 **/

type refundResult struct {
	Response refundResponse `json:"alipay_trade_refund_response"`
	data.Sign
}

type refundResponse struct {
	data.AlipayResponse
	TradeNo              string  `json:"trade_no"`
	OutTradeNo           string  `json:"out_trade_no"`
	BuyerLogonID         string  `json:"buyer_logon_id"`
	FundChange           string  `json:"fund_change"`
	RefundFee            float64 `json:"refund_fee"`
	RefundDetailItemList []struct {
		FundChannel string  `json:"fund_channel"`
		Amount      int     `json:"amount"`
		RealAmount  float64 `json:"real_amount"`
		FundType    string  `json:"fund_type"`
	} `json:"refund_detail_item_list"`
	StoreName   string `json:"store_name"`
	BuyerUserID string `json:"buyer_user_id"`
	SendBackFee string `json:"send_back_fee"`
}

func (rest *Common) Refund(bizContent map[string]string) (res *refundResult, err error) {
	config := data.GetOptions()
	options := new(Common)
	err = copier.Copy(&options, &config)
	if err != nil {
		return
	}
	options.AppAuthToken = rest.AppAuthToken
	options.Method = "alipay.trade.refund"
	bizBytes, _ := json.Marshal(bizContent)
	options.BizContent = string(bizBytes)
	resp, err := util.Post(options)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}
	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 关闭交易
 * @Date 2022/6/27 8:54:27
 * @Param
 * @return 可前往alipay.trade.close查看更加详细的参数说明。
 **/

type closeResult struct {
	Response closeResponse `json:"alipay_trade_close_response"`
	data.Sign
}

type closeResponse struct {
	data.AlipayResponse
	OutTradeNo string `json:"out_trade_no"`
	TradeNo    string `json:"trade_no"`
}

func (rest *Common) Close(bizContent map[string]string) (res *closeResult, err error) {
	config := data.GetOptions()
	options := new(Common)
	err = copier.Copy(&options, &config)
	if err != nil {
		return
	}
	options.AppAuthToken = rest.AppAuthToken
	options.Method = "alipay.trade.close"
	bizBytes, _ := json.Marshal(bizContent)
	options.BizContent = string(bizBytes)
	resp, err := util.Post(options)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}
	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 撤销交易
 * @Date 2022/6/27 8:55:25
 * @Param
 * @return 可前往alipay.trade.cancel查看更加详细的参数说明。
 **/

type cancelResult struct {
	Response cancelResponse `json:"alipay_trade_cancel_response"`
	data.Sign
}

type cancelResponse struct {
	TradeNo    string `json:"trade_no"`
	OutTradeNo string `json:"out_trade_no"`
	RetryFlag  string `json:"retry_flag"`
	Action     string `json:"action"`
}

func (rest *Common) Cancel(bizContent map[string]string) (res *cancelResult, err error) {
	config := data.GetOptions()
	options := new(Common)
	err = copier.Copy(&options, &config)
	if err != nil {
		return
	}
	options.AppAuthToken = rest.AppAuthToken
	options.Method = "alipay.trade.close"
	bizBytes, _ := json.Marshal(bizContent)
	options.BizContent = string(bizBytes)
	resp, err := util.Post(options)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}
	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 交易退款查询
 * @Date 2022/6/27 8:56:13
 * @Param
 * @return 可前往alipay.trade.fastpay.refund.query查看更加详细的参数说明。
 **/

type queryRefundResult struct {
	Response queryRefundResponse `json:"alipay_trade_cancel_response"`
	data.Sign
}

type queryRefundResponse struct {
	data.AlipayResponse
	TradeNo        string  `json:"trade_no"`
	OutTradeNo     string  `json:"out_trade_no"`
	OutRequestNo   string  `json:"out_request_no"`
	TotalAmount    float64 `json:"total_amount"`
	RefundAmount   float64 `json:"refund_amount"`
	RefundStatus   string  `json:"refund_status"`
	RefundRoyaltys []struct {
		RefundAmount  int    `json:"refund_amount"`
		RoyaltyType   string `json:"royalty_type"`
		ResultCode    string `json:"result_code"`
		TransOut      string `json:"trans_out"`
		TransOutEmail string `json:"trans_out_email"`
		TransIn       string `json:"trans_in"`
		TransInEmail  string `json:"trans_in_email"`
	} `json:"refund_royaltys"`
	GmtRefundPay         string `json:"gmt_refund_pay"`
	RefundDetailItemList []struct {
		FundChannel string  `json:"fund_channel"`
		Amount      int     `json:"amount"`
		RealAmount  float64 `json:"real_amount"`
		FundType    string  `json:"fund_type"`
	} `json:"refund_detail_item_list"`
	SendBackFee     string `json:"send_back_fee"`
	DepositBackInfo struct {
		HasDepositBack     string  `json:"has_deposit_back"`
		DbackStatus        string  `json:"dback_status"`
		DbackAmount        float64 `json:"dback_amount"`
		BankAckTime        string  `json:"bank_ack_time"`
		EstBankReceiptTime string  `json:"est_bank_receipt_time"`
	} `json:"deposit_back_info"`
}

func (rest *Common) QueryRefund(bizContent map[string]string) (res *queryRefundResult, err error) {
	config := data.GetOptions()
	options := new(Common)
	err = copier.Copy(&options, &config)
	if err != nil {
		return
	}
	options.AppAuthToken = rest.AppAuthToken
	options.Method = "alipay.trade.fastpay.refund.query"
	bizBytes, _ := json.Marshal(bizContent)
	options.BizContent = string(bizBytes)
	resp, err := util.Post(options)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}
	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 查询对账单
 * @Date 2022/6/27 8:56:59
 * @Param
 * @return 可前往alipay.data.dataservice.bill.downloadurl.query查看更加详细的参数说明。
 **/

type downloadResult struct {
	Response downloadResponse `json:"alipay_data_dataservice_bill_downloadurl_query_response"`
}

type downloadResponse struct {
	data.AlipayResponse
	BillDownloadUrl string `json:"bill_download_url"`
}

func (rest *Common) DownloadBill(bizContent map[string]string) (res *downloadResult, err error) {
	config := data.GetOptions()
	options := new(Common)
	err = copier.Copy(&options, &config)
	if err != nil {
		return
	}
	options.AppAuthToken = rest.AppAuthToken
	options.Method = "alipay.data.dataservice.bill.downloadurl.query"
	bizBytes, _ := json.Marshal(bizContent)
	options.BizContent = string(bizBytes)
	resp, err := util.Post(options)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}
	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 异步通知验签
 * @Date 2022/6/27 21:59:8
 * @Param
 * @return
 **/

func (rest *Common) VerifyNotify(parameters map[string]string) (err error) {
	sign := strings.ReplaceAll(parameters["sign"], " ", "+")
	delete(parameters, "sign")
	delete(parameters, "sign_type")
	encode := util.SortEncode(parameters)
	return util.VerifyNotify(encode, sign, "alipay")
}
