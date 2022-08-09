package slo

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"strings"
	"text/template"
	"time"

	openslov1 "github.com/alexandreLamarre/oslo/pkg/manifest/v1"
	"github.com/alexandreLamarre/sloth/core/prometheus"
	"gopkg.in/yaml.v2"
)

var errorRatioRawQueryTmpl = template.Must(template.New("").Parse(`
  1 - (
    (
      {{ .good }}
    )
    /
    (
      {{ .total }}
    )
  )
`))

// Returns the same number of SLO Groups as the number of OpenSLO specs
//(Same number as the number of services in the request)
func ParseToPrometheusModel(slo openslov1.SLO) (*prometheus.SLOGroup, error) {
	//TODO
	y := NewYAMLSpecLoader(time.Hour * 24 * 30) // FIXME: hardcoded window period
	m, err := y.MapSpecToModel(slo)
	if err != nil {
		return nil, fmt.Errorf("could not map SLO: %w", err)
	}
	return m, nil
}

type YAMLSpecLoader struct {
	windowPeriod time.Duration
}

// YAMLSpecLoader knows how to load YAML specs and converts them to a model.
func NewYAMLSpecLoader(windowPeriod time.Duration) YAMLSpecLoader {
	return YAMLSpecLoader{
		windowPeriod: windowPeriod,
	}
}

func (YAMLSpecLoader) ValidateTimeWindow(spec openslov1.SLO) error {
	if len(spec.Spec.TimeWindow) == 0 {
		return nil
	}

	if len(spec.Spec.TimeWindow) > 1 {
		return fmt.Errorf("only 1 time window is supported")
	}

	t := spec.Spec.TimeWindow[0]
	if strings.ToLower(t.Duration) != "day" {
		return fmt.Errorf("only days based time windows are supported")
	}

	return nil
}

func (y YAMLSpecLoader) MapSpecToModel(spec openslov1.SLO) (*prometheus.SLOGroup, error) {
	slos, err := y.GetSLOs(spec)
	if err != nil {
		return nil, fmt.Errorf("could not map SLOs correctly: %w", err)
	}
	return &prometheus.SLOGroup{SLOs: slos}, nil
}

func (y YAMLSpecLoader) LoadSpec(ctx context.Context, data []byte) (*prometheus.SLOGroup, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("spec is required")
	}

	s := openslov1.SLO{}
	err := yaml.Unmarshal(data, &s)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall YAML spec correctly: %w", err)
	}

	// Check version.
	if s.APIVersion != openslov1.APIVersion {
		return nil, fmt.Errorf("invalid spec version, should be %q", openslov1.APIVersion)
	}

	// Check at least we have one SLO.
	if len(s.Spec.Objectives) == 0 {
		return nil, fmt.Errorf("at least one SLO is required")
	}

	// Validate time windows are correct.
	err = y.ValidateTimeWindow(s)
	if err != nil {
		return nil, fmt.Errorf("invalid SLO time windows: %w", err)
	}

	m, err := y.MapSpecToModel(s)
	if err != nil {
		return nil, fmt.Errorf("could not map to model: %w", err)
	}

	return m, nil
}

func (y YAMLSpecLoader) GetSLI(spec openslov1.SLOSpec, slo openslov1.Objective) (*prometheus.SLI, error) {
	if spec.Indicator.Spec.RatioMetric == nil {
		return nil, fmt.Errorf("missing ratioMetrics")
	}

	good := spec.Indicator.Spec.RatioMetric.Good
	total := spec.Indicator.Spec.RatioMetric.Total

	if good.MetricSource.Type != "prometheus" {
		return nil, fmt.Errorf("prometheus source required in spec.Spec.Indicator.Spec.Good")
	}

	if total.MetricSource.Type != "prometheus" {
		return nil, fmt.Errorf("prometheus or sloth query ratio 'total' source is required")
	}

	if good.MetricSource.MetricSourceSpec["queryType"] != "promql" {
		return nil, fmt.Errorf("promql query type required in spec.Spec.Indicator.Spec.Total.MetricSource.MetricSourceSpec")
	}
	if total.MetricSource.MetricSourceSpec["queryType"] != "promql" {
		return nil, fmt.Errorf("promql query type required in spec.Spec.Indicator.Spec.Total.MetricSource.MetricSourceSpec")
	}

	// Map as good and total events as a raw query.
	var b bytes.Buffer
	goodQuery := good.MetricSource.MetricSourceSpec["query"]
	totalQuery := total.MetricSource.MetricSourceSpec["query"]
	err := errorRatioRawQueryTmpl.Execute(&b, map[string]string{"good": goodQuery, "total": totalQuery})
	if err != nil {
		return nil, fmt.Errorf("could not execute mapping SLI template: %w", err)
	}

	return &prometheus.SLI{Raw: &prometheus.SLIRaw{
		ErrorRatioQuery: b.String(),
	}}, nil
}

// getSLOs will try getting all the objectives as individual SLOs, this way we can map
// to what Sloth understands as an SLO, that OpenSLO understands as a list of objectives
// for the same SLO.
func (y YAMLSpecLoader) GetSLOs(spec openslov1.SLO) ([]prometheus.SLO, error) {
	res := []prometheus.SLO{}

	for idx, slo := range spec.Spec.Objectives {
		sli, err := y.GetSLI(spec.Spec, slo)
		if err != nil {
			return nil, fmt.Errorf("could not map SLI: %w", err)
		}

		timeWindow := y.windowPeriod
		if len(spec.Spec.TimeWindow) > 0 {
			// The time window is validated to be one of '7d', '28d' or '30d'
			newTime, err := strconv.Atoi(strings.ReplaceAll(spec.Spec.TimeWindow[0].Duration, "d", ""))
			if err != nil {
				return nil, err
			}
			timeWindow = (time.Duration(newTime) * 24) * time.Hour
		}

		res = append(res, prometheus.SLO{
			ID:              fmt.Sprintf("%s-%d", spec.Metadata.Name, idx),
			Name:            fmt.Sprintf("%s-%d", spec.Metadata.Name, idx),
			Service:         spec.Spec.Service,
			Description:     spec.Spec.Description,
			TimeWindow:      timeWindow,
			SLI:             *sli,
			Objective:       slo.Target,
			PageAlertMeta:   prometheus.AlertMeta{Disable: true},
			TicketAlertMeta: prometheus.AlertMeta{Disable: true},
		})
	}

	return res, nil
}
