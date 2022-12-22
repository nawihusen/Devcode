package helper

type Empty struct {
}

type Respnse struct {
	Status  any `json:"status"`
	Message any `json:"message"`
	Data    any `json:"data"`
}

func FailedResponseHelper(msg interface{}) Respnse {
	return Respnse{
		Status:  "Failed",
		Message: msg,
	}
}

func NotFoundHelper(msg interface{}) Respnse {
	return Respnse{
		Status:  "Not Found",
		Message: msg,
		Data:    Empty{},
	}
}

func BadRequest() Respnse {
	return Respnse{
		Status:  "Bad Request",
		Message: "title cannot be null",
		Data:    Empty{},
	}
}

func BadRequestTodo(msg string) Respnse {
	return Respnse{
		Status:  "Bad Request",
		Message: msg,
		Data:    Empty{},
	}
}

func SuccessResponseHelper(msg interface{}) Respnse {
	return Respnse{
		Status:  "Success",
		Message: msg,
		Data:    Empty{},
	}
}

func SuccessDataResponseHelper(msg, data interface{}) Respnse {
	return Respnse{
		Status:  "Success",
		Message: msg,
		Data:    data,
	}
}
