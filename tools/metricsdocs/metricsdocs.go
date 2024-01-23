package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/machadovilaca/operator-observability/pkg/operatormetrics"

	sspMetrics "kubevirt.io/ssp-operator/pkg/monitoring/metrics/ssp-operator"
	validatorMetrics "kubevirt.io/ssp-operator/pkg/monitoring/metrics/template-validator"
	"kubevirt.io/ssp-operator/pkg/monitoring/rules"
)

const (
	title      = "# SSP Operator metrics\n"
	background = "This document aims to help users that are not familiar with metrics exposed by the SSP Operator.\n" +
		"All metrics documented here are auto-generated by the utility tool `tools/metricsdocs` and reflects exactly what is being exposed.\n\n"

	KVSpecificMetrics = "## SSP Operator Metrics List\n"

	opening = title +
		background +
		KVSpecificMetrics

	footerHeading = "## Developing new metrics\n"
	footerContent = "After developing new metrics or changing old ones, please run `make generate-doc` to regenerate this document.\n"

	footer = footerHeading + footerContent
)

func main() {
	metricsList := recordRulesDescToMetricList(rules.RecordRulesWithDescriptions())

	sspMetrics.SetupMetrics()
	validatorMetrics.SetupMetrics()

	for _, m := range operatormetrics.ListMetrics() {
		metricsList = append(metricsList, metric{
			name:        m.GetOpts().Name,
			description: m.GetOpts().Help,
			mtype:       strings.TrimSuffix(string(m.GetType()), "Vec"),
		})
	}

	sort.Sort(metricsList)
	printMetrics(metricsList)
}

func printMetrics(metricsList metricList) {
	fmt.Print(opening)
	metricsList.writeOut()
	fmt.Print(footer)
}

type metric struct {
	name        string
	description string
	mtype       string
}

func recordRulesDescToMetricList(mdl []rules.RecordRulesDesc) metricList {
	res := make([]metric, len(mdl))
	for i, md := range mdl {
		res[i] = metricDescriptionToMetric(md)
	}

	return res
}

func metricDescriptionToMetric(rrd rules.RecordRulesDesc) metric {
	return metric{
		name:        rrd.Name,
		description: rrd.Description,
		mtype:       rrd.Type,
	}
}

func (m metric) writeOut() {
	fmt.Println("###", m.name)
	fmt.Println(m.description + ". Type: " + m.mtype + ".")
}

type metricList []metric

var _ sort.Interface = metricList{}

// Len implements sort.Interface.Len
func (m metricList) Len() int {
	return len(m)
}

// Less implements sort.Interface.Less
func (m metricList) Less(i, j int) bool {
	return m[i].name < m[j].name
}

// Swap implements sort.Interface.Swap
func (m metricList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m metricList) writeOut() {
	for _, met := range m {
		met.writeOut()
	}
}
