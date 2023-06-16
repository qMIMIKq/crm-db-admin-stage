package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"time"
)

func CheckAuthMw(c *gin.Context) {
	//url := c.Request.Referer()
	//port := strings.Split(url[:len(url)-1], ":")[2]

	//store := ginSession.FromContext(c)
	//
	//_, ok := store.Get("userID")
	//if !ok {
	//	c.Redirect(http.StatusFound, "/login")
	//	return
	//}
}

func LoggerMiddleware(c *gin.Context) {
	log.Info().Msgf("Method [%s] - time [%s]", c.Request.Method, time.Now().Format(time.DateTime))
	c.Next()
}
