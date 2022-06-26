/**
** @创建时间: 2022/6/26 09:01
** @作者　　: return
** @描述　　:
 */

package base

import (
	"bytes"
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

func (rest *Image) Upload(imageName string, imageFilePath string, imagePid ...string) {
	rest.upload("jpg", imageName, imageFilePath, imagePid...)
}

func (rest *Video) Upload(videoName string, videoFilePath string, imagePid ...string) {
	rest.upload("mp4", videoName, videoFilePath, imagePid...)
}

func (rest *Image) upload(imageType string, imageName string, imageFilePath string, imagePid ...string) {
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

	resp, err := util.Post(options, util.WithBody(bodyBuf), util.WithHeader(header))
	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println("resp", string(resp))

}
