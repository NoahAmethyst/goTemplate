package model

import (
	"github.com/kataras/iris/v12"
	"github.com/rs/zerolog/log"
	"goTemplate/constant"
	"goTemplate/entity"
)

type DemoVo struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}

func BuildDemoVoByContext(requestId string, ctx iris.Context) (DemoVo, ErrorResponse) {
	var thisModel DemoVo
	res := BuildSuccess()
	err := ctx.ReadBody(&thisModel)
	log.Info().Fields(map[string]interface{}{
		constant.REQUEST_ID: requestId,
		"action":            "parsing requestBody",
		"requestBody":       thisModel,
	}).Send()
	if err != nil {
		log.Error().Fields(map[string]interface{}{
			constant.REQUEST_ID: requestId,
			"action":            "parsing requestBody",
			"error":             err.Error(),
		}).Send()
		res = BuildError(constant.SYSTEM_ERROR, err.Error())
	}
	return thisModel, res
}

type DemoDto struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Number     int    `json:"number"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
}

func BuildDemoDtoByEntity(thisEntity entity.Demo) DemoDto {
	return DemoDto{
		Id:         thisEntity.Id,
		Name:       thisEntity.Name,
		Number:     thisEntity.Number,
		CreateTime: thisEntity.CreatedTime.Unix(),
		UpdateTime: thisEntity.UpdatedTime.Unix(),
	}
}
