/**
** @创建时间: 2020/9/5 10:51 下午
** @作者　　: return
** @描述　　:
 */
package alipayEasySdk

func main()  {

	options := map[string]string{
		"protocol":           "https",
		"gatewayHost":        "openapi.alipay.com/gateway.do",
		"signType":           "RSA2",
		"appId":              "2021001199695134",
		"version":            "1.0",
		"charset":            "utf-8",
		"merchantPrivateKey": "",
		"alipayCertPath":     "",
		"alipayRootCertPath": "",
		"merchantCertPath":   "",
		"notifyUrl":          "",
		"encryptKey":         "",
		"appAuthToken":       "202011BB7180118e46ed488489e442aba3047A61",
	}

	NewOptions(options)

}
