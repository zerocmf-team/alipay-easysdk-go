/**
** @创建时间: 2021/4/3 5:07 下午
** @作者　　: return
** @描述　　:
 */
package logistics

import (
	"encoding/json"
	"fmt"
	"github.com/gincmf/alipayEasySdk/util"
)

type Order struct {
}

type ServiceCodes struct {
	ServiceCode string `json:"service_code"`
}

type Companies struct {
	LogisticsCode string         `json:"logistics_code"`
	ServiceCodes  []ServiceCodes `json:"service_codes,omitempty"`
}

// 商品汇总信息
type GoodsInfo struct {
	Price       float64 `json:"price"`
	Weight      string  `json:"weight"`
	FirstClass  string  `json:"first_class"`
	SecondClass string  `json:"second_class"`
}

// 商品明细
type GoodsDetail struct {
	Count int     `json:"count"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Unit  string  `json:"unit"`
}

// 收件人信息
type ReceiverIstd struct {
	Name           string `json:"name"`
	City           string `json:"city"`
	Address        string `json:"address"`
	AddressDetail  string `json:"address_detail"`
	MobileNo       string `json:"mobile_no"`
	Lng            string `json:"lng"`
	Lat            string `json:"lat"`
	CoordinateType int    `json:"coordinate_type"`
}

// 即使配送订单扩展
type OrderExtIstd struct {
}

// 消费者通知明细
type ConsumerNotify struct {
	TinyAppId      string `json:"tiny_app_id"`
	TinyAppUrl     string `json:"tiny_app_url"`
	GoodsImg       string `json:"goods_img"`
	GoodsName      string `json:"goods_name"`
	GoodsCount     string `json:"goods_count"`
	MerchantName   string `json:"merchant_name"`
	MerchantMobile string `json:"merchant_mobile"`
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 预下单
 * @Date 2021/4/3 17:7:35
 * @Param
 * @return
 **/

func (rest *Order) PreCreate(bizContent map[string]interface{}) {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.commerce.logistics.order.instantdelivery.precreate", b)
	data := util.GetResult(params)
	fmt.Println("data", string(data))
}
