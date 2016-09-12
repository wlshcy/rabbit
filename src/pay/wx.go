package pay

const (
	APP_ID     = "57d12695421aa92ece939d64"
	MCH_ID     = "57d12695421aa92ece939d64"
	NONCE_STR  = "5K8264ILTKCH16CQ2502SI8ZNMTM67VS"
	NOTIFY_URL = "http://freshtest.me:8080/v1/wxpay"
)

type WXPay struct {
}

func (wx *WXPay) SendUnifiedOrderReq(totalFee int, orderNumber string) (string, error) {
	req := &UnifiedOrderReq{
		AppId:          APP_ID,
		Body:           "支付",
		MchId:          MCH_ID,
		NonceStr:       NONCE_STR,
		NotifyUrl:      NOTIFY_URL,
		TradeType:      "APP",
		SpbillCreateIp: "0.0.0.0",
		TotalFee:       totalFee,
		OutTradeNo:     orderNumber,
	}
}
