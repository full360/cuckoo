package cloudwatch

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

// CloudCloudWatch provides a Session to the CloudWatch API
type CloudWatch struct {
	Session *cloudwatch.CloudWatch
}

// DefaultSessionConfig creates a Session configuration based on our custom
// defaults
func DefaultSessionConfig() *aws.Config {
	config := &aws.Config{
		Region: aws.String("us-east-1"),
	}

	awsDefaultRegion := os.Getenv("AWS_DEFAULT_REGION")
	if awsDefaultRegion != "" {
		config.Region = &awsDefaultRegion
	}

	awsRegion := os.Getenv("AWS_REGION")
	if awsRegion != "" {
		config.Region = &awsRegion
	}
	return config
}

// NewSession creates and returns a cloudwatch Session session
func NewSession(config *aws.Config) *CloudWatch {
	return &CloudWatch{Session: cloudwatch.New(session.New(config))}
}
