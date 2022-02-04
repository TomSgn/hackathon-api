package metrics

import "github.com/prometheus/client_golang/prometheus"

const (
	HostName      = "FinanceServicesGo"
	HostGroupName = "PrometheusDemo"
)

var responseLabels = prometheus.Labels{
	"resource": HostName,
	"group":    HostGroupName,
	"warning":  "2.5",
	"critical": "2.8",
}

var ResponseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "response_status",
		Help: "The status code of the response.",
	},
	[]string{"code", "method", "endpoint"},
)

var ResponseTime = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name:        "response_time",
		Help:        "Finance Services http response time average over 1 minute",
		ConstLabels: responseLabels,
	},
	[]string{"code", "method", "endpoint"},
)

func RegisterMetrics() {
	prometheus.MustRegister(ResponseStatus)
}
