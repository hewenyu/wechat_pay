# wechat_pay
用于微信支付的

go get -u github.com/hewenyu/wechat_pay@v0.0.1

## 使用范例

注意配置payconfig

```go
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
```