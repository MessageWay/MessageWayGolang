package MessageWay

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	SendUrl          = ApiBaseUrl + "/send"
	MethodMessenger  = "messenger"
	MethodSMS        = "sms"
	MethodIVR        = "ivr"
	ProviderIVR      = 1
	ProviderWhatsapp = 1
	ProviderGap      = 2
	ProviderSMS3000x = 1
	ProviderSMS2000x = 2
	ProviderSMS9000x = 3
)

var (
	MobileIsRequiredErr     = errors.New("mobile not found")
	MethodIsRequiredErr     = errors.New("method not found")
	TemplateIDIsRequiredErr = errors.New("templateID not found")
	InvalidCountryCodeErr   = errors.New("countryCode is invalid")
)

type MessageBuilder interface {
	SetMobile(mobile string, countryCode ...int) MessageBuilder
	SetCountryCode(countryCode int) MessageBuilder
	SetTemplateID(templateID int) MessageBuilder
	ViaGap() MessageBuilder
	ViaWhatsapp() MessageBuilder
	ViaSMS3000x() MessageBuilder
	ViaSMS2000x() MessageBuilder
	ViaSMS9000x() MessageBuilder
	ViaIVR() MessageBuilder
	SetParams(params ...string) MessageBuilder
	Build() Message
}

type Message struct {
	//Method is required.
	Method string

	//Mobile is required.
	Mobile string

	//if you set TemplateID to 0 then default templateID use.
	TemplateID int

	//Params is optional.
	Params []string

	//CountryCode is optional.
	CountryCode int

	//Provider is optional.
	Provider int

	//Length is optional.
	Length int

	//Code is optional.
	Code string

	//if you want your otp could be expired, set ExpireTime as second.
	ExpireTime int64
}

type Builder struct {
	message Message
}

type SendResponse struct {
	Status      string                 `json:"status"`
	Error       map[string]interface{} `json:"error"`
	ReferenceID string                 `json:"referenceID"`
	Sender      string                 `json:"sender"`
}

func NewBuilder() Builder {
	return Builder{}
}

func (b Builder) SetMobile(mobile string, countryCode ...int) MessageBuilder {
	b.message.Mobile = mobile
	if len(countryCode) > 0 {
		b.message.CountryCode = countryCode[0]
	}
	return b
}

func (b Builder) SetCountryCode(countryCode int) MessageBuilder {
	b.message.CountryCode = countryCode
	return b
}

func (b Builder) SetTemplateID(templateID int) MessageBuilder {
	b.message.TemplateID = templateID
	return b
}

func (b Builder) ViaGap() MessageBuilder {
	b.message.Method = MethodMessenger
	b.message.Provider = ProviderGap
	return b
}

func (b Builder) ViaWhatsapp() MessageBuilder {
	b.message.Method = MethodMessenger
	b.message.Provider = ProviderWhatsapp
	return b
}

func (b Builder) ViaSMS3000x() MessageBuilder {
	b.message.Method = MethodSMS
	b.message.Provider = ProviderSMS3000x
	return b
}

func (b Builder) ViaSMS2000x() MessageBuilder {
	b.message.Method = MethodSMS
	b.message.Provider = ProviderSMS2000x
	return b
}

func (b Builder) ViaSMS9000x() MessageBuilder {
	b.message.Method = MethodSMS
	b.message.Provider = ProviderSMS9000x
	return b
}

func (b Builder) ViaIVR() MessageBuilder {
	b.message.Method = MethodIVR
	b.message.Provider = ProviderIVR
	return b
}

func (b Builder) SetParams(params ...string) MessageBuilder {
	if len(params) > 0 {
		b.message.Params = params
	}
	return b
}

func (b Builder) Build() Message {
	return b.message
}

func (m *Message) validate() error {
	if m.Mobile == "" {
		return MobileIsRequiredErr
	}
	if m.Method == "" {
		return MethodIsRequiredErr
	}
	if m.CountryCode < 0 {
		return InvalidCountryCodeErr
	}
	if m.TemplateID <= 0 {
		return TemplateIDIsRequiredErr
	}
	return nil
}

func (r *SendResponse) ToString() string {
	return fmt.Sprintf("%+v", r)
}

func (app *App) Send(req Request) (*SendResponse, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	buf.Write(data)
	request, err := http.NewRequest("POST", SendUrl, buf)
	if err != nil {
		return nil, err
	}
	request.Header.Add("apiKey", app.config.ApiKey)
	if app.config.AcceptLanguage == "" {
		app.config.AcceptLanguage = "fa"
	}
	request.Header.Add("accept-language", app.config.AcceptLanguage)
	r, err := app.client.Do(request)
	if err != nil {
		return nil, err
	}
	data, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	res := &SendResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
