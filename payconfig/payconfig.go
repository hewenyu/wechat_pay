package payconfig

import "os"

type PayConfig struct {
	appId                      string // 小程序或者公众号的appid
	appSecret                  string // 小程序或者公众号的Secret
	mchId                      string // 微信支付的商户id
	notifyUrl                  string // 微信支付的回调接口
	mchCertificateSerialNumber string // 私钥证书号
	hashkey                    string // 微信支付 APIV3 key
	privateKeyPath             string // 商户私钥证书
	payKeyPath                 string // 微信支付证书
}

func (pcf *PayConfig) AppId() string {
	return pcf.appId
}

func (pcf *PayConfig) AppSecret() string {
	return pcf.appSecret
}

func (pcf *PayConfig) MchId() string {
	return pcf.mchId
}

func (pcf *PayConfig) NotifyUrl() string {
	return pcf.notifyUrl
}

func (pcf *PayConfig) HashKey() string {
	return pcf.hashkey
}

func (pcf *PayConfig) MchCertificateSerial() string {
	return pcf.mchCertificateSerialNumber
}

func (pcf *PayConfig) PrivateKeyPath() string {
	return pcf.privateKeyPath
}

func (pcf *PayConfig) PayKeyPath() string {
	return pcf.payKeyPath
}

/**
 * NewPayConfig 微信相关参数
 * 可以自己实现
 */
func NewPayConfig() *PayConfig {

	return &PayConfig{
		appId:                      os.Getenv("WX_APPID"),
		appSecret:                  os.Getenv("WX_SECRET"),
		mchId:                      os.Getenv("WX_MCHID"),
		notifyUrl:                  os.Getenv("WX_NOTIFYURL"),
		mchCertificateSerialNumber: os.Getenv("WX_CERT_NUMBER"),
		hashkey:                    os.Getenv("WX_HASHKEY"),
		privateKeyPath:             os.Getenv("WX_PRIVATEKEYPATH"),
		payKeyPath:                 os.Getenv("WX_PAYKEYPATH"),
	}
}
