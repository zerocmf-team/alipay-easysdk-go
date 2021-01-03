/**
** @创建时间: 2020/12/8 4:40 下午
** @作者　　: return
** @描述　　:
 */
package marketing

import (
	"encoding/json"
	"fmt"
	"github.com/gincmf/alipayEasySdk/data"
	"github.com/gincmf/alipayEasySdk/util"
)

type Card struct {
}

type TemplateStyleInfo struct {
	CardShowName string `json:"card_show_name,omitempty"`
	LogoId       string `json:"logo_id,omitempty"`
	Color        string `json:"color,omitempty"`
	BackgroundId string `json:"background_id,omitempty"`
	BgColor      string `json:"bg_color,omitempty"`
}

type TemplateBenefitInfo struct {
	Title       string   `json:"title,omitempty"`
	BenefitDesc []string `json:"benefit_desc,omitempty"`
	StartDate   string   `json:"start_date,omitempty"`
	EndDate     string   `json:"end_date,omitempty"`
}

type ColumnInfoList struct {
	Code        string    `json:"code,omitempty"`
	MoreInfo    *MoreInfo `json:"more_info,omitempty"`
	Title       string    `json:"title,omitempty"`
	OperateType string    `json:"operate_type,omitempty"`
	Value       string    `json:"value,omitempty"`
}

type MoreInfo struct {
	Title  string            `json:"title,omitempty"`
	Url    string            `json:"url,omitempty"`
	Params map[string]string `json:"params,omitempty"`
	Desc   []string          `json:"desc,omitempty"`
}

type FiledRuleList struct {
	FieldName string `json:"field_name,omitempty"`
	RuleName  string `json:"rule_name,omitempty"`
	RuleValue string `json:"rule_value,omitempty"`
}

type Fields struct {
	Required string `json:"required,omitempty"`
	Optional string `json:"optional,omitempty"`
}

