/**
** @创建时间: 2022/6/21 16:33
** @作者　　: return
** @描述　　:
 */

package data

// 公共请求参数

type PublicParams struct {
	AppId        string `json:"app_id" sign:"app_id"`
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
 * @Description 配置代调用token
 * @Date 2022/6/22 0:53:25
 * @Param
 * @return
 **/

func (rest *Options) Agent(appAuthToken string) *Options {
	rest.AppAuthToken = appAuthToken
	return rest
}
