/**
** @创建时间: 2022/6/21 21:33
** @作者　　: return
** @描述　　:
 */

package util

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/daifuyang/alipayEasySdkGo/data"
	"github.com/jinzhu/copier"
	"io"
	"io/ioutil"
	"net/http"
)

type Options struct {
	f func(*reqOptions)
}

type reqOptions struct {
	body   io.Reader
	header map[string]string
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 请求库封装
 * @Date 2022/6/21 21:33:31
 * @Param
 * @return
 **/

func Request(method string, url string, body io.Reader, header map[string]string) (response []byte, err error) {
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

	res, err := client.Do(r)

	if err != nil {
		return
	}

	defer res.Body.Close()

	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ := gzip.NewReader(res.Body)
		for {
			buf := make([]byte, 1024)
			n, err := reader.Read(buf)
			if err != nil && err != io.EOF {
				panic(err)
			}
			if n == 0 {
				break
			}
			response = append(response, buf...)
		}
	default:
		response, err = ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("err", err.Error())
		}
	}

	index := bytes.IndexByte(response, 0)

	if index > 0 {
		response = response[:index]
	}
	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 统一封装post请求
 * @Date 2022/6/25 16:18:51
 * @Param
 * @return
 **/

func Post(options interface{}, ops ...Options) (resp []byte, err error) {

	params := ReflectPtr(options, "sign")
	dOps := data.GetOptions()
	copier.Copy(&dOps, &options)
	encode := EncodeAndSign(dOps.MerchantPrivateKey, params)
	protocol := dOps.Protocol   // 协议
	baseUrl := dOps.GatewayHost // 网关
	url := protocol + "://" + baseUrl + "?" + encode

	reqOps := new(reqOptions)
	for _, v := range ops {
		v.f(reqOps)
	}

	body := reqOps.body
	header := reqOps.header

	resp, err = Request("POST", url, body, header)
	return
}

func WithBody(body io.Reader) Options {
	return Options{func(o *reqOptions) {
		o.body = body
	}}
}

func WithHeader(header map[string]string) Options {
	return Options{func(o *reqOptions) {
		o.header = header
	}}
}