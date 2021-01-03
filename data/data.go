/**
** @创建时间: 2020/12/11 4:30 下午
** @作者　　: return
** @描述　　:
 */
package data

type AlipayResponse struct {
	Code       string `json:"code"`
	Msg        string `json:"msg"`
	SubCode    string `json:"sub_code"`
	SubMsg     string `json:"sub_msg"`
}