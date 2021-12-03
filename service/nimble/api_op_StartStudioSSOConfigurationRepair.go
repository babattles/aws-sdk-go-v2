// Code generated by smithy-go-codegen DO NOT EDIT.

package nimble

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/nimble/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Repairs the Amazon Web Services SSO configuration for a given studio. If the
// studio has a valid Amazon Web Services SSO configuration currently associated
// with it, this operation will fail with a validation error. If the studio does
// not have a valid Amazon Web Services SSO configuration currently associated with
// it, then a new Amazon Web Services SSO application is created for the studio and
// the studio is changed to the READY state. After the Amazon Web Services SSO
// application is repaired, you must use the Amazon Nimble Studio console to add
// administrators and users to your studio.
func (c *Client) StartStudioSSOConfigurationRepair(ctx context.Context, params *StartStudioSSOConfigurationRepairInput, optFns ...func(*Options)) (*StartStudioSSOConfigurationRepairOutput, error) {
	if params == nil {
		params = &StartStudioSSOConfigurationRepairInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartStudioSSOConfigurationRepair", params, optFns, c.addOperationStartStudioSSOConfigurationRepairMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartStudioSSOConfigurationRepairOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartStudioSSOConfigurationRepairInput struct {

	// The studio ID.
	//
	// This member is required.
	StudioId *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. If you don’t specify a client token, the AWS SDK automatically
	// generates a client token and uses it for the request to ensure idempotency.
	ClientToken *string

	noSmithyDocumentSerde
}

//
type StartStudioSSOConfigurationRepairOutput struct {

	// Information about a studio.
	//
	// This member is required.
	Studio *types.Studio

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartStudioSSOConfigurationRepairMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartStudioSSOConfigurationRepair{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartStudioSSOConfigurationRepair{}, middleware.After)
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
	if err = addIdempotencyToken_opStartStudioSSOConfigurationRepairMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartStudioSSOConfigurationRepairValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartStudioSSOConfigurationRepair(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartStudioSSOConfigurationRepair struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartStudioSSOConfigurationRepair) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartStudioSSOConfigurationRepair) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartStudioSSOConfigurationRepairInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartStudioSSOConfigurationRepairInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartStudioSSOConfigurationRepairMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartStudioSSOConfigurationRepair{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartStudioSSOConfigurationRepair(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "nimble",
		OperationName: "StartStudioSSOConfigurationRepair",
	}
}
