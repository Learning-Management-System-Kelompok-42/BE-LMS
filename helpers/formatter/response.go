package formatter

type BaseResponseFailed struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BaseResponseSuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(data interface{}) BaseResponseSuccess {
	return BaseResponseSuccess{
		Code:    200,
		Message: "Success",
		Data:    data,
	}
}

func CreateSuccessResponse(data interface{}) BaseResponseSuccess {
	return BaseResponseSuccess{
		Code:    201,
		Message: "Success",
		Data:    data,
	}
}

func BadRequestResponse(message string) BaseResponseFailed {
	return BaseResponseFailed{
		Code:    400,
		Message: message,
	}
}

func UnauthorizedResponse(message string) BaseResponseFailed {
	return BaseResponseFailed{
		Code:    401,
		Message: message,
	}
}

func ForbiddenResponse(message string) BaseResponseFailed {
	return BaseResponseFailed{
		Code:    403,
		Message: message,
	}
}

func NotFoundResponse(message string) BaseResponseFailed {
	return BaseResponseFailed{
		Code:    404,
		Message: message,
	}
}

func ConflictResponse(message string) BaseResponseFailed {
	return BaseResponseFailed{
		Code:    409,
		Message: message,
	}
}

func InternalServerErrorResponse(message string) BaseResponseFailed {
	return BaseResponseFailed{
		Code:    500,
		Message: message,
	}
}
