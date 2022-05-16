package MessageWay

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	StatusUrl = ApiBaseUrl + "/status"
)

type StatusRequest struct {
	ReferenceID string `json:"OTPReferenceID"`
}

type StatusResponse struct {
	Status      string                 `json:"status"`
	Error       map[string]interface{} `json:"error"`
	OTPStatus   string                 `json:"OTPStatus"`
	OTPVerified bool                   `json:"OTPVerified"`
	OTPMethod   string                 `json:"OTPMethod"`
}

func (m *StatusRequest) validate() error {
	if m.ReferenceID == "" {
		return InvalidReferenceID
	}
	return nil
}

func (r *StatusResponse) ToString() string {
	if r.Status == "success" {
		return "sms status: success\nOTP Status : " + r.OTPStatus + "\nOTP Verified : " + strconv.FormatBool(r.OTPVerified) + "\nOTP Method : " + r.OTPMethod
	}
	return "sms status: failed\nmessage : " + r.Error["message"].(string) + "\nerror code : " + strconv.FormatFloat(r.Error["code"].(float64), 'E', -1, 64)
}

func (app *App) GetStatus(req StatusRequest) (*StatusResponse, error) {
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
	request, err := http.NewRequest("POST", StatusUrl, buf)
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
	res := &StatusResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
