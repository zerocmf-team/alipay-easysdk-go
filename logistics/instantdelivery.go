/**
** @创建时间: 2020/12/7 7:06 下午
** @作者　　: return
** @描述　　:
 */
package logistics

import (
	"encoding/json"
	"fmt"
	"github.com/gincmf/alipayEasySdk"
	"github.com/gincmf/alipayEasySdk/util"
)

type InstantDelivery struct {
}

// 查询配送公司
func (rest *InstantDelivery) Query(bizContent map[string]interface{}) {
	op := alipayEasySdk.Options()

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	paramsMap := make(map[string]string, 0)
	paramsMap["method"] = "alipay.commerce.logistics.logisticscompany.instantdelivery.query"
	paramsMap["biz_content"] = b

	if op.AppAuthToken != "" {
		paramsMap["app_auth_token"] = op.AppAuthToken
	}

	data := util.GetResult(paramsMap)

	fmt.Println("data",string(data))
}