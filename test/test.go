package main

import (
	//"fmt"
	"github.com/livenowhy/wxpay"
	"fmt"
)

func main() {
	//初始化
	cfg := &wxpay.WxConfig{
		AppId:         "10000100",
		AppKey:        "10000100",
		MchId:         "10000100",
		NotifyUrl:     "10000100",
		PlaceOrderUrl: "https://api.mch.weixin.qq.com/pay/unifiedorder",
		QueryOrderUrl: "https://api.mch.weixin.qq.com/pay/orderquery",
		TradeType:     "APP",
	}
	appTrans, err := wxpay.NewAppTrans(cfg)
	if err != nil {
		panic(err)
	}

	//获取prepay id，手机端得到prepay id后加上验证就可以使用这个id发起支付调用
	prepayId, err := appTrans.Submit("10000100", 1, "10000100", "114.25.139.11")
	if err != nil {
		fmt.Printf("appTrans.Submit is error: %s \n", err.Error())
		panic(err)
	}
	fmt.Println(prepayId)
	//
	////加上Sign，已方便手机直接调用
	//payRequest := appTrans.NewPaymentRequest(prepayId)
	//fmt.Println(payRequest)
	//
	////查询订单接口
	//queryResult, err := appTrans.Query("1008450740201411110005820873")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(queryResult)
}
