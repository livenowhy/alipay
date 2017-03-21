package wxpay

type WxConfig struct {
	AppId         string // 应用程序Id, 从https://open.weixin.qq.com上可以看得到
	AppKey        string // API密钥, 在 商户平台->账户设置->API安全 中设置
	MchId         string // 商户号
	NotifyUrl     string
	PlaceOrderUrl string
	QueryOrderUrl string
	TradeType     string
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
