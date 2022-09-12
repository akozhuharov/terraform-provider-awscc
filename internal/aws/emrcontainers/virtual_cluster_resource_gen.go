// Code generated by generators/resource/main.go; DO NOT EDIT.

package emrcontainers

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_emrcontainers_virtual_cluster", virtualClusterResourceType)
}

// virtualClusterResourceType returns the Terraform awscc_emrcontainers_virtual_cluster resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EMRContainers::VirtualCluster resource type.
func virtualClusterResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"container_provider": {
			// Property: ContainerProvider
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Container provider of the virtual cluster.",
			//   "properties": {
			//     "Id": {
			//       "description": "The ID of the container cluster",
			//       "maxLength": 100,
			//       "minLength": 1,
			//       "pattern": "^[0-9A-Za-z][A-Za-z0-9\\-_]*",
			//       "type": "string"
			//     },
			//     "Info": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "EksInfo": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Namespace": {
			//               "maxLength": 63,
			//               "minLength": 1,
			//               "pattern": "[a-z0-9]([-a-z0-9]*[a-z0-9])?",
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "Namespace"
			//           ],
			//           "type": "object"
			//         }
			//       },
			//       "required": [
			//         "EksInfo"
			//       ],
			//       "type": "object"
			//     },
			//     "Type": {
			//       "description": "The type of the container provider",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Type",
			//     "Id",
			//     "Info"
			//   ],
			//   "type": "object"
			// }
			Description: "Container provider of the virtual cluster.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"id": {
						// Property: Id
						Description: "The ID of the container cluster",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 100),
							validate.StringMatch(regexp.MustCompile("^[0-9A-Za-z][A-Za-z0-9\\-_]*"), ""),
						},
					},
					"info": {
						// Property: Info
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"eks_info": {
									// Property: EksInfo
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"namespace": {
												// Property: Namespace
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 63),
													validate.StringMatch(regexp.MustCompile("[a-z0-9]([-a-z0-9]*[a-z0-9])?"), ""),
												},
											},
										},
									),
									Required: true,
								},
							},
						),
						Required: true,
					},
					"type": {
						// Property: Type
						Description: "The type of the container provider",
						Type:        types.StringType,
						Required:    true,
					},
				},
			),
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Id of the virtual cluster.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Id of the virtual cluster.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the virtual cluster.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "[\\.\\-_/#A-Za-z0-9]+",
			//   "type": "string"
			// }
			Description: "Name of the virtual cluster.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
				validate.StringMatch(regexp.MustCompile("[\\.\\-_/#A-Za-z0-9]+"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this virtual cluster.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "An arbitrary set of tags (key-value pairs) for this virtual cluster.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this virtual cluster.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Schema of AWS::EMRContainers::VirtualCluster Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EMRContainers::VirtualCluster").WithTerraformTypeName("awscc_emrcontainers_virtual_cluster")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                "Arn",
		"container_provider": "ContainerProvider",
		"eks_info":           "EksInfo",
		"id":                 "Id",
		"info":               "Info",
		"key":                "Key",
		"name":               "Name",
		"namespace":          "Namespace",
		"tags":               "Tags",
		"type":               "Type",
		"value":              "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
