package alipay

// https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.t3A8tQ&docType=4&apiId=757
// alipay.trade.query (统一收单线下交易查询)

// 该接口提供所有支付宝支付订单的查询,商户可以通过该接口主动查询订单状态,完成下一步的业务逻辑.
// 需要调用查询接口的情况: 当商户后台、网络、服务器等出现异常,商户系统最终未接收到支付通知;
// 调用支付接口后, 返回系统错误或未知交易状态情况;
// 调用alipay.trade.pay, 返回INPROCESS的状态;
// 调用alipay.trade.cancel之前, 需确认支付状态


// 请求参数; 私有参数
type AliPayTradeQuery struct {
	AppAuthToken string `json:"-"`                      // 可选
	OutTradeNo   string `json:"out_trade_no,omitempty"` // 与 TradeNo 二选一
	TradeNo      string `json:"trade_no,omitempty"`
}

func (this AliPayTradeQuery) APIName() string {
	return "alipay.trade.query"
}

func (this AliPayTradeQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this AliPayTradeQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayTradeQuery) ExtJSONParamValue() string {
	return marshal(this)
}

type AliPayTradeQueryResponse struct {
	AliPayTradeQuery struct {
		Code                string `json:"code"`
		Msg                 string `json:"msg"`
		SubCode             string `json:"sub_code"`
		SubMsg              string `json:"sub_msg"`
		BuyerLogonId        string `json:"buyer_logon_id"`        // 买家支付宝账号
		BuyerPayAmount      string `json:"buyer_pay_amount"`      // 买家实付金额，单位为元，两位小数。
		BuyerUserId         string `json:"buyer_user_id"`         // 买家在支付宝的用户id
		InvoiceAmount       string `json:"invoice_amount"`        // 交易中用户支付的可开具发票的金额，单位为元，两位小数。
		Openid              string `json:"open_id"`               // 买家支付宝用户号，该字段将废弃，不要使用
		OutTradeNo          string `json:"out_trade_no"`          // 商家订单号
		PointAmount         string `json:"point_amount"`          // 积分支付的金额，单位为元，两位小数。
		ReceiptAmount       string `json:"receipt_amount"`        // 实收金额，单位为元，两位小数
		SendPayDate         string `json:"send_pay_date"`         // 本次交易打款给卖家的时间
		TotalAmount         string `json:"total_amount"`          // 交易的订单金额
		TradeNo             string `json:"trade_no"`              // 支付宝交易号
		TradeStatus         string `json:"trade_status"`          // 交易状态
		AliPayStoreId       string `json:"alipay_store_id"`       // 支付宝店铺编号
		StoreId             string `json:"store_id"`              // 商户门店编号
		TerminalId          string `json:"terminal_id"`           // 商户机具终端编号
		StoreName           string `json:"store_name"`            // 请求交易支付中的商户店铺的名称
		DiscountGoodsDetail string `json:"discount_goods_detail"` // 本次交易支付所使用的单品券优惠的商品优惠信息
		IndustrySepcDetail  string `json:"industry_sepc_detail"`  // 行业特殊信息（例如在医保卡支付业务中，向用户返回医疗信息）。
		FundBillList        []struct {
			FundChannel string `json:"fund_channel"` // 交易使用的资金渠道，详见 支付渠道列表
			Amount      string `json:"amount"`       // 该支付工具类型所使用的金额
			RealAmount  string `json:"real_amount"`  // 渠道实际付款金额
		} `json:"fund_bill_list"` // 交易支付使用的资金渠道
		voucher_detail_list []VoucherDetail `json:"voucher_detail_list"` // 本交易支付时使用的所有优惠券信息
	} `json:"alipay_trade_query_response"`
	Sign string `json:"sign"`
}

type VoucherDetail struct {
	Id                 string `json:"id"`                  // 券id
	Name               string `json:"name"`                // 券名称
	Type               string `json:"type"`                // 当前有三种类型： ALIPAY_FIX_VOUCHER - 全场代金券, ALIPAY_DISCOUNT_VOUCHER - 折扣券, ALIPAY_ITEM_VOUCHER - 单品优惠
	Amount             string `json:"amount"`              // 优惠券面额，它应该会等于商家出资加上其他出资方出资
	MerchantContribute string `json:"merchant_contribute"` // 商家出资（特指发起交易的商家出资金额）
	OtherContribute    string `json:"other_contribute"`    // 其他出资方出资金额，可能是支付宝，可能是品牌商，或者其他方，也可能是他们的一起出资
	Memo               string `json:"memo"`                // 优惠券备注信息
}

func (this *AliPayTradeQueryResponse) IsSuccess() bool {
	if this.AliPayTradeQuery.Msg == "Success" {
		return true
	}
	return false
}

