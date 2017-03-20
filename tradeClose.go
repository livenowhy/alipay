package alipay

// https://doc.open.alipay.com/doc2/apiDetail.htm?apiId=1058&docType=4
// 用于交易创建后，用户在一定时间内未进行支付，可调用该接口直接将未付款的交易进行关闭。
// alipay.trade.close (统一收单交易关闭接口)
type AliPayTradeClose struct {
	AppAuthToken string `json:"-"`                      // 可选
	NotifyURL    string `json:"-"`                      // 可选
	OutTradeNo   string `json:"out_trade_no,omitempty"` // 与 TradeNo 二选一
	TradeNo      string `json:"trade_no,omitempty"`     // 与 OutTradeNo 二选一
	OperatorId   string `json:"operator_id,omitempty"`  // 可选
}

func (this AliPayTradeClose) APIName() string {
	return "alipay.trade.close"
}

func (this AliPayTradeClose) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["notify_url"] = this.NotifyURL
	return m
}

func (this AliPayTradeClose) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayTradeClose) ExtJSONParamValue() string {
	return marshal(this)
}

type AliPayTradeCloseResponse struct {
	AliPayTradeClose struct {
		Code       string `json:"code"`
		Msg        string `json:"msg"`
		SubCode    string `json:"sub_code"`
		SubMsg     string `json:"sub_msg"`
		OutTradeNo string `json:"out_trade_no"`
		TradeNo    string `json:"trade_no"`
	} `json:"alipay_trade_close_response"`
	Sign string `json:"sign"`
}

