/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package metrics

var (
	registries = map[string]func(*ReporterConfig) MetricRegistry{
		"prometheus": func(c *ReporterConfig) MetricRegistry {
			return NewPrometheusRegistry(c)
		},
	}

	collectors = make([]CollectorFunc, 0)

	registry MetricRegistry
)

func NewPrometheusRegistry(c *ReporterConfig) MetricRegistry {
	// todo implement it
	return nil
}

// CollectorFunc is a function that adds collectors to a MetricRegistry.
type CollectorFunc func(MetricRegistry, *ReporterConfig)

// Init initializes the Metrics module.
func Init(c *ReporterConfig) {
	// config.extention = prometheus
	regFunc, ok := registries["config.extention"]
	if !ok {
		regFunc = registries["prometheus"] // default
	}
	registry = regFunc(c)
	for _, v := range collectors {
		v(registry, c)
	}
	registry.Export()
}

// SetRegistry adds a metric registry to the registries map.
func SetRegistry(name string, v func(*ReporterConfig) MetricRegistry) {
	registries[name] = v
}

// MetricRegistry is an interface for a container that holds metrics.
type MetricRegistry interface {
	Counter(*MetricId) CounterMetric
	Gauge(*MetricId) GaugeMetric
	Histogram(*MetricId) HistogramMetric
	Summary(*MetricId) SummaryMetric
	Export()
}

// Type is a constant that represents a type of metric.
type Type uint8

const (
	Counter Type = iota
	Gauge
	LongTaskTimer
	Timer
	DistributionSummary
	Other
)

// MetricId represents a metric identifier.
type MetricId struct {
	Name string
	Desc string
	Tags map[string]string
	Type Type
}

func (m *MetricId) TagKeys() []string {
	keys := make([]string, 0, len(m.Tags))
	for k := range m.Tags {
		keys = append(keys, k)
	}
	return keys
}

// CounterMetric is an interface for a counter metric.
type CounterMetric interface {
	Inc()
	Add(float64)
}

// GaugeMetric is an interface for a gauge metric.
type GaugeMetric interface {
	Set(float64)
}

// HistogramMetric is an interface for a histogram metric.
type HistogramMetric interface {
	Record(float64)
}

// SummaryMetric is an interface for a summary metric.
type SummaryMetric interface {
	Record(float64)
}

// StatesMetrics is an interface for a set of metrics that represent states.
type StatesMetrics interface {
	Success()
	AddSuccess(float64)
	Fail()
	AddFailed(float64)
}

// DefaultStatesMetric is a default implementation of StatesMetrics.
type DefaultStatesMetric struct {
	r     MetricRegistry
	total *MetricId
	succ  *MetricId
	fail  *MetricId
}

func (c DefaultStatesMetric) Success() {
	c.r.Counter(c.total).Inc()
	c.r.Counter(c.succ).Inc()
}

func (c DefaultStatesMetric) AddSuccess(v float64) {
	c.r.Counter(c.total).Add(v)
	c.r.Counter(c.succ).Add(v)
}

func (c DefaultStatesMetric) Fail() {
	c.r.Counter(c.total).Inc()
	c.r.Counter(c.fail).Inc()
}

func (c DefaultStatesMetric) AddFailed(v float64) {
	c.r.Counter(c.total).Add(v)
	c.r.Counter(c.fail).Add(v)
}

// TimeMetrics is an interface for aset of metrics that represent time-related values.
type TimeMetrics interface {
	Record(float64)
}

// NewTimeMetrics init and write all data to registry
func NewTimeMetrics(min *MetricId, avg *MetricId, max *MetricId, last *MetricId, sum *MetricId) TimeMetrics {
	// todo implement it
	return nil
}

// AddCollector adds a collector function to the given MetricRegistry.
func AddCollector(fun CollectorFunc) {
	collectors = append(collectors, fun)
}

// ExportMetrics exports all metrics to the registry.
func ExportMetrics() {
	registry.Export()
}

// CompositeRegistry is a registry that combines multiple registries.
type CompositeRegistry struct {
	rs []MetricRegistry
}

func (c *CompositeRegistry) Counter(id *MetricId) CounterMetric {
	return &CompositeCounter{rs: c.rs, id: id}
}

func (c *CompositeRegistry) Gauge(id *MetricId) GaugeMetric {
	return &CompositeGauge{rs: c.rs, id: id}
}

func (c *CompositeRegistry) Histogram(id *MetricId) HistogramMetric {
	return &CompositeHistogram{rs: c.rs, id: id}
}

func (c *CompositeRegistry) Summary(id *MetricId) SummaryMetric {
	return &CompositeSummary{rs: c.rs, id: id}
}

func (c *CompositeRegistry) Export() {
	for _, r := range c.rs {
		r.Export()
	}
}

// NewCompositeRegistry returns a new CompositeRegistry that combines multiple registries.
func NewCompositeRegistry(registries ...MetricRegistry) MetricRegistry {
	return &CompositeRegistry{rs: registries}
}

// CompositeCounter is a counter that combines multiple counters from different registries.
type CompositeCounter struct {
	rs []MetricRegistry
	id *MetricId
}

func (c *CompositeCounter) Inc() {
	for _, r := range c.rs {
		r.Counter(c.id).Inc()
	}
}

func (c *CompositeCounter) Add(v float64) {
	for _, r := range c.rs {
		r.Counter(c.id).Add(v)
	}
}

// CompositeGauge is a gauge that combines multiple gauges from different registries.
type CompositeGauge struct {
	rs []MetricRegistry
	id *MetricId
}

func (c *CompositeGauge) Set(v float64) {
	for _, r := range c.rs {
		r.Gauge(c.id).Set(v)
	}
}

// CompositeHistogram is a histogram that combines multiple histograms from different registries.
type CompositeHistogram struct {
	rs []MetricRegistry
	id *MetricId
}

func (c *CompositeHistogram) Record(v float64) {
	for _, r := range c.rs {
		r.Histogram(c.id).Record(v)
	}
}

// CompositeSummary is a summary that combines multiple summaries from different registries.
type CompositeSummary struct {
	rs []MetricRegistry
	id *MetricId
}

func (c *CompositeSummary) Record(v float64) {
	for _, r := range c.rs {
		r.Summary(c.id).Record(v)
	}
}
