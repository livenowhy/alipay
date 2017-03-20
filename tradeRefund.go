package alipay

// https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.5zkPUI&docType=4&apiId=759
// alipay.trade.refund (统一收单交易退款接口)
type AliPayTradeRefund struct {
	AppAuthToken string `json:"-"`                      // 可选
	OutTradeNo   string `json:"out_trade_no,omitempty"` // 与 TradeNo 二选一
	TradeNo      string `json:"trade_no,omitempty"`     // 与 OutTradeNo 二选一
	RefundAmount string `json:"refund_amount"`          // 必须 需要退款的金额，该金额不能大于订单金额,单位为元，支持两位小数
	RefundReason string `json:"refund_reason"`          // 可选 退款的原因说明
	OutRequestNo string `json:"out_request_no"`         // 可选 标识一次退款请求，同一笔交易多次退款需要保证唯一，如需部分退款，则此参数必传。
	OperatorId   string `json:"operator_id"`            // 可选 商户的操作员编号
	StoreId      string `json:"store_id"`               // 可选 商户的门店编号
	TerminalId   string `json:"terminal_id"`            // 可选 商户的终端编号
}

func (this AliPayTradeRefund) APIName() string {
	return "alipay.trade.refund"
}

func (this AliPayTradeRefund) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this AliPayTradeRefund) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayTradeRefund) ExtJSONParamValue() string {
	return marshal(this)
}

type AliPayTradeRefundResponse struct {
	AliPayTradeRefund struct {
		Code                 string `json:"code"`
		Msg                  string `json:"msg"`
		SubCode              string `json:"sub_code"`
		SubMsg               string `json:"sub_msg"`
		TradeNo              string `json:"trade_no"`       // 支付宝交易号
		OutTradeNo           string `json:"out_trade_no"`   // 商户订单号
		BuyerLogonId         string `json:"buyer_logon_id"` // 用户的登录id
		BuyerUserId          string `json:"buyer_user_id"`  // 买家在支付宝的用户id
		FundChange           string `json:"fund_change"`    // 本次退款是否发生了资金变化
		RefundFee            string `json:"refund_fee"`     // 退款总金额
		GmtRefundPay         string `json:"gmt_refund_pay"` // 退款支付时间
		StoreName            string `json:"store_name"`     // 交易在支付时候的门店名称
		RefundDetailItemList []struct {
			FundChannel string `json:"fund_channel"` // 交易使用的资金渠道，详见 支付渠道列表
			Amount      string `json:"amount"`       // 该支付工具类型所使用的金额
			RealAmount  string `json:"real_amount"`  // 渠道实际付款金额
		} `json:"refund_detail_item_list"` // 退款使用的资金渠道
	} `json:"alipay_trade_refund_response"`
	Sign string `json:"sign"`
}

func (this *AliPayTradeRefundResponse) IsSuccess() bool {
	if this.AliPayTradeRefund.Msg == "Success" {
		return true
	}
	return false
}

