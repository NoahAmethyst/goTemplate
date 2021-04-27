package testservice

import (
	"github.com/rs/zerolog/log"
	"goTemplate/constant"
	"goTemplate/dao/demodao"
	"goTemplate/entity"
	"goTemplate/model"
	"goTemplate/utils/datetime"
	"goTemplate/utils/encrypt"
	"goTemplate/utils/system"
	"strconv"
	"time"
)

func GetHello(requestId string, message string) (map[string]string, model.ErrorResponse) {
	log.Info().Fields(map[string]interface{}{
		constant.REQUEST_ID: requestId,
		"action":            "get hello",
	}).Send()

	res := model.BuildSuccess()

	response := make(map[string]string)
	response["nowTime"] = datetime.TimeToString(time.Now(), datetime.STANDRD_TIME)
	response["cpuNumber"] = strconv.Itoa(system.GetCpuNumber())
	response["message"] = message

	return response, res
}

func PostHello(requestId string, thisModel model.DemoVo) (model.DemoDto, model.ErrorResponse) {
	log.Info().Fields(map[string]interface{}{
		constant.REQUEST_ID: requestId,
		"action":            "post hello",
	}).Send()

	thisEntity, res := CreateDemo(requestId, thisModel)

	dto := model.BuildDemoDtoByEntity(thisEntity)

	return dto, res

}

func CreateDemo(requestId string, thisModel model.DemoVo) (entity.Demo, model.ErrorResponse) {
	log.Info().Fields(map[string]interface{}{
		constant.REQUEST_ID: requestId,
		"action":            "create demo",
	}).Send()

	res := model.BuildSuccess()

	thisEntity := buildDemoEntityByVo(thisModel)
	err := demodao.InsertDemo(&thisEntity)
	if err != nil {
		res = model.BuildError(constant.SYSTEM_ERROR, err.Error())
	}
	return thisEntity, res

}

func buildDemoEntityByVo(thisModel model.DemoVo) entity.Demo {
	return entity.Demo{
		Id:         encrypt.GenerateUuid(false),
		Name:       thisModel.Name,
		Number:     thisModel.Number,
		BaseEntity: entity.InitBaseEntity(),
	}

}
