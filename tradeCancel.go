package alipay

// https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.o1inS0&docType=4&apiId=866
// 支付交易返回失败或支付系统超时，调用该接口撤销交易。
// 如果此订单用户支付失败，支付宝系统会将此订单关闭；如果用户支付成功，支付宝系统会将此订单资金退还给用户。
// 注意：只有发生支付系统超时或者支付结果未知时可调用撤销，其他正常支付的单如需实现相同功能请调用申请退款API。
// 提交支付交易后调用【查询订单API】，没有明确的支付结果再调用【撤销订单API】
// alipay.trade.cancel (统一收单交易撤销接口)
type AliPayTradeCancel struct {
	AppAuthToken string `json:"-"`                      // 可选
	NotifyURL    string `json:"-"`                      // 可选
	OutTradeNo   string `json:"out_trade_no,omitempty"` // 原支付请求的商户订单号,和支付宝交易号不能同时为空
	TradeNo      string `json:"trade_no,omitempty"`     // 支付宝交易号，和商户订单号不能同时为空; 与 OutTradeNo 二选一
}

func (this AliPayTradeCancel) APIName() string {
	return "alipay.trade.cancel"
}

func (this AliPayTradeCancel) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["notify_url"] = this.NotifyURL
	return m
}

func (this AliPayTradeCancel) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayTradeCancel) ExtJSONParamValue() string {
	return marshal(this)
}

type AliPayTradeCancelResponse struct {
	AliPayTradeCancel struct {
		Code       string `json:"code"`
		Msg        string `json:"msg"`
		SubCode    string `json:"sub_code"`
		SubMsg     string `json:"sub_msg"`
		OutTradeNo string `json:"out_trade_no"`
		TradeNo    string `json:"trade_no"`
		RetryFlag  string `json:"retry_flag"`
		Action     string `json:"action"`
	} `json:"alipay_trade_close_response"`
	Sign string `json:"sign"`
}


func (this *AliPayTradeCancelResponse) IsSuccess() bool {
	if this.AliPayTradeCancel.Msg == "Success" {
		return true
	}
	return false
}