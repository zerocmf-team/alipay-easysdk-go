/**
** @创建时间: 2022/6/21 23:28
** @作者　　: return
** @描述　　:
 */

package util

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"time"
)

/**
 * @Author return <1140444693@qq.com>
 * @Description 序列化参数并进行签名操作
 * @Date 2022/6/21 22:20:56
 * @Param
 * @return
 **/

func EncodeAndSign(merchantPrivateKey string, params map[string]string) (encode string) {
	unix := time.Now().Unix() // 时间戳
	time := time.Unix(unix, 0).Format("2006-01-02 15:04:05")
	params["timestamp"] = time
	sign, _ := GenerateSign(merchantPrivateKey, params)
	params["sign"] = sign
	// 获取提交的参数列表
	urls := url.Values{}
	for k, v := range params {
		value := strings.TrimSpace(v)
		if string(value) != "" {
			urls.Set(k, string(value))
		}
	}
	encode = urls.Encode()
	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 对配置项进行签名操作
 * @Date 2022/6/21 22:23:19
 * @Param
 * @return
 **/

func GenerateSign(merchantPrivateKey string, params map[string]string) (sign string, err error) {
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
		v := params[k]
		v = strings.TrimSpace(v)
		if string(v) != "" {
			pStr = append(pStr, k+"="+params[k])
		}
	}

	block := []byte(merchantPrivateKey)
	blocks, _ := pem.Decode(block)
	privateKey, err := x509.ParsePKCS8PrivateKey(blocks.Bytes)
	if err != nil {
		return
	}

	encode := strings.Join(pStr, "&")
	h := sha256.New()
	h.Write([]byte(encode))
	digest := h.Sum(nil)
	s, _ := rsa.SignPKCS1v15(nil, privateKey.(*rsa.PrivateKey), crypto.SHA256, digest)
	sign = base64.StdEncoding.EncodeToString(s)
	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 递归结构体转map[string]string
 * @Date 2022/6/22 0:52:15
 * @Param
 * @return
 **/

func ReflectPtr(ptr interface{}, tag string) (json map[string]string) {

	if json == nil {
		json = make(map[string]string, 0)
	}

	t := reflect.TypeOf(ptr)
	v := reflect.ValueOf(ptr)

	kd := t.Kind()

	if kd == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	num := v.NumField()

	for i := 0; i < num; i++ {
		if v.Field(i).Kind() == reflect.Struct || (v.Field(i).Kind() == reflect.Ptr && v.Field(i).Elem().Kind() == reflect.Struct) {
			jsonChildren := ReflectPtr(v.Field(i).Interface(), "sign")
			for k, v := range jsonChildren {
				json[k] = v
			}
		} else {
			tags := t.Field(i).Tag.Get(tag)
			if tags == "" {
				continue
			}
			tagArr := strings.Split(tags, ",")
			tag := tagArr[0]
			omitempty := false
			if len(tagArr) > 1 {
				if tagArr[1] == "omitempty" {
					omitempty = true
				}
			}
			val := v.Field(i).String()
			if omitempty == true && val == "" {
				continue
			}
			json[tag] = val
		}
	}
	return
}
