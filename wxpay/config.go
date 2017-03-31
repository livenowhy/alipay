package wxpay

type WxConfig struct {
	AppId               string // 应用程序Id, 从https://open.weixin.qq.com上可以看得到
	AppKey              string // API密钥, 在 商户平台->账户设置->API安全 中设置
	MchId               string // 商户号

	PlaceOrderUrl       string // 统一下单
	QueryOrderUrl       string // 查询订单
	CloseOrderUrl       string // 关闭订单
	RefundOrderUrl      string // 申请退款
	RefundQueryOrderUrl string // 查询退款
	DownloadBillUrl     string // 下载对账单
	ReportUrl           string // 交易保障
	NotifyUrl           string // 支付结果通用通知

	TradeType           string
}

type PaymentRequest struct {
	AppId     string
	PartnerId string
	PrepayId  string
	Package   string
	NonceStr  string
	Timestamp string
	Sign      string
}

// 1. 统一下单   https://api.mch.weixin.qq.com/pay/unifiedorder
// 2. 查询订单   https://api.mch.weixin.qq.com/pay/orderquery
// 3. 关闭订单   https://api.mch.weixin.qq.com/pay/closeorder
// 4. 申请退款   https://api.mch.weixin.qq.com/secapi/pay/refund
// 5. 查询退款   https://api.mch.weixin.qq.com/pay/refundquery
// 6. 下载对账单 https://api.mch.weixin.qq.com/pay/downloadbill
// 7. 支付结果通用通知 notify_url: "https://pay.weixin.qq.com/wxpay/pay.action"
// 8. 交易保障   https://api.mch.weixin.qq.com/payitil/report


// 1. 微信支付接口签名校验工具
// https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=20_1

// 2. 支付验收指引
// https://pay.weixin.qq.com/wiki/doc/api/native.php?chapter=23_1