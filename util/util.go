/**
** @创建时间: 2020/9/7 9:46 上午
** @作者　　: return
** @描述　　:
 */
package util

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/gincmf/alipayEasySdk"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//封装请求库
func Request(method string, url string, body io.Reader, header map[string]string) (int, []byte) {
	client := &http.Client{}
	switch method {
	case "get", "GET":
		method = "GET"
	case "post", "POST":
		method = "POST"
	case "put", "PUT":
		method = "PUT"
	case "delete", "DELETE":
		method = "POST"
	}
	r, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println("http错误", err)
	}

	r.Header.Add("Host", "")
	r.Header.Add("Connection", "keep-alive")
	r.Header.Add("Accept-Encoding", "gzip, deflate, br")
	r.Header.Add("Content-Length", "0")
	r.Header.Add("Cache-Control", "no-cache")
	for k, v := range header {
		r.Header.Add(k, v)
	}

	response, err := client.Do(r)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer response.Body.Close()

	var data []byte = nil

	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ := gzip.NewReader(response.Body)
		for {
			buf := make([]byte, 1024)
			n, err := reader.Read(buf)
			if err != nil && err != io.EOF {
				panic(err)
			}
			if n == 0 {
				break
			}
			data = append(data, buf...)
		}
	default:
		data, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("err", err.Error())
		}
	}

	index := bytes.IndexByte(data, 0)

	if index > 0 {
		data = data[:index]
	}

	return response.StatusCode, data
}

// 自动获取序列话签名操作
func EncodeParams(paramsMap map[string]string) string {
	// 获取签名
	sign, _ := Sign(paramsMap)
	paramsMap["sign"] = sign // 追加参数

	// 获取提交的参数列表
	tempParams := url.Values{}
	for keys, value := range paramsMap {
		v := []byte(value)
		v = bytes.TrimSpace(v)
		if string(v) != "" {
			tempParams.Set(keys, value)
		}
	}

	paramsEncode := tempParams.Encode()
	return paramsEncode
}

func GetParams(method string, b string) map[string]string {
	op := alipayEasySdk.Options()
	paramsMap := make(map[string]string, 0)
	paramsMap["method"] = method
	if b != "" {
		paramsMap["biz_content"] = b
	}

	if op.AppAuthToken != "" {
		paramsMap["app_auth_token"] = op.AppAuthToken
	}

	return paramsMap
}

// 封装统一的支付宝签名和公共参数
/*
	date: 2020-09-07
	Author: frank_dai
	desc: 通过授权码获取商户信息
*/
func GetResult(paramsMap map[string]string) []byte {
	return request(paramsMap, nil, nil)
}

func GetAppIdResult(paramsMap map[string]string) []byte {
	paramsMap["noAuthToken"] = "1"
	return request(paramsMap, nil, nil)
}

func GetRequestUrl(paramsMap map[string]string) string {

	options := alipayEasySdk.Options()

	if paramsMap["noAuthToken"] == "1" {
		delete(paramsMap, "noAuthToken")
		delete(paramsMap, "app_auth_token")
	}

	unix := time.Now().Unix() // 时间戳
	time := time.Unix(unix, 0).Format("2006-01-02 15:04:05")
	protocol := options.Protocol   // 协议
	baseUrl := options.GatewayHost // 网关

	paramsMap["app_id"] = options.AppId
	paramsMap["notify_url"] = options.NotifyUrl
	if paramsMap["notify"] != "" {
		options.NotifyUrl = paramsMap["notify"]
	}
	paramsMap["timestamp"] = time
	paramsMap["sign_type"] = options.SignType
	paramsMap["charset"] = options.Charset
	paramsMap["version"] = options.Version

	/*if options.EncryptType != "" {
		paramsMap["encrypt_type"] = options.EncryptType
	}*/

	paramsEncode := EncodeParams(paramsMap)

	url := protocol + "://" + baseUrl + "?" + paramsEncode
	return url
}

func GetUploadResult(paramsMap map[string]string, body io.Reader, header map[string]string) []byte {
	return request(paramsMap, body, header)
}

func request(paramsMap map[string]string, body io.Reader, header map[string]string) []byte {

	url := GetRequestUrl(paramsMap)
	encode := strings.Split(url,"?")
	if body == nil {
		body = strings.NewReader(encode[1])
	}
	_, data := Request("POST", url, body, header)
	return data

}

// AES加密
//PKCS7 填充模式
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个,然后合并成新的字节切片返回
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//PKCS7填充的反向操作，删除填充的字符串
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	paddLen := int(origData[length-1])

	end := length - paddLen

	text := origData[:length]
	if end > 0 {
		text = origData[:end]
	}

	return text
}

//aes加密操作
func AesEncrypt(origData []byte, key []byte) ([]byte, error) {

	key, _ = base64.StdEncoding.DecodeString(string(key))

	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块大小
	blockSize := block.BlockSize()
	//对数据进行填充，让数据长度满足需求
	origData = PKCS7Padding(origData, blockSize)

	iv := make([]byte, blockSize)

	//采用AES加密方法中的CBC加密模式
	blocMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	//执行加密
	blocMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//aes解密操作
func AesDeCrypt(cypted []byte, key []byte) ([]byte, error) {

	key, _ = base64.StdEncoding.DecodeString(string(key))

	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	iv := make([]byte, blockSize)

	//采用AES加密方法中的CBC加密模式 创建加密客户端实例
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cypted))
	//这个函数还可以用来解密
	blockMode.CryptBlocks(origData, cypted)
	//去除填充字符串
	origData = PKCS7UnPadding(origData)
	return origData, err
}
