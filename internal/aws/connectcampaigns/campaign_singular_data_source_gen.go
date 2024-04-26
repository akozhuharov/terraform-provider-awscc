// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package connectcampaigns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_connectcampaigns_campaign", campaignDataSource)
}

// campaignDataSource returns the Terraform awscc_connectcampaigns_campaign data source.
// This Terraform data source corresponds to the CloudFormation AWS::ConnectCampaigns::Campaign resource.
func campaignDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Connect Campaign Arn",
		//	  "maxLength": 256,
		//	  "minLength": 0,
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect-campaigns:[-a-z0-9]*:[0-9]{12}:campaign/[-a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Connect Campaign Arn",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ConnectInstanceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Connect Instance Arn",
		//	  "maxLength": 256,
		//	  "minLength": 0,
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"connect_instance_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Connect Instance Arn",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DialerConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The possible types of dialer config parameters",
		//	  "oneOf": [
		//	    {
		//	      "required": [
		//	        "ProgressiveDialerConfig"
		//	      ]
		//	    },
		//	    {
		//	      "required": [
		//	        "PredictiveDialerConfig"
		//	      ]
		//	    },
		//	    {
		//	      "required": [
		//	        "AgentlessDialerConfig"
		//	      ]
		//	    }
		//	  ],
		//	  "properties": {
		//	    "AgentlessDialerConfig": {
		//	      "additionalProperties": false,
		//	      "description": "Agentless Dialer config",
		//	      "properties": {
		//	        "DialingCapacity": {
		//	          "description": "Allocates dialing capacity for this campaign between multiple active campaigns.",
		//	          "maximum": 1,
		//	          "minimum": 0.01,
		//	          "type": "number"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "PredictiveDialerConfig": {
		//	      "additionalProperties": false,
		//	      "description": "Predictive Dialer config",
		//	      "properties": {
		//	        "BandwidthAllocation": {
		//	          "description": "The bandwidth allocation of a queue resource.",
		//	          "maximum": 1,
		//	          "minimum": 0,
		//	          "type": "number"
		//	        },
		//	        "DialingCapacity": {
		//	          "description": "Allocates dialing capacity for this campaign between multiple active campaigns.",
		//	          "maximum": 1,
		//	          "minimum": 0.01,
		//	          "type": "number"
		//	        }
		//	      },
		//	      "required": [
		//	        "BandwidthAllocation"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "ProgressiveDialerConfig": {
		//	      "additionalProperties": false,
		//	      "description": "Progressive Dialer config",
		//	      "properties": {
		//	        "BandwidthAllocation": {
		//	          "description": "The bandwidth allocation of a queue resource.",
		//	          "maximum": 1,
		//	          "minimum": 0,
		//	          "type": "number"
		//	        },
		//	        "DialingCapacity": {
		//	          "description": "Allocates dialing capacity for this campaign between multiple active campaigns.",
		//	          "maximum": 1,
		//	          "minimum": 0.01,
		//	          "type": "number"
		//	        }
		//	      },
		//	      "required": [
		//	        "BandwidthAllocation"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"dialer_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AgentlessDialerConfig
				"agentless_dialer_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DialingCapacity
						"dialing_capacity": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Description: "Allocates dialing capacity for this campaign between multiple active campaigns.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Agentless Dialer config",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: PredictiveDialerConfig
				"predictive_dialer_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: BandwidthAllocation
						"bandwidth_allocation": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Description: "The bandwidth allocation of a queue resource.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: DialingCapacity
						"dialing_capacity": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Description: "Allocates dialing capacity for this campaign between multiple active campaigns.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Predictive Dialer config",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ProgressiveDialerConfig
				"progressive_dialer_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: BandwidthAllocation
						"bandwidth_allocation": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Description: "The bandwidth allocation of a queue resource.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: DialingCapacity
						"dialing_capacity": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Description: "Allocates dialing capacity for this campaign between multiple active campaigns.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Progressive Dialer config",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The possible types of dialer config parameters",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Connect Campaign Name",
		//	  "maxLength": 127,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Connect Campaign Name",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OutboundCallConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration used for outbound calls.",
		//	  "properties": {
		//	    "AnswerMachineDetectionConfig": {
		//	      "additionalProperties": false,
		//	      "description": "The configuration used for answering machine detection during outbound calls",
		//	      "properties": {
		//	        "AwaitAnswerMachinePrompt": {
		//	          "description": "Enables detection of prompts (e.g., beep after after a voicemail greeting)",
		//	          "type": "boolean"
		//	        },
		//	        "EnableAnswerMachineDetection": {
		//	          "description": "Flag to decided whether outbound calls should have answering machine detection enabled or not",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "required": [
		//	        "EnableAnswerMachineDetection"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "ConnectContactFlowArn": {
		//	      "description": "The identifier of the contact flow for the outbound call.",
		//	      "maxLength": 500,
		//	      "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/contact-flow/[-a-zA-Z0-9]*$",
		//	      "type": "string"
		//	    },
		//	    "ConnectQueueArn": {
		//	      "description": "The queue for the call. If you specify a queue, the phone displayed for caller ID is the phone number specified in the queue. If you do not specify a queue, the queue defined in the contact flow is used. If you do not specify a queue, you must specify a source phone number.",
		//	      "maxLength": 500,
		//	      "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/queue/[-a-zA-Z0-9]*$",
		//	      "type": "string"
		//	    },
		//	    "ConnectSourcePhoneNumber": {
		//	      "description": "The phone number associated with the Amazon Connect instance, in E.164 format. If you do not specify a source phone number, you must specify a queue.",
		//	      "maxLength": 100,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "ConnectContactFlowArn"
		//	  ],
		//	  "type": "object"
		//	}
		"outbound_call_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AnswerMachineDetectionConfig
				"answer_machine_detection_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: AwaitAnswerMachinePrompt
						"await_answer_machine_prompt": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "Enables detection of prompts (e.g., beep after after a voicemail greeting)",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: EnableAnswerMachineDetection
						"enable_answer_machine_detection": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "Flag to decided whether outbound calls should have answering machine detection enabled or not",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The configuration used for answering machine detection during outbound calls",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ConnectContactFlowArn
				"connect_contact_flow_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The identifier of the contact flow for the outbound call.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ConnectQueueArn
				"connect_queue_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The queue for the call. If you specify a queue, the phone displayed for caller ID is the phone number specified in the queue. If you do not specify a queue, the queue defined in the contact flow is used. If you do not specify a queue, you must specify a source phone number.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ConnectSourcePhoneNumber
				"connect_source_phone_number": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The phone number associated with the Amazon Connect instance, in E.164 format. If you do not specify a source phone number, you must specify a queue.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The configuration used for outbound calls.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "One or more tags.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that's 1 to 256 characters in length.",
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
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that's 1 to 256 characters in length.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "One or more tags.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ConnectCampaigns::Campaign",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ConnectCampaigns::Campaign").WithTerraformTypeName("awscc_connectcampaigns_campaign")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"agentless_dialer_config":         "AgentlessDialerConfig",
		"answer_machine_detection_config": "AnswerMachineDetectionConfig",
		"arn":                             "Arn",
		"await_answer_machine_prompt":     "AwaitAnswerMachinePrompt",
		"bandwidth_allocation":            "BandwidthAllocation",
		"connect_contact_flow_arn":        "ConnectContactFlowArn",
		"connect_instance_arn":            "ConnectInstanceArn",
		"connect_queue_arn":               "ConnectQueueArn",
		"connect_source_phone_number":     "ConnectSourcePhoneNumber",
		"dialer_config":                   "DialerConfig",
		"dialing_capacity":                "DialingCapacity",
		"enable_answer_machine_detection": "EnableAnswerMachineDetection",
		"key":                             "Key",
		"name":                            "Name",
		"outbound_call_config":            "OutboundCallConfig",
		"predictive_dialer_config":        "PredictiveDialerConfig",
		"progressive_dialer_config":       "ProgressiveDialerConfig",
		"tags":                            "Tags",
		"value":                           "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
