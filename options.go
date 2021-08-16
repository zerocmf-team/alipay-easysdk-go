/**
** @创建时间: 2020/9/5 10:41 下午
** @作者　　: return
** @描述　　:
 */
package alipayEasySdk

import (
	"io/ioutil"
	"reflect"
)

// 配置基本信息
type baseOptions struct {
	Protocol           string `json:"protocol"`
	GatewayHost        string `json:"gateway_host"`
	SignType           string `json:"sign_type"`
	Charset            string `json:"charset"`
	AppId              string `json:"appId"`
	Version            string `json:"version"`
	PrivateKey         string `json:"private_key"`
	PublicKey          string `json:"public_key"`
	AliPublicKey       string `json:"ali_public_key"`
	AppCertPath        string `json:"app_cert_path"`
	AlipayCertPath     string `json:"alipay_cert_path"`
	AlipayRootCertPath string `json:"alipay_root_cert_path"`
	NotifyUrl          string `json:"notify_url"`
	EncryptType        string `json:"encrypt_type"`
	EncryptKey         string `json:"encrypt_key"`
	AppAuthToken       string `json:"app_auth_token"`
}

var options *baseOptions

func NewOptions(params map[string]string) baseOptions {

	options = &baseOptions{
		Protocol:           params["protocol"],
		GatewayHost:        params["gatewayHost"],
		SignType:           params["signType"],
		Charset:            params["charset"],
		AppId:              params["appId"],
		Version:            params["version"],
		AlipayRootCertPath: params["alipayRootCertPath"],
		NotifyUrl:          params["notifyUrl"],
		EncryptType:        params["encryptType"],
		EncryptKey:         params["encryptKey"],
		AppAuthToken:       params["appAuthToken"],
	}

	if params["appCertPath"] != "" {
		privateData, err := ioutil.ReadFile(params["appCertPath"] + "/private_key.pem")
		if err != nil {
			panic("读取私钥出错，文件不存在！")
		}

		options.PrivateKey = string(privateData)

		publicData, err := ioutil.ReadFile(params["appCertPath"] + "/public_key.pem")
		if err != nil {
			panic("读取公钥钥出错，文件不存在！")
		}

		options.PublicKey = string(publicData)

		options.AppCertPath = params["appCertPath"]

	}

	if params["aliCertPath"] != "" {
		publicData, err := ioutil.ReadFile(params["aliCertPath"] + "/ali_public_key.pem")
		if err != nil {
			panic("读取公钥钥出错，文件不存在！")
		}

		options.AliPublicKey = string(publicData)
	}

	return *options
}

func SetOption(key string, val string) {
	oPoint := reflect.ValueOf(options)
	field := oPoint.Elem().FieldByName(key)
	field.SetString(val)
}

func Options() baseOptions {
	if options != nil {
		return *options
	}
	panic("请先初始化options")
}
