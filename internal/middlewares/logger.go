package middlewares

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// capture the response writer
		blw := &bodyLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = blw

		startTime := time.Now()

		c.Next()

		statusCode := c.Writer.Status()
		if statusCode >= 300 {
			log.Error().
				Int("status", statusCode).
				Str("method", c.Request.Method).
				Str("path", c.Request.URL.Path).
				Dur("restime", time.Since(startTime)).
				Str("body", blw.body.String()).
				Msg("oh no")
		}
	}
}
