package redisutil

import (
	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog/log"
	"time"
)

const (
	EXPIRE_SECOND = 1
	EXPIRE_MINITE = 60 * EXPIRE_SECOND
	EXPIRE_HOUR   = 60 * EXPIRE_MINITE
)

var redisClient *redis.Pool

var debug = false

func Connect(host string) {
	maxIdle := 20
	maxActive := 100
	// 建立连接池
	redisClient = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: 3600 * time.Second,
		Wait:        false,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", host,
				redis.DialDatabase(0),
				redis.DialConnectTimeout(10*time.Second),
				redis.DialReadTimeout(1*time.Second),
				redis.DialWriteTimeout(1*time.Second))
			if err != nil {
				log.Error().Err(err).Send()
				return nil, err
			}
			return con, nil
		},
	}
}

func GetClient() redis.Conn {
	if debug {
		log.Debug().Str("type", "redis").Msgf("status %v", redisClient.Stats())
	}
	return redisClient.Get()
}

func Do(cmd string, args ...interface{}) (interface{}, error) {
	if debug {
		log.Debug().Str("type", "redis").Msgf("status %v", redisClient.Stats())
	}

	rc := redisClient.Get()

	defer rc.Close()

	if debug {
		log.Debug().Str("type", "redis").Msgf("redis[%s], [%v]", cmd, args)
	}

	reply, err := rc.Do(cmd, args...)

	if err != nil {
		log.Error().Err(err).Send()
	}

	return reply, err
}

func DoStr(cmd string, args ...interface{}) string {
	c := GetClient()
	defer c.Close()
	reply, err := Do(cmd, args...)

	if str, err := redis.String(reply, err); err == nil {
		return str
	}

	return ""
}

func GetString(key string) string {
	value := DoStr("get", key)
	log.Info().Fields(map[string]interface{}{
		"action": "get redis string value",
		"key":    key,
		"value":  value,
	}).Send()
	return value
}

func DelKey(key string) string {
	value := DoStr("del", key)
	log.Info().Fields(map[string]interface{}{
		"action": "delete redis key",
		"key":    key,
	}).Send()
	return value
}

func SetString(key string, value string, expireSeconds int) {
	DoStr("set", key, value)
	log.Info().Fields(map[string]interface{}{
		"action":         "set redis string value",
		"key":            key,
		"value":          value,
		"expire seconds": expireSeconds,
	}).Send()
	_, err := Do("EXPIRE", key, expireSeconds)
	if err != nil {
		log.Error().Fields(map[string]interface{}{
			"action":         "set redis string value",
			"key":            "key",
			"value":          value,
			"expire seconds": expireSeconds,
			"error":          err.Error(),
		}).Send()
	}
}
