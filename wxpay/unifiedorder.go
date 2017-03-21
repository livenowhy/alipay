package wxpay

//  统一下单
import (
	"encoding/xml"
	"github.com/livenowhy/goTools/xmltools"
)

// PlaceOrderResult represent place order reponse message from weixin pay.
// For field explanation refer to: http://pay.weixin.qq.com/wiki/doc/api/app.php?chapter=9_1
// 统一下单;返回
type PlaceOrderResult struct {
	XMLName     xml.Name `xml:"xml"`
	ReturnCode  string   `xml:"return_code"`
	ReturnMsg   string   `xml:"return_msg"`

	// 以下字段在 return_code 为 SUCCESS 的时候有返回
	AppId       string   `xml:"appid"`
	MchId       string   `xml:"mch_id"`
	DeviceInfo  string   `xml:"device_info"`
	NonceStr    string   `xml:"nonce_str"`
	Sign        string   `xml:"sign"`
	ResultCode  string   `xml:"result_code"`
	ErrCode     string   `xml:"err_code"`
	ErrCodeDesc string   `xml:"err_code_des"`

	// 以下字段在return_code 和result_code都为SUCCESS的时候有返回
	TradeType   string   `xml:"trade_type"`
	PrepayId    string   `xml:"prepay_id"`
	CodeUrl     string   `xml:"code_url"`
}


// struct 转换成 map
func (this *PlaceOrderResult) ToMap() map[string]string {
	retMap, err := xmltools.ToMap(this)
	if err != nil {
		panic(err)
	}
	return retMap
}

// Parse the reponse message from weixin pay to struct of PlaceOrderResult, xml字符串转换成结构体
func ParsePlaceOrderResult(resp []byte) (PlaceOrderResult, error) {
	placeOrderResult := PlaceOrderResult{}
	err := xml.Unmarshal(resp, &placeOrderResult)
	if err != nil {
		return placeOrderResult, err
	}

	return placeOrderResult, nil
}

