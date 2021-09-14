package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"time"

	"github.com/eiiches/go-gen-proxy/pkg/interceptor"
	"github.com/foo/bar/pkg/greeter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Interceptor

type PrometheusInterceptor struct {
	CallDurations *prometheus.HistogramVec

	CallsTotal      *prometheus.CounterVec
	CallErrorsTotal *prometheus.CounterVec
}

func NewPrometheusInterceptor(namespace string) *PrometheusInterceptor {
	return &PrometheusInterceptor{
		CallDurations: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Name:      "call_duration_seconds",
				Help:      "time took to complete the method call",
				Buckets:   prometheus.DefBuckets,
			},
			[]string{"method"},
		),
		CallsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "calls_total",
				Help:      "the number of total calls to the method. incremented before the actual method call.",
			},
			[]string{"method"},
		),
		CallErrorsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "call_errors_total",
				Help:      "the number of total errors returned from the method. incremented after the method call is ended.",
			},
			[]string{"method"},
		),
	}
}

func (this *PrometheusInterceptor) RegisterTo(registerer prometheus.Registerer) {
	registerer.Register(this.CallDurations)
	registerer.Register(this.CallsTotal)
	registerer.Register(this.CallErrorsTotal)
}

func canReturnError(method reflect.Value) bool {
	if method.Type().NumOut() == 0 {
		return false
	}
	lastReturnType := method.Type().Out(method.Type().NumOut() - 1)
	return lastReturnType.Name() == "error" && lastReturnType.PkgPath() == ""
}

func (this *PrometheusInterceptor) Intercept(receiver interface{}, method string, args []interface{}, delegate func([]interface{}) []interface{}) []interface{} {
	r := reflect.ValueOf(receiver)
	m := r.MethodByName(method)

	this.CallsTotal.WithLabelValues(method).Inc()

	t0 := time.Now()

	rets := delegate(args)

	if canReturnError(m) && rets[m.Type().NumOut()-1] != nil {
		this.CallErrorsTotal.WithLabelValues(method).Inc()
	}

	seconds := time.Since(t0).Seconds()

	this.CallDurations.WithLabelValues(method).Observe(seconds)

	return rets
}

// Greeter

type GreeterImpl struct {
	Random *rand.Rand
}

func (this *GreeterImpl) SayHello(name string) (string, error) {
	nanos := this.Random.Float32() * float32(time.Second.Nanoseconds())
	time.Sleep(time.Duration(nanos * float32(time.Nanosecond)))
	if this.Random.Float32() < 0.5 {
		return "", fmt.Errorf("failed to say hello")
	}
	return fmt.Sprintf("Hello %s!", name), nil
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	prom := NewPrometheusInterceptor("greeter")
	prom.RegisterTo(prometheus.DefaultRegisterer)

	p := &greeter.GreeterProxy{
		Handler: &interceptor.InterceptingInvocationHandler{
			Delegate:    &GreeterImpl{Random: r},
			Interceptor: prom,
		},
	}

	go func() {
		for {
			msg, err := p.SayHello("James")
			fmt.Println(msg, err)
			time.Sleep(1 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	))
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
