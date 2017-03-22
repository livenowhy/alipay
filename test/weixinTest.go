package main

import (
	//"fmt"
	"github.com/livenowhy/alipay/wxpay"
	"fmt"
)

func TestTrans()  {
	//初始化
	cfg := &wxpay.WxConfig{
		AppId:         "wxd678efh567hg6787",
		AppKey:        "10000100",
		MchId:         "10000100",
		NotifyUrl:     "10000100",
		PlaceOrderUrl: "https://api.mch.weixin.qq.com/pay/unifiedorder",
		QueryOrderUrl: "https://api.mch.weixin.qq.com/pay/orderquery",
		TradeType:     "WEB",
	}
	appTrans, err := wxpay.NewAppTrans(cfg)
	if err != nil {
		panic(err)
	}

	//获取 prepay id, 手机端得到 prepay id 后加上验证就可以使用这个 id 发起支付调用
	prepayId, err := appTrans.Submit("10000100", 1, "10000100", "114.25.139.11")
	if err != nil {
		fmt.Printf("appTrans.Submit is error: %s \n", err.Error())
		//panic(err)
	}
	fmt.Println(prepayId)
}

func TestQuery() {

		//初始化
	cfg := &wxpay.WxConfig{
		AppId:         "wxd678efh567hg6787",
		AppKey:        "10000100",
		MchId:         "10000100",
		NotifyUrl:     "10000100",
		PlaceOrderUrl: "https://api.mch.weixin.qq.com/pay/unifiedorder",
		QueryOrderUrl: "https://api.mch.weixin.qq.com/pay/orderquery",
		TradeType:     "WEB",
	}
	appTrans, err := wxpay.NewAppTrans(cfg)
	if err != nil {
		panic(err)
	}

	// 查询订单接口
	queryResult, err := appTrans.Query("1008450740201411110005820873")
	if err != nil {
		fmt.Printf("appTrans.Submit is error: %s \n", err.Error())
		//panic(err)
	}
	fmt.Println(queryResult)
}

func main() {
	TestQuery()

	//
	////加上Sign，已方便手机直接调用
	//payRequest := appTrans.NewPaymentRequest(prepayId)
	//fmt.Println(payRequest)
	//
}
