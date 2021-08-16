/**
** @创建时间: 2021/2/21 1:08 下午
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

type Version struct {
}

type VersionUploadResult struct {
	Response VersionUploadResponse `json:"alipay_open_mini_version_upload_response"`
	Sign     string                `json:"sign"`
}

type VersionUploadResponse struct {
	data.AlipayResponse
	CreateStatus string `json:"create_status"`
	NeedRotation string `json:"need_rotation"`
	BuildStatus  string `json:"build_status"`
}

type VersionExperienceCreateResult struct {
	Response data.AlipayResponse `json:"alipay_open_mini_experience_create_response"`
	Sign     string              `json:"sign"`
}

type VersionExperienceQueryResult struct {
	Response VersionExperienceQueryResponse `json:"alipay_open_mini_experience_query_response"`
	Sign     string                         `json:"sign"`
}

type VersionExperienceQueryResponse struct {
	data.AlipayResponse
	ExpQrCodeUrl string `json:"exp_qr_code_url"`
	Status       string `json:"status"`
}

/**
 * @Author return <1140444693@qq.com>
 * @Description // 上传小程序
 * @Date 2021/2/21 13:17:59
 * @Param
 * @return
 **/
func (rest *Version) Upload(bizContent map[string]interface{}) VersionUploadResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.mini.version.upload", b)
	data := util.GetResult(params)

	upload := VersionUploadResult{}

	json.Unmarshal(data, &upload)

	return upload

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 查询小程序版本列表
 * @Date 2021/3/1 13:14:10
 * @Param
 * @return
 **/

type VersionListQueryResult struct {
	Response VersionListQueryResponse `json:"alipay_open_mini_version_list_query_response"`
	Sign     string                   `json:"sign"`
}

type VersionListQueryResponse struct {
	data.AlipayResponse
	AppVersionInfos AppVersionInfos `json:"app_version_infos"`
	AppVersions     []string        `json:"app_versions"`
}

type AppVersionInfos struct {
	BundleId           string `json:"bundle_id"`
	AppVersion         string `json:"app_version"`
	VersionDescription string `json:"version_description,omitempty"`
	VersionStatus      string `json:"version_status"`
	CreateTime         string `json:"create_time"`
}

func (rest *Version) VersionListQuery(bizContent map[string]interface{}) VersionListQueryResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.mini.version.list.query", b)
	data := util.GetResult(params)

	result := VersionListQueryResult{}
	json.Unmarshal(data, &result)

	return result

}

/**
 * @Author return <1140444693@qq.com>
 * @Description // 查询构建版本
 * @Date 2021/2/21 13:17:49
 * @Param
 * @return
 **/
func (rest *Version) BuildQuery(bizContent map[string]interface{}) {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.mini.version.build.query", b)
	data := util.GetResult(params)

	fmt.Println(string(data))

}

/**
 * @Author return <1140444693@qq.com>
 * @Description // 删除版本
 * @Date 2021/2/21 13:46:49
 * @Param
 * @return
 **/

func (rest *Version) Delete(bizContent map[string]interface{}) {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.mini.version.delete", b)
	data := util.GetResult(params)

	fmt.Println(string(data))

}

/**
 * @Author return <1140444693@qq.com>
 * @Description // 生成商家小程序体验版
 * @Date 2021/2/21 13:38:49
 * @Param
 * @return
 **/

func (rest *Version) ExperienceCreate(bizContent map[string]interface{}) VersionExperienceCreateResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.mini.experience.create", b)
	data := util.GetResult(params)

	result := VersionExperienceCreateResult{}

	json.Unmarshal(data, &result)

	fmt.Println("result", result)

	return result

}

// 取消体验版
func (rest *Version) ExperienceCancel(bizContent map[string]interface{}) {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.mini.experience.cancel", b)
	data := util.GetResult(params)

	fmt.Println(string(data))

}

/**
 * @Author return <1140444693@qq.com>
 * @Description // 查看体验版状态
 * @Date 2021/2/21 13:55:49
 * @Param
 * @return
 **/
func (rest *Version) ExperienceQuery(bizContent map[string]interface{}) VersionExperienceQueryResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.mini.experience.query", b)
	data := util.GetResult(params)

	result := VersionExperienceQueryResult{}

	json.Unmarshal(data, &result)

	return result

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 上线小程序
 * @Date 2021/4/28 11:31:8
 * @Param
 * @return
 **/

type AlipayOpenMiniVersionOnlineResult struct {
	Response data.AlipayResponse `json:"response"`
	Sign     string              `json:"sign"`
}

func (rest *Version) Online(bizContent map[string]interface{}) AlipayOpenMiniVersionOnlineResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.mini.version.online", b)
	data := util.GetResult(params)

	result := AlipayOpenMiniVersionOnlineResult{}

	json.Unmarshal(data, &result)

	return result

}
