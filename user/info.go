/**
** @创建时间: 2020/9/7 8:27 下午
** @作者　　: return
** @描述　　:
 */
package user

import (
	"github.com/gincmf/alipayEasySdk/util"
)

type Info struct {}

func (rest *Info) Share(appId string,authToken string) []byte {
	// 参数集合
	paramsMap := map[string]string{
		"method":      "alipay.user.info.share",
		"app_id":      appId,
		"auth_token":authToken,
	}

	data := util.GetResult(paramsMap)
	return data
}
