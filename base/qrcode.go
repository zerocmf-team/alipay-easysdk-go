/**
** @创建时间: 2022/6/24 12:47
** @作者　　: return
** @描述　　: 小程序二维码 Qrcode
 */

package base

import (
	"encoding/json"
	"github.com/daifuyang/alipayEasySdkGo/data"
	"github.com/daifuyang/alipayEasySdkGo/util"
	"github.com/jinzhu/copier"
)

type QrOption struct {
	f func(*optional)
}

type optional struct {
	Color string `json:"color,omitempty" sign:"color,omitempty"`
	Size  string `json:"size,omitempty" sign:"size,omitempty"`
}

type Qrcode struct {
	data.Options
}

type qrcodeResult struct {
	Response qrcodeResponse `json:"alipay_open_app_qrcode_create_response"`
	data.Sign
}

type qrcodeResponse struct {
	data.AlipayResponse
	QrCodeUrl            string `json:"qr_code_url"`
	QrCodeUrlCircleWhite string `json:"qr_code_url_circle_white,omitempty"`
	QrCodeUrlCircleBlue  string `json:"qr_code_url_circle_blue,omitempty"`
}

const S = "s"
const M = "m"
const L = "l"

func (rest *Qrcode) WithColor(color string) QrOption {
	return QrOption{func(o *optional) {
		o.Color = color
	}}
}

func (rest *Qrcode) WithSize(size string) QrOption {
	return QrOption{func(o *optional) {
		o.Size = size
	}}
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 创建小程序二维码
 * @Date 2022/6/24 12:50:18
 * @Param urlParam 小程序中能访问到的页面路径，例如：page/component/component-pages/view/view
 * @Param queryParam 小程序的启动参数，打开小程序的query ，在小程序 onLaunch的方法中获取
 * @Param describe 二维码描述
 * @return 可前往alipay.open.app.qrcode.create查看更加详细的参数说明。
 **/

func (rest *Qrcode) Create(urlParam string, queryParam string, describe string, ops ...QrOption) (res *qrcodeResult, err error) {

	config := data.GetOptions()
	options := new(Qrcode)
	err = copier.Copy(&options, &config)
	if err != nil {
		return
	}

	options.AppAuthToken = rest.AppAuthToken
	// 组合请求参数
	options.Method = "alipay.open.app.qrcode.create"

	bizContent := make(map[string]interface{}, 0)
	bizContent["query_param"] = describe
	bizContent["describe"] = queryParam
	bizContent["url_param"] = urlParam

	opt := new(optional)
	for _, o := range ops {
		o.f(opt)
	}
	if opt.Color != "" {
		bizContent["color"] = opt.Color
	}
	if opt.Size != "" {
		bizContent["size"] = opt.Size
	}

	bizBytes, _ := json.Marshal(bizContent)
	options.BizContent = string(bizBytes)

	resp, err := util.Post(options)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &resp)
	if err != nil {
		return
	}
	return
}
