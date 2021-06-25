package main

import (
	"fmt"
	"log"

	payclientgo "github.com/hewenyu/wechat_pay/payclient.go"
)

func main() {
	payCli := payclientgo.NewPayClient("order string", 10, "openid string")

	// 发起支付
	res, err := payCli.JSAPI()

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(res)
}
