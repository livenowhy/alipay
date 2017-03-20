package alipay

// https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.SPYWcR&docType=4&apiId=862
// alipay.trade.precreate (统一收单线下交易预创建)
type AliPayTradePrecreate struct {
	AppAuthToken         string             `json:"-"`                       // 可选
	OutTradeNo           string             `json:"out_trade_no,omitempty"`  // 必须;商户订单号
	SellerId             string             `json:"seller_id,omitempty"`     // 可选; 卖家支付宝用户ID
	TotalAmount          string             `json:"total_amount,omitempty"`  // 必须; 订单标题
	DiscountableAmount   string             `json:"discountable_amount"`     // 可打折金额. 参与优惠计算的金额
	UndiscountableAmount string             `json:"undiscountable_amount"`   // 不可打折金额. 不参与优惠计算的金额
	BuyerLogonId         string             `json:"buyer_logon_id"`          // 买家支付宝账号
	Subject              string             `json:"subject,omitempty"`       // 订单标题
	Body                 string             `json:"body,omitempty"`          // 对交易或商品的描述
	GoodsDetail          []*GoodsDetailItem `json:"goods_detail,omitempty"`  // 订单包含的商品列表信息.Json格式
	OperatorId           string             `json:"operator_id"`             // 商户操作员编号
	StoreId              string             `json:"store_id"`                // 商户操作员编号
	TerminalId           string             `json:"terminal_id"`             // 商户操作员编号
	ExtendParams         *ExtendParamsItem  `json:"extend_params,omitempty"` // 业务扩展参数
	TimeoutExpress       string             `json:"timeout_express"`         // 该笔订单允许的最晚付款时间
	RoyaltyInfo          *RoyaltyInfo       `json:"royalty_info"`            // 描述分账信息，json格式。
	SubMerchant          *SubMerchantItem   `json:"sub_merchant"`            // 二级商户信息
	AlipayStoreId        string             `json:"alipay_store_id"`         // 支付宝店铺的门店ID。
}

func (this AliPayTradePrecreate) APIName() string {
	return "alipay.trade.precreate"
}

func (this AliPayTradePrecreate) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this AliPayTradePrecreate) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayTradePrecreate) ExtJSONParamValue() string {
	return marshal(this)
}

type AliPayTradePrecreateResponse struct {
	AliPayTradePrecreate struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`

		OutTradeNo string `json:"out_trade_no"`
		QrCode     string `json:"qr_code"`
	} `json:"alipay_trade_precreate_response"`
	Sign string `json:"sign"`
}

func (this *AliPayTradePrecreateResponse) IsSuccess() bool {
	if this.AliPayTradePrecreate.Msg == "Success" {
		return true
	}
	return false
}
