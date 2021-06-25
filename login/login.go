package login

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/hewenyu/toolspackage/release/zlog"
	"github.com/hewenyu/wechat_pay/payconfig"
)

type WxClient struct {
	httpClient *http.Client
}

/**
 * NewPayClient
 * 创建登陆客户端
 */
func NewPayClient(httpClient *http.Client) *WxClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	if httpClient == nil {

		httpClient = http.DefaultClient
		httpClient.Timeout = time.Second * 5
		httpClient.Transport = tr
	}

	return &WxClient{
		httpClient: httpClient,
	}
}

/**
 * Login
 * 微信发起登陆请求
 */
func (pc *WxClient) Login(jscode string) (*PayData, error) {

	payConfig := payconfig.NewPayConfig()

	loginUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=" + url.QueryEscape(payConfig.AppId()) +
		"&secret=" + url.QueryEscape(payConfig.AppSecret()) +
		"&js_code=" + url.QueryEscape(jscode) +
		"&grant_type=authorization_code"

	httpResp, err := pc.httpClient.Get(loginUrl)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http.Status: %s", httpResp.Status)
	}

	respData := NewPayData()

	err = respData.FromJson(httpResp.Body)
	if err != nil {
		return nil, err
	}

	zlog.Zap().Info("-----微信登陆结果输出-----")

	zlog.Zap().Info(respData.ToJson())

	return respData, nil
}
