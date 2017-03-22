package wxpay


import (
	"encoding/xml"
	"github.com/livenowhy/goTools/xmltools"
	"github.com/livenowhy/goTools/httptools"
	"fmt"
	"github.com/livenowhy/goTools/tokentools"
)


// 查询订单
// 该接口提供所有微信支付订单的查询,商户可以通过该接口主动查询订单状态,完成下一步的业务逻辑.
// 需要调用查询接口的情况:
// 1. 当商户后台、网络、服务器等出现异常,商户系统最终未接收到支付通知;
// 2. 调用支付接口后,返回系统错误或未知交易状态情况;
// 3. 调用被扫支付API,返回 USERPAYING 的状态
// 4. 调用关单或撤销接口API之前, 需确认支付状态;

// QueryOrder Result represent query response message from weixin pay
// Refer to http://pay.weixin.qq.com/wiki/doc/api/app.php?chapter=9_2&index=4
// https://api.mch.weixin.qq.com/pay/orderquery
type QueryOrderResult struct {
	XMLName        xml.Name `xml:"xml"`
	ReturnCode     string   `xml:"return_code"`
	ReturnMsg      string   `xml:"return_msg"`

	// 以下字段在 return_code 为 SUCCESS 的时候有返回
	AppId          string   `xml:"appid"`
	MchId          string   `xml:"mch_id"`
	NonceStr       string   `xml:"nonce_str"`
	Sign           string   `xml:"sign"`
	ResultCode     string   `xml:"result_code"`
	ErrCode        string   `xml:"err_code"`
	ErrCodeDesc    string   `xml:"err_code_des"`

	// 以下字段在return_code 和 result_code 都为 SUCCESS 的时候有返回
	DeviceInfo     string   `xml:"device_info"`
	OpenId         string   `xml:"open_id"`
	IsSubscribe    string   `xml:"is_subscribe"`
	TradeType      string   `xml:"trade_type"`
	TradeState     string   `xml:"trade_state"`
	TradeStateDesc string   `xml:"trade_state_desc"`
	BankType       string   `xml:"bank_type"`
	TotalFee       string   `xml:"total_fee"`
	FeeType        string   `xml:"fee_type"`
	CashFee        string   `xml:"cash_fee"`
	CashFeeType    string   `xml:"cash_fee_type"`
	CouponFee      string   `xml:"coupon_fee"`
	CouponCount    string   `xml:"coupon_count"`
	TransactionId  string   `xml:"transaction_id"`
	OrderId        string   `xml:"out_trade_no"`
	Attach         string   `xml:"attach"`
	TimeEnd        string   `xml:"time_end"`
}

func (this *QueryOrderResult) ToMap() map[string]string {
	retMap, err := xmltools.ToMap(this)
	if err != nil {
		panic(err)
	}

	return retMap
}

func ParseQueryOrderResult(resp []byte) (QueryOrderResult, error) {
	queryOrderResult := QueryOrderResult{}
	err := xml.Unmarshal(resp, &queryOrderResult)
	if err != nil {
		return queryOrderResult, err
	}

	return queryOrderResult, nil
}




// https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_2&index=4
func (this *AppTrans) newQueryXml(transId string) string {
	param := make(map[string]string)
	param["appid"] = this.Config.AppId
	param["mch_id"] = this.Config.MchId
	param["transaction_id"] = transId
	param["nonce_str"] = tokentools.NewNonceString()

	sign := httptools.Sign(param, this.Config.AppKey)
	param["sign"] = sign
	fmt.Printf("sign : %s \n", sign)

	return xmltools.ToXmlString(param)
}

// Query the order from weixin pay server by transaction id of weixin pay
func (this *AppTrans) Query(transId string) (QueryOrderResult, error) {
	queryOrderResult := QueryOrderResult{}

	queryXml := this.newQueryXml(transId)
	fmt.Printf("queryXml: %s \n", queryXml)
	// fmt.Println(queryXml)
	resp, err := httptools.DoHttpPost(this.Config.QueryOrderUrl, []byte(queryXml))
	if err != nil {
		return queryOrderResult, nil
	}

	queryOrderResult, err = ParseQueryOrderResult(resp)
	if err != nil {
		return queryOrderResult, err
	}

	//verity sign of response
	resultInMap := queryOrderResult.ToMap()
	wantSign := httptools.Sign(resultInMap, this.Config.AppKey)
	gotSign := resultInMap["sign"]
	if wantSign != gotSign {
		return queryOrderResult, fmt.Errorf("sign not match, want:%s, got:%s", wantSign, gotSign)
	}

	return queryOrderResult, nil
}

