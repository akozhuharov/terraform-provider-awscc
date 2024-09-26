// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_transit_gateway", transitGatewayResource)
}

// transitGatewayResource returns the Terraform awscc_ec2_transit_gateway resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::TransitGateway resource.
func transitGatewayResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AmazonSideAsn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "int64",
		//	  "type": "integer"
		//	}
		"amazon_side_asn": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
				int64planmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AssociationDefaultRouteTableId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"association_default_route_table_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AutoAcceptSharedAttachments
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"auto_accept_shared_attachments": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DefaultRouteTableAssociation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"default_route_table_association": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DefaultRouteTablePropagation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"default_route_table_propagation": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DnsSupport
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"dns_support": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"transit_gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MulticastSupport
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"multicast_support": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PropagationDefaultRouteTableId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"propagation_default_route_table_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SecurityGroupReferencingSupport
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"security_group_referencing_support": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TransitGatewayArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"transit_gateway_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TransitGatewayCidrBlocks
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"transit_gateway_cidr_blocks": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpnEcmpSupport
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"vpn_ecmp_support": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
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
		Description: "Resource Type definition for AWS::EC2::TransitGateway",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::TransitGateway").WithTerraformTypeName("awscc_ec2_transit_gateway")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"amazon_side_asn":                    "AmazonSideAsn",
		"association_default_route_table_id": "AssociationDefaultRouteTableId",
		"auto_accept_shared_attachments":     "AutoAcceptSharedAttachments",
		"default_route_table_association":    "DefaultRouteTableAssociation",
		"default_route_table_propagation":    "DefaultRouteTablePropagation",
		"description":                        "Description",
		"dns_support":                        "DnsSupport",
		"key":                                "Key",
		"multicast_support":                  "MulticastSupport",
		"propagation_default_route_table_id": "PropagationDefaultRouteTableId",
		"security_group_referencing_support": "SecurityGroupReferencingSupport",
		"tags":                               "Tags",
		"transit_gateway_arn":                "TransitGatewayArn",
		"transit_gateway_cidr_blocks":        "TransitGatewayCidrBlocks",
		"transit_gateway_id":                 "Id",
		"value":                              "Value",
		"vpn_ecmp_support":                   "VpnEcmpSupport",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
