package appmetrics

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitPrometheus(r *gin.Engine) {
	// Create a new Prometheus registry
	reg := prometheus.NewRegistry()
	httpRequestsTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)
	// Create new prometheus registry
	reg.MustRegister(httpRequestsTotal)
	// Create a middleware that updates the HTTP requests counter
	r.Use(func(ctx *gin.Context) {
		// Call the next handler
		ctx.Next()
		// Update the HTTP requests counter
		httpRequestsTotal.WithLabelValues(
			ctx.Request.Method,
			ctx.Request.URL.Path,
			http.StatusText(ctx.Writer.Status()),
		).Inc()
	})
	// Register Prometheus metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.HandlerFor(reg, promhttp.HandlerOpts{})))

}
