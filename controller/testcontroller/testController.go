package testcontroller

import (
	"github.com/kataras/iris/v12"
	"goTemplate/constant"
	"goTemplate/model"
	"goTemplate/service/testservice"
)

func GetHello(ctx iris.Context) {

	defer ctx.Next()

	response, requestId := model.ParseRequestValues(ctx)

	message := ctx.URLParam("message")

	responseData, res := testservice.GetHello(requestId, message)
	if constant.SUCCESS != res.Code {
		response.BuildError(requestId, res)
		model.BuildResponse(ctx, &response)
	}

	response.BuildSuccess(requestId, responseData)

	model.BuildResponse(ctx, &response)

	return
}

func PostHello(ctx iris.Context) {

	defer ctx.Next()
	response, requestId := model.ParseRequestValues(ctx)

	thisModel, res := model.BuildDemoVoByContext(requestId, ctx)
	if constant.SUCCESS != res.Code {
		response.BuildError(requestId, res)
		model.BuildResponse(ctx, &response)
	}

	responseData, res := testservice.PostHello(requestId, thisModel)
	if constant.SUCCESS != res.Code {
		response.BuildError(requestId, res)
		model.BuildResponse(ctx, &response)
	}

	response.BuildSuccess(requestId, responseData)

	model.BuildResponse(ctx, &response)

	return
}
