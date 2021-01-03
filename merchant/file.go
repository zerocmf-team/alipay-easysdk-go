/**
** @创建时间: 2020/12/20 10:29 下午
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

type File struct{}

type uploadResponse struct {
	data.AlipayResponse
	MaterialId  string `json:"material_id"`
	MaterialKey string `json:"material_key"`
}

type UploadResult struct {
	Response uploadResponse `json:"alipay_merchant_item_file_upload_response"`
	Sign     string         `json:"sign"`
}

func (rest *File) Upload(bizContent map[string]string, filename string) (UploadResult, error) {

	params := util.GetParams("alipay.merchant.item.file.upload", "")

	for k, v := range params {
		bizContent[k] = v
	}

	imageType := strings.Split(filename, ".")

	if len(imageType) < 2 {
		return UploadResult{}, errors.New("图片格式非法")
	}

	bizContent["scene"] = "SYNC_ORDER"
	header := make(map[string]string, 0)

	// 文件上传
	bodyBuf := &bytes.Buffer{}
	bodyWrite := multipart.NewWriter(bodyBuf)

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		fmt.Println("加载文件失败", err)
		return UploadResult{}, err
	}

	fileWrite, err := bodyWrite.CreateFormFile("file_content", filename)

	_, err = io.Copy(fileWrite, file)
	if err != nil {
		fmt.Println("io Copy error", err)
		return UploadResult{}, err
	}

	header["Content-Type"] = bodyWrite.FormDataContentType()
	bodyWrite.Close()
	data := util.GetUploadResult(bizContent, bodyBuf, header)

	r := UploadResult{}
	json.Unmarshal(data, &r)
	return r, nil
}
