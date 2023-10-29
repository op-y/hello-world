package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	called = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_called_total",
			Help: "Total number of API called by path.",
		},
		[]string{"path"})

	income = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "income",
			Help: "Last income number ",
		})

	duration = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "request_duration_ms",
			Help:    "Income API request duration ms",
			Buckets: []float64{10, 100, 200, 500, 1000},
		})

	summary = promauto.NewSummary(
		prometheus.SummaryOpts{
			Name: "request_duration_pct",
			Help: "Income API request duration pct",
		})
)

func main() {
	//gin.DisableConsoleColor()
	gin.ForceConsoleColor()

	r := gin.Default()
	//r := gin.New()

	r.GET("/ping", func(c *gin.Context) {
		called.WithLabelValues("ping").Inc()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/income/:num", func(c *gin.Context) {
		p := c.Param("num")

		start := time.Now().UnixMilli()
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		called.WithLabelValues("income").Inc()
		num, err := strconv.Atoi(p)
		if err != nil {
			income.Set(0)
			summary.Observe(0)
		} else {
			income.Set(float64(num))
			summary.Observe(float64(num))
		}
		d := time.Now().UnixMilli() - start
		duration.Observe(float64(d))
		c.JSON(200, gin.H{
			"profit": p,
		})
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	pprof.Register(r)
	r.Run()
}
