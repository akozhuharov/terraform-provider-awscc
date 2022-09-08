// Code generated by generators/resource/main.go; DO NOT EDIT.

package m2_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSM2Environment_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::M2::Environment", "awscc_m2_environment", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
