package MessageWay

import "errors"

var (
	CodeIsRequiredErr       = errors.New("code not found")
	MobileIsRequiredErr     = errors.New("mobile not found")
	InvalidCountryCodeErr   = errors.New("countryCode is invalid")
	InvalidReferenceID      = errors.New("referenceID not found")
	MethodIsRequiredErr     = errors.New("method not found")
	TemplateIDIsRequiredErr = errors.New("templateID not found")
	InvalidHashFormatErr    = errors.New("hash format is invalid")
)
