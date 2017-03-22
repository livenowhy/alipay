package wxpay

import (
	"errors"
	"github.com/livenowhy/goTools/httptools"
	"github.com/livenowhy/goTools/xmltools"
	"github.com/livenowhy/goTools/tokentools"
	"fmt"
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

func (this *AppTrans) signedOrderRequestXmlString(orderId, amount, desc, clientIp string) string {
	order := this.newOrderRequest(orderId, amount, desc, clientIp)
	sign := httptools.Sign(order, this.Config.AppKey)

	fmt.Printf("order: %s \n", order)
	fmt.Printf("sign: %s \n", sign)
	order["sign"] = sign

	return xmltools.ToXmlString(order)
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
	//param["nonce_str"] = tokentools.NewNonceString()
	param["nonce_str"] = "10000100"
	param["notify_url"] = this.Config.NotifyUrl
	param["out_trade_no"] = orderId
	param["spbill_create_ip"] = clientIp
	param["total_fee"] = amount
	param["trade_type"] = this.Config.TradeType

	return param
}
