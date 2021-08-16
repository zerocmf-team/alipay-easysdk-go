/**
** @创建时间: 2021/5/14 8:02 下午
** @作者　　: return
** @描述　　: 对账单下载
 */
package payment

import (
	"encoding/json"
	"fmt"
	"github.com/gincmf/alipayEasySdk/util"
)

type Bill struct {
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 对账单下载
 * @Date 2021/5/14 20:3:34
 * @Param
 * @return
 **/
func (rest *Bill) DownloadUrl(bizContent map[string]interface{}) {
	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.data.dataservice.bill.downloadurl.query", b)
	data := util.GetResult(params)
	fmt.Println("data",string(data))
}
