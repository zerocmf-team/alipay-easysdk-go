/**
** @创建时间: 2022/6/22 11:52
** @作者　　: return
** @描述　　:
 */

package alipayEasySdkGo

import (
	"github.com/zerocmf/alipayEasySdkGo/data"
	"io/ioutil"
)

func GetOptions() *data.Options {
	config := new(data.Options)
	config.AppId = "2021001192664075"
	privateData, err := ioutil.ReadFile("../pem/private_key.pem")
	if err != nil {
		panic("读取私钥出错，文件不存在！")
	}
	config.MerchantPrivateKey = string(privateData)

	publicData, err := ioutil.ReadFile("../pem/public_key.pem")
	if err != nil {
		panic("读取公钥钥出错，文件不存在！")
	}
	config.MerchantPublicKey = string(publicData)

	alipayPublicData, err := ioutil.ReadFile("../pem/ali_public_key.pem")
	if err != nil {
		panic("读取公钥钥出错，文件不存在！")
	}
	config.AlipayCertPath = string(alipayPublicData)

	config.NotifyUrl = "http://www.codecloud.ltd/api/v1/app/alipay/receive_notify"
	return config
}
