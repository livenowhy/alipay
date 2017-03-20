package alipay

import (
	"testing"
	//"os"
	"fmt"
	//"os"
)

var publicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA5nHjhVek/ubz6iAlPmJyIA1YQkzWIXgF6Hm2Ddzwkod+GeXpIpM2qlfDu3ENunSemrXj/LNWcvzUHaC/9dl3yES7+OEWhpIAoxFv33mPEsQ3U/YPjAgLO7uZ7VI5Mb26o4xwNh97Vt30W8A0GGmzWyI7EooLq7Tx340b4hGiOHp52LHqb3eV0HuKWr5BPpRZn7JrkkiI6MP+Z6s+U4VHdR/g0PYEHycDL5etN4CPApO+IoNB9ey0/FtAqpBQkTPjubN/p9DL4MO//L8j4X8nHaphK7CXAJLwXwl8YYdlqbmyYzBDR8LgwbEcGTPpeUdGnDFtHIJcG04A7imH0nu9MQIDAQAB
-----END PUBLIC KEY-----`)

var privateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA5nHjhVek/ubz6iAlPmJyIA1YQkzWIXgF6Hm2Ddzwkod+GeXpIpM2qlfDu3ENunSemrXj/LNWcvzUHaC/9dl3yES7+OEWhpIAoxFv33mPEsQ3U/YPjAgLO7uZ7VI5Mb26o4xwNh97Vt30W8A0GGmzWyI7EooLq7Tx340b4hGiOHp52LHqb3eV0HuKWr5BPpRZn7JrkkiI6MP+Z6s+U4VHdR/g0PYEHycDL5etN4CPApO+IoNB9ey0/FtAqpBQkTPjubN/p9DL4MO//L8j4X8nHaphK7CXAJLwXwl8YYdlqbmyYzBDR8LgwbEcGTPpeUdGnDFtHIJcG04A7imH0nu9MQIDAQABAoIBADtr1JI2lloQLYcKgPAELI9tQXvfGjwJGeTnNXV/qhzuiSqeaS48LZVhChyUO0j/90HHcyFfEQSXgw6cu8LmL4fZiWrUh51tmnJx9Zn3W1godmfVBA9Ep7jgWAGk+RCFbtHrTaf3GDO/Cdp6xjQhVVlDi13tcLL29oOBfFYf2+eZ1/P/by2uTNjQljIy0C3Y7HKHC9FIMayb0qr1ogtbrXpv0X+SJDGnkzZ0abD+hb6m6nCPFuWNlgug/woDNbz2ERUud046apvT8QcgKf8g+wkxSFmBBhyzOSdlfgPg3IRANHdwraH0giWjgl0ATwYivcmcee5CKTIJJti/9OxJGV0CgYEA+J8bvYZAjOay/WTbyvHeSFL0LYOe0mpOUqCRFkbXuu2NfaVQgY6M66/P93BJsnvywTCzYK8NqtuTd5lT05Z65VvARHaY70TVsbhHHGL7uPW4zl+w24cgvJspW32idT11V/rpyxa4TAYLz4Sc7DWWqd2gS19PR1gmTWLUpbY4C6MCgYEA7UivI1XPRcOpkjSKPqnE1U24Z6amo2jR0S75HeZiyr57WNVF/C7GOOrAa9iC0lP90ZTZz4sA25pNvIJCBXPbCEkkeQEMlYMP9gNfF6swCJSd0o+/FTxBQjFNAxJbjGiez6JiS/pKlU5xahUlz+MPIIDoPpYhG5OJfSGUomvLoRsCgYEArgFNvPu1OehM5aQtYICffibw1toD0misoI/Ye6rhWebg3EtTaRmH73yBJUDlk5HEeQnIwaZC/o3nTBF35eKv6M6qZkBLEUsv9M5wlqGAZBc2XNoEQeD1PIusCmiHx15YPldSQd2I50xj/71kxr1K0WBXy+jM57X+OLMu+TvrRdkCgYBe2WAQ2ts4wubEp7XfKJOixCeDX01bIqpEWOeBuTciXhQGRXJtRPE9jTEPqwNYOfTK0xph6OmV3s9WHCdx8IbMgBhROkcieLpnl59qYaaYt3eC1K0IOwbHA9gJ80ufC673xqxOPKWVmz6grgVoXRYN4xly8yYW1Nnvp/23EynZowKBgAUo6DkHj8oTySCJ2MT06Zevmj1H4pFZWPJDtGV90DKdNcOFTqDqJ+R8gTuufy0IOtk9vZvqlxMn+oI5cgPohG8GpzGhElRRiejq9KQS1ASlSbhBdpLI088dvUo/X+xbvof1P1vsJbQWgpBHPnX3S7o2g6z2Odw1rR7ExQPNE1vb
-----END RSA PRIVATE KEY-----
`)

var appID = "2016080200149684"

func TestSign(t *testing.T) {

	var client = New(appID, "2088102169227503", publicKey, privateKey, false)

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

	var p = AliPayTradePrecreate{}
	p.OutTradeNo = "10.00"
	p.Subject = "ssdsds"
	p.TotalAmount = "1111111"
	p.TimeoutExpress = "15d"

	var r, err = client.TradePrecreate(p)
	fmt.Println(err, r)
	fmt.Println(r.IsSuccess(), err, r.AliPayTradePrecreate.QrCode)
}
