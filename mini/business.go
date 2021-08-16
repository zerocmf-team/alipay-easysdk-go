/**
** @创建时间: 2021/4/28 12:17 下午
** @作者　　: return
** @描述　　:
 */
package mini

import (
	"fmt"
	"github.com/gincmf/alipayEasySdk/util"
)

type Business struct {

}

func (rest *Business) Certify(bizContent map[string]string)  {

	params := util.GetParams("alipay.open.mini.individual.business.certify", "")

	for k,v := range bizContent{
		params[k] = v
	}

	data := util.GetResult(params)

	fmt.Println(string(data))


}
