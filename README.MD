欢迎使用 Alipay **Easy** SDK Go。本SDK非官方实现  
打造**最好用**的支付宝开放平台**Golang服务端SDK**
遵循官方的设计理念，化繁为简

dev1.0已完成，master版本正在重构中。代码更加规范，享受极简编程

## 代码示例
* Alipay Easy Go SDK :smiley:

```
faceToFace := new(payment.FaceToFace)
faceToFace.Agent("202106BB21c6e8a21c7846caa2c39362a20ccX15")

bizContent := make(map[string]interface{}, 0)
bizContent["out_trade_no"] = "20240320010101029"
bizContent["subject"] = "test"
bizContent["total_amount"] = "0.01"
bizContent["trans_currency"] = "CNY"
bizContent["settle_currency"] = "CNY"

// 订单信息

goodsDetail := make(map[string]interface{}, 0)
goodsDetail["goods_id"] = "apple-01"
goodsDetail["goods_name"] = "iphone12"
goodsDetail["quantity"] = 1
goodsDetail["price"] = 2

bizContent["goods_detail"] = []map[string]interface{}{goodsDetail}

result, err := faceToFace.PreCrete(bizContent)
if err != nil {
	fmt.Println("result", result)
}
```

<a name="apiList"/>

## 支持的API列表

