/**
** @创建时间: 2020/12/9 10:19 上午
** @作者　　: return
** @描述　　:
 */
package base

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

type Image struct{}

type AlipayResponse struct {
	data.AlipayResponse
	ImageId    string `json:"image_id"`
	ImageUrl   string `json:"image_url"`
}

type Result struct {
	Response AlipayResponse `json:"alipay_offline_material_image_upload_response"`
	Sign     string         `json:"sign"`
}

func (rest *Image) Upload(bizContent map[string]string, filename string) (Result,error) {

	params := util.GetParams("alipay.offline.material.image.upload", "")

	for k, v := range params {
		bizContent[k] = v
	}

	header := make(map[string]string, 0)

	// 文件上传
	bodyBuf := &bytes.Buffer{}
	bodyWrite := multipart.NewWriter(bodyBuf)

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		fmt.Println("加载文件失败", err)
		return Result{},err
	}

	fileWrite, err := bodyWrite.CreateFormFile("image_content", filename)

	_, err = io.Copy(fileWrite, file)
	if err != nil {
		fmt.Println("io Copy error", err)
		return Result{},err
	}

	header["Content-Type"] = bodyWrite.FormDataContentType()
	bodyWrite.Close()
	data := util.GetUploadResult(bizContent, bodyBuf, header)

	r := Result{}
	json.Unmarshal(data, &r)
	return r,nil
}
