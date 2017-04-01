# alipay-go-sdk
支付宝支付 新版 go sdk

###usage
go to /tools/config.go

input your config infomation

```
  	outTradeNo := "creat your out trade no"
	charge := new(pay.Charge)
	charge.MoneyFee = totalFee
	charge.Describe = "describe"
	charge.TradeNum = outTradeNo
	reqStr := pay.Pay(charge)
```
then return reqStr to client
