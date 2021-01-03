/**
** @创建时间: 2020/12/9 10:19 上午
** @作者　　: return
** @描述　　:
 */
package merchant

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gincmf/alipayEasySdk/data"
	"github.com/gincmf/alipayEasySdk/util"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

type Image struct{}

type alipayResponse struct {
	data.AlipayResponse
	ImageId    string `json:"image_id"`
}

type Result struct {
	Response alipayResponse `json:"ant_merchant_expand_indirect_image_upload_response"`
	Sign     string         `json:"sign"`
}

func (rest *Image) Upload(bizContent map[string]string, filename string) (Result,error) {

	params := util.GetParams("ant.merchant.expand.indirect.image.upload", "")

	for k, v := range params {
		bizContent[k] = v
	}

	imageType := strings.Split(filename,".")

	if len(imageType) < 2 {
		return Result{},errors.New("图片格式非法")
	}

	bizContent["image_type"] = imageType[1]

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
	err = json.Unmarshal(data, &r)
	if err != nil {
		fmt.Println(err.Error())
	}
	return r,nil
}
