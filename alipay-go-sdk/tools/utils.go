package tools

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net"
	"strconv"
	"time"
)

//RandomStr 获取一个随机字符串
func RandomStr() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// LocalIP 获取机器的IP
func LocalIP() string {
	info, _ := net.InterfaceAddrs()
	for _, addr := range info {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}
	return ""
}

func MapStringToStruct(m map[string]string, i interface{}) error {
	bin, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bin, i)
	if err != nil {
		return err
	}
	return nil
}

func Bytes2RSAPrivateKey(priKey []byte) interface{} {
	block, _ := pem.Decode(priKey)
	if block == nil {
		fmt.Println("Sign private key decode error")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
	}
	return privateKey
}

func Bytes2RSAPublicKey(pubKey []byte) interface{} {
	block, _ := pem.Decode(pubKey)
	if block == nil {
		fmt.Println("Sign pubilc key decode error")
	}
	pubilcKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
	}
	return pubilcKey
}

func NewOutTradeNo(uid int64, tType string) string {
	tn := ""
	tn += tType
	tn += strconv.Itoa(int(uid))
	tn += NewUUID().Hex()
	return tn
}
