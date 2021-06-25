package router

import (
	"encoding/json"
	"io"
	"strings"
	"time"

	"github.com/hewenyu/gin-simple/logger"
	"github.com/hewenyu/wechat_pay/payconfig"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// 微信回调信息处理
type WX_NOTIFY struct {
	ID           string    `json:"id"`
	CreateTime   time.Time `json:"create_time"`
	ResourceType string    `json:"resource_type"`
	EventType    string    `json:"event_type"`
	Summary      string    `json:"summary"`
	Resource     struct {
		OriginalType   string `json:"original_type"`
		Algorithm      string `json:"algorithm"`
		Ciphertext     string `json:"ciphertext"`
		AssociatedData string `json:"associated_data"`
		Nonce          string `json:"nonce"`
	} `json:"resource"`
}

// 微信回调信息解析
type LoadPayInfo struct {
	Mchid          string    `json:"mchid"`
	Appid          string    `json:"appid"`
	OutTradeNo     string    `json:"out_trade_no"`
	TransactionID  string    `json:"transaction_id"`
	TradeType      string    `json:"trade_type"`
	TradeState     string    `json:"trade_state"`
	TradeStateDesc string    `json:"trade_state_desc"`
	BankType       string    `json:"bank_type"`
	Attach         string    `json:"attach"`
	SuccessTime    time.Time `json:"success_time"`
	Payer          struct {
		Openid string `json:"openid"`
	} `json:"payer"`
	Amount struct {
		Total         int    `json:"total"`
		PayerTotal    int    `json:"payer_total"`
		Currency      string `json:"currency"`
		PayerCurrency string `json:"payer_currency"`
	} `json:"amount"`
}

/**
 * UnSignature
 * 加密数据解密
 */
func (s *WX_NOTIFY) UnSignature() (*LoadPayInfo, error) {

	payconfig := payconfig.NewPayConfig()

	return_info, err := utils.DecryptAES256GCM(payconfig.HashKey(), s.Resource.AssociatedData, s.Resource.Nonce, s.Resource.Ciphertext)

	if err != nil {
		logger.Zap().Info(err.Error())

		return nil, err
	}

	logger.Zap().Info(return_info)

	var load LoadPayInfo

	dec := json.NewDecoder(strings.NewReader(return_info))

	if err := dec.Decode(&load); err == io.EOF {
		return nil, err
	} else if err != nil {
		return nil, err
	}
	return &load, err
}