type CardUserInfo struct {
	UserUniId     string `json:"user_uni_id"`
	UserUniIdType string `json:"user_uni_id_type"`
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 卡面文案列表
 * @Date 2020/12/10 11:1:32
 * @Param
 * @return
 **/
type FrontTextList struct {
	Label        string `json:"label,omitempty"`
	Value        string `json:"value,omitempty"`
	FrontImageId string `json:"front_image_id,omitempty"`
}

type MemberExtInfo struct {
	Name              string `json:"name,omitempty"`
	Gende             string `json:"gende,omitempty"`
	Birth             string `json:"birth,omitempty"`
	Cell              string `json:"cell,omitempty"`
}

type CardExtInfo struct {
	BizCardNo      string          `json:"biz_card_no,omitempty"`
	ExternalCardNo string          `json:"external_card_no,omitempty"`
	OpenDate       string          `json:"open_date"`
	ValidDate      string          `json:"valid_date"`
	Level          string          `json:"level,omitempty"`
	Point          string          `json:"point,omitempty"`
	Balance        string          `json:"balance,omitempty"`
	TemplateId     string          `json:"template_id,omitempty"`
	FrontTextList  []FrontTextList `json:"front_text_list,omitempty"`
}

type CardActionList struct {
	Code       string      `json:"code"`
	Text       string      `json:"text"`
	UrlType    string      `json:"url_type,omitempty"`
	Url        string      `json:"url,omitempty"`
	MiniAppUrl *MiniAppUrl `json:"mini_app_url,omitempty"`
}

type MiniAppUrl struct {
	MiniAppId      string `json:"mini_app_id"`
	MiniPageParam  string `json:"mini_page_param,omitempty"`
	MiniQueryParam string `json:"mini_query_param,omitempty"`
	DisplayOnList  string `json:"display_on_list"`
}

type CardCreateResult struct {
	Response alipayResponse `json:"alipay_marketing_card_template_create_response"`
	Sign     string         `json:"sign"`
}

type cardApplyResponse struct {
	data.AlipayResponse
	ApplyCardUrl string `json:"apply_card_url"`
}

type CardApplyResult struct {
	Response cardApplyResponse `json:"alipay_marketing_card_activateurl_apply_response"`
	Sign     string         `json:"sign"`
}


type cardSendResponse struct {
	data.AlipayResponse
}

type CardSendResult struct {
	Response cardSendResponse `json:"alipay_marketing_card_open_response"`
	Sign     string         `json:"sign"`
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 卡券模板创建
 * @Date 2020/12/8 16:41:32
 * @Param
 * @return
 **/
func (rest *Card) CreateTemplate(bizContent map[string]interface{}) CardCreateResult {
	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.marketing.card.template.create", b)
	data := util.GetResult(params)

	var cardCreateResult CardCreateResult
	json.Unmarshal(data,&cardCreateResult)

	return cardCreateResult

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 修改模板
 * @Date 2020/12/10 15:56:52
 * @Param
 * @return
 **/
func (rest *Card) ModifyTemplate(bizContent map[string]interface{}) {
	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
		fmt.Println(b)
	}

	params := util.GetParams("alipay.marketing.card.template.modify", b)
	data := util.GetResult(params)
	fmt.Println("data", string(data))
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 设置开卡表单
 * @Date 2020/12/9 23:28:30
 * @Param
 * @return
 **/

func (rest *Card) SetForm(bizContent map[string]interface{}) {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.marketing.card.formtemplate.set", b)
	data := util.GetResult(params)
	fmt.Println("data", string(data))

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取会员卡小程序开卡链接
 * @Date 2020/12/9 23:41:48
 * @Param
 * @return
 **/

func (rest *Card) Apply(bizContent map[string]interface{}) (CardApplyResult,error) {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.marketing.card.activateurl.apply", b)
	data := util.GetResult(params)

	var cardApplyResult CardApplyResult
	json.Unmarshal(data,&cardApplyResult)
	return cardApplyResult,nil

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 查询开卡表单
 * @Date 2020/12/10 10:44:58
 * @Param
 * @return
 **/
func (rest *Card) QueryForm(bizContent map[string]interface{}, authToken string) {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
		fmt.Println(b)
	}

	params := util.GetParams("alipay.marketing.card.activateform.query", b)

	params["auth_token"] = authToken

	data := util.GetResult(params)
	fmt.Println("data", string(data))

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 发卡给用户
 * @Date 2020/12/10 10:46:3
 * @Param
 * @return
 **/

func (rest *Card) Send(bizContent map[string]interface{}, authToken string) CardSendResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
		fmt.Println(b)
	}

	params := util.GetParams("alipay.marketing.card.open", b)

	params["auth_token"] = authToken

	data := util.GetResult(params)

	result := CardSendResult{}
	_ = json.Unmarshal(data, &result)

	return result
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 查询会员卡
 * @Date 2020/12/10 17:30:20
 * @Param
 * @return
 **/

func (rest *Card) Query(bizContent map[string]interface{}) {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
		fmt.Println(b)
	}

	params := util.GetParams("alipay.marketing.card.query", b)

	data := util.GetResult(params)
	fmt.Println("data", string(data))

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 更新会员卡
 * @Date 2020/12/10 16:48:2
 * @Param
 * @return
 **/

func (rest *Card) Update(bizContent map[string]interface{}) {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
		fmt.Println(b)
	}

	params := util.GetParams("alipay.marketing.card.update", b)

	data := util.GetResult(params)
	fmt.Println("data", string(data))

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 删除会员卡
 * @Date 2020/12/10 14:3:28
 * @Param
 * @return
 **/

func (rest *Card) Delete(bizContent map[string]interface{}) {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
		fmt.Println(b)
	}

	params := util.GetParams("alipay.marketing.card.delete", b)

	data := util.GetResult(params)
	fmt.Println("data", string(data))

}
