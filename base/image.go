/**
** @创建时间: 2022/6/26 09:01
** @作者　　: return
** @描述　　:
 */

package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/daifuyang/alipayEasySdkGo/data"
	"github.com/daifuyang/alipayEasySdkGo/util"
	"github.com/jinzhu/copier"
	"io"
	"mime/multipart"
	"os"
)

type Video struct {
	Image
}

type Image struct {
	data.Options
	AssetsParams
}

type AssetsParams struct {
	ImageType    string `json:"image_type" sign:"image_type"`
	ImageName    string `json:"image_name" sign:"image_name"`
	ImageContent string `json:"image_content" sign:"image_content"`
	ImagePid     string `json:"image_pid,omitempty" sign:"image_type,omitempty"`
}

type uploadResult struct {
	Response uploadResponse `json:"alipay_offline_material_image_upload_response"`
	data.Sign
}

type uploadResponse struct {
	data.PublicParams
	ImageId  string `json:"image_id"`
	ImageUrl string `json:"image_url"`
}

func (rest *Image) Upload(imageName string, imageFilePath string, imagePid ...string) (res uploadResult, err error) {
	res,err = rest.upload("jpg", imageName, imageFilePath, imagePid...)
	return
}

func (rest *Video) Upload(videoName string, videoFilePath string, imagePid ...string) (res uploadResult, err error) {
	res,err = rest.upload("mp4", videoName, videoFilePath, imagePid...)
	return
}

func (rest *Image) upload(imageType string, imageName string, imageFilePath string, imagePid ...string) (res uploadResult, err error) {
	config := data.GetOptions()
	options := new(Image)
	copier.Copy(&options, &config)
	options.AppAuthToken = rest.AppAuthToken
	// 组合请求参数
	options.Method = "alipay.offline.material.image.upload"

	options.ImageType = imageType
	options.ImageName = imageName

	for _, v := range imagePid {
		options.ImagePid = v
	}

	header := make(map[string]string, 0)
	// 文件上传
	bodyBuf := &bytes.Buffer{}
	bodyWrite := multipart.NewWriter(bodyBuf)

	file, err := os.Open(imageFilePath)
	defer file.Close()
	if err != nil {
		fmt.Println("加载文件失败", err)
		return
	}

	fileWrite, err := bodyWrite.CreateFormFile("image_content", imageFilePath)

	_, err = io.Copy(fileWrite, file)
	if err != nil {
		fmt.Println("io Copy error", err)
		return
	}

	header["Content-Type"] = bodyWrite.FormDataContentType()
	bodyWrite.Close()

	data, err := util.Post(options, util.WithBody(bodyBuf), util.WithHeader(header))
	if err != nil {
		fmt.Println("err", err)
		return
	}

	json.Unmarshal(data, &res)

	return

}
