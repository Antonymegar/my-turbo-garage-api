package sms

import (
	"myturbogarage/helpers"

	"github.com/ochom/gutils/gttp"
)

// SendSMS ...
func SendSMS(phone, text string) error {
	apiKey := helpers.GetEnv("SMS_GATEWAY_API_KEY", "")
	senderID := helpers.GetEnv("SMS_SENDER_ID", "")

	url := "https://portal.paylifesms.com/sms/api?action=send-sms&api_key=" + apiKey + "&to=" + phone + "&from=" + senderID + "&sms=" + text
	res, err := gttp.NewRequest(url, nil, nil).Get()
	if err != nil {
		return err
	}

	if res.Status != 200 {
		return err
	}

	return nil
}
