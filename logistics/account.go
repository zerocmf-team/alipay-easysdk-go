/**
** @创建时间: 2020/12/8 2:16 下午
** @作者　　: return
** @描述　　:
 */
package logistics

import (
	"encoding/json"
	"fmt"
	"github.com/gincmf/alipayEasySdk/util"
)

type Account struct {

}

func (rest *Account) Create(bizContent map[string]interface{}) {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.instantdelivery.account.create",b)
	data := util.GetResult(params)
	fmt.Println("data",string(data))

}


func (rest *Account) Query(bizContent map[string]interface{}) {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.instantdelivery.account.query",b)
	data := util.GetResult(params)

	fmt.Println("data",string(data))

}
