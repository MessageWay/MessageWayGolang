package MessageWay_test

import (
	"os"
	"testing"
	"time"

	MessageWay "github.com/MessageWay/MessageWayGolang"
)

var ApiKey string

func init() {
	ApiKey = os.Getenv("MESSAGE_WAY_APIKEY")
}

func TestSend(t *testing.T) {
	app := MessageWay.New(MessageWay.Config{
		ApiKey:  ApiKey,
		Timeout: 5 * time.Second,
	})
	message := MessageWay.NewBuilder().SetMobile("09123456789").SetParams("foo", "doo", "loo").ViaWhatsapp().Build()
	res, err := app.Send(message)
	if err != nil {
		t.Errorf("sending sms failed, [%v]", err)
		return
	}
	if res.ReferenceID == "" {
		t.Errorf("sending sms failed: " + res.ToString())
	}
}

func TestGetStatus(t *testing.T) {
	app := MessageWay.New(MessageWay.Config{ApiKey: ApiKey})
	res, err := app.GetStatus(MessageWay.StatusRequest{
		ReferenceID: "123",
	})
	if err != nil {
		t.Errorf("verify error accurred %e", err)
		return
	}
	if res.OTPStatus != "error" {
		t.Errorf("verify not failed")
		return
	}
}

func TestVerify(t *testing.T) {
	app := MessageWay.New(MessageWay.Config{ApiKey: ApiKey})
	res, err := app.Verify(MessageWay.OTPVerifyRequest{
		CountryCode: 0,
		Mobile:      "9123456789",
		OTP:         "xxx",
	})
	if err != nil {
		t.Errorf("verify error accurred%e", err)
		return
	}
	if res.Status != "error" {
		t.Errorf("verify not failed")
		t.Log(res)
		return
	}
}
