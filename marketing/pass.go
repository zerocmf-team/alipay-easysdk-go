/**
** @创建时间: 2022/6/27 08:36
** @作者　　: return
** @描述　　:
 */

package marketing

import (
	"encoding/json"
	"github.com/daifuyang/alipayEasySdkGo/data"
	"github.com/daifuyang/alipayEasySdkGo/util"
)

type Pass struct {
	data.Options
}

type passResponse struct {
	Result    string `json:"result"`
	ErrorCode string `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 支付宝卡包 Pass 卡券模板创建
 * @Date 2022/6/30 9:16:34
 * @Param uniqueId 商户用于控制模版的唯一性（可以使用时间戳保证唯一性）
 * @Param tplContent 模板内容信息，遵循JSON规范，详情参见tpl_content
 * @return 可前往alipay.pass.template.add查看更加详细的参数说明。
 **/

type createResult struct {
	Response createResponse `json:"alipay_pass_template_add_response"`
	data.Sign
}

type createResponse struct {
	data.AlipayResponse
	passResponse
	TplID     string   `json:"tpl_id"`
	TplParams []string `json:"tpl_params"`
}

func (rest *Pass) CreateTemplate(uniqueId string, tplContent string) (res *createResult, err error) {
	options := data.GetOptions()
	options.AppAuthToken = rest.AppAuthToken
	// 组合请求参数
	options.Method = "alipay.pass.template.add"
	bizContent := make(map[string]interface{}, 0)
	bizContent["unique_id"] = uniqueId
	bizContent["tpl_content"] = tplContent
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

	err = json.Unmarshal([]byte(res.Response.Result), &res.Response)
	if err != nil {
		return
	}

	return

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 支付宝卡包 Pass 卡券模板更新
 * @Date 2022/6/30 10:29:9
 * @Param
 * @return 可前往alipay.pass.template.update查看更加详细的参数说明。
 **/

type updateResult struct {
	Response updateResponse `json:"alipay_pass_template_update_response"`
	data.Sign
}

type updateResponse struct {
	data.AlipayResponse
	passResponse
	TplID     string   `json:"tpl_id"`
	TplParams []string `json:"tpl_params"`
}

func (rest *Pass) UpdateTemplate(uniqueId string, tplContent string) (res *updateResult, err error) {
	options := data.GetOptions()
	options.AppAuthToken = rest.AppAuthToken
	// 组合请求参数
	options.Method = "alipay.pass.template.update"
	bizContent := make(map[string]interface{}, 0)
	bizContent["unique_id"] = uniqueId
	bizContent["tpl_content"] = tplContent
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
	err = json.Unmarshal([]byte(res.Response.Result), &res.Response)
	if err != nil {
		return
	}
	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 卡券实例发放接口
 * @Date 2022/6/30 10:33:18
 * @Param
 * @return 可前往alipay.pass.template.add查看更加详细的参数说明。
 **/

type addInstanceResult struct {
	Response addInstanceResponse `json:"alipay_pass_instance_add_response"`
	data.Sign
}

type addInstanceResponse struct {
	data.AlipayResponse
	passResponse
	SerialNumber string `json:"serialNumber"`
	PassID       string `json:"passId"`
	Operate      string `json:"operate"`
}

func (rest *Pass) AddInstance(tplId string, tplParams string, recognitionType string, recognitionInfo string) (res *addInstanceResult, err error) {
	options := data.GetOptions()
	options.AppAuthToken = rest.AppAuthToken
	// 组合请求参数
	options.Method = "alipay.pass.instance.add"
	bizContent := make(map[string]interface{}, 0)
	bizContent["tpl_id"] = tplId
	bizContent["tpl_content"] = tplParams
	bizContent["recognition_type"] = recognitionType
	bizContent["recognition_info"] = recognitionInfo
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
	err = json.Unmarshal([]byte(res.Response.Result), &res.Response)
	if err != nil {
		return
	}
	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 卡券实例更新
 * @Date 2022/6/30 11:22:57
 * @Param
 * @return 可前往alipay.pass.instance.update查看更加详细的参数说明。
 **/

type updateOption struct {
	f func(*updateOptional)
}

type updateOptional struct {
	TplParams  string `json:"tpl_params"`
	Status     string `json:"status"`
	VerifyCode string `json:"verify_code"`
	VerifyType string `json:"verify_type"`
}

type updateInsResult struct {
	Response updateInsResponse `json:"alipay_pass_instance_update_response"`
	data.Sign
}

type updateInsResponse struct {
	data.AlipayResponse
	passResponse
	SerialNumber string `json:"serialNumber"`
	PassID       string `json:"passId"`
	Operate      string `json:"operate"`
}

func (rest *Pass) UpdateInstance(userId string, serialNumber string, channelId string, opts ...updateOption) (res *updateInsResult, err error) {
	options := data.GetOptions()
	options.AppAuthToken = rest.AppAuthToken
	// 组合请求参数
	options.Method = "alipay.pass.instance.add"
	bizContent := make(map[string]interface{}, 0)

	bizContent["user_id"] = userId
	bizContent["serial_number"] = serialNumber
	bizContent["channel_id"] = channelId

	up := new(updateOptional)
	for _, o := range opts {
		o.f(up)
	}

	if up.TplParams != "" {
		bizContent["tpl_params"] = up.TplParams
	}

	if up.Status != "" {
		bizContent["status"] = up.Status
	}

	if up.VerifyCode != "" {
		bizContent["verify_code"] = up.VerifyCode
	}

	if up.VerifyType != "" {
		bizContent["verify_code"] = up.VerifyType
	}

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
	err = json.Unmarshal([]byte(res.Response.Result), &res.Response)
	if err != nil {
		return
	}
	return
}
