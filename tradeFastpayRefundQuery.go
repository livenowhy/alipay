package alipay


// https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.chC7PJ&docType=4&apiId=1049
// alipay.trade.fastpay.refund.query (统一收单交易退款查询)
// 商户可使用该接口查询自已通过alipay.trade.refund提交的退款请求是否执行成功。
type AliPayFastpayTradeRefundQuery struct {
	AppAuthToken string `json:"-"`                      // 可选
	OutTradeNo   string `json:"out_trade_no,omitempty"` // 与 TradeNo 二选一
	TradeNo      string `json:"trade_no,omitempty"`     // 与 OutTradeNo 二选一
	OutRequestNo string `json:"out_request_no"`         // 必须 请求退款接口时，传入的退款请求号，如果在退款请求时未传入，则该值为创建交易时的外部交易号
}

func (this AliPayFastpayTradeRefundQuery) APIName() string {
	return "alipay.trade.fastpay.refund.query"
}

func (this AliPayFastpayTradeRefundQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this AliPayFastpayTradeRefundQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayFastpayTradeRefundQuery) ExtJSONParamValue() string {
	return marshal(this)
}

type AliPayFastpayTradeRefundQueryResponse struct {
	AliPayTradeFastpayRefundQueryResponse struct {
		Code         string `json:"code"`
		Msg          string `json:"msg"`
		SubCode      string `json:"sub_code"`
		SubMsg       string `json:"sub_msg"`
		OutRequestNo string `json:"out_request_no"` // 本笔退款对应的退款请求号
		OutTradeNo   string `json:"out_trade_no"`   // 创建交易传入的商户订单号
		RefundReason string `json:"refund_reason"`  // 发起退款时，传入的退款原因
		TotalAmount  string `json:"total_amount"`   // 发该笔退款所对应的交易的订单金额
		RefundAmount string `json:"refund_amount"`  // 本次退款请求，对应的退款金额
		TradeNo      string `json:"trade_no"`       // 支付宝交易号
	} `json:"alipay_trade_fastpay_refund_query_response"`
	Sign string `json:"sign"`
}

func (this *AliPayFastpayTradeRefundQueryResponse) IsSuccess() bool {
	if this.AliPayTradeFastpayRefundQueryResponse.Msg == "Success" {
		return true
	}
	return false
}


