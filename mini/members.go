/**
** @创建时间: 2021/2/21 10:19 下午
** @作者　　: return
** @描述　　:
 */
package mini

import (
	"encoding/json"
	"fmt"
	"github.com/gincmf/alipayEasySdk/data"
	"github.com/gincmf/alipayEasySdk/util"
)

type Member struct {
}

type MemberQueryResult struct {
	Response MemberQueryResponse `json:"alipay_open_app_members_query_response"`
	Sign     string              `json:"sign"`
}

type AppMemberInfoList struct {
	UserId    string `json:"user_id"`
	NickName  string `json:"nick_name"`
	Portrait  string `json:"portrait"`
	Status    string `json:"status"`
	GmtJoin   string `json:"gmt_join"`
	LogonId   string `json:"logon_id"`
	GmtInvite string `json:"gmt_invite"`
	Role      string `json:"role"`
}

type MemberQueryResponse struct {
	data.AlipayResponse
	AppMemberInfoList []AppMemberInfoList `json:"app_member_info_list"`
}

type MemberCreateResult struct {
	Response MemberCreateResponse `json:"alipay_open_app_members_create_response"`
	Sign     string               `json:"sign"`
}

type MemberDeleteResult struct {
	Response MemberCreateResponse `json:"alipay_open_app_members_delete_response"`
	Sign     string               `json:"sign"`
}

type MemberCreateResponse struct {
	data.AlipayResponse
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 查询成员列表
 * @Date 2021/2/21 22:27:2
 * @Param
 * @return
 **/
func (rest *Member) Query(bizContent map[string]interface{}) MemberQueryResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.app.members.query", b)
	data := util.GetResult(params)

	result := MemberQueryResult{}
	json.Unmarshal(data, &result)

	return result

}

/**
 * @Author return <1140444693@qq.com>
 * @Description // 添加成员
 * @Date 2021/2/21 22:27:22
 * @Param
 * @return
 **/

func (rest *Member) Create(bizContent map[string]interface{}) MemberCreateResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.app.members.create", b)
	data := util.GetResult(params)

	result := MemberCreateResult{}
	json.Unmarshal(data, &result)

	fmt.Println(result)

	return result

}

/**
 * @Author return <1140444693@qq.com>
 * @Description // 删除成员
 * @Date 2021/2/21 23:38:9
 * @Param
 * @return
 **/

func (rest *Member) Delete(bizContent map[string]interface{}) MemberDeleteResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.app.members.delete", b)
	data := util.GetResult(params)

	result := MemberDeleteResult{}
	json.Unmarshal(data, &result)

	return result

}
