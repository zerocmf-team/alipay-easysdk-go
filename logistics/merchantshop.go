/**
** @创建时间: 2020/12/8 1:43 下午
** @作者　　: return
** @描述　　: 商家门店
 */
package logistics

import (
	"encoding/json"
	"fmt"
	"github.com/gincmf/alipayEasySdk/util"
)

type MerchantShop struct {
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 创建门店
 * @Date 2020/12/8 13:54:46
 * @Param
 * @return
 **/
func (rest *MerchantShop) Create(bizContent map[string]interface{}) {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.instantdelivery.merchantshop.create", b)
	data := util.GetResult(params)
	fmt.Println("data", string(data))
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 更新门店
 * @Date 2021/4/6 14:50:33
 * @Param
 * @return
 **/
func (rest *MerchantShop) Update(bizContent map[string]interface{}) {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.instantdelivery.merchantshop.modify", b)
	data := util.GetResult(params)
	fmt.Println("data", string(data))
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 查询门店
 * @Date 2021/4/3 17:4:20
 * @Param
 * @return
 **/

func (rest *MerchantShop) Query(bizContent map[string]interface{}) {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.instantdelivery.merchantshop.query", b)
	data := util.GetResult(params)
	fmt.Println("data", string(data))
}

func (rest *MerchantShop) BatchQuery(bizContent map[string]interface{}) {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.instantdelivery.merchantshop.batchquery", b)
	data := util.GetResult(params)
	fmt.Println("data", string(data))
}
