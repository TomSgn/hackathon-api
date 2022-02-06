package metrics

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func PrometheusMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		status := c.Writer.Status()
		TotalRequests.WithLabelValues(c.Request.URL.Path).Inc()
		if status == 404 || status == 401 || status == 403 {
			return
		}
		timer := prometheus.NewTimer(HttpResponseTime.WithLabelValues(c.Request.URL.Path))
		c.Next()

		ResponseStatus.WithLabelValues(fmt.Sprintf("%d", c.Writer.Status()), c.Request.Method, c.Request.URL.Path).Inc()

		timer.ObserveDuration()
	}
}
