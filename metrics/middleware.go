package metrics

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func PrometheusMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		status := c.Writer.Status()
		if status == 404 || status == 401 || status == 403 || status == 200 {
			return
		}
		timer := prometheus.NewTimer(HttpResponseTime.WithLabelValues(fmt.Sprintf("%d", c.Writer.Status()), c.Request.Method, c.Request.URL.Path))
		fmt.Println(c.Next())

		ResponseStatus.WithLabelValues(fmt.Sprintf("%d", c.Writer.Status()), c.Request.Method, c.Request.URL.Path).Inc()
		TotalRequests.WithLabelValues(c.Request.URL.Path).Inc()

		timer.ObserveDuration()
	}
}
