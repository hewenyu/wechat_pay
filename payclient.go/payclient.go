package payclient

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"time"

	"github.com/hewenyu/toolspackage/release/zlog"
	"github.com/hewenyu/wechat_pay/payconfig"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

type PayClient struct {
	Order  string // 商户的订单
	Pays   int32  // 订单号的金额,单位分
	Openid string // 支付用户的openid
}

/**
 * NewPayClient
 * 微信支付二次封装
 */
func NewPayClient(order string, pays int32, openid string) *PayClient {
	return &PayClient{
		Order:  order,
		Pays:   pays,
		Openid: openid,
	}
}

/**
 * JSAPI
 * 微信调用微信支付
 */
func (p *PayClient) JSAPI() (resp *jsapi.PrepayResponse, err error) {

	payConfig := payconfig.NewPayConfig()

	// 通过私钥的文件路径内容加载私钥
	privateKey, err := utils.LoadPrivateKeyWithPath(payConfig.PrivateKeyPath())
	if err != nil {
		zlog.Zap().Info(err.Error())
		return nil, err
	}
	public, err := utils.LoadCertificateWithPath(payConfig.PayKeyPath())
	if err != nil {
		zlog.Zap().Info(err.Error())

		return nil, err
	}

	// 增加客户端配置
	ctx := context.Background()
	// 忽略证书的报错
	clientSS := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	opts := []core.ClientOption{
		core.WithHTTPClient(clientSS),     // 可以不设置
		core.WithTimeout(time.Second * 2), // 自行进行超时时间配置
		core.WithMerchantCredential(payConfig.MchId(), payConfig.MchCertificateSerial(), privateKey), // 设置商户信息，用于生成签名信息
		core.WithWechatPayValidator([]*x509.Certificate{public}),                                     // 设置微信支付平台证书信息，对回包进行校验
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		zlog.Zap().Info(err.Error())
		return nil, err
	}
	svc := jsapi.JsapiApiService{Client: client}

	resp, result, err := svc.Prepay(ctx,
		jsapi.PrepayRequest{
			Appid:       core.String(payConfig.AppId()),
			Mchid:       core.String(payConfig.MchId()),
			Description: core.String("商品支付"),
			OutTradeNo:  core.String(p.Order),
			Attach:      core.String("自定义数据说明"),
			NotifyUrl:   core.String(payConfig.NotifyUrl()),
			Amount: &jsapi.Amount{
				Total: core.Int32(p.Pays),
			},
			Payer: &jsapi.Payer{
				Openid: core.String(p.Openid),
			},
		},
	)

	defer result.Request.Body.Close()

	return
}
