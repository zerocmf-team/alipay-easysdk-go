/**
** @创建时间: 2020/12/20 6:28 下午
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

type Order struct{}

type OrderSyncResult struct {
	Response syncResponse `json:"alipay_merchant_order_sync_response"`
	Sign     string       `json:"sign"`
}

type syncResponse struct {
	data.AlipayResponse
	RecordId string `json:"record_id"`
	OrderId  string `json:"order_id"`
	OrderStatus string `json:"order_status"`
}

func (rest *Order) Sync(bizContent map[string]interface{}) OrderSyncResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.merchant.order.sync", b)
	data := util.GetResult(params)

	fmt.Println(string(data))

	var orderSyncResult OrderSyncResult
	json.Unmarshal(data, &orderSyncResult)

	return orderSyncResult

}
