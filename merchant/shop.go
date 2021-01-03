/**
** @创建时间: 2020/12/20 10:14 上午
** @作者　　: return
** @描述　　:
 */
package merchant

import (
	"encoding/json"
	"fmt"
	"github.com/gincmf/alipayEasySdk/data"
	"github.com/gincmf/alipayEasySdk/util"
)

type Shop struct{}

type shopQueryResult struct {
	Response data.AlipayResponse `json:"ant_merchant_expand_shop_query_response"`
	Sign     string              `json:"sign"`
}

type shopResponse struct {
	data.AlipayResponse
	OrderId string `json:"order_id"`
}

type shopCreateResult struct {
	Response shopResponse `json:"ant_merchant_expand_shop_create_response"`
	Sign     string       `json:"sign"`
}

type shopModifyResult struct {
	Response shopResponse `json:"ant_merchant_expand_shop_modify_response"`
	Sign     string       `json:"sign"`
}

type statusResponse struct {
	data.AlipayResponse
	IpRoleId     []string `json:"ip_role_id"`
	MerchantName string   `json:"merchant_name"`
	Status       string   `json:"status"`
	ApplyTime    string   `json:"apply_time"`
	ExtInfo      string   `json:"ext_info"`
}

type shopStatusResult struct {
	Response statusResponse `json:"ant_merchant_expand_order_query_response"`
	Sign     string       `json:"sign"`
}

// 企业信息
type BusinessAddress struct {
	ProvinceCode string `json:"province_code"`
	CityCode     string `json:"city_code"`
	DistrictCode string `json:"district_code"`
	Address      string `json:"address"`
	Poiid        string `json:"poiid,omitempty"`
	Longitude    string `json:"longitude,omitempty"`
	Latitude     string `json:"latitude,omitempty"`
	Type         string `json:"type,omitempty"`
}

// 营业时间
type BusinessTime struct {
	WeekDay   int    `json:"week_day"`
	OpenTime  string `json:"open_time"`
	CloseTime string `json:"close_time"`
}

//联系人
type ContactInfo struct {
	Name   string `json:"name"`
	Phone  string `json:"phone,omitempty"`
	Mobile string `json:"mobile,omitempty"`
	Email  string `json:"email,omitempty"`
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 查询门店
 * @Date 2020/12/30 12:49:6
 * @Param
 * @return
 **/
func (rest *Shop) Query(bizContent map[string]interface{}) shopQueryResult {
	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("ant.merchant.expand.shop.query", b)

	data := util.GetResult(params)

	var result shopQueryResult

	err := json.Unmarshal(data, &result)

	if err != nil {
		fmt.Println("query", err)
	}

	return result
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 创建门店
 * @Date 2020/12/8 16:41:32
 * @Param
 * @return
 **/
func (rest *Shop) Create(bizContent map[string]interface{}) shopCreateResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("ant.merchant.expand.shop.create", b)
	data := util.GetResult(params)

	var result shopCreateResult
	json.Unmarshal(data, &result)

	return result

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 修改门店
 * @Date 2020/12/8 16:41:32
 * @Param
 * @return
 **/
func (rest *Shop) Modify(bizContent map[string]interface{}) shopModifyResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("ant.merchant.expand.shop.modify", b)
	data := util.GetResult(params)

	var result shopModifyResult
	json.Unmarshal(data, &result)

	return result

}

// 查询门店审核状态
func (rest *Shop) QueryStatus(bizContent map[string]interface{}) shopStatusResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("ant.merchant.expand.order.query", b)
	data := util.GetResult(params)

	fmt.Println(string(data))

	var result shopStatusResult
	json.Unmarshal(data, &result)

	return result
}
