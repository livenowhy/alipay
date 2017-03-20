package alipay

// https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.IGVsS6&docType=4&apiId=850
// alipay.trade.pay (统一收单交易支付接口)
// 收银员使用扫码设备读取用户手机支付宝“付款码”/声波获取设备（如麦克风）读取用户手机支付宝的声波信息后，将二维码或条码信息/声波信息通过本接口上送至支付宝发起支付。
type AliPayTradePay struct {
	AppAuthToken string `json:"-"` // 可选
	NotifyURL    string `json:"-"` // 可选

	OutTradeNo           string             `json:"out_trade_no"`           // 必须 商户订单号,64个字符以内、可包含字母、数字、下划线；需保证在商户端不重复
	Scene                string             `json:"scene"`                  // 必须 支付场景 条码支付，取值：bar_code 声波支付，取值：wave_code	bar_code,wave_code
	AuthCode             string             `json:"auth_code"`              // 必须 支付授权码
	Subject              string             `json:"subject"`                // 必须 订单标题
	BuyerId              string             `json:"buyer_id"`               // 可选 家的支付宝用户id，如果为空，会从传入了码值信息中获取买家ID
	SellerId             string             `json:"seller_id"`              // 可选 如果该值为空，则默认为商户签约账号对应的支付宝用户ID
	TotalAmount          string             `json:"total_amount"`           // 可选 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]。 如果同时传入【可打折金额】和【不可打折金额】，该参数可以不用传入； 如果同时传入了【可打折金额】，【不可打折金额】，【订单总金额】三者，则必须满足如下条件：【订单总金额】=【可打折金额】+【不可打折金额】
	DiscountableAmount   string             `json:"discountable_amount"`    // 可选 参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]。 如果该值未传入，但传入了【订单总金额】和【不可打折金额】，则该值默认为【订单总金额】-【不可打折金额】
	UnDiscountableAmount string             `json:"undiscountable_amount"`  // 可选 不参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]。如果该值未传入，但传入了【订单总金额】和【可打折金额】，则该值默认为【订单总金额】-【可打折金额】
	Body                 string             `json:"body"`                   // 可选 订单描述
	GoodsDetail          []*GoodsDetailItem `json:"goods_detail,omitempty"` // 可选 订单包含的商品列表信息，Json格式，其它说明详见商品明细说明
	OperatorId           string             `json:"operator_id"`            // 可选 商户操作员编号
	StoreId              string             `json:"store_id"`               // 可选 商户门店编号
	TerminalId           string             `json:"terminal_id"`            // 可选 商户机具终端编号
	AliPayStoreId        string             `json:"alipay_store_id"`        // 可选 支付宝的店铺编号
	TimeoutExpress       string             `json:"timeout_express"`        // 该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m
	AuthNo               string             `json:"auth_no"`                // 预授权号，预授权转交易请求中传入
}

func (this AliPayTradePay) APIName() string {
	return "alipay.trade.pay"
}

func (this AliPayTradePay) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["notify_url"] = this.NotifyURL
	return m
}

func (this AliPayTradePay) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayTradePay) ExtJSONParamValue() string {
	return marshal(this)
}

type AliPayTradePayResponse struct {
	AliPayTradePay struct {
		Code                string `json:"code"`
		Msg                 string `json:"msg"`
		SubCode             string `json:"sub_code"`
		SubMsg              string `json:"sub_msg"`
		BuyerLogonId        string `json:"buyer_logon_id"`        // 买家支付宝账号
		BuyerPayAmount      string `json:"buyer_pay_amount"`      // 买家实付金额，单位为元，两位小数。
		BuyerUserId         string `json:"buyer_user_id"`         // 买家在支付宝的用户id
		CardBalance         string `json:"card_balance"`          // 支付宝卡余额
		DiscountGoodsDetail string `json:"discount_goods_detail"` // 本次交易支付所使用的单品券优惠的商品优惠信息
		FundBillList        []struct {
			FundChannel string `json:"fund_channel"` // 交易使用的资金渠道，详见 支付渠道列表
			Amount      string `json:"amount"`       // 该支付工具类型所使用的金额
			RealAmount  string `json:"real_amount"`  // 渠道实际付款金额
		} `json:"fund_bill_list"` // 交易支付使用的资金渠道
		GmtPayment        string          `json:"gmt_payment"`
		InvoiceAmount     string          `json:"invoice_amount"`      // 交易中用户支付的可开具发票的金额，单位为元，两位小数。
		OutTradeNo        string          `json:"out_trade_no"`        // 创建交易传入的商户订单号
		TradeNo           string          `json:"trade_no"`            // 支付宝交易号
		PointAmount       string          `json:"point_amount"`        // 积分支付的金额，单位为元，两位小数。
		ReceiptAmount     string          `json:"receipt_amount"`      // 实收金额，单位为元，两位小数
		StoreName         string          `json:"store_name"`          // 发生支付交易的商户门店名称
		TotalAmount       string          `json:"total_amount"`        // 发该笔退款所对应的交易的订单金额
		VoucherDetailList []VoucherDetail `json:"voucher_detail_list"` // 本交易支付时使用的所有优惠券信息
	} `json:"alipay_trade_pay_response"`
	Sign string `json:"sign"`
}

func (this *AliPayTradePayResponse) IsSuccess() bool {
	if this.AliPayTradePay.Msg == "Success" {
		return true
	}
	return false
}
