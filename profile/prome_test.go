package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func Test_Count(t *testing.T) {
	recordCounterMetrics()
	recordGaugeMetrics()
	recordHistogramMetrics()

	client := http.Client{}
	client.Post()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":12345", nil)
}

func recordCounterMetrics() {
	opsProcessed := promauto.NewCounter(prometheus.CounterOpts{
		Name:        "hanyu_processed_ops_total",
		Help:        "The total number of processed events",
		ConstLabels: map[string]string{"hanyu_test": "bbb"},
	})
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func recordGaugeMetrics() {
	gauge := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hanyu_gauge_total",
		Help: "gauge",
	}, []string{"pool"})
	go func() {
		labels := []string{"A", "B", "C"}
		poolWeights := []float64{10, 20, 40}
		for {
			index := rand.Intn(len(labels))
			randW := rand.Intn(int(poolWeights[index]))
			gauge.WithLabelValues(labels[index]).Set(float64(randW))
			time.Sleep(1 * time.Second)
		}
	}()
}

func recordHistogramMetrics() {
	myTestHistogram := prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "hanyu_pond_temperature_celsius",
		Help:    "hanyu_pond_temperature_celsius",   // Sorry, we can't measure how badly it smells.
		Buckets: prometheus.LinearBuckets(20, 5, 5), // 5 buckets, each 5 centigrade wide.
	})
	prometheus.MustRegister(myTestHistogram)
	go func() {
		i := 0
		for {
			i++
			myTestHistogram.Observe(30 + math.Floor(120*math.Sin(float64(i)*0.1))/10) //每次观察一个18 - 42的量
			time.Sleep(100 * time.Millisecond)
		}
	}()
}
