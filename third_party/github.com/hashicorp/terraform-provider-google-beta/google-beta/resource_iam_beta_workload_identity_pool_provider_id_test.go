package google

import (
	"strings"
	"testing"
)

func TestValidateIAMBetaWorkloadIdentityPoolProviderId(t *testing.T) {
	x := []StringValidationTestCase{
		// No errors
		{TestName: "basic", Value: "foobar"},
		{TestName: "with numbers", Value: "foobar123"},
		{TestName: "short", Value: "foos"},
		{TestName: "long", Value: "12345678901234567890123456789012"},
		{TestName: "has a hyphen", Value: "foo-bar"},

		// With errors
		{TestName: "empty", Value: "", ExpectError: true},
		{TestName: "starts with a gcp-", Value: "gcp-foobar", ExpectError: true},
		{TestName: "with uppercase", Value: "fooBar", ExpectError: true},
		{TestName: "has an slash", Value: "foo/bar", ExpectError: true},
		{TestName: "has an backslash", Value: "foo\bar", ExpectError: true},
		{TestName: "too short", Value: "foo", ExpectError: true},
		{TestName: "too long", Value: strings.Repeat("f", 33), ExpectError: true},
	}

	es := testStringValidationCases(x, ValidateWorkloadIdentityPoolProviderId)
	if len(es) > 0 {
		t.Errorf("Failed to validate WorkloadIdentityPoolProvider names: %v", es)
	}
}
