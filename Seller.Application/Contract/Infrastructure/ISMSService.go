package Infrastructure

type ISMSService interface {
	SendSms(destPhone string, message string)
	SendVerificationCode(destphone string, verificationCode string)
}
