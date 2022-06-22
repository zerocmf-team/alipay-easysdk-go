/**
** @创建时间: 2022/6/21 21:33
** @作者　　: return
** @描述　　:
 */

package utils

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

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

