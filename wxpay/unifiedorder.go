package wxpay

//  统一下单
import (
	"encoding/xml"
	"github.com/livenowhy/goTools/xmltools"
	"github.com/livenowhy/goTools/httptools"
	"fmt"
)


type WeixinTradePrecreate struct {
	AppAuthToken   string         `json:"-"`                          // 可选
	Appid          string         `json:"appid,omitempty"`            // 必须; 应用ID;微信开放平台审核通过的应用APPID
	MchId          string         `json:"mch_id,omitempty"`           // 必须; 商户号; 微信支付分配的商户号
	NonceStr       string         `json:"nonce_str,omitempty"`        // 必须; 随机字符串; 随机字符串, 不长于32位
	Sign           string         `json:"sign,omitempty"`             // 必须;签名
	Body           string         `json:"body,omitempty"`             // 必须;商品描述
	OutTradeNo     string         `json:"out_trade_no,omitempty"`     // 必须;商户订单号;商户系统内部的订单号
	TotalFee       int            `json:"total_fee,omitempty"`        // 必须;总金额
	SpbillCreateIp string         `json:"spbill_create_ip,omitempty"` // 必须;终端IP
	NotifyUrl      string         `json:"notify_url,omitempty"`       // 必须;通知地址
	TradeType      string         `json:"trade_type,omitempty"`       // 必须;交易类型 (web, app)
	Detail         []*GoodsDetail `json:"goods_detail,omitempty"`     // 商品详情
	SignType       string         `json:"sign_type,omitempty"`        // 签名类型, 目前支持HMAC-SHA256和MD5,默认为MD5
	DeviceInfo     string         `json:"device_info,omitempty"`      // 终端设备号(门店号或收银设备ID), 默认请传"WEB"
	Attach         string         `json:"attach,omitempty"`           // 附加数据
	FeeType        string         `json:"fee_type,omitempty"`         // CNY 符合ISO 4217标准的三位字母代码,默认人民币:CNY, 其他值列表详见货币类型
	TimeStart      string         `json:"time_start,omitempty"`       // 订单生成时间
	TimeExpire     string         `json:"time_start,omitempty"`       // 交易结束时间
	GoodsTag       string         `json:"goods_tag,omitempty"`        // 商品标记
	LimitPay       string         `json:"limit_pay,omitempty"`        // 指定支付方式; no_credit--指定不能使用信用卡支付
}

type GoodsDetail struct {
	GoodsId       string `json:"goods_id,omitempty"`       // 必须; 商品的编号
	WxpayGoodsId  string `json:"wxpay_goods_id,omitempty"` // 微信支付定义的统一商品编号
	GoodsName     string `json:"goods_name,omitempty"`     // 必须; 商品名称
	Quantity      string `json:"quantity,omitempty"`       // 必须; 商品数量
	Price         string `json:"price,omitempty"`          // 必须; 商品单价; 单位为分
	GoodsCategory string `json:"goods_category,omitempty"` // 商品类目ID
	Body          string `json:"body,omitempty"`           //  商品描述信息
}

//////////////////////////

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




// Submit the order to weixin pay and return the prepay id if success,
// Prepay id is used for app to start a payment
// If fail, error is not nil, check error for more information
// 提交
func (this *AppTrans) Submit(orderId string, amount float64, desc string, clientIp string) (string, error) {


	// 签名下单的xml报文,请求报文中添加 sign 内容
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