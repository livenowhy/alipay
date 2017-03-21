package main

import (
	"fmt"
	"github.com/livenowhy/alipay/wxpay"
)



func TestPrecreate() {

	//var client = wxpay.GoodsDetail{}.New(alipay.AppID, "2088102169227503", alipay.PublicKey, alipay.PrivateKey, false)
	var p = wxpay.WeixinTradePrecreate{}
	p.Appid = "1230000109"
	p.MchId= "1230000109"
	p.NonceStr = "1230000109"
	p.Sign = "1230000109"
	p.Body = "1230000109"
	p.OutTradeNo = "1230000109"
	p.TotalFee = 123
	p.SpbillCreateIp = "123.12.12.123"
	p.NotifyUrl = "123.12.12.123"
	p.TradeType = "web"


	//var r, err = client.TradePrecreate(p)
	//fmt.Println(err, r)
	//fmt.Println(r.IsSuccess(), err, r.AliPayTradePrecreate.QrCode)
}


func main() {
	fmt.Printf("sss")
}
