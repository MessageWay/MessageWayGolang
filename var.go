package MessageWay

type Method int

const (
	SendUrl      = ApiBaseUrl + "/send"
	OTPVerifyUrl = ApiBaseUrl + "/otp/verify"
	StatusUrl    = ApiBaseUrl + "/status"

	MethodMessenger  = "messenger"
	MethodSMS        = "sms"
	MethodIVR        = "ivr"
	ProviderIVR      = 1
	ProviderWhatsapp = 1
	ProviderGap      = 2
	ProviderSMS3000x = 1
	ProviderSMS2000x = 2
	ProviderSMS9000x = 3

	SMS = iota + 1
	WHATSAPP
	GAP
	IVR
	SMS3000x
	SMS2000x
	SMS9000x
)
