/**
** @创建时间: 2020/9/7 1:32 下午
** @作者　　: return
** @描述　　:
 */
package base

import (
	"encoding/json"
	"fmt"
	"github.com/gincmf/alipayEasySdk"
	"github.com/gincmf/alipayEasySdk/util"
)

type App struct {

}

func (rest *App) upload() {

	options := alipayEasySdk.Options()

	templateVersion := "0.0.1"
	templateId := "2021001192675085"
	appVersion := "0.0.1"

	bizContent := make(map[string]string, 0)
	bizContent["template_version"] = templateVersion
	bizContent["template_id"] = templateId
	bizContent["app_version"] = appVersion

	b, _ := json.Marshal(bizContent)
	fmt.Println("bizContent", string(b))

	// 参数集合
	paramsMap := map[string]string{
		"method":         "alipay.open.mini.version.upload",
		"app_auth_token": options.AppAuthToken,
		"biz_content": string(b),
	}

	data := util.GetResult(paramsMap)
	fmt.Println("data",data)
}
