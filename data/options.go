/**
** @创建时间: 2022/6/21 16:33
** @作者　　: return
** @描述　　:
 */

package data

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"github.com/zerocmf/alipayEasySdkGo/utils"
	"net/url"
	"sort"
	"strings"
	"time"
)

// 公共请求参数

type PublicParams struct {
	AppId        string `json:"appId" sign:"app_id"'`
	Method       string `json:"method" sign:"method"`
	Format       string `json:"format" sign:"format"`
	Charset      string `json:"charset" sign:"charset"`
	SignType     string `json:"sign_type" sign:"sign_type"`
	Sign         string `json:"sign" sign:"sign"`
	Timestamp    string `json:"timestamp" sign:"timestamp"`
	Version      string `json:"version" sign:"version"`
	NotifyUrl    string `json:"notify_url,omitempty" sign:"notify_url,omitempty"`
	AppAuthToken string `json:"app_auth_token,omitempty" sign:"app_auth_token,omitempty"`
	BizContent   string `json:"biz_content" sign:"biz_content,omitempty"`
}

// 配置基本信息

type Options struct {
	Protocol           string `json:"protocol"`
	GatewayHost        string `json:"gateway_host"`
	MerchantPrivateKey string `json:"merchant_private_key"`
	MerchantPublicKey  string `json:"merchant_public_key"`
	AlipayPublicKey    string `json:"alipay_public_key"`
	MerchantCertPath   string `json:"merchant_cert_path,omitempty"`    // <-- 请填写您的应用公钥证书文件路径，例如：/foo/appCertPublicKey_2019051064521003.crt -->
	AlipayCertPath     string `json:"alipay_cert_path,omitempty"`      // <-- 请填写您的支付宝公钥证书文件路径，例如：/foo/alipayCertPublicKey_RSA2.crt -->
	AlipayRootCertPath string `json:"alipay_root_cert_path,omitempty"` // <-- 请填写您的支付宝根证书文件路径，例如：/foo/alipayRootCert.crt -->
	EncryptType        string `json:"encrypt_type,omitempty"`
	EncryptKey         string `json:"encrypt_key,omitempty"`
	PublicParams
}

var options *Options

/**
 * @Author return <1140444693@qq.com>
 * @Description 第一次设置参数
 * @Date 2022/6/22 8:30:54
 * @Param
 * @return
 **/

func SetOptions(in *Options) {
	in.Protocol = "https"
	in.GatewayHost = "openapi.alipay.com/gateway.do"
	in.Format = "JSON"
	in.Charset = "utf-8"
	in.SignType = "RSA2"
	in.Version = "1.0"
	options = in
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取默认配置项
 * @Date 2022/6/22 0:53:10
 * @Param
 * @return
 **/

func GetOptions() *Options {
	return options
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 序列化参数并进行签名操作
 * @Date 2022/6/21 22:20:56
 * @Param
 * @return
 **/

func (rest *Options) EncodeAndSign() (encode string) {

	unix := time.Now().Unix() // 时间戳
	time := time.Unix(unix, 0).Format("2006-01-02 15:04:05")
	rest.Timestamp = time
	sign, _ := rest.GenerateSign()
	rest.Sign = sign

	// 获取提交的参数列表
	params := url.Values{}
	json := utils.ReflectPtr(rest, "sign")

	for k, v := range json {
		value := []byte(v)
		value = bytes.TrimSpace(value)
		if string(value) != "" {
			params.Set(k, string(value))
		}
	}
	encode = params.Encode()
	return

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 对配置项进行签名操作
 * @Date 2022/6/21 22:23:19
 * @Param
 * @return
 **/

func (rest *Options) GenerateSign() (sign string, encode string) {
	//ksort 对参数进行排序

	var keys []string
	json := utils.ReflectPtr(rest, "sign")
	for k := range json {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 对参数进行序列化
	pStr := make([]string, 0)
	//拼接
	for _, k := range keys {
		v := []byte(json[k])
		v = bytes.TrimSpace(v)
		if string(v) != "" {
			pStr = append(pStr, k+"="+json[k])
		}
	}

	// 序列化结果
	encode = strings.Join(pStr, "&")
	h := sha256.New()
	h.Write([]byte(encode))

	block := []byte(rest.MerchantPrivateKey)
	blocks, _ := pem.Decode(block)
	privateKey, err := x509.ParsePKCS8PrivateKey(blocks.Bytes)
	if err != nil {
		return "", ""
	}

	digest := h.Sum(nil)
	s, _ := rsa.SignPKCS1v15(nil, privateKey.(*rsa.PrivateKey), crypto.SHA256, digest)
	sign = base64.StdEncoding.EncodeToString(s)
	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 配置代调用token
 * @Date 2022/6/22 0:53:25
 * @Param
 * @return
 **/

func (rest *PublicParams) Agent(appAuthToken string) *PublicParams {
	rest.AppAuthToken = appAuthToken
	return rest
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 基于请求库的阿里post请求封装
 * @Date 2022/6/22 8:19:18
 * @Param
 * @return
 **/

func (rest *Options) Post(encode string) (data []byte, err error) {
	protocol := rest.Protocol   // 协议
	baseUrl := rest.GatewayHost // 网关
	url := protocol + "://" + baseUrl + "?" + encode
	body := strings.NewReader(encode)
	data, err = utils.Request("POST", url, body, nil)
	return
}
