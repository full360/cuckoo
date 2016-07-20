package cloudwatch

import (
	"reflect"
	"testing"
)

func TestDefaultMetricConfig(t *testing.T) {
	cases := []struct {
		expected struct {
			config *MetricConfig
		}
	}{
		{
			expected: struct {
				config *MetricConfig
			}{&MetricConfig{
				Name:      "service_monitoring",
				Namespace: "microservices",
				Service: &Service{
					Name: "Default",
					Env:  "Default",
				},
				Value: 0},
			},
		},
	}

	for _, c := range cases {
		config := DefaultMetricConfig()

		if !reflect.DeepEqual(config, c.expected.config) {
			t.Errorf("expected %q to be %q", config, c.expected.config)
		}
	}
}

func TestNewMetric(t *testing.T) {
	cases := []struct {
		expected struct {
			config      *MetricConfig
			serviceName string
			serviceEnv  string
			unit        string
		}
	}{
		{
			expected: struct {
				config      *MetricConfig
				serviceName string
				serviceEnv  string
				unit        string
			}{&MetricConfig{
				Name:      "service_monitoring",
				Namespace: "microservices",
				Service: &Service{
					Name: "users",
					Env:  "dev",
				},
				Value: 0},
				"ServiceName",
				"Environment",
				"Count",
			},
		},
	}

	for _, c := range cases {
		metric := NewMetric(c.expected.config)

		if *metric.Data.Namespace != c.expected.config.Namespace {
			t.Errorf("expected %q to be %q", *metric.Data.Namespace, c.expected.config.Namespace)
		}

		if *metric.Data.MetricData[0].MetricName != c.expected.config.Name {
			t.Errorf("expected %q to be %q", *metric.Data.MetricData[0].MetricName, c.expected.config.Name)
		}

		if *metric.Data.MetricData[0].Dimensions[0].Value != c.expected.config.Service.Name {
			t.Errorf("expected %q to be %q", *metric.Data.MetricData[0].Dimensions[0].Value, c.expected.config.Service.Name)
		}

		if *metric.Data.MetricData[0].Dimensions[0].Name != c.expected.serviceName {
			t.Errorf("expected %q to be %q", *metric.Data.MetricData[0].Dimensions[0].Name, c.expected.serviceName)
		}

		if *metric.Data.MetricData[0].Dimensions[1].Value != c.expected.config.Service.Env {
			t.Errorf("expected %q to be %q", *metric.Data.MetricData[0].Dimensions[1].Value, c.expected.config.Service.Env)
		}

		if *metric.Data.MetricData[0].Dimensions[1].Name != c.expected.serviceEnv {
			t.Errorf("expected %q to be %q", *metric.Data.MetricData[0].Dimensions[1].Name, c.expected.serviceEnv)
		}

		if *metric.Data.MetricData[0].Value != c.expected.config.Value {
			t.Errorf("expected %q to be %q", *metric.Data.MetricData[0].Value, c.expected.config.Value)
		}

		if *metric.Data.MetricData[0].Unit != c.expected.unit {
			t.Errorf("expected %q to be %q", *metric.Data.MetricData[0].Unit, c.expected.unit)
		}

		if !reflect.DeepEqual(metric.config, c.expected.config) {
			t.Errorf("expected %q to be %q", metric.config, c.expected.config)
		}
	}
}
