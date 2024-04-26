// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package bedrock

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_bedrock_guardrail", guardrailResource)
}

// guardrailResource returns the Terraform awscc_bedrock_guardrail resource.
// This Terraform resource corresponds to the CloudFormation AWS::Bedrock::Guardrail resource.
func guardrailResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: BlockedInputMessaging
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Messaging for when violations are detected in text",
		//	  "maxLength": 500,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"blocked_input_messaging": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Messaging for when violations are detected in text",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 500),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: BlockedOutputsMessaging
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Messaging for when violations are detected in text",
		//	  "maxLength": 500,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"blocked_outputs_messaging": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Messaging for when violations are detected in text",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 500),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: ContentPolicyConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Content policy config for a guardrail.",
		//	  "properties": {
		//	    "FiltersConfig": {
		//	      "description": "List of content filter configs in content policy.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Content filter config in content policy.",
		//	        "properties": {
		//	          "InputStrength": {
		//	            "description": "Strength for filters",
		//	            "enum": [
		//	              "NONE",
		//	              "LOW",
		//	              "MEDIUM",
		//	              "HIGH"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "OutputStrength": {
		//	            "description": "Strength for filters",
		//	            "enum": [
		//	              "NONE",
		//	              "LOW",
		//	              "MEDIUM",
		//	              "HIGH"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "Type": {
		//	            "description": "Type of filter in content policy",
		//	            "enum": [
		//	              "SEXUAL",
		//	              "VIOLENCE",
		//	              "HATE",
		//	              "INSULTS",
		//	              "MISCONDUCT",
		//	              "PROMPT_ATTACK"
		//	            ],
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "InputStrength",
		//	          "OutputStrength",
		//	          "Type"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 6,
		//	      "minItems": 1,
		//	      "type": "array"
		//	    }
		//	  },
		//	  "required": [
		//	    "FiltersConfig"
		//	  ],
		//	  "type": "object"
		//	}
		"content_policy_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: FiltersConfig
				"filters_config": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: InputStrength
							"input_strength": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Strength for filters",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"NONE",
										"LOW",
										"MEDIUM",
										"HIGH",
									),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: OutputStrength
							"output_strength": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Strength for filters",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"NONE",
										"LOW",
										"MEDIUM",
										"HIGH",
									),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: Type
							"type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Type of filter in content policy",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"SEXUAL",
										"VIOLENCE",
										"HATE",
										"INSULTS",
										"MISCONDUCT",
										"PROMPT_ATTACK",
									),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "List of content filter configs in content policy.",
					Required:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeBetween(1, 6),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Content policy config for a guardrail.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Time Stamp",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "Time Stamp",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Description of the guardrail or its version",
		//	  "maxLength": 200,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Description of the guardrail or its version",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FailureRecommendations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of failure recommendations",
		//	  "items": {
		//	    "description": "Recommendation for guardrail failure status",
		//	    "maxLength": 200,
		//	    "minLength": 1,
		//	    "type": "string"
		//	  },
		//	  "maxItems": 100,
		//	  "type": "array"
		//	}
		"failure_recommendations": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "List of failure recommendations",
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GuardrailArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Arn representation for the guardrail",
		//	  "maxLength": 2048,
		//	  "pattern": "^arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:[0-9]{12}:guardrail/[a-z0-9]+$",
		//	  "type": "string"
		//	}
		"guardrail_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Arn representation for the guardrail",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GuardrailId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Unique id for the guardrail",
		//	  "maxLength": 64,
		//	  "pattern": "^[a-z0-9]+$",
		//	  "type": "string"
		//	}
		"guardrail_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Unique id for the guardrail",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The KMS key with which the guardrail was encrypted at rest",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws(-[^:]+)?:kms:[a-zA-Z0-9-]*:[0-9]{12}:key/[a-zA-Z0-9-]{36}$",
		//	  "type": "string"
		//	}
		"kms_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The KMS key with which the guardrail was encrypted at rest",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 2048),
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws(-[^:]+)?:kms:[a-zA-Z0-9-]*:[0-9]{12}:key/[a-zA-Z0-9-]{36}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of the guardrail",
		//	  "maxLength": 50,
		//	  "minLength": 1,
		//	  "pattern": "^[0-9a-zA-Z-_]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of the guardrail",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 50),
				stringvalidator.RegexMatches(regexp.MustCompile("^[0-9a-zA-Z-_]+$"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: SensitiveInformationPolicyConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Sensitive information policy config for a guardrail.",
		//	  "properties": {
		//	    "PiiEntitiesConfig": {
		//	      "description": "List of entities.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Pii entity configuration.",
		//	        "properties": {
		//	          "Action": {
		//	            "description": "Options for sensitive information action.",
		//	            "enum": [
		//	              "BLOCK",
		//	              "ANONYMIZE"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "Type": {
		//	            "description": "The currently supported PII entities",
		//	            "enum": [
		//	              "ADDRESS",
		//	              "AGE",
		//	              "AWS_ACCESS_KEY",
		//	              "AWS_SECRET_KEY",
		//	              "CA_HEALTH_NUMBER",
		//	              "CA_SOCIAL_INSURANCE_NUMBER",
		//	              "CREDIT_DEBIT_CARD_CVV",
		//	              "CREDIT_DEBIT_CARD_EXPIRY",
		//	              "CREDIT_DEBIT_CARD_NUMBER",
		//	              "DRIVER_ID",
		//	              "EMAIL",
		//	              "INTERNATIONAL_BANK_ACCOUNT_NUMBER",
		//	              "IP_ADDRESS",
		//	              "LICENSE_PLATE",
		//	              "MAC_ADDRESS",
		//	              "NAME",
		//	              "PASSWORD",
		//	              "PHONE",
		//	              "PIN",
		//	              "SWIFT_CODE",
		//	              "UK_NATIONAL_HEALTH_SERVICE_NUMBER",
		//	              "UK_NATIONAL_INSURANCE_NUMBER",
		//	              "UK_UNIQUE_TAXPAYER_REFERENCE_NUMBER",
		//	              "URL",
		//	              "USERNAME",
		//	              "US_BANK_ACCOUNT_NUMBER",
		//	              "US_BANK_ROUTING_NUMBER",
		//	              "US_INDIVIDUAL_TAX_IDENTIFICATION_NUMBER",
		//	              "US_PASSPORT_NUMBER",
		//	              "US_SOCIAL_SECURITY_NUMBER",
		//	              "VEHICLE_IDENTIFICATION_NUMBER"
		//	            ],
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Action",
		//	          "Type"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "minItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "RegexesConfig": {
		//	      "description": "List of regex.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "A regex configuration.",
		//	        "properties": {
		//	          "Action": {
		//	            "description": "Options for sensitive information action.",
		//	            "enum": [
		//	              "BLOCK",
		//	              "ANONYMIZE"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "Description": {
		//	            "description": "The regex description.",
		//	            "maxLength": 1000,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Name": {
		//	            "description": "The regex name.",
		//	            "maxLength": 100,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Pattern": {
		//	            "description": "The regex pattern.",
		//	            "minLength": 1,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Action",
		//	          "Name",
		//	          "Pattern"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "minItems": 1,
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"sensitive_information_policy_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: PiiEntitiesConfig
				"pii_entities_config": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Action
							"action": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Options for sensitive information action.",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"BLOCK",
										"ANONYMIZE",
									),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: Type
							"type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The currently supported PII entities",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"ADDRESS",
										"AGE",
										"AWS_ACCESS_KEY",
										"AWS_SECRET_KEY",
										"CA_HEALTH_NUMBER",
										"CA_SOCIAL_INSURANCE_NUMBER",
										"CREDIT_DEBIT_CARD_CVV",
										"CREDIT_DEBIT_CARD_EXPIRY",
										"CREDIT_DEBIT_CARD_NUMBER",
										"DRIVER_ID",
										"EMAIL",
										"INTERNATIONAL_BANK_ACCOUNT_NUMBER",
										"IP_ADDRESS",
										"LICENSE_PLATE",
										"MAC_ADDRESS",
										"NAME",
										"PASSWORD",
										"PHONE",
										"PIN",
										"SWIFT_CODE",
										"UK_NATIONAL_HEALTH_SERVICE_NUMBER",
										"UK_NATIONAL_INSURANCE_NUMBER",
										"UK_UNIQUE_TAXPAYER_REFERENCE_NUMBER",
										"URL",
										"USERNAME",
										"US_BANK_ACCOUNT_NUMBER",
										"US_BANK_ROUTING_NUMBER",
										"US_INDIVIDUAL_TAX_IDENTIFICATION_NUMBER",
										"US_PASSPORT_NUMBER",
										"US_SOCIAL_SECURITY_NUMBER",
										"VEHICLE_IDENTIFICATION_NUMBER",
									),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "List of entities.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeAtLeast(1),
						listvalidator.UniqueValues(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: RegexesConfig
				"regexes_config": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Action
							"action": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Options for sensitive information action.",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"BLOCK",
										"ANONYMIZE",
									),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: Description
							"description": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The regex description.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 1000),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The regex name.",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 100),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: Pattern
							"pattern": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The regex pattern.",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthAtLeast(1),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "List of regex.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeAtLeast(1),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Sensitive information policy config for a guardrail.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Status of the guardrail",
		//	  "enum": [
		//	    "CREATING",
		//	    "UPDATING",
		//	    "VERSIONING",
		//	    "READY",
		//	    "FAILED",
		//	    "DELETING"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Status of the guardrail",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StatusReasons
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of status reasons",
		//	  "items": {
		//	    "description": "Reason for guardrail status",
		//	    "maxLength": 200,
		//	    "minLength": 1,
		//	    "type": "string"
		//	  },
		//	  "maxItems": 100,
		//	  "type": "array"
		//	}
		"status_reasons": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "List of status reasons",
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of Tags",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Definition of the key/value pair for a tag",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "Tag Key",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "^[a-zA-Z0-9\\s._:/=+@-]*$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "Tag Value",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "pattern": "^[a-zA-Z0-9\\s._:/=+@-]*$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Tag Key",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9\\s._:/=+@-]*$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Tag Value",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
							stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9\\s._:/=+@-]*$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "List of Tags",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(0, 200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TopicPolicyConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Topic policy config for a guardrail.",
		//	  "properties": {
		//	    "TopicsConfig": {
		//	      "description": "List of topic configs in topic policy.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Topic config in topic policy.",
		//	        "properties": {
		//	          "Definition": {
		//	            "description": "Definition of topic in topic policy",
		//	            "maxLength": 200,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Examples": {
		//	            "description": "List of text examples",
		//	            "items": {
		//	              "description": "Text example in topic policy",
		//	              "maxLength": 100,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "minItems": 0,
		//	            "type": "array"
		//	          },
		//	          "Name": {
		//	            "description": "Name of topic in topic policy",
		//	            "maxLength": 100,
		//	            "minLength": 1,
		//	            "pattern": "^[0-9a-zA-Z-_ !?.]+$",
		//	            "type": "string"
		//	          },
		//	          "Type": {
		//	            "description": "Type of topic in a policy",
		//	            "enum": [
		//	              "DENY"
		//	            ],
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Definition",
		//	          "Name",
		//	          "Type"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "minItems": 1,
		//	      "type": "array"
		//	    }
		//	  },
		//	  "required": [
		//	    "TopicsConfig"
		//	  ],
		//	  "type": "object"
		//	}
		"topic_policy_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: TopicsConfig
				"topics_config": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Definition
							"definition": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Definition of topic in topic policy",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 200),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: Examples
							"examples": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "List of text examples",
								Optional:    true,
								Computed:    true,
								Validators: []validator.List{ /*START VALIDATORS*/
									listvalidator.SizeAtLeast(0),
									listvalidator.ValueStringsAre(
										stringvalidator.LengthBetween(1, 100),
									),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
									listplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Name of topic in topic policy",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 100),
									stringvalidator.RegexMatches(regexp.MustCompile("^[0-9a-zA-Z-_ !?.]+$"), ""),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: Type
							"type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Type of topic in a policy",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"DENY",
									),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "List of topic configs in topic policy.",
					Required:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeAtLeast(1),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Topic policy config for a guardrail.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Time Stamp",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "Time Stamp",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Version
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Guardrail version",
		//	  "pattern": "^(([1-9][0-9]{0,7})|(DRAFT))$",
		//	  "type": "string"
		//	}
		"version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Guardrail version",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: WordPolicyConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Word policy config for a guardrail.",
		//	  "properties": {
		//	    "ManagedWordListsConfig": {
		//	      "description": "A config for the list of managed words.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "A managed words config.",
		//	        "properties": {
		//	          "Type": {
		//	            "description": "Options for managed words.",
		//	            "enum": [
		//	              "PROFANITY"
		//	            ],
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Type"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "WordsConfig": {
		//	      "description": "List of custom word configs.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "A custom word config.",
		//	        "properties": {
		//	          "Text": {
		//	            "description": "The custom word text.",
		//	            "minLength": 1,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Text"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "minItems": 1,
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"word_policy_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ManagedWordListsConfig
				"managed_word_lists_config": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Type
							"type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Options for managed words.",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"PROFANITY",
									),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "A config for the list of managed words.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: WordsConfig
				"words_config": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Text
							"text": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The custom word text.",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthAtLeast(1),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "List of custom word configs.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeAtLeast(1),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Word policy config for a guardrail.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
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
		Description: "Definition of AWS::Bedrock::Guardrail Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Bedrock::Guardrail").WithTerraformTypeName("awscc_bedrock_guardrail")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                              "Action",
		"blocked_input_messaging":             "BlockedInputMessaging",
		"blocked_outputs_messaging":           "BlockedOutputsMessaging",
		"content_policy_config":               "ContentPolicyConfig",
		"created_at":                          "CreatedAt",
		"definition":                          "Definition",
		"description":                         "Description",
		"examples":                            "Examples",
		"failure_recommendations":             "FailureRecommendations",
		"filters_config":                      "FiltersConfig",
		"guardrail_arn":                       "GuardrailArn",
		"guardrail_id":                        "GuardrailId",
		"input_strength":                      "InputStrength",
		"key":                                 "Key",
		"kms_key_arn":                         "KmsKeyArn",
		"managed_word_lists_config":           "ManagedWordListsConfig",
		"name":                                "Name",
		"output_strength":                     "OutputStrength",
		"pattern":                             "Pattern",
		"pii_entities_config":                 "PiiEntitiesConfig",
		"regexes_config":                      "RegexesConfig",
		"sensitive_information_policy_config": "SensitiveInformationPolicyConfig",
		"status":                              "Status",
		"status_reasons":                      "StatusReasons",
		"tags":                                "Tags",
		"text":                                "Text",
		"topic_policy_config":                 "TopicPolicyConfig",
		"topics_config":                       "TopicsConfig",
		"type":                                "Type",
		"updated_at":                          "UpdatedAt",
		"value":                               "Value",
		"version":                             "Version",
		"word_policy_config":                  "WordPolicyConfig",
		"words_config":                        "WordsConfig",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
