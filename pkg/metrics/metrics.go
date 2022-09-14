/*
Copyright © 2020 Enrico Stahn <enrico.stahn@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

var (
	metricLabels  = []string{"namespace", "registry", "repo"}
	swapperErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "k8s_image_swapper",
			Subsystem: "swapping",
			Name:      "errors",
			Help:      "Number of errors",
		},
		metricLabels,
	)
	numberOfSwaps = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "k8s_image_swapper",
			Subsystem: "swaping",
			Name:      "swaps",
			Help:      "Number of swaps for the repo",
		},
		metricLabels,
	)
)

var PromReg *prometheus.Registry

func init() {
	PromReg = prometheus.NewRegistry()
	PromReg.MustRegister(collectors.NewGoCollector())
	PromReg.MustRegister(swapperErrors)
	PromReg.MustRegister(numberOfSwaps)

}

// PrometheusMetricServer the type of MetricsServer
type PrometheusMetricServer struct{}

// RecordSwapError counts the number of swapping errors
func (metricsServer PrometheusMetricServer) RecordSwapError(namespace string, registry string, repo string) {
	labels := getLabels(namespace, registry, repo)
	swapperErrors.With(labels).Inc()
}

func getLabels(namespace string, registry string, repo string) prometheus.Labels {
	return prometheus.Labels{"namespace": namespace, "registry": registry, "repo": repo}
}
