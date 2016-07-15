package cloudwatch

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

// Service is used to represent a service
type Service struct {
	// Name of the service
	Name string
	// Env of the service, could also be represented by the tag
	Env string
}

// MetricConfig is used to represent a cloudwatch metric configuration data
type MetricConfig struct {
	// Name of the metric
	Name string
	// Namespace of the metric
	Namespace string
	// Service representation for the dimension
	Service *Service
	// Value of the metric that will be posted
	Value float64
}

// Metric is used to represent a cloudwatch metric
type Metric struct {
	// cw represents the cloudwatch session
	cw *CloudWatch
	// Data of the metric to be posted
	Data *cloudwatch.PutMetricDataInput
	// config of the metric
	config *MetricConfig
}

// DefaultClientConfig sets the default metric configuration
func DefaultMetricConfig() *MetricConfig {
	config := &MetricConfig{
		Name:      "service_monitoring",
		Namespace: "microservices",
		Service: &Service{
			Name: "Default",
			Env:  "Default",
		},
		Value: 0,
	}
	return config
}

// NewMetric returns a new metric
func NewMetric(config *MetricConfig) *Metric {
	session := NewSession(DefaultSessionConfig())
	metric := &cloudwatch.PutMetricDataInput{
		MetricData: []*cloudwatch.MetricDatum{
			{
				MetricName: aws.String(config.Name),
				Dimensions: []*cloudwatch.Dimension{
					{
						Name:  aws.String("ServiceName"),
						Value: aws.String(config.Service.Name),
					},
					{
						Name:  aws.String("Environment"),
						Value: aws.String(config.Service.Env),
					},
				},
				Timestamp: aws.Time(time.Now().UTC()),
				Unit:      aws.String("Count"),
				Value:     aws.Float64(config.Value),
			},
		},
		Namespace: aws.String(config.Namespace),
	}
	return &Metric{
		cw:     session,
		Data:   metric,
		config: config,
	}
}

// Put posts to cloudwatch the metric defined
func (m *Metric) Put(value float64) (*cloudwatch.PutMetricDataOutput, error) {
	m.Data.MetricData[0].Value = aws.Float64(value)
	m.Data.MetricData[0].Timestamp = aws.Time(time.Now().UTC())

	out, err := m.cw.Session.PutMetricData(m.Data)
	if err != nil {
		return nil, err
	}
	return out, nil
}
