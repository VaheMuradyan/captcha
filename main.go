package main

import (
	//"context"
	//	"log"
	//"net/http"
	//"time"

	//"github.com/dchest/captcha"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//"github.com/go-redis/redis/v8"
)

// var ctx = context.Background()
// var rdb *redis.Client

func main() {
	// initRedis()

	// redisStore := NewRedisStore(rdb, time.Minute*5, ctx)
	// captcha.SetCustomStore(redisStore)

	captchaController := NewCaptchaController()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// r.GET("/captcha", func(c *gin.Context) {
	// 	captchaId := captcha.New()

	// 	captchaDigiths := captcha.RandomDigits(6)

	// 	redisStore.Set(captchaId, captchaDigiths)

	// 	c.IndentedJSON(http.StatusOK, gin.H{"id": captchaId})
	// })

	// r.GET("/captcha/image/:id", func(c *gin.Context) {
	// 	captchaId := c.Param("id")
	// 	log.Printf("Requested CAPTCHA ID: %s\n", captchaId)

	// 	digits := redisStore.Get(captchaId, false)
	// 	if digits == nil {
	// 		c.JSON(http.StatusNotFound, gin.H{"error": "captcha not found"})
	// 		return
	// 	}

	// 	c.Header("Content-Type", "image/png")
	// 	if err := captcha.WriteImage(c.Writer, captchaId, captcha.StdWidth, captcha.StdHeight); err != nil {
	// 		log.Printf("Failed to render CAPTCHA: %v\n", err)
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to render captcha"})
	// 	}
	// })

	// r.POST("/captcha/verify", func(c *gin.Context) {

	// 	var req struct {
	// 		Id     string `json:"id"`
	// 		Answer string `json:"answer"`
	// 	}

	// 	if err := c.ShouldBindJSON(&req); err != nil {
	// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
	// 		return
	// 	}

	// 	storedCode, err := rdb.Get(ctx, req.Id).Result()
	// 	if err == redis.Nil {
	// 		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "captcha not found or expired"})
	// 		return
	// 	} else if err != nil {
	// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "redis error"})
	// 		return
	// 	}

	// 	if storedCode == req.Answer {
	// 		rdb.Del(ctx, req.Id)
	// 		c.IndentedJSON(http.StatusOK, gin.H{"status": "success"})
	// 	} else {
	// 		c.IndentedJSON(http.StatusUnauthorized, gin.H{"status": "failure"})
	// 	}
	// })

	r.GET("/captcha", captchaController.CreateCaptcha)
	r.GET("/captcha/image/:id", captchaController.RenderCaptchaImage)
	r.POST("/captcha/verify", captchaController.VerifyCaptcha)

	r.Run(":8080")
}

// func initRedis() {
// 	rdb = redis.NewClient(&redis.Options{
// 		Addr: "localhost:6379",
// 	})

// 	pong, err := rdb.Ping(ctx).Result()
// 	if err != nil {
// 		panic("Faild to connect ot Redis: " + err.Error())
// 	}

// 	log.Println("Redis connected:", pong)
// }
