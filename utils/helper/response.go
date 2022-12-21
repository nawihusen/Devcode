package helper

import "skyshi/features/activity"

func FailedResponseHelper(msg interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "Failed",
		"message": msg,
	}
}

func NotFoundHelper(msg interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "Not Found",
		"message": msg,
		"data":    activity.Activity{},
	}
}

func BadRequest() map[string]interface{} {
	return map[string]interface{}{
		"status":  "Bad Request",
		"message": "title cannot be null",
		"data":    activity.Activity{},
	}
}

func BadRequestTodo(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "Bad Request",
		"message": msg,
		"data":    activity.Activity{},
	}
}

func SuccessResponseHelper(msg interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "Success",
		"message": msg,
		"data":    activity.Activity{},
	}
}

func SuccessDataResponseHelper(msg, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "Success",
		"message": msg,
		"data":    data,
	}
}
