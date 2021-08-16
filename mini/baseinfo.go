/**
** @创建时间: 2021/2/21 9:03 下午
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

type BaseInfo struct {
}

type QueryResult struct {
	Response QueryResponse `json:"alipay_open_mini_baseinfo_query_response"`
	Sign     string        `json:"sign"`
}

type QueryResponse struct {
	data.AlipayResponse
	AppName        string   `json:"app_name"`
	AppEnglishName string   `json:"app_english_name"`
	AppSlogan      string   `json:"app_slogan"`
	AppLogo        string   `json:"app_logo"`
	CategoryNames  string   `json:"category_names"`
	AppDesc        string   `json:"app_desc"`
	ServicePhone   string   `json:"service_phone"`
	ServiceEmail   string   `json:"service_email"`
	SafeDomains    []string   `json:"safe_domains,omitempty"`
	PackageNames   []string `json:"package_names,omitempty"`
}

func (rest *BaseInfo) Query(bizContent map[string]interface{}) QueryResult {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.mini.baseinfo.query", b)
	data := util.GetResult(params)

	fmt.Println(string(data))

	result := QueryResult{}

	json.Unmarshal(data, &result)

	return result

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 基本信息修改
 * @Date 2021/5/25 10:11:43
 * @Param
 * @return
 **/

func (rest *BaseInfo) Modify(bizContent map[string]interface{},appLogo string)  {

	var b string = ""
	if len(bizContent) > 0 {
		bytes, _ := json.Marshal(bizContent)
		b = string(bytes)
	}

	params := util.GetParams("alipay.open.mini.baseinfo.modify", b)

	header := make(map[string]string, 0)

	// 文件上传
	bodyBuf := &bytes.Buffer{}
	bodyWrite := multipart.NewWriter(bodyBuf)

	file, err := os.Open(appLogo)
	defer file.Close()
	if err != nil {
		fmt.Println("加载文件失败", err)
	}

	fileWrite, err := bodyWrite.CreateFormFile("app_logo", appLogo)

	_, err = io.Copy(fileWrite, file)
	if err != nil {
		fmt.Println("io Copy error", err)
	}

	header["Content-Type"] = bodyWrite.FormDataContentType()
	bodyWrite.Close()
	data := util.GetUploadResult(params, bodyBuf, header)


	fmt.Println(string(data))

}
