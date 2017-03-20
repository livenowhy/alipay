package alipay

// https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.SPYWcR&docType=4&apiId=862
// alipay.trade.precreate (统一收单线下交易预创建)
type AliPayTradePrecreate struct {
	AppAuthToken string `json:"-"`                      // 可选
	OutTradeNo   string `json:"out_trade_no,omitempty"` // 商户订单号
	Subject   string `json:"subject,omitempty"` // 订单标题
	TotalAmount   string `json:"total_amount,omitempty"` // 订单标题

	TimeoutExpress   string `json:"timeout_express"` // 订单标题
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
	AliPayTradePrecreate struct{
		Code       string `json:"code"`
		Msg        string `json:"msg"`
		SubCode    string `json:"sub_code"`
		SubMsg     string `json:"sub_msg"`

		OutTradeNo string `json:"out_trade_no"`
		QrCode    string `json:"qr_code"`
	} `json:"alipay_trade_precreate_response"`
	Sign string `json:"sign"`
}

func (this *AliPayTradePrecreateResponse) IsSuccess() bool {
	if this.AliPayTradePrecreate.Msg == "Success" {
		return true
	}
	return false
}

