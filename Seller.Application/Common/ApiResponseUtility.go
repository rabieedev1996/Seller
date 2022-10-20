package Common

type ResponseModel[data any] struct {
	StatusCode int
	Message    string
	Data       *data
}

const (
	STATUS_SUCCESS = iota
	STATUS_INVALID_ACTIVATION_CODE
	STATUS_FILE_EMPETY
	STATUS_Exception
)

func (model ResponseModel[data]) ApiResponse(resultData *data, statusCode int) ResponseModel[data] {
	return ResponseModel[data]{
		Data:       resultData,
		StatusCode: statusCode,
		Message:    getErrorText(statusCode),
	}
}

func getErrorText(id int) string {
	switch id {
	case STATUS_SUCCESS:
		return "عملیات با موفقیت انجام شد."
	case STATUS_INVALID_ACTIVATION_CODE:
		return "کد فعالسازی وارد شده صحیح نیست."
	case STATUS_FILE_EMPETY:
		return "فایل وارد نشده است."
	case STATUS_Exception:
		return "خطایی روی داده است. لطفا دقایقی دیگر مجددا اثدام نمایید."
	default:
		return "خطایی روی داده است."
	}
}
