/**
** @创建时间: 2020/9/7 9:10 上午
** @作者　　: return
** @描述　　:
 */
package base

import (
	"encoding/json"
	"fmt"
	"github.com/gincmf/alipayEasySdk"
	"github.com/gincmf/alipayEasySdk/data"
	"github.com/gincmf/alipayEasySdk/util"
)

type Oauth struct{}

type systemTokenResponse struct {
	data.AlipayResponse
	AlipayUserId string `json:"alipay_user_id"`
	UserId       string `json:"user_id"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	ReExpiresIn  string `json:"re_expires_in"`
}

type systemTokenResult struct {
	Response      systemTokenResponse `json:"alipay_system_oauth_token_response"`
	ErrorResponse systemTokenResponse `json:"error_response"`
	Sign          string              `json:"sign"`
}

type tokens struct {
	UserId          string `json:"user_id"`
	AuthAppId       string `json:"auth_app_id"`
	AppAuthToken    string `json:"app_auth_token"`
	AppRefreshToken string `json:"app_refresh_token"`
	ExpiresIn       string `json:"expires_in"`
	ReExpiresIn     string `json:"re_expires_in"`
}

type openTokenResponse struct {
	data.AlipayResponse
	Tokens          []tokens `json:"tokens,omitempty"`
	UserId          string   `json:"user_id,omitempty"`
	AuthAppId       string   `json:"auth_app_id,omitempty"`
	AppAuthToken    string   `json:"app_auth_token,omitempty"`
	AppRefreshToken string   `json:"app_refresh_token,omitempty"`
	ExpiresIn       string   `json:"expires_in,omitempty"`
	ReExpiresIn     string   `json:"re_expires_in,omitempty"`
}

type openTokenResult struct {
	Response openTokenResponse `json:"alipay_open_oauth_token_response"`
	Sign     string            `json:"sign"`
}

/*
	date: 2020-09-07 19:35
	Author: frank_dai
	desc: 通过isv的授权码获取商户信息
*/
func (rest *Oauth) GetSystemToken(code string) systemTokenResult {

	grantType := "authorization_code" // 授权码模式
	bizContent := make(map[string]string, 0)
	bizContent["grant_type"] = grantType
	bizContent["code"] = code
	b, _ := json.Marshal(bizContent)

	// 参数集合
	paramsMap := map[string]string{
		"method":      "alipay.system.oauth.token",
		"code":        code,
		"grant_type":  grantType,
		"biz_content": string(b),
	}

	appAuthToken := alipayEasySdk.Options()
	if appAuthToken != nil && appAuthToken.AppAuthToken != "" {
		paramsMap["app_auth_token"] = appAuthToken.AppAuthToken
	}
	data := util.GetResult(paramsMap)

	fmt.Println(string(data))

	result := systemTokenResult{}

	json.Unmarshal(data, &result)

	return result
}

/*
	date: 2020-09-07
	Author: frank_dai
	desc: 通过isv的授权码获取开放平台token信息
*/
func (rest *Oauth) GetOpenToken(code string) openTokenResult {

	grantType := "authorization_code" // 授权码模式
	bizContent := make(map[string]string, 0)
	bizContent["grant_type"] = grantType
	bizContent["code"] = code
	b, _ := json.Marshal(bizContent)
	fmt.Println("bizContent", string(b))

	// 参数集合
	paramsMap := map[string]string{
		"method":      "alipay.open.auth.token.app",
		"code":        code,
		"grant_type":  grantType,
		"biz_content": string(b),
	}

	appAuthToken := alipayEasySdk.Options()
	if appAuthToken != nil && appAuthToken.AppAuthToken != "" {
		paramsMap["app_auth_token"] = appAuthToken.AppAuthToken
	}

	data := util.GetResult(paramsMap)

	result := openTokenResult{}

	json.Unmarshal(data, &result)

	return result
}
