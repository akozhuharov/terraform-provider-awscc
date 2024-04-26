// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package workspacesthinclient

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_workspacesthinclient_environment", environmentResource)
}

// environmentResource returns the Terraform awscc_workspacesthinclient_environment resource.
// This Terraform resource corresponds to the CloudFormation AWS::WorkSpacesThinClient::Environment resource.
func environmentResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ActivationCode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Activation code for devices associated with environment.",
		//	  "pattern": "^[a-z]{2}[a-z0-9]{6}$",
		//	  "type": "string"
		//	}
		"activation_code": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Activation code for devices associated with environment.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The environment ARN.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:[\\w+=\\/,.@-]+:[a-zA-Z0-9\\-]+:[a-zA-Z0-9\\-]*:[0-9]{0,12}:[a-zA-Z0-9\\-\\/\\._]+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The environment ARN.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp in unix epoch format when environment was created.",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The timestamp in unix epoch format when environment was created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DesiredSoftwareSetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the software set to apply.",
		//	  "pattern": "^[0-9]{1,9}$",
		//	  "type": "string"
		//	}
		"desired_software_set_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the software set to apply.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^[0-9]{1,9}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DesktopArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the desktop to stream from Amazon WorkSpaces, WorkSpaces Web, or AppStream 2.0.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:[\\w+=\\/,.@-]+:[a-zA-Z0-9\\-]+:[a-zA-Z0-9\\-]*:[0-9]{0,12}:[a-zA-Z0-9\\-\\/\\._]+$",
		//	  "type": "string"
		//	}
		"desktop_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the desktop to stream from Amazon WorkSpaces, WorkSpaces Web, or AppStream 2.0.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(20, 2048),
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:[\\w+=\\/,.@-]+:[a-zA-Z0-9\\-]+:[a-zA-Z0-9\\-]*:[0-9]{0,12}:[a-zA-Z0-9\\-\\/\\._]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DesktopEndpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The URL for the identity provider login (only for environments that use AppStream 2.0).",
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "pattern": "^(https:\\/\\/)[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,32}(:[0-9]{1,5})?(\\/.*)?$",
		//	  "type": "string"
		//	}
		"desktop_endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The URL for the identity provider login (only for environments that use AppStream 2.0).",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 1024),
				stringvalidator.RegexMatches(regexp.MustCompile("^(https:\\/\\/)[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,32}(:[0-9]{1,5})?(\\/.*)?$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DesktopType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of VDI.",
		//	  "enum": [
		//	    "workspaces",
		//	    "appstream",
		//	    "workspaces-web"
		//	  ],
		//	  "type": "string"
		//	}
		"desktop_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of VDI.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Unique identifier of the environment.",
		//	  "pattern": "^[a-z0-9]{9}$",
		//	  "type": "string"
		//	}
		"environment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Unique identifier of the environment.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the AWS Key Management Service key used to encrypt the environment.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:[\\w+=\\/,.@-]+:kms:[a-zA-Z0-9\\-]*:[0-9]{0,12}:key\\/[a-zA-Z0-9-]+$",
		//	  "type": "string"
		//	}
		"kms_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the AWS Key Management Service key used to encrypt the environment.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(20, 2048),
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:[\\w+=\\/,.@-]+:kms:[a-zA-Z0-9\\-]*:[0-9]{0,12}:key\\/[a-zA-Z0-9-]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MaintenanceWindow
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A specification for a time window to apply software updates.",
		//	  "properties": {
		//	    "ApplyTimeOf": {
		//	      "description": "The desired time zone maintenance window.",
		//	      "enum": [
		//	        "UTC",
		//	        "DEVICE"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "DaysOfTheWeek": {
		//	      "description": "The date of maintenance window.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "enum": [
		//	          "MONDAY",
		//	          "TUESDAY",
		//	          "WEDNESDAY",
		//	          "THURSDAY",
		//	          "FRIDAY",
		//	          "SATURDAY",
		//	          "SUNDAY"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "maxItems": 7,
		//	      "minItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "EndTimeHour": {
		//	      "description": "The hour end time of maintenance window.",
		//	      "maximum": 23,
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "EndTimeMinute": {
		//	      "description": "The minute end time of maintenance window.",
		//	      "maximum": 59,
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "StartTimeHour": {
		//	      "description": "The hour start time of maintenance window.",
		//	      "maximum": 23,
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "StartTimeMinute": {
		//	      "description": "The minute start time of maintenance window.",
		//	      "maximum": 59,
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "Type": {
		//	      "description": "The type of maintenance window.",
		//	      "enum": [
		//	        "SYSTEM",
		//	        "CUSTOM"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Type"
		//	  ],
		//	  "type": "object"
		//	}
		"maintenance_window": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ApplyTimeOf
				"apply_time_of": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The desired time zone maintenance window.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"UTC",
							"DEVICE",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: DaysOfTheWeek
				"days_of_the_week": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "The date of maintenance window.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Set{ /*START VALIDATORS*/
						setvalidator.SizeBetween(1, 7),
						setvalidator.ValueStringsAre(
							stringvalidator.OneOf(
								"MONDAY",
								"TUESDAY",
								"WEDNESDAY",
								"THURSDAY",
								"FRIDAY",
								"SATURDAY",
								"SUNDAY",
							),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
						setplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: EndTimeHour
				"end_time_hour": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The hour end time of maintenance window.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.Between(0, 23),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: EndTimeMinute
				"end_time_minute": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The minute end time of maintenance window.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.Between(0, 59),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: StartTimeHour
				"start_time_hour": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The hour start time of maintenance window.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.Between(0, 23),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: StartTimeMinute
				"start_time_minute": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The minute start time of maintenance window.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.Between(0, 59),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The type of maintenance window.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"SYSTEM",
							"CUSTOM",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "A specification for a time window to apply software updates.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the environment.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^.+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the environment.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
				stringvalidator.RegexMatches(regexp.MustCompile("^.+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PendingSoftwareSetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the software set that is pending to be installed.",
		//	  "pattern": "^[0-9]{1,9}$",
		//	  "type": "string"
		//	}
		"pending_software_set_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the software set that is pending to be installed.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PendingSoftwareSetVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The version of the software set that is pending to be installed.",
		//	  "type": "string"
		//	}
		"pending_software_set_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The version of the software set that is pending to be installed.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RegisteredDevicesCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Number of devices registered to the environment.",
		//	  "minimum": 0,
		//	  "type": "integer"
		//	}
		"registered_devices_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Number of devices registered to the environment.",
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SoftwareSetComplianceStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Describes if the software currently installed on all devices in the environment is a supported version.",
		//	  "enum": [
		//	    "COMPLIANT",
		//	    "NOT_COMPLIANT",
		//	    "NO_REGISTERED_DEVICES"
		//	  ],
		//	  "type": "string"
		//	}
		"software_set_compliance_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Describes if the software currently installed on all devices in the environment is a supported version.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SoftwareSetUpdateMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An option to define which software updates to apply.",
		//	  "enum": [
		//	    "USE_LATEST",
		//	    "USE_DESIRED"
		//	  ],
		//	  "type": "string"
		//	}
		"software_set_update_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An option to define which software updates to apply.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"USE_LATEST",
					"USE_DESIRED",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SoftwareSetUpdateSchedule
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An option to define if software updates should be applied within a maintenance window.",
		//	  "enum": [
		//	    "USE_MAINTENANCE_WINDOW",
		//	    "APPLY_IMMEDIATELY"
		//	  ],
		//	  "type": "string"
		//	}
		"software_set_update_schedule": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An option to define if software updates should be applied within a maintenance window.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"USE_MAINTENANCE_WINDOW",
					"APPLY_IMMEDIATELY",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthAtMost(256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp in unix epoch format when environment was last updated.",
		//	  "type": "string"
		//	}
		"updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The timestamp in unix epoch format when environment was last updated.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource type definition for AWS::WorkSpacesThinClient::Environment.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::WorkSpacesThinClient::Environment").WithTerraformTypeName("awscc_workspacesthinclient_environment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"activation_code":                "ActivationCode",
		"apply_time_of":                  "ApplyTimeOf",
		"arn":                            "Arn",
		"created_at":                     "CreatedAt",
		"days_of_the_week":               "DaysOfTheWeek",
		"desired_software_set_id":        "DesiredSoftwareSetId",
		"desktop_arn":                    "DesktopArn",
		"desktop_endpoint":               "DesktopEndpoint",
		"desktop_type":                   "DesktopType",
		"end_time_hour":                  "EndTimeHour",
		"end_time_minute":                "EndTimeMinute",
		"environment_id":                 "Id",
		"key":                            "Key",
		"kms_key_arn":                    "KmsKeyArn",
		"maintenance_window":             "MaintenanceWindow",
		"name":                           "Name",
		"pending_software_set_id":        "PendingSoftwareSetId",
		"pending_software_set_version":   "PendingSoftwareSetVersion",
		"registered_devices_count":       "RegisteredDevicesCount",
		"software_set_compliance_status": "SoftwareSetComplianceStatus",
		"software_set_update_mode":       "SoftwareSetUpdateMode",
		"software_set_update_schedule":   "SoftwareSetUpdateSchedule",
		"start_time_hour":                "StartTimeHour",
		"start_time_minute":              "StartTimeMinute",
		"tags":                           "Tags",
		"type":                           "Type",
		"updated_at":                     "UpdatedAt",
		"value":                          "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
