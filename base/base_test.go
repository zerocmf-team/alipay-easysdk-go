/**
** @创建时间: 2022/6/22 21:43
** @作者　　: return
** @描述　　:
 */

package base

import (
	"fmt"
	"github.com/daifuyang/alipayEasySdkGo"
	_ "github.com/daifuyang/alipayEasySdkGo"
	"github.com/daifuyang/alipayEasySdkGo/data"
	"net/url"
	"testing"
)

func TestOauth_GetToken(t *testing.T) {
	options := alipayEasySdkGo.GetOptions()
	options.AppId = "2021001192664075"
	data.SetOptions(options)
	oauth := new(Oauth)
	oauth.Agent("202105BBcf3de5e5472d4111acc38b70ca40bX61")
	resp, err := oauth.GetToken("9bda6912b19c424bbae1059292b8XB71")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("resp", resp)
}

func TestQrcode_Create(t *testing.T) {
	options := alipayEasySdkGo.GetOptions()
	options.AppId = "2021001192664075"
	data.SetOptions(options)
	qrcode := new(Qrcode)
	qrcode.Agent("202105BBcf3de5e5472d4111acc38b70ca40bX61")

	queryParam := "pages/store/index"

	urls := url.Values{}
	urls.Add("store_number", "487934091")
	urls.Add("desk_id", "3")
	urlParam := urls.Encode()

	resp, err := qrcode.Create(queryParam, urlParam, "扫码点单", qrcode.WithColor("0x00BFFF"), qrcode.WithSize(S))

	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println("resp", resp)

}

func TestImage_Upload(t *testing.T) {
	options := alipayEasySdkGo.GetOptions()
	options.AppId = "2021001192664075"
	data.SetOptions(options)
	image := new(Image)
	image.Agent("202105BBcf3de5e5472d4111acc38b70ca40bX61")
	filepath := "/Users/return/workspace/yijing/changsanjiao/src/assets/images/stepSecond/1.png"
	image.Upload("testName",filepath)
}
