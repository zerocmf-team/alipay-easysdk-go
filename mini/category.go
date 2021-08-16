/**
** @创建时间: 2021/3/1 11:59 上午
** @作者　　: return
** @描述　　:
 */
package mini

import (
	"encoding/json"
	"github.com/gincmf/alipayEasySdk/data"
	"github.com/gincmf/alipayEasySdk/util"
)

type Category struct{}

type CategoryQueryResult struct {
	Response CategoryQueryResponse `json:"alipay_open_mini_category_query_response"`
	Sign     string                `json:"sign"`
}

type CategoryQueryResponse struct {
	data.AlipayResponse
	CategoryList []CategoryList `json:"mini_category_list"`
}

type CategoryList struct {
	HasChild           bool   `json:"has_child"`
	ParentCategoryId   string `json:"parent_category_id"`
	CategoryName       string `json:"category_name"`
	CategoryId         string `json:"category_id"`
	NeedLicense        bool   `json:"need_license"`
	NeedOutDoorPic     bool   `json:"need_out_door_pic"`
	NeedSpecialLicense bool   `json:"need_special_license"`
}

func (rest *Category) Query(bizContent map[string]interface{}) CategoryQueryResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.mini.category.query", b)
	data := util.GetResult(params)

	result := CategoryQueryResult{}

	json.Unmarshal(data, &result)

	return result

}
