/**
** @创建时间: 2022/6/22 21:26
** @作者　　: return
** @描述　　: 可前往https://opendocs.alipay.com/open/02xtla查看更加详细的参数说明。
 */

package base

import (
	"encoding/json"
	"errors"
	"github.com/daifuyang/alipayEasySdkGo/data"
	"github.com/daifuyang/alipayEasySdkGo/util"
	"github.com/jinzhu/copier"
)

type Oauth struct {
	data.Options
	GrantType    string `json:"grant_type" sign:"grant_type"`
	Code         string `json:"code" sign:"code"`
	RefreshToken string `json:"refresh_token" sign:"refresh_token,omitempty"`
}

type oauthResult struct {
	Response oauthResponse `json:"alipay_system_oauth_token_response"`
	data.Sign
}

type errResult struct {
	Response data.AlipayResponse `json:"error_response"`
	data.Sign
}

type oauthResponse struct {
	data.AlipayResponse
	AlipayUserId string `json:"alipay_user_id"`
	UserId       string `json:"user_id"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	ReExpiresIn  string `json:"re_expires_in"`
	AuthStart    string `json:"auth_start"`
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取授权访问令牌 alipay.system.oauth.token
 * @Date 2022/6/22 21:28:14
 * @Param
 * @return
 **/

// alipay.system.oauth.token

func (rest *Oauth) GetToken(code string) (resp *oauthResult, err error) {
	grantType := "authorization_code"
	resp, err = rest.token(grantType, code)
	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 刷新授权访问令牌
 * @Date 2022/6/23 8:51:18
 * @Param
 * @return
 **/

func (rest *Oauth) Refresh(refreshToken string) (res *oauthResult, err error) {
	grantType := "refresh_token"
	res, err = rest.token(grantType, refreshToken)
	return
}

// 封装统一token请求
func (rest *Oauth) token(grantType string, code string) (res *oauthResult, err error) {
	//  获取通用公共参数
	config := data.GetOptions()
	options := new(Oauth)
	err = copier.Copy(&options, &config)
	if err != nil {
		return
	}
	options.AppAuthToken = rest.AppAuthToken
	// 组合请求参数
	options.Method = "alipay.system.oauth.token"
	options.GrantType = grantType
	options.Code = code
	resp, err := util.Post(options)
	if err != nil {
		return
	}
	errResp := new(errResult)
	err = json.Unmarshal(resp, &errResp)
	if err != nil {
		return
	}
	if errResp.Response.Code != "" {
		err = errors.New(errResp.Response.SubCode + "：" + errResp.Response.SubMsg)
		return
	}
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}
	return

}
