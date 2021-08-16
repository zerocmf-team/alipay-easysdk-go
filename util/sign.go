/**
** @创建时间: 2020/9/7 9:36 上午
** @作者　　: return
** @描述　　:
 */
package util

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/gincmf/alipayEasySdk"
	"sort"
	"strings"
)

// 对参数签名，获取签名参数
func Sign(params map[string]string) (sign string, encode string) {

	//ksort 对参数进行排序
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 对参数进行序列化
	pStr := make([]string, 0)
	//拼接
	for _, k := range keys {
		v := []byte(params[k])
		v = bytes.TrimSpace(v)
		if string(v) != "" {
			pStr = append(pStr, k+"="+params[k])
		}
	}

	// 序列化结果
	encode = strings.Join(pStr, "&")

	h := sha256.New()
	h.Write([]byte(encode))
	// hashed := h.Sum(nil)
	// 加密生成sign
	// block 私钥

	options := alipayEasySdk.Options()
	block := []byte(options.PrivateKey)

	blocks, _ := pem.Decode(block)
	privateKey, err := x509.ParsePKCS8PrivateKey(blocks.Bytes)
	if err != nil {
		fmt.Println("err",err.Error())
		return "",""
	}

	digest := h.Sum(nil)
	s, _ := rsa.SignPKCS1v15(nil, privateKey.(*rsa.PrivateKey), crypto.SHA256, digest)
	sign = base64.StdEncoding.EncodeToString(s)

	return sign, encode
}

// 验参
func VerifySign(params string, sign string) error {
	return vSign(params,sign,"")
}

func AliVerifySign(params string, sign string) error {
	return vSign(params,sign,"alipay")
}

func vSign(params string, sign string,t string) error {

	options := alipayEasySdk.Options()

	block := []byte(options.PublicKey)
	if t =="alipay" {
		block = []byte(options.AliPublicKey)
	}
	signByte, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		fmt.Println("sign err",err.Error())
	}

	blocks, _ := pem.Decode(block)
	pub, err := x509.ParsePKIXPublicKey(blocks.Bytes)
	if err != nil {
		fmt.Println("err", err.Error())
		return err
	}

	h := sha256.New()
	h.Write([]byte(params))

	digest := h.Sum(nil)

	err = rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA256, digest, signByte)

	if err != nil {
		fmt.Println("err", err)
	}

	return err

}

func SortParam (params map[string]string ) string {

	//ksort 对参数进行排序
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 对参数进行序列化
	pStr := make([]string, 0)
	//拼接
	for _, k := range keys {
		v := []byte(params[k])
		v = bytes.TrimSpace(v)
		if string(v) != "" {
			pStr = append(pStr, k+"="+params[k])
		}
	}

	// 序列化结果
	encode := strings.Join(pStr, "&")

	return encode
}
