package pay

import (
	"github.com/livenowhy/alipay-go-sdk/tools"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"
)

func init() {
	InitAliAppClient(&AliAppClient{
		AppID:      tools.APPID,
		PrivateKey: tools.Bytes2RSAPrivateKey(tools.PRIVATE_KEY).(*rsa.PrivateKey),
		PublicKey:  tools.Bytes2RSAPublicKey(tools.PUBILC_KEY).(*rsa.PublicKey),
	})
}

var defaultAliAppClient *AliAppClient

type AliAppClient struct {
	AppID      string // 应用ID
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func InitAliAppClient(c *AliAppClient) {
	defaultAliAppClient = c
}

// DefaultAliAppClient 得到默认支付宝app客户端
func DefaultAliAppClient() *AliAppClient {
	return defaultAliAppClient
}

// 获取支付接口
func Pay(charge *Charge) string {
	err := checkCharge(charge)
	checkErr(err)

	ct := DefaultAliAppClient()
	re, err := ct.BuildPayMap(charge)
	checkErr(err)
	reqStr := ct.BuildPayRequestString(re)
	return reqStr
}

// 验证内容
func checkCharge(charge *Charge) error {
	if charge.PayMethod < 0 {
		return errors.New("payMethod less than 0")
	}
	if charge.MoneyFee < 0 {
		return errors.New("totalFee less than 0")
	}
	return nil
}

func buildBizContent(charge *Charge) string {
	var mc = make(map[string]string)
	mc["subject"] = charge.Describe
	mc["out_trade_no"] = charge.TradeNum
	mc["product_code"] = "QUICK_MSECURITY_PAY"
	mc["total_amount"] = fmt.Sprintf("%.2f", float64(charge.MoneyFee)/float64(100))
	byteBiz, _ := json.Marshal(mc)
	biz := string(byteBiz)
	return biz
}

func (this *AliAppClient) BuildPayRequestString(m map[string]string) string {
	reqStr := ""
	reqStr += "app_id=" + url.QueryEscape(m["app_id"])
	reqStr += "&biz_content=" + url.QueryEscape(m["biz_content"])
	reqStr += "&charset=" + url.QueryEscape(m["charset"])
	reqStr += "&format=" + url.QueryEscape(m["format"])
	reqStr += "&method=" + url.QueryEscape(m["method"])
	reqStr += "&notify_url=" + url.QueryEscape(m["notify_url"])
	reqStr += "&sign_type=" + url.QueryEscape(m["sign_type"])
	reqStr += "&timestamp=" + url.QueryEscape(m["timestamp"])
	reqStr += "&version=" + url.QueryEscape(m["version"])
	reqStr += "&sign=" + url.QueryEscape(m["sign"])

	return reqStr
}

func (this *AliAppClient) BuildPayMap(charge *Charge) (map[string]string, error) {
	var m = make(map[string]string)
	m["app_id"] = this.AppID
	m["method"] = "alipay.trade.app.pay"
	m["charset"] = "utf-8"
	m["format"] = "JSON"
	m["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	m["version"] = "1.0"
	m["notify_url"] = tools.NOTIFY_URL
	m["sign_type"] = "RSA2"

	m["biz_content"] = buildBizContent(charge)

	sign, err := this.GenSign(m)
	checkErr(err)
	m["sign"] = sign
	fmt.Println("sign:", sign)
	return m, nil
}

// GenSign 产生签名
func (this *AliAppClient) GenSign(m map[string]string) (string, error) {
	delete(m, "sign")
	var data []string
	for k, v := range m {
		if v == "" {
			continue
		}
		data = append(data, fmt.Sprintf(`%s=%s`, k, v))
	}
	sort.Strings(data)
	signData := strings.Join(data, "&")
	fmt.Println("signData", signData)
	s := sha256.New()
	_, err := s.Write([]byte(signData))
	checkErr(err)
	hashByte := s.Sum(nil)
	signByte, err := this.PrivateKey.Sign(rand.Reader, hashByte, crypto.SHA256)
	checkErr(err)
	return base64.StdEncoding.EncodeToString(signByte), nil
}

// CheckSign 检测签名
func (this *AliAppClient) CheckSign(signData, sign string) {
	signByte, err := base64.StdEncoding.DecodeString(sign)
	checkErr(err)
	s := sha256.New()
	_, err = s.Write([]byte(signData))
	checkErr(err)
	hash := s.Sum(nil)
	err = rsa.VerifyPKCS1v15(this.PublicKey, crypto.SHA256, hash, signByte)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
