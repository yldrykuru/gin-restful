package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Loglama işlemini gerçekleştirin
		logrus.WithFields(logrus.Fields{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
		}).Info("Request received")

		// Sonraki middleware veya işlemleri devam ettirin
		c.Next()

		// Loglama işlemini gerçekleştirin
		logrus.WithFields(logrus.Fields{
			"status": c.Writer.Status(),
		}).Info("Response sent")
	}
}
