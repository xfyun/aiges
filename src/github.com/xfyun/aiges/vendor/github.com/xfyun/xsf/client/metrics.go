package xsf

import (
	"fmt"
	metricCollector "github.com/xfyun/xsf/client/internal/metric_collector"
	"github.com/xfyun/xsf/client/internal/rolling"
	"sync"
	"time"
)

type commandExecution struct {
	Types            []string      `json:"types"`
	Start            time.Time     `json:"start_time"`
	RunDuration      time.Duration `json:"run_duration"`
	ConcurrencyInUse float64       `json:"concurrency_inuse"`
}

type metricExchange struct {
	parent  *CircuitBreaker
	Name    string
	Updates chan *commandExecution
	Mutex   *sync.RWMutex

	metricCollectors []metricCollector.MetricCollector
}

func newMetricExchange(parent *CircuitBreaker, name string) *metricExchange {
	m := &metricExchange{}
	m.Name = name
	m.parent = parent
	m.Updates = make(chan *commandExecution, 10000) //todo hard code not complete
	m.Mutex = &sync.RWMutex{}
	m.metricCollectors = metricCollector.Registry.InitializeMetricCollectors(name, m.getSettings(name).SleepWindow)
	m.Reset()

	go m.Monitor()

	return m
}

// The Default Collector function will panic if collectors are not setup to specification.
func (m *metricExchange) DefaultCollector() *metricCollector.DefaultMetricCollector {
	if len(m.metricCollectors) < 1 {
		panic("No Metric Collectors Registered.")
	}
	collection, ok := m.metricCollectors[0].(*metricCollector.DefaultMetricCollector)
	if !ok {
		panic("Default metric collector is not registered correctly. The default metric collector must be registered first.")
	}
	return collection
}

func (m *metricExchange) Monitor() {
	for update := range m.Updates {
		// we only grab a read lock to make sure Reset() isn't changing the numbers.
		m.Mutex.RLock()

		totalDuration := time.Since(update.Start)
		wg := &sync.WaitGroup{}
		for _, collector := range m.metricCollectors {
			wg.Add(1)
			//todo 潜在的优化点
			go m.IncrementMetrics(wg, collector, update, totalDuration)
		}
		wg.Wait()

		m.Mutex.RUnlock()
	}
}

func (m *metricExchange) IncrementMetrics(wg *sync.WaitGroup, collector metricCollector.MetricCollector, update *commandExecution, totalDuration time.Duration) {
	// granular metrics
	r := metricCollector.MetricResult{
		Attempts:         1,
		TotalDuration:    totalDuration,
		RunDuration:      update.RunDuration,
		ConcurrencyInUse: update.ConcurrencyInUse,
	}

	switch update.Types[0] {
	case "success":
		r.Successes = 1
	case "failure":
		r.Failures = 1
		r.Errors = 1
	case "rejected":
		r.Rejects = 1
		r.Errors = 1
	case "short-circuit":
		r.ShortCircuits = 1
		r.Errors = 1
	case "timeout":
		r.Timeouts = 1
		r.Errors = 1
	case "context_canceled":
		r.ContextCanceled = 1
	case "context_deadline_exceeded":
		r.ContextDeadlineExceeded = 1
	}

	if len(update.Types) > 1 {
		// fallback metrics
		if update.Types[1] == "fallback-success" {
			r.FallbackSuccesses = 1
		}
		if update.Types[1] == "fallback-failure" {
			r.FallbackFailures = 1
		}
	}

	collector.Update(r)

	wg.Done()
}

func (m *metricExchange) Reset() {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	for _, collector := range m.metricCollectors {
		collector.Reset()
	}
}

func (m *metricExchange) Requests() *rolling.Number {
	m.Mutex.RLock()
	defer m.Mutex.RUnlock()
	return m.requestsLocked()
}

func (m *metricExchange) requestsLocked() *rolling.Number {
	return m.DefaultCollector().NumRequests()
}

func (m *metricExchange) ErrorPercent(now time.Time) int {
	m.Mutex.RLock()
	defer m.Mutex.RUnlock()

	var errPct float64
	reqs := m.requestsLocked().Sum(now)
	errs := m.DefaultCollector().Errors().Sum(now)

	if reqs > 0 {
		errPct = (errs / reqs) * 100
	}
	return int(errPct + 0.5)
}
func (m *metricExchange) getSettings(name string) *Settings {
	return m.parent.parent.parent.getSettings(name)
}
func (m *metricExchange) IsHealthy(now time.Time) (bool, error) {
	errorPercent := m.ErrorPercent(now)
	errorPercentThreshold := m.getSettings(m.Name).ErrorPercentThreshold

	if errorPercent > errorPercentThreshold {
		return false, fmt.Errorf("the errorPercent(%d) exceed the ErrorPercentThreshold(%d)", errorPercent, errorPercentThreshold)
	}
	return errorPercent < errorPercentThreshold, nil
}
