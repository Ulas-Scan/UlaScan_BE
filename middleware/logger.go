package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log incoming request
		logrus.WithFields(logrus.Fields{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
		}).Info("Incoming request")

		c.Next()

		// Log outgoing response
		logrus.WithField("status", c.Writer.Status()).Info("Outgoing response")
	}
}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log panic
				logrus.WithField("error", err).Error("Panic recovered")

				// Abort request
				c.AbortWithStatus(500)
			}
		}()
		c.Next()
	}
}
