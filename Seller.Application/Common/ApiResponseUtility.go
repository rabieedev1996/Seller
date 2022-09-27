package Common

type ResponseModel[data any] struct {
	StatusCode int
	Message    string
	Data       data
}

func (model ResponseModel[data]) ApiResponse(resultData data, statusCode int) ResponseModel[data] {
	return ResponseModel[data]{
		Data:       resultData,
		StatusCode: statusCode,
		Message:    getErrorText(statusCode),
	}
}
func getErrorText(id int) string {
	switch id {
	case 0:
		return "عملیات با موفقیت انجام شد."
	default:
		return "خطایی روی داده است."
	}
}
