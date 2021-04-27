package model

import (
	"github.com/kataras/iris/v12"
	"goTemplate/constant"
)

type ResponseModel struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	RequestId string      `json:"requestId"`
	Data      interface{} `json:"data"`
}

func (response *ResponseModel) BuildSuccess(requestId string, data interface{}) {
	response.Code = constant.SUCCESS
	response.Message = ""
	response.RequestId = requestId
	if data != nil {
		response.Data = data
	}

}

func (response *ResponseModel) BuildError(requestId string, errorResponse ErrorResponse) {
	var message string
	if 0 == len(errorResponse.Message) {
		message = constant.GetResMsg(errorResponse.Code)
	} else {
		message = errorResponse.Message
	}
	response.Code = errorResponse.Code
	response.Message = message
	response.RequestId = requestId
	response.Data = nil
}

type ErrorResponse struct {
	Code    int
	Message string
}

func BuildError(code int, message string) ErrorResponse {
	errorResponse := ErrorResponse{
		Code:    code,
		Message: message,
	}
	return errorResponse
}

func BuildSuccess() ErrorResponse {
	errorResponse := ErrorResponse{
		Code:    constant.SUCCESS,
		Message: "",
	}
	return errorResponse
}

func ParseRequestValues(ctx iris.Context) (ResponseModel, string) {
	response := ResponseModel{}
	requestId := ctx.Values().Get(constant.REQUEST_ID).(string)
	return response, requestId
}

func BuildResponse(ctx iris.Context, response *ResponseModel) {
	ctx.Values().Set(constant.RESPONSE, response)
}
