package metrics

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func PrometheusMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		status := c.Writer.Status()
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w
		TotalRequests.WithLabelValues(c.Request.URL.Path).Inc()
		if status == 404 || status == 401 || status == 403 {
			return
		}
		timer := prometheus.NewTimer(HttpResponseTime.WithLabelValues(c.Request.URL.Path))
		c.Next()

		if c.Request.URL.Path == "/statistics" {
			var myBody map[string]interface{}
			json.Unmarshal(w.body.Bytes(), &myBody)
			log.Println(myBody)
			t := myBody["stats"].(map[string]interface{})
			log.Println(t)
			log.Println(myBody["total"])
			CurrencyValue.WithLabelValues(
				myBody["stats"].(map[string]interface{})["money"].(string),
			).Set(myBody["total"].(float64))
		}

		ResponseStatus.WithLabelValues(fmt.Sprintf("%d", c.Writer.Status()), c.Request.Method, c.Request.URL.Path).Inc()

		timer.ObserveDuration()
	}
}
