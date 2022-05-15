package MessageWay_test

import (
	MessageWay "ehsansabet/MessageWayGolang"
	"os"
	"testing"
)

var ApiKey string

func init() {
	ApiKey = os.Getenv("MESSAGE_WAY_APIKEY")
}

func TestApp_Send1(t *testing.T) {
	config := MessageWay.Config{
		ApiKey:ApiKey,
	}
	app := MessageWay.New(config)
	message := MessageWay.NewBuilder().SetMobile("09123456789").SetParams("foo","doo","loo").ViaWhatsapp().Build()

	res, err := app.Send(&message)
	if err != nil {
		t.Errorf("sending sms failed, [%v]", err)
		return
	}
	if res.ReferenceID == "" {
		t.Errorf("sending sms failed: " + res.ToString())
	}
}
