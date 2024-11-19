package main

import (
	"net/http"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/captcha", func(c *gin.Context) {
		captchaId := captcha.New()
		c.IndentedJSON(http.StatusOK, gin.H{"id": captchaId})
	})

	r.GET("/captcha/image/:id", func(c *gin.Context) {
		captchaId := c.Param("id")
		c.Header("Content-Type", "image/png")
		captcha.WriteImage(c.Writer, captchaId, captcha.StdWidth, captcha.StdHeight)
	})

	r.POST("/captcha/verify", func(c *gin.Context) {
		var req struct {
			Id     string `json:"id"`
			Answer string `json:"answer"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		if captcha.VerifyString(req.Id, req.Answer) {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "CAPTCHA correct"})
		} else {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "CAPTCHA incorrect"})
		}
	})

	r.Run(":8080")
}
