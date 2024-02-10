package responsebuilder

//Response is used for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type Response_table struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Count   int64       `json:"count"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type Response_login struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
	Token   string      `json:"token"`
}

type ResponseBpjs struct {
	Metadata struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"metadata"`
	Response interface{} `json:"response"`
}

//EmptyObj object is used when data doesnt want to be null on json
type EmptyObj struct{}

//BuildResponse method is to inject data value to dynamic success response
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

func BuildResponse_table(status bool, message string, count int64, data interface{}) Response_table {
	res := Response_table{
		Status:  status,
		Message: message,
		Count:   count,
		Errors:  nil,
		Data:    data,
	}
	return res
}

func BuildResponseLogin(status bool, message string, token string, data interface{}) Response_login {
	res := Response_login{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
		Token:   token,
	}
	return res
}

//BuildErrorResponse method is to inject data value to dynamic failed response
func BuildErrorResponse(message string, err interface{}, data interface{}) Response {
	// splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  err,
		Data:    data,
	}
	return res
}

func BuildResponseBpjs(message_ string, Code_ int, Response interface{}) ResponseBpjs {
	res := ResponseBpjs{
		Metadata: struct {
			Message string `json:"message"`
			Code    int    `json:"code"`
		}{Message: message_,
			Code: Code_},
		Response: Response,
	}
	return res
}
