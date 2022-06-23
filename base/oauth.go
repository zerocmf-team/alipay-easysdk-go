/**
** @创建时间: 2022/6/22 21:26
** @作者　　: return
** @描述　　: 可前往https://opendocs.alipay.com/open/02xtla查看更加详细的参数说明。
 */

package base

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zerocmf/alipayEasySdkGo/data"
	"github.com/zerocmf/alipayEasySdkGo/util"
)

type Oauth struct {
	data.Options
	GrantType    string `json:"grant_type" sign:"grant_type"`
	Code         string `json:"code" sign:"code"`
	RefreshToken string `json:"refresh_token" sign:"refresh_token,omitempty"`
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取授权访问令牌 alipay.system.oauth.token
 * @Date 2022/6/22 21:28:14
 * @Param
 * @return
 **/

// alipay.system.oauth.token

func (rest *Oauth) GetToken(code string) {
	grantType := "authorization_code"
	rest.token(grantType,code)
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 刷新授权访问令牌
 * @Date 2022/6/23 8:51:18
 * @Param
 * @return
 **/

func (rest *Oauth) Refresh(refreshToken string) {
	grantType := "refresh_token"
	rest.token(grantType,refreshToken)
}

// 封装统一token请求
func (rest *Oauth) token(grantType string ,code string) {
	//  获取通用公共参数
	config := data.GetOptions()
	options := new(Oauth)
	copier.Copy(&options, &config)

	options.AppAuthToken = rest.AppAuthToken
	// 组合请求参数
	options.Method = "alipay.system.oauth.token"
	options.GrantType = grantType
	options.Code = code

	// 获取签名参数
	params := util.ReflectPtr(options, "sign")
	encode := util.EncodeAndSign(options.MerchantPrivateKey, params)

	data, err := options.Post(encode)
	if err != nil {
		fmt.Println("err", err.Error())
	}

	fmt.Println("data", string(data))
}
