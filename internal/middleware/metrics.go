package middleware

import (
	"clean-architecture-golang/internal/metric"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Metrics(metricService metric.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		appMetric := metric.NewHTTP(c.Request.URL.Path, c.Request.Method)
		appMetric.Started()
		c.Next()
		appMetric.Finished()
		appMetric.StatusCode = strconv.Itoa(c.Writer.Status())
		metricService.SaveHTTP(appMetric)
	}
}
