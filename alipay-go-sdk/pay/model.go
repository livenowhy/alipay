package pay

// PayClient 支付客户端接口
type PayClient interface {
	Pay(charge *Charge) (map[string]string, error)
	//检查签名
	//CheckSign(data []byte, sign []byte) error
}

// Charge 支付参数
type Charge struct {
	TradeNum    string `json:"tradeNum,omitempty"`
	Origin      string `json:"origin,omitempty"`
	UserID      string `json:"userId,omitempty"`
	PayMethod   int64  `json:"payMethod,omitempty"`
	MoneyFee    int64  `json:"MoneyFee,omitempty"`
	CallbackURL string `json:"callbackURL,omitempty"`
	ReturnURL   string `json:"returnURL,omitempty"`
	ShowURL     string `json:"showURL,omitempty"`
	Describe    string `json:"describe,omitempty"`
	OpenID      string `json:"openid,omitempty"`
}

//PayCallback 支付返回
type PayCallback struct {
	Origin      string `json:"origin"`
	TradeNum    string `json:"trade_num"`
	OrderNum    string `json:"order_num"`
	CallBackURL string `json:"callback_url"`
	Status      int64  `json:"status"`
}

// CallbackReturn 回调业务代码时的参数
type CallbackReturn struct {
	IsSucceed     bool   `json:"isSucceed"`
	OrderNum      string `json:"orderNum"`
	TradeNum      string `json:"tradeNum"`
	UserID        string `json:"userID"`
	MoneyFee      int64  `json:"moneyFee"`
	Sign          string `json:"sign"`
	ThirdDiscount int64  `json:"thirdDiscount"`
}

// BaseResult 支付结果
type BaseResult struct {
	IsSucceed     bool   // 是否交易成功
	TradeNum      string // 交易流水号
	MoneyFee      int64  // 支付金额
	TradeTime     string // 交易时间
	ContractNum   string // 交易单号
	UserInfo      string // 支付账号信息(有可能有，有可能没有)
	ThirdDiscount int64  // 第三方优惠
}

type AliWebPayResult struct {
	NotifyTime       string `json:"notify_time"`
	NotifyType       string `json:"notify_type"`
	NotifyID         string `json:"notify_id"`
	AppId            string `json:"app_id"`
	Charset          string `json:"charset"`
	Version          string `json:"version"`
	SignType         string `json:"sign_type"`
	Sign             string `json:"sign"`
	TradeNum         string `json:"trade_no"`
	OutTradeNum      string `json:"out_trade_no"`
	OutBizNo         string `json:"out_biz_no"`
	BuyerID          string `json:"buyer_id"`
	buyerPayAmount   string `json:"buyer_pay_amount"`
	RefundFee        string `json:"refund_fee"`
	Subject          string `json:"subject"`
	PayMentType      string `json:"payment_type"`
	GmtPayMent       string `json:"gmt_payment"`
	GmtClose         string `json:"gmt_close"`
	SellerEmail      string `json:"seller_email"`
	BuyerEmail       string `json:"buyer_email"`
	SellerID         string `json:"seller_id"`
	Price            string `json:"price"`
	TotalFee         string `json:"total_fee"`
	Quantity         string `json:"quantity"`
	Body             string `json:"body"`
	Discount         string `json:"discount"`
	IsTotalFeeAdjust string `json:"is_total_fee_adjust"`
	UseCoupon        string `json:"use_coupon"`
	RefundStatus     string `json:"refund_status"`
	GmtRefund        string `json:"gmt_refund"`
	AliQueryResult
}

type AliQueryResult struct {
	TradeNo        string `json:"trade_no"`
	OutTradeNo     string `json:"out_trade_no"`
	OpenID         string `json:"open_id"`
	BuyerLogonID   string `json:"buyer_logon_id"`
	TradeStatus    string `json:"trade_status"`
	TotalAmount    string `json:"total_amount"`
	ReceiptAmount  string `json:"receipt_amount"`
	BuyerPayAmount string `json:"BuyerPayAmount"`
	PointAmount    string `json:"point_amount"`
	InvoiceAmount  string `json:"invoice_amount"`
	SendPayDate    string `json:"send_pay_date"`
	AlipayStoreID  string `json:"alipay_store_id"`
	StoreID        string `json:"store_id"`
	TerminalID     string `json:"terminal_id"`
	FundBillList   []struct {
		FundChannel string `json:"fund_channel"`
		Amount      string `json:"amount"`
	} `json:"fund_bill_list"`
	StoreName           string `json:"store_name"`
	BuyerUserID         string `json:"buyer_user_id"`
	DiscountGoodsDetail string `json:"discount_goods_detail"`
	IndustrySepcDetail  string `json:"industry_sepc_detail"`
	PassbackParams      string `json:"passback_params"`
}