| 开发状态      | 能力类别      | 场景类别            | 接口方法名称                 | 调用的OpenAPI名称                                              |
|-----------|-----------|-----------------|------------------------|-----------------------------------------------------------|
|  已完成 | Base<br/>基础能力      | OAuth<br/>用户授权           | getToken<br/>获取授权访问令牌和用户user_id               | alipay\.system\.oauth\.token                              |
|  已完成 | Base<br/>基础能力     | OAuth<br/>用户授权           | refreshToken<br/>刷新授权访问令牌           | alipay\.system\.oauth\.token                              |
|  已完成 | Base<br/>基础能力      | Qrcode<br/>小程序二维码          | create<br/>创建小程序二维码                 | alipay\.open\.app\.qrcode\.create                         |
|  已完成 | Base<br/>基础能力      | Image<br/>图片           | upload<br/>上传门店照片                 | alipay\.offline\.material\.image\.upload                  |
|  已完成 | Base<br/>基础能力      | Video<br/>视频           | upload<br/>上传门店视频                 | alipay\.offline\.material\.image\.upload                  |
| * [ ] 开发中... | Member<br/>会员能力    | Identification<br/>支付宝身份认证  | init<br/>身份认证初始化                   | alipay\.user\.certify\.open\.initialize                   |
| * [ ] 开发中... | Member<br/>会员能力    | Identification<br/>支付宝身份认证  | certify<br/>生成认证链接                | alipay\.user\.certify\.open\.certify                      |
| * [ ] 开发中... | Member<br/>会员能力    | Identification<br/>支付宝身份认证  | query<br/>身份认证记录查询                  | alipay\.user\.certify\.open\.query                        |
| 已完成 | Payment<br/>支付能力   | Common<br/>通用          | create<br/>创建交易                 | alipay\.trade\.create                                     |
| 已完成 | Payment<br/>支付能力   | Common<br/>通用          | query<br/>查询交易                  | alipay\.trade\.query                                      |
| 已完成 | Payment<br/>支付能力   | Common<br/>通用          | refund<br/>交易退款                 | alipay\.trade\.refund                                     |
| 已完成 | Payment<br/>支付能力   | Common<br/>通用          | close<br/>关闭交易                  | alipay\.trade\.close                                      |
| 已完成 | Payment<br/>支付能力   | Common<br/>通用          | cancel<br/>撤销交易                 | alipay\.trade\.cancel                                     |
| 已完成 | Payment<br/>支付能力   | Common<br/>通用          | queryRefund<br/>交易退款查询            | alipay\.trade\.fastpay\.refund\.query                     |
| 已完成 | Payment<br/>支付能力   | Common<br/>通用          | downloadBill<br/>查询对账单下载地址           | alipay\.data\.dataservice\.bill\.downloadurl\.query       |
| 已完成 | Payment<br/>支付能力   | Common<br/>通用          | verifyNotify<br/>异步通知验签           | -                                                         |
| 已完成 | Payment<br/>支付能力   | Huabei<br/>花呗分期          | create<br/>创建花呗分期交易                 | alipay\.trade\.create                                     |
| 已完成 | Payment<br/>支付能力   | FaceToFace<br/>当面付      | pay<br/>扫用户出示的付款码，完成付款                    | alipay\.trade\.pay                                        |
| 已完成 | Payment<br/>支付能力   | FaceToFace<br/>当面付      | precreate<br/>生成交易付款码，待用户扫码付款              | alipay\.trade\.precreate                                  |
| * [ ] 开发中... | Payment<br/>支付能力   | App<br/>手机APP             | pay<br/>生成订单串，再使用客户端 SDK 凭此串唤起支付宝收银台                    | alipay\.trade\.app\.pay                                   |
| * [ ] 开发中... | Payment<br/>支付能力   | Page<br/>电脑网站            | pay<br/>生成交易表单，渲染后自动跳转支付宝网站引导用户完成支付                    | alipay\.trade\.page\.pay                                  |
| * [ ] 开发中... | Payment<br/>支付能力   | Wap<br/>手机网站             | pay<br/>生成交易表单，渲染后自动跳转支付宝网站引导用户完成支付                    | alipay\.trade\.wap\.pay                                   |
| * [ ] 开发中... | Security<br/>安全能力  | TextRisk<br/>文本内容安全        | detect<br/>检测内容风险                 | alipay\.security\.risk\.content\.detect                   |
| * [ ] 开发中... | Marketing<br/>营销能力 | Pass<br/>支付宝卡包            | createTemplate<br/>卡券模板创建         | alipay\.pass\.template\.add                               |
| * [ ] 开发中... | Marketing<br/>营销能力 | Pass<br/>支付宝卡包            | updateTemplate<br/>卡券模板更新         | alipay\.pass\.template\.update                            |
| * [ ] 开发中... | Marketing<br/>营销能力 | Pass<br/>支付宝卡包            | addInstance<br/>卡券实例发放            | alipay\.pass\.instance\.add                               |
| * [ ] 开发中... | Marketing<br/>营销能力 | Pass<br/>支付宝卡包            | updateInstance<br/>卡券实例更新         | alipay\.pass\.instance\.update                            |
| * [ ] 开发中... | Marketing<br/>营销能力 | TemplateMessage<br/>小程序模板消息 | send <br/>发送模板消息| alipay\.open\.app\.mini\.templatemessage\.send            |
| * [ ] 开发中... | Marketing<br/>营销能力 | OpenLife<br/>生活号        | createImageTextContent<br/>创建图文消息内容 | alipay\.open\.public\.message\.content\.create            |
| * [ ] 开发中... | Marketing<br/>营销能力 | OpenLife<br/>生活号        | modifyImageTextContent<br/>更新图文消息内容 | alipay\.open\.public\.message\.content\.modify            |
| * [ ] 开发中... | Marketing<br/>营销能力 | OpenLife<br/>生活号        | sendText<br/>群发本文消息               | alipay\.open\.public\.message\.total\.send                |
| * [ ] 开发中... | Marketing<br/>营销能力 | OpenLife<br/>生活号        | sendImageText<br/>群发图文消息          | alipay\.open\.public\.message\.total\.send                |
| * [ ] 开发中... | Marketing<br/>营销能力 | OpenLife<br/>生活号        | sendSingleMessage<br/>单发模板消息      | alipay\.open\.public\.message\.single\.send               |
| * [ ] 开发中... | Marketing<br/>营销能力 | OpenLife<br/>生活号        | recallMessage<br/>生活号消息撤回          | alipay\.open\.public\.life\.msg\.recall                   |
| * [ ] 开发中... | Marketing<br/>营销能力 | OpenLife<br/>生活号        | setIndustry<br/>模板消息行业设置            | alipay\.open\.public\.template\.message\.industry\.modify |
| * [ ] 开发中... | Marketing<br/>营销能力 | OpenLife<br/>生活号        | getIndustry<br/>生活号查询行业设置            | alipay\.open\.public\.setting\.category\.query            |
| * [ ] 开发中... | Util<br/>辅助工具      | AES<br/>加解密             | decrypt<br/>解密，常用于会员手机号解密                | -                                                         |
| * [ ] 开发中... | Util<br/>辅助工具      | AES<br/>加解密             | encrypt<br/>加密                | -                                                         |
| * [ ] 开发中... | Util<br/>辅助工具      | Generic<br/>通用接口         | execute<br/>自行拼接参数，执行OpenAPI调用                | -                                                         |

> 注：更多高频场景的API持续更新中，敬请期待。

 # 变更日志

 # 相关
* [支付宝开放平台](https://open.alipay.com/platform/home.htm)
* [支付宝开放平台文档中心](https://docs.open.alipay.com/catalog)

# 交流与技术支持
请通过github本身的issue功能进行反馈。