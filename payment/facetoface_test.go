/**
** @创建时间: 2022/6/21 16:12
** @作者　　: return
** @描述　　:
 */

package payment

import (
	"fmt"
	"github.com/daifuyang/alipayEasySdkGo"
	"github.com/daifuyang/alipayEasySdkGo/data"
	"testing"
)

func TestNewFaceToFace(t *testing.T) {
	options := alipayEasySdkGo.GetOptions()
	options.AppId = "2021001192664075"
	data.SetOptions(options)

	faceToFace := new(FaceToFace)
	faceToFace.Agent("202105BBcf3de5e5472d4111acc38b70ca40bX61")

	bizContent := make(map[string]interface{}, 0)
	bizContent["out_trade_no"] = "20240320010101029"
	bizContent["subject"] = "test"
	bizContent["total_amount"] = "0.01"
	bizContent["trans_currency"] = "CNY"
	bizContent["settle_currency"] = "CNY"

	goodsDetail := make(map[string]interface{}, 0)
	goodsDetail["goods_id"] = "apple-01"
	goodsDetail["goods_name"] = "iphone12"
	goodsDetail["quantity"] = 1
	goodsDetail["price"] = 2

	bizContent["goods_detail"] = []map[string]interface{}{goodsDetail}

	result, err := faceToFace.PreCrete(bizContent)
	if err != nil {
		fmt.Println("err",err.Error())
	}
	fmt.Println("result", result)
}
