package main

import (
	captchaimage "github.com/VaheMuradyan/captcha/captcha-image"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main() {
	captchaController := captchaimage.NewCaptchaController()

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



