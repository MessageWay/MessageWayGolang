package MessageWay

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type OTPVerifyRequest struct {
	OTP         string `json:"otp"`
	CountryCode int    `json:"countryCode"`
	Mobile      string `json:"mobile"`
}

type OTPVerifyResponse struct {
	Status string                 `json:"status"`
	Error  map[string]interface{} `json:"error"`
}

func (m *OTPVerifyRequest) validate() error {
	if m.Mobile == "" {
		return MobileIsRequiredErr
	}
	if m.OTP == "" {
		return CodeIsRequiredErr
	}
	if m.CountryCode < 0 {
		return InvalidCountryCodeErr
	}
	return nil
}

func (r *OTPVerifyResponse) ToString() string {
	if r.Status == "success" {
		return "verify status: success\nsent : true"
	}
	return "verify status: failed\nmessage : " + r.Error["message"].(string) + "\nerror code : " + strconv.FormatFloat(r.Error["code"].(float64), 'E', -1, 64)
}

func (app *App) Verify(req OTPVerifyRequest) (*OTPVerifyResponse, error) {
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
	request, err := http.NewRequest("POST", OTPVerifyUrl, buf)
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
	res := &OTPVerifyResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
