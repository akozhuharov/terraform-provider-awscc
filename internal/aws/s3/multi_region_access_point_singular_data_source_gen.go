// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package s3

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_s3_multi_region_access_point", multiRegionAccessPointDataSource)
}

// multiRegionAccessPointDataSource returns the Terraform awscc_s3_multi_region_access_point data source.
// This Terraform data source corresponds to the CloudFormation AWS::S3::MultiRegionAccessPoint resource.
func multiRegionAccessPointDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"alias": {
			// Property: Alias
			// CloudFormation resource type schema:
			// {
			//   "description": "The alias is a unique identifier to, and is part of the public DNS name for this Multi Region Access Point",
			//   "type": "string"
			// }
			Description: "The alias is a unique identifier to, and is part of the public DNS name for this Multi Region Access Point",
			Type:        types.StringType,
			Computed:    true,
		},
		"created_at": {
			// Property: CreatedAt
			// CloudFormation resource type schema:
			// {
			//   "description": "The timestamp of the when the Multi Region Access Point is created",
			//   "type": "string"
			// }
			Description: "The timestamp of the when the Multi Region Access Point is created",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name you want to assign to this Multi Region Access Point.",
			//   "maxLength": 50,
			//   "minLength": 3,
			//   "pattern": "^[a-z0-9][-a-z0-9]{1,48}[a-z0-9]$",
			//   "type": "string"
			// }
			Description: "The name you want to assign to this Multi Region Access Point.",
			Type:        types.StringType,
			Computed:    true,
		},
		"public_access_block_configuration": {
			// Property: PublicAccessBlockConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The PublicAccessBlock configuration that you want to apply to this Multi Region Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.",
			//   "properties": {
			//     "BlockPublicAcls": {
			//       "description": "Specifies whether Amazon S3 should block public access control lists (ACLs) for buckets in this account. Setting this element to TRUE causes the following behavior:\n- PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.\n - PUT Object calls fail if the request includes a public ACL.\n. - PUT Bucket calls fail if the request includes a public ACL.\nEnabling this setting doesn't affect existing policies or ACLs.",
			//       "type": "boolean"
			//     },
			//     "BlockPublicPolicy": {
			//       "description": "Specifies whether Amazon S3 should block public bucket policies for buckets in this account. Setting this element to TRUE causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access. Enabling this setting doesn't affect existing bucket policies.",
			//       "type": "boolean"
			//     },
			//     "IgnorePublicAcls": {
			//       "description": "Specifies whether Amazon S3 should ignore public ACLs for buckets in this account. Setting this element to TRUE causes Amazon S3 to ignore all public ACLs on buckets in this account and any objects that they contain. Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set.",
			//       "type": "boolean"
			//     },
			//     "RestrictPublicBuckets": {
			//       "description": "Specifies whether Amazon S3 should restrict public bucket policies for this bucket. Setting this element to TRUE restricts access to this bucket to only AWS services and authorized users within this account if the bucket has a public policy.\nEnabling this setting doesn't affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked.",
			//       "type": "boolean"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The PublicAccessBlock configuration that you want to apply to this Multi Region Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"block_public_acls": {
						// Property: BlockPublicAcls
						Description: "Specifies whether Amazon S3 should block public access control lists (ACLs) for buckets in this account. Setting this element to TRUE causes the following behavior:\n- PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.\n - PUT Object calls fail if the request includes a public ACL.\n. - PUT Bucket calls fail if the request includes a public ACL.\nEnabling this setting doesn't affect existing policies or ACLs.",
						Type:        types.BoolType,
						Computed:    true,
					},
					"block_public_policy": {
						// Property: BlockPublicPolicy
						Description: "Specifies whether Amazon S3 should block public bucket policies for buckets in this account. Setting this element to TRUE causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access. Enabling this setting doesn't affect existing bucket policies.",
						Type:        types.BoolType,
						Computed:    true,
					},
					"ignore_public_acls": {
						// Property: IgnorePublicAcls
						Description: "Specifies whether Amazon S3 should ignore public ACLs for buckets in this account. Setting this element to TRUE causes Amazon S3 to ignore all public ACLs on buckets in this account and any objects that they contain. Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set.",
						Type:        types.BoolType,
						Computed:    true,
					},
					"restrict_public_buckets": {
						// Property: RestrictPublicBuckets
						Description: "Specifies whether Amazon S3 should restrict public bucket policies for this bucket. Setting this element to TRUE restricts access to this bucket to only AWS services and authorized users within this account if the bucket has a public policy.\nEnabling this setting doesn't affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked.",
						Type:        types.BoolType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"regions": {
			// Property: Regions
			// CloudFormation resource type schema:
			// {
			//   "description": "The list of buckets that you want to associate this Multi Region Access Point with.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "The name of the bucket that represents of the region belonging to this Multi Region Access Point.",
			//     "properties": {
			//       "Bucket": {
			//         "maxLength": 63,
			//         "minLength": 3,
			//         "pattern": "^[a-z0-9][a-z0-9//.//-]*[a-z0-9]$",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Bucket"
			//     ],
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "The list of buckets that you want to associate this Multi Region Access Point with.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"bucket": {
						// Property: Bucket
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::S3::MultiRegionAccessPoint",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3::MultiRegionAccessPoint").WithTerraformTypeName("awscc_s3_multi_region_access_point")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alias":                             "Alias",
		"block_public_acls":                 "BlockPublicAcls",
		"block_public_policy":               "BlockPublicPolicy",
		"bucket":                            "Bucket",
		"created_at":                        "CreatedAt",
		"ignore_public_acls":                "IgnorePublicAcls",
		"name":                              "Name",
		"public_access_block_configuration": "PublicAccessBlockConfiguration",
		"regions":                           "Regions",
		"restrict_public_buckets":           "RestrictPublicBuckets",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
