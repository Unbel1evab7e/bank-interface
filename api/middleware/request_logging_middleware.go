package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

type RequestLoggingMiddleware struct {
}

func (m *RequestLoggingMiddleware) Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.URL.Path, "swagger") {
			logrus.Info("Go to swagger, logging disable")
			return
		}

		logrus.Printf("Before %s request on %s", c.Request.Method, c.Request.URL.Path)

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		t := time.Now()
		c.Next()
		latency := time.Since(t)
		status := c.Writer.Status()

		logrus.Printf("Response on %s with status %s is %s with latency: %s", c.Request.URL.Path, status, blw.body.String(), latency)
	}
}
func NewLoggerMiddleware() *RequestLoggingMiddleware {
	return &RequestLoggingMiddleware{}
}
