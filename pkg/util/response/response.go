package response

type Response struct {
	StatusCode int         `json:"status-code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      interface{} `json:"error"`
}

func ResponseMessage(statusCode int, message string, data interface{}, err interface{}) Response {

	return Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Error:      err,
	}

}
