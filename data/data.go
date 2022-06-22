/**
** @创建时间: 2022/6/22 11:41
** @作者　　: return
** @描述　　:
 */

package data

type AlipayResponse struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty"`
}

type Sign struct {
	Sign string `json:"sign"`
}
