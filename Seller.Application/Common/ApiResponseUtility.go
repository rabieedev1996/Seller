package Common

type ResponseModel[data any] struct {
	StatusCode int
	Message    string
	Data       data
}

const (
	SUCCESS = iota
	INVALID_ACTIVATION_CODE
)

func (model ResponseModel[data]) ApiResponse(resultData data, statusCode int) ResponseModel[data] {
	return ResponseModel[data]{
		Data:       resultData,
		StatusCode: statusCode,
		Message:    getErrorText(statusCode),
	}
}
func getErrorText(id int) string {
	switch id {
	case SUCCESS:
		return "عملیات با موفقیت انجام شد."
	case INVALID_ACTIVATION_CODE:
		return "کد فعالسازی وارد شده صحیح نیست."
	default:
		return "خطایی روی داده است."
	}
}
