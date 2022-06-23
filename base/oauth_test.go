/**
** @创建时间: 2022/6/22 21:43
** @作者　　: return
** @描述　　:
 */

package base

import (
	"github.com/zerocmf/alipayEasySdkGo"
	_ "github.com/zerocmf/alipayEasySdkGo"
	"github.com/zerocmf/alipayEasySdkGo/data"
	"testing"
)

func TestOauth_GetToken(t *testing.T) {
	options := alipayEasySdkGo.GetOptions()
	data.SetOptions(options)
	oauth := new(Oauth)
	oauth.Agent("202105BBcf3de5e5472d4111acc38b70ca40bX61")
	oauth.GetToken("9bda6912b19c424bbae1059292b8XB71")
}
