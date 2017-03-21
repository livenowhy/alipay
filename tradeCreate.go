package alipay



// https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.CkYNiG&docType=4&apiId=1046
// alipay.trade.create (统一收单交易创建接口)
// 商户通过该接口进行交易的创建下单
func (this AliPayTradeCreate) APIName() string {
	return "alipay.trade.create"
}

func (this AliPayTradeCreate) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this AliPayTradeCreate) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayTradeCreate) ExtJSONParamValue() string {
	return marshal(this)
}

type AliPayTradeCreate struct {
	AppAuthToken         string             `json:"-"`                      // 可选
	OutTradeNo           string             `json:"out_trade_no,omitempty"` // 与 TradeNo 二选一
	SellerId             string             `json:"seller_id,omitempty"`    // 卖家支付宝用户ID
	TotalAmount          string             `json:"total_amount"`           // 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000] 如果同时传入了【打折金额】，【不可打折金额】，【订单总金额】三者，则必须满足如下条件：【订单总金额】=【打折金额】+【不可打折金额】
	DiscountableAmount   string             `json:"discountable_amount"`    // 可打折金额. 参与优惠计算的金额，单位为元，精确到小数点后两位
	UndiscountableAmount string             `json:"undiscountable_amount"`
	BuyerLogonId         string             `json:"buyer_logon_id"`
	Subject              string             `json:"subject"`
	Body                 string             `json:"body"`
	BuyerId              string             `json:"buyer_id"`
	GoodsDetail          []*GoodsDetailItem `json:"goods_detail,omitempty"`
	OperatorId           string             `json:"operator_id"`
	StoreId              string             `json:"store_id"`
	TerminalId           string             `json:"terminal_id"`
	ExtendParams         *ExtendParamsItem  `json:"extend_params,omitempty"`
	TimeoutExpress       string             `json:"timeout_express"`
	RoyaltyInfo          *RoyaltyInfo       `json:"royalty_info,omitempty"`
	AliPayStoreId        string             `json:"alipay_store_id"`
	SubMerchant          []SubMerchantItem  `json:"sub_merchant"`
	MerchantOrderNo      string             `json:"merchant_order_no"`
}

type AliPayTradeCreateResponse struct {
	AliPayTradeCreateResponse struct {
		Code       string `json:"code"`
		Msg        string `json:"msg"`
		SubCode    string `json:"sub_code"`
		SubMsg     string `json:"sub_msg"`
		TradeNo    string `json:"trade_no"` // 支付宝交易号
		OutTradeNo string `json:"out_trade_no"`
	} `json:"alipay_trade_create_response"`
	Sign string `json:"sign"`
}

type ExtendParamsItem struct {
	SysServiceProviderId string `json:"sys_service_provider_id"`
	HbFqNum              string `json:"hb_fq_num"`
	HbFqSellerPercent    string `json:"hb_fq_seller_percent"`
	TimeoutExpress       string `json:"timeout_express"`
}

type RoyaltyInfo struct {
	RoyaltyType       string                   `json:"royalty_type"`
	RoyaltyDetailInfo []*RoyaltyDetailInfoItem `json:"royalty_detail_infos,omitempty"`
}

type RoyaltyDetailInfoItem struct {
	SerialNo         string `json:"serial_no"`
	TransInType      string `json:"trans_in_type"`
	BatchNo          string `json:"batch_no"`
	OutRelationId    string `json:"out_relation_id"`
	TransOutType     string `json:"trans_out_type"`
	TransOut         string `json:"trans_out"`
	TransIn          string `json:"trans_in"`
	Amount           string `json:"amount"`
	Desc             string `json:"desc"`
	AmountPercentage string `json:"amount_percentage"`
	AliPayStoreId    string `json:"alipay_store_id"`
}

type SubMerchantItem struct {
	MerchantId string `json:"merchant_id"`
}

type GoodsDetailItem struct {
	GoodsId       string `json:"goods_id"`
	AliPayGoodsId string `json:"alipay_goods_id"`
	GoodsName     string `json:"goods_name"`
	Quantity      string `json:"quantity"`
	Price         string `json:"price"`
	GoodsCategory string `json:"goods_category"`
	Body          string `json:"body"`
	ShowUrl       string `json:"show_url"`
}
