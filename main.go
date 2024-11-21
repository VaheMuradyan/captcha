package main

import (
	"context"
	"log"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func main() {
	initRedis()

	redisStore := NewRedisStore(rdb, time.Minute*5, ctx)
	captchaController := NewCaptchaController(&redisStore)
	captcha.SetCustomStore(redisStore)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/captcha", captchaController.CreateCaptcha)

	r.GET("/captcha/image/:id", captchaController.RenderCaptchaImage)

	r.POST("/captcha/verify", captchaController.VerifyCaptcha)

	r.Run(":8080")
}

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic("Faild to connect ot Redis: " + err.Error())
	}

	log.Println("Redis connected:", pong)
}
