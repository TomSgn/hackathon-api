package metrics

import "github.com/prometheus/client_golang/prometheus"

var ResponseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "hackathon_5_response_status",
		Help: "The status code of the response.",
	},
	[]string{"code", "method", "endpoint"},
)

var TotalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "hackathon_5_total_requests",
		Help: "The total number of requests.",
	},
	[]string{"endpoint"},
)

var HttpResponseTime = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "hackathon_5_http_response_time_seconds",
		Help: "Finance Services http response time average over 1 minute",
	},
	[]string{"code", "method", "endpoint"},
)

func RegisterMetrics() {
	prometheus.MustRegister(ResponseStatus, HttpResponseTime, TotalRequests)
}
