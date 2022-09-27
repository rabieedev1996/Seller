package ServiceModel

type FarazSmsVerificationSMSRequest struct {
	PatternCode string                        `json:"pattern_code"`
	Originator  string                        `json:"originator"`
	Recipient   string                        `json:"recipient"`
	Value       FarazSmsVerificationSMSValues `json:"values"`
}
type FarazSmsVerificationSMSValues struct {
	VerificationCode string `json:"verification-code"`
}
