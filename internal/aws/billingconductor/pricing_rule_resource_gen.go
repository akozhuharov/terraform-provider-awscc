// Code generated by generators/resource/main.go; DO NOT EDIT.

package billingconductor

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_billingconductor_pricing_rule", pricingRuleResource)
}

// pricingRuleResource returns the Terraform awscc_billingconductor_pricing_rule resource.
// This Terraform resource corresponds to the CloudFormation AWS::BillingConductor::PricingRule resource.
func pricingRuleResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Pricing rule ARN",
			//	  "pattern": "arn:aws(-cn)?:billingconductor::[0-9]{12}:pricingrule/[a-zA-Z0-9]{10}",
			//	  "type": "string"
			//	}
			Description: "Pricing rule ARN",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"associated_pricing_plan_count": {
			// Property: AssociatedPricingPlanCount
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The number of pricing plans associated with pricing rule",
			//	  "minimum": 0,
			//	  "type": "integer"
			//	}
			Description: "The number of pricing plans associated with pricing rule",
			Type:        types.Int64Type,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"billing_entity": {
			// Property: BillingEntity
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The seller of services provided by AWS, their affiliates, or third-party providers selling services via AWS Marketplaces. Supported billing entities are AWS, AWS Marketplace, and AISPL.",
			//	  "enum": [
			//	    "AWS",
			//	    "AWS Marketplace",
			//	    "AISPL"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "The seller of services provided by AWS, their affiliates, or third-party providers selling services via AWS Marketplaces. Supported billing entities are AWS, AWS Marketplace, and AISPL.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"AWS",
					"AWS Marketplace",
					"AISPL",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Creation timestamp in UNIX epoch time format",
			//	  "type": "integer"
			//	}
			Description: "Creation timestamp in UNIX epoch time format",
			Type:        types.Int64Type,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Pricing rule description",
			//	  "maxLength": 1024,
			//	  "type": "string"
			//	}
			Description: "Pricing rule description",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(1024),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"last_modified_time": {
			// Property: LastModifiedTime
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Latest modified timestamp in UNIX epoch time format",
			//	  "type": "integer"
			//	}
			Description: "Latest modified timestamp in UNIX epoch time format",
			Type:        types.Int64Type,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"modifier_percentage": {
			// Property: ModifierPercentage
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Pricing rule modifier percentage",
			//	  "minimum": 0,
			//	  "type": "number"
			//	}
			Description: "Pricing rule modifier percentage",
			Type:        types.Float64Type,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.FloatAtLeast(0.000000),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Pricing rule name",
			//	  "maxLength": 128,
			//	  "minLength": 1,
			//	  "pattern": "[a-zA-Z0-9_\\+=\\.\\-@]+",
			//	  "type": "string"
			//	}
			Description: "Pricing rule name",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 128),
				validate.StringMatch(regexp.MustCompile("[a-zA-Z0-9_\\+=\\.\\-@]+"), ""),
			},
		},
		"scope": {
			// Property: Scope
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A term used to categorize the granularity of a Pricing Rule.",
			//	  "enum": [
			//	    "GLOBAL",
			//	    "SERVICE",
			//	    "BILLING_ENTITY"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "A term used to categorize the granularity of a Pricing Rule.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"GLOBAL",
					"SERVICE",
					"BILLING_ENTITY",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"service": {
			// Property: Service
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The service which a pricing rule is applied on",
			//	  "maxLength": 128,
			//	  "minLength": 1,
			//	  "pattern": "[a-zA-Z0-9\\.\\-]+",
			//	  "type": "string"
			//	}
			Description: "The service which a pricing rule is applied on",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 128),
				validate.StringMatch(regexp.MustCompile("[a-zA-Z0-9\\.\\-]+"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "Key": {
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "maxLength": 256,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Key",
			//	      "Value"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "One of MARKUP or DISCOUNT that describes the direction of the rate that is applied to a pricing plan.",
			//	  "enum": [
			//	    "MARKUP",
			//	    "DISCOUNT"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "One of MARKUP or DISCOUNT that describes the direction of the rate that is applied to a pricing plan.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"MARKUP",
					"DISCOUNT",
				}),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "A markup/discount that is defined for a specific set of services that can later be associated with a pricing plan.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::BillingConductor::PricingRule").WithTerraformTypeName("awscc_billingconductor_pricing_rule")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                           "Arn",
		"associated_pricing_plan_count": "AssociatedPricingPlanCount",
		"billing_entity":                "BillingEntity",
		"creation_time":                 "CreationTime",
		"description":                   "Description",
		"key":                           "Key",
		"last_modified_time":            "LastModifiedTime",
		"modifier_percentage":           "ModifierPercentage",
		"name":                          "Name",
		"scope":                         "Scope",
		"service":                       "Service",
		"tags":                          "Tags",
		"type":                          "Type",
		"value":                         "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
