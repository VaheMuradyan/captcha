package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

func main() {
	r := gin.Default()

	// CORS Middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"}, // Թույլատրված Origin-ը
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// CAPTCHA ստեղծող API
	r.GET("/captcha", func(c *gin.Context) {
		driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
		captcha := base64Captcha.NewCaptcha(driver, store)

		id, b64s, _, err := captcha.Generate()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("CAPTCHA ID:", id)       // Debug
		fmt.Println("CAPTCHA Base64:", b64s) // Debug

		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"captcha": b64s,
		})
	})

	// CAPTCHA ստուգող API
	r.POST("/verify", func(c *gin.Context) {
		var request struct {
			ID     string `json:"id"`
			Answer string `json:"answer"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if store.Verify(request.ID, request.Answer, true) {
			c.JSON(http.StatusOK, gin.H{"message": "CAPTCHA ճիշտ է"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "CAPTCHA սխալ է"})
		}
	})

	// Սերվերի գործարկում
	r.Run(":8080")
}
