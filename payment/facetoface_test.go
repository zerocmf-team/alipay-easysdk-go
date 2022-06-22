/**
** @创建时间: 2022/6/21 16:12
** @作者　　: return
** @描述　　:
 */

package payment

import (
	"fmt"
	"github.com/zerocmf/alipayEasySdkGo"
	"github.com/zerocmf/alipayEasySdkGo/data"
	"testing"
)

func TestNewFaceToFace(t *testing.T) {
	options := alipayEasySdkGo.GetOptions()
	data.SetOptions(options)

	faceToFace := new(FaceToFace)
	faceToFace.Agent("202106BB21c6e8a21c7846caa2c39362a20ccX15")

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
		fmt.Println("result", result)
	}
}
