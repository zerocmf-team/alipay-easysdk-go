/**
** @创建时间: 2021/2/21 2:42 下午
** @作者　　: return
** @描述　　:
 */
package mini

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gincmf/alipayEasySdk/data"
	"github.com/gincmf/alipayEasySdk/util"
	"io"
	"mime/multipart"
	"os"
)

type Audit struct{}

type DetailQueryResult struct {
	Response DetailQueryResponse `json:"alipay_open_mini_version_detail_query_response"`
	Sign     string              `json:"sign"`
}

type AppCategoryInfoList struct {
	FirstCategoryId    string `json:"first_category_id"`
	FirstCategoryName  string `json:"first_category_name"`
	SecondCategoryId   string `json:"second_category_id"`
	SecondCategoryName string `json:"second_category_name"`
	ThirdCategoryId    string `json:"third_category_id"`
	ThirdCategoryName  string `json:"third_category_name"`
}

type DetailQueryResponse struct {
	data.AlipayResponse
	AppVersion              string                `json:"app_version"`
	AppName                 string                `json:"app_name"`
	AppEnglishName          string                `json:"app_english_name"`
	AppLogo                 string                `json:"app_logo"`
	VersionDesc             string                `json:"version_desc"`
	GrayStrategy            string                `json:"gray_strategy,omitempty"`
	Status                  string                `json:"status"`
	RejectReason            string                `json:"reject_reason,omitempty"`
	ScanResult              string                `json:"scan_result"`
	GmtCreate               string                `json:"gmt_create"`
	MiniAppCategoryInfoList []AppCategoryInfoList `json:"mini_app_category_info_list"`
}

func (rest *Audit) DetailQuery(bizContent map[string]interface{}) DetailQueryResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.mini.version.detail.query", b)
	data := util.GetResult(params)

	result := DetailQueryResult{}
	json.Unmarshal(data, &result)

	return result

}

type AuditedCancelResult struct {
	Response data.AlipayResponse `json:"alipay_open_mini_version_audited_cancel_response"`
	Sign     string              `json:"sign"`
}

func (rest *Audit) AuditedCancel(bizContent map[string]interface{}) AuditedCancelResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.mini.version.audited.cancel", b)
	data := util.GetResult(params)

	result := AuditedCancelResult{}
	json.Unmarshal(data, &result)

	return result

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 提审小程序
 * @Date 2021/3/2 18:0:53
 * @Param
 * @return
 **/

type ApplyResult struct {
	Response data.AlipayResponse `json:"alipay_open_mini_version_audit_apply_response"`
	Sign     string              `json:"sign"`
}

func (rest *Audit) Apply(bizContent map[string]string, files map[string]string) (ApplyResult, error) {

	params := util.GetParams("alipay.open.mini.version.audit.apply", "")
	for k, v := range params {
		bizContent[k] = v
	}

	header := make(map[string]string, 0)

	// 文件上传
	bodyBuf := &bytes.Buffer{}
	bodyWrite := multipart.NewWriter(bodyBuf)

	for k, v := range files {
		file, err := os.Open(v)
		defer file.Close()
		if err != nil {
			fmt.Println("加载文件失败", err)
			return ApplyResult{}, err
		}

		fileWrite, err := bodyWrite.CreateFormFile(k, v)

		_, err = io.Copy(fileWrite, file)
		if err != nil {
			fmt.Println("io Copy error", err)
			return ApplyResult{}, err
		}

	}

	header["Content-Type"] = bodyWrite.FormDataContentType()
	bodyWrite.Close()
	data := util.GetUploadResult(bizContent, bodyBuf, header)

	fmt.Println("data", string(data))

	result := ApplyResult{}
	json.Unmarshal(data, &result)

	return result, nil
}
