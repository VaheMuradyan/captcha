package main

import (
	"net/http"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type captchaController struct {
	redisStore *RedisStore
}

func NewCaptchaController(redisStore **RedisStore) *captchaController {
	return &captchaController{
		redisStore: *redisStore,
	}
}

func (cc *captchaController) CreateCaptcha(c *gin.Context) {
	captchaId := captcha.New()

	captchaDigiths := captcha.RandomDigits(6)

	cc.redisStore.Set(captchaId, captchaDigiths)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "captcha created"})
}

// func (cc *captchaController) RenderCaptchaImage(c *gin.Context) {
// 	captchaId := c.Param("id")

// 	digits := cc.redisStore.Get(captchaId, false)
// 	if digits == nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "captcha not found"})
// 		return
// 	}

// 	c.Header("Content-Type", "image/png")
// 	if err := captcha.WriteImage(c.Writer, captchaId, captcha.StdWidth, captcha.StdHeight); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to render captcha"})
// 	}
// }

// func (cc *captchaController) RenderCaptchaImage(c *gin.Context) {
// 	captchaId := c.Param("id")

// 	// Fetch digits from Redis
// 	digits := cc.redisStore.Get(captchaId, false)
// 	if digits == nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "captcha not found"})
// 		return
// 	}

// 	// Sync the digits back to the CAPTCHA store
// 	captcha.SetCustomStore(cc.redisStore) // Ensure RedisStore is registered as the custom store
// 	cc.redisStore.Set(captchaId, digits)  // Re-set the digits to make them available to captcha.WriteImage

// 	// Render the CAPTCHA image
// 	c.Header("Content-Type", "image/png")
// 	if err := captcha.WriteImage(c.Writer, captchaId, captcha.StdWidth, captcha.StdHeight); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to render captcha"})
// 	}
// }

func (cc *captchaController) RenderCaptchaImage(c *gin.Context) {
	captchaId := c.Param("id")

	digits := cc.redisStore.Get(captchaId, false)
	if digits == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "captcha not found"})
		return
	}

	c.Header("Content-Type", "image/png")
	if err := captcha.WriteImage(c.Writer, captchaId, captcha.StdWidth, captcha.StdHeight); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to render captcha"})
	}
}

func (cc *captchaController) VerifyCaptcha(c *gin.Context) {
	var req struct {
		Id     string `json:"id"`
		Answer string `json:"answer"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	storedCode, err := rdb.Get(ctx, req.Id).Result()
	if err == redis.Nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "captcha not found or expired"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "redis error"})
		return
	}

	if storedCode == req.Answer {
		rdb.Del(ctx, req.Id)
		c.IndentedJSON(http.StatusOK, gin.H{"status": "success"})
	} else {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"status": "failure"})
	}
}
