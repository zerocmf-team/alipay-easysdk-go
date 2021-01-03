/**
** @创建时间: 2020/12/11 10:56 上午
** @作者　　: return
** @描述　　:
 */
package marketing

import (
	"bytes"
	"fmt"
	"github.com/gincmf/alipayEasySdk/util"
	"io"
	"mime/multipart"
	"os"
)

type Image struct {

}


//
func (rest *Image) Upload (bizContent map[string]string,filename string) {

	params := util.GetParams("alipay.marketing.material.image.upload","")

	for k,v := range params{
		bizContent[k] = v
	}

	header := make(map[string]string,0)

	// 文件上传
	bodyBuf := &bytes.Buffer{}
	bodyWrite := multipart.NewWriter(bodyBuf)


	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		fmt.Println("加载文件失败", err)
		return
	}

	fileWrite, err := bodyWrite.CreateFormFile("file_content", filename)

	_,err = io.Copy(fileWrite,file)
	if err != nil {
		fmt.Println("io Copy error",err)
		return
	}

	header["Content-Type"] = bodyWrite.FormDataContentType()
	bodyWrite.Close()
	data := util.GetUploadResult(bizContent,bodyBuf,header)
	fmt.Println(data)
}
