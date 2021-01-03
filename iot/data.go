/**
** @创建时间: 2020/12/28 12:48 下午
** @作者　　: return
** @描述　　:
 */
package iot

type AlipayResponse struct {
	Code    string `json:"code"`
	Msg     string `json:"msg,omitempty"`
	SubCode string `json:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty"`
}

type AlipayResult struct {
	Response interface{} `json:"response"`
	Sign     string      `json:"sign"`
}
