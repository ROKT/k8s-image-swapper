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
	errors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "k8s_image_swapper",
			Subsystem: "main",
			Name:      "errors",
			Help:      "Number of errors",
		},
		[]string{"error_type"},
	)
	metricLabels = []string{"resource_namespace", "registry", "repo", "error_type"}
	ecrErrors    = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "k8s_image_swapper",
			Subsystem: "ecr",
			Name:      "errors",
			Help:      "Number of ecr errors",
		},
		metricLabels,
	)
)

var PromReg *prometheus.Registry

func init() {
	PromReg = prometheus.NewRegistry()
	PromReg.MustRegister(collectors.NewGoCollector())
	PromReg.MustRegister(ecrErrors)
}

// Increments the counter of errors
func IncrementError(errType string) {
	ecrErrors.With(
		prometheus.Labels{
			"error_type": errType,
		},
	).Inc()
}

// Increments the counter of ecr errors
func IncrementEcrError(resource_namespace string, registry string, repo string, errType string) {
	errors.With(
		prometheus.Labels{
			"resource_namespace": resource_namespace,
			"registry":           registry,
			"repo":               repo,
			"error_type":         errType,
		},
	).Inc()
}
