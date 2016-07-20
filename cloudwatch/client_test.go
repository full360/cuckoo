package cloudwatch

import (
	"os"
	"testing"
)

func TestDefaultSessionConfig(t *testing.T) {
	// Cleanup before the test
	os.Unsetenv("AWS_DEFAULT_REGION")
	os.Unsetenv("AWS_REGION")

	cases := []struct {
		expected  string
		export    bool
		exportVar string
		exportVal string
	}{
		{
			expected:  "us-east-1",
			export:    false,
			exportVar: "",
			exportVal: "",
		},
		{
			expected:  "ap-southeast-2",
			export:    true,
			exportVar: "AWS_DEFAULT_REGION",
			exportVal: "ap-southeast-2",
		},
		{
			expected:  "us-west-2",
			export:    true,
			exportVar: "AWS_REGION",
			exportVal: "us-west-2",
		},
	}

	for _, c := range cases {
		if c.export == true {
			os.Setenv(c.exportVar, c.exportVal)
		}

		config := DefaultSessionConfig()

		if *config.Region != c.expected {
			t.Errorf("expected %q to be %q", *config.Region, c.expected)
		}

		if c.export == true {
			os.Unsetenv(c.exportVar)
		}
	}
}
