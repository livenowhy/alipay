package main

import (
	"fmt"
	"github.com/livenowhy/alipay"
)

func TestTradePrecreate() {

	var client = alipay.New(alipay.AppID, "2088102169227503", alipay.PublicKey, alipay.PrivateKey, false)
	var p = alipay.AliPayTradePrecreate{}
	p.OutTradeNo = "12666323fgfhgjhgj23232w"
	p.Subject = "ssdsds"
	p.TotalAmount = "10"
	p.TimeoutExpress = "15d"

	var r, err = client.TradePrecreate(p)
	fmt.Println(err, r)
	fmt.Println(r.IsSuccess(), err, r.AliPayTradePrecreate.QrCode)
}


func TestTradeQuery() {

	var client = alipay.New(alipay.AppID, "2088102169227503", alipay.PublicKey, alipay.PrivateKey, false)
	var p = alipay.AliPayTradeQuery{}
	p.OutTradeNo = "12666323fgfhgjhgj23232w"

	var r, err = client.TradeQuery(p)
	fmt.Println(err, r)

	if r.IsSuccess() {
		fmt.Printf("%v \n", r.AliPayTradeQuery)
	} else {
		fmt.Printf("%v \n", r.AliPayTradeQuery)
	}
}


func TestTradeCancel() {

	var client = alipay.New(alipay.AppID, "2088102169227503", alipay.PublicKey, alipay.PrivateKey, false)
	var p = alipay.AliPayTradeCancel{}
	p.OutTradeNo = "1266632323232"

	var r, err = client.TradeCanncel(p)
	fmt.Println(err, r)

	if r.IsSuccess() {
		fmt.Printf("%v \n", r.AliPayTradeCancel)
		fmt.Printf("%v \n", r)
	} else {
		fmt.Printf("%v \n", r)
	}
}

func main() {
	TestTradePrecreate()
	TestTradeQuery()
	//TestTradeCancel()
}





	//var p = AliPayFastpayTradeRefundQuery{}
	//p.OutTradeNo = "1111111"
	//p.OutRequestNo = "1111111"
	//
	//var r, err = client.TradeFastpayRefundQuery(p)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(r.IsSuccess(), r)

	//var p = AliPayTradeRefund{}
	//p.RefundAmount = "10.00"
	//p.OutTradeNo = "1111111"
	//
	//var r, err = client.TradeRefund(p)
	//fmt.Println(r.IsSuccess(), err, r)



	//var p = AliPayTradeWapPay{}
	//p.NotifyURL = "http://203.86.24.181:3000/alipay"
	//p.ReturnURL = "http://203.86.24.181:3000"
	//p.Subject = "修正了中文的 Bug"
	//p.OutTradeNo = "trade_no_1234"
	//p.TotalAmount = "10.00"
	//p.ProductCode = "eeeeee"
	//
	//var html, url, _ = client.TradeWapPay(p)
	//
	//fmt.Println("url--->")
	//fmt.Println(url)
	//
	//var f, _ = os.Create("test.html")
	//
	//f.WriteString(html)
	//f.Close()