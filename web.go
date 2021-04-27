package goTemplate

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/rs/zerolog/log"
	"goTemplate/constant"
	"goTemplate/controller/testcontroller"
	"goTemplate/utils/db"
	"goTemplate/utils/encrypt"
	"goTemplate/utils/redisutil"
	"os"
	"time"
)

const (
	REDIS_HOST_PORT = "REDIS_HOST_PORT"
	DB_HOST         = "DB_HOST"
	DB_PORT         = "DB_PORT"
	DB_NAME         = "DB_NAME"
	DB_USER         = "DB_USER"
	DB_PASSWORD     = "DB_PASSWORD"
	SERVER_PORT     = "SERVER_PORT"
)

func beforeProcess(ctx iris.Context) {
	ctx.Values().Set(constant.REQUEST_ID, encrypt.GenerateUuid(false))
	ctx.Values().Set("nowTime", time.Now().Unix())
	ctx.Next()
}

func endProcess(ctx iris.Context) {
	endTime := time.Now().Unix()
	nowTime := ctx.Values().Get("nowTime").(int64)
	requestId := ctx.Values().Get(constant.REQUEST_ID).(string)

	log.Info().Fields(map[string]interface{}{
		constant.REQUEST_ID: requestId,
		"path":              ctx.Path(),
		"time consuming":    endTime - nowTime,
	}).Send()
	ctx.JSON(ctx.Values().Get(constant.RESPONSE))
	return
}

func main() {

	initializeConfig()
	go initializeTask()
	app := iris.New()

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //允许通过的主机名称
		AllowCredentials: true,
	})

	pokerChainAPI := app.Party("/template", crs).AllowMethods(iris.MethodOptions)
	{
		pokerChainAPI.Use(beforeProcess)

		pokerChainAPI.Get("/test/getHello", testcontroller.GetHello, endProcess)

		pokerChainAPI.Post("/test/postHello", testcontroller.PostHello, endProcess)

	}

	server_port := os.Getenv(SERVER_PORT)
	_ = app.Listen(server_port)

}

func initializeConfig() {
	// connect to redis
	redisHostPort := os.Getenv(REDIS_HOST_PORT)
	redisutil.Connect(redisHostPort)

	// connect to ccpay database
	dbHost := os.Getenv(DB_HOST)
	dbPort := os.Getenv(DB_PORT)
	dbName := os.Getenv(DB_NAME)
	dbUser := os.Getenv(DB_USER)
	dbPassword := os.Getenv(DB_PASSWORD)
	err := db.ConnectDb(dbHost, dbPort, dbName, dbUser, dbPassword)
	if err != nil {
		log.Error().Fields(map[string]interface{}{
			"action": "connect to db",
			"error":  err.Error(),
		})
	}
}

func initializeTask() {

}
