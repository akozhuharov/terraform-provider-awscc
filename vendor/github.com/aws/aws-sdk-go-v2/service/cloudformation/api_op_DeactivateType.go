// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deactivates a public extension that was previously activated in this account and
// region. Once deactivated, an extension cannot be used in any CloudFormation
// operation. This includes stack update operations where the stack template
// includes the extension, even if no updates are being made to the extension. In
// addition, deactivated extensions are not automatically updated if a new version
// of the extension is released.
func (c *Client) DeactivateType(ctx context.Context, params *DeactivateTypeInput, optFns ...func(*Options)) (*DeactivateTypeOutput, error) {
	if params == nil {
		params = &DeactivateTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeactivateType", params, optFns, c.addOperationDeactivateTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeactivateTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeactivateTypeInput struct {

	// The Amazon Resource Name (ARN) for the extension, in this account and region.
	// Conditional: You must specify either Arn, or TypeName and Type.
	Arn *string

	// The extension type. Conditional: You must specify either Arn, or TypeName and
	// Type.
	Type types.ThirdPartyType

	// The type name of the extension, in this account and region. If you specified a
	// type name alias when enabling the extension, use the type name alias.
	// Conditional: You must specify either Arn, or TypeName and Type.
	TypeName *string
}

type DeactivateTypeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDeactivateTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeactivateType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeactivateType{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeactivateType(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDeactivateType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DeactivateType",
	}
}
