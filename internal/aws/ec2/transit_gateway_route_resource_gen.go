// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_transit_gateway_route", transitGatewayRouteResource)
}

// transitGatewayRouteResource returns the Terraform awscc_ec2_transit_gateway_route resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::TransitGatewayRoute resource.
func transitGatewayRouteResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Blackhole
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether to drop traffic that matches this route.",
		//	  "type": "boolean"
		//	}
		"blackhole": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether to drop traffic that matches this route.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
				boolplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DestinationCidrBlock
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The CIDR range used for destination matches. Routing decisions are based on the most specific match.",
		//	  "type": "string"
		//	}
		"destination_cidr_block": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The CIDR range used for destination matches. Routing decisions are based on the most specific match.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TransitGatewayAttachmentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of transit gateway attachment.",
		//	  "type": "string"
		//	}
		"transit_gateway_attachment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of transit gateway attachment.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TransitGatewayRouteTableId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of transit gateway route table.",
		//	  "type": "string"
		//	}
		"transit_gateway_route_table_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of transit gateway route table.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
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
		Description: "Resource Type definition for AWS::EC2::TransitGatewayRoute",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::TransitGatewayRoute").WithTerraformTypeName("awscc_ec2_transit_gateway_route")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"blackhole":                      "Blackhole",
		"destination_cidr_block":         "DestinationCidrBlock",
		"transit_gateway_attachment_id":  "TransitGatewayAttachmentId",
		"transit_gateway_route_table_id": "TransitGatewayRouteTableId",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}