package wxpay

import (
	"errors"
	"github.com/livenowhy/goTools/httptools"
	"github.com/livenowhy/goTools/xmltools"
	"github.com/livenowhy/goTools/tokentools"
	"fmt"
	"github.com/livenowhy/alipay/wxpay-old"
)

// AppTrans is abstact of Transaction handler.
// With AppTrans, we can get prepay id
type AppTrans struct {
	Config *WxConfig
}

// Initialized the AppTrans with specific config
func NewAppTrans(cfg *WxConfig) (*AppTrans, error) {
	if cfg.AppId == "" ||
		cfg.MchId == "" ||
		cfg.AppKey == "" ||
		cfg.NotifyUrl == "" ||
		cfg.QueryOrderUrl == "" ||
		cfg.PlaceOrderUrl == "" ||
		cfg.TradeType == "" {
		return &AppTrans{Config: cfg}, errors.New("config field canot empty string")
	}

	return &AppTrans{Config: cfg}, nil
}

// Submit the order to weixin pay and return the prepay id if success,
// Prepay id is used for app to start a payment
// If fail, error is not nil, check error for more information
func (this *AppTrans) Submit(orderId string, amount float64, desc string, clientIp string) (string, error) {

	odrInXml := this.signedOrderRequestXmlString(orderId, fmt.Sprintf("%.0f", amount), desc, clientIp)

	fmt.Printf("odrInXml :%s \n", odrInXml)
	resp, err := httptools.DoHttpPost(this.Config.PlaceOrderUrl, []byte(odrInXml))
	if err != nil {
		return "", err
	}

	fmt.Printf("resp: %s \n", resp)
	placeOrderResult, err := ParsePlaceOrderResult(resp)
	if err != nil {
		return "", err
	}

	//Verify the sign of response
	resultInMap := placeOrderResult.ToMap()
	wantSign := httptools.Sign(resultInMap, this.Config.AppKey)
	gotSign := resultInMap["sign"]

	fmt.Printf("placeOrderResult: %v\n", placeOrderResult)
	fmt.Printf("resultInMap: %v\n", resultInMap)

	fmt.Printf("wantSign: %s\n", wantSign)
	fmt.Printf(" gotSign: %s\n", gotSign)
	if wantSign != gotSign {
		return "", fmt.Errorf("sign not match, want:%s, got:%s", wantSign, gotSign)
	}

	if placeOrderResult.ReturnCode != "SUCCESS" {
		return "", fmt.Errorf("return code:%s, return desc:%s", placeOrderResult.ReturnCode, placeOrderResult.ReturnMsg)
	}

	if placeOrderResult.ResultCode != "SUCCESS" {
		return "", fmt.Errorf("resutl code:%s, result desc:%s", placeOrderResult.ErrCode, placeOrderResult.ErrCodeDesc)
	}

	return placeOrderResult.PrepayId, nil
}

func (this *AppTrans) newQueryXml(transId string) string {
	param := make(map[string]string)
	param["appid"] = this.Config.AppId
	param["mch_id"] = this.Config.MchId
	param["transaction_id"] = transId
	param["nonce_str"] = tokentools.NewNonceString()

	sign := httptools.Sign(param, this.Config.AppKey)
	param["sign"] = sign

	return xmltools.ToXmlString(param)
}

// Query the order from weixin pay server by transaction id of weixin pay
func (this *AppTrans) Query(transId string) (QueryOrderResult, error) {
	queryOrderResult := QueryOrderResult{}

	queryXml := this.newQueryXml(transId)
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

// NewPaymentRequest build the payment request structure for app to start a payment.
// Return stuct of PaymentRequest,
// please refer to http://pay.weixin.qq.com/wiki/doc/api/app.php?chapter=9_12&index=2
func (this *AppTrans) NewPaymentRequest(prepayId string) PaymentRequest {
	noncestr := tokentools.NewNonceString()
	timestamp := tokentools.NewTimestampString()

	param := make(map[string]string)
	param["appid"] = this.Config.AppId
	param["partnerid"] = this.Config.MchId
	param["prepayid"] = prepayId
	param["package"] = "Sign=WXPay"
	param["noncestr"] = noncestr
	param["timestamp"] = timestamp

	sign := httptools.Sign(param, this.Config.AppKey)

	payRequest := PaymentRequest{
		AppId:     this.Config.AppId,
		PartnerId: this.Config.MchId,
		PrepayId:  prepayId,
		Package:   "Sign=WXPay",
		NonceStr:  noncestr,
		Timestamp: timestamp,
		Sign:      sign,
	}

	return payRequest
}

func (this *AppTrans) newOrderRequest(orderId, amount, desc, clientIp string) map[string]string {
	param := make(map[string]string)
	param["appid"] = this.Config.AppId
	param["attach"] = "10000100" //optional
	param["body"] = desc
	param["mch_id"] = this.Config.MchId
	//param["nonce_str"] = NewNonceString()
	param["nonce_str"] = "10000100"
	param["notify_url"] = this.Config.NotifyUrl
	param["out_trade_no"] = orderId
	param["spbill_create_ip"] = clientIp
	param["total_fee"] = amount
	param["trade_type"] = this.Config.TradeType

	return param
}

func (this *AppTrans) signedOrderRequestXmlString(orderId, amount, desc, clientIp string) string {
	order := this.newOrderRequest(orderId, amount, desc, clientIp)
	sign := httptools.Sign(order, this.Config.AppKey)

	fmt.Printf("order: %s \n", order)
	fmt.Printf("sign: %s \n", sign)

	order["sign"] = sign

	return xmltools.ToXmlString(order)
}
