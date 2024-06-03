// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the backend environments for an Amplify app.
//
// This API is available only to Amplify Gen 1 applications where the backend is
// created using Amplify Studio or the Amplify command line interface (CLI). This
// API isn’t available to Amplify Gen 2 applications. When you deploy an
// application with Amplify Gen 2, you provision the app's backend infrastructure
// using Typescript code.
func (c *Client) ListBackendEnvironments(ctx context.Context, params *ListBackendEnvironmentsInput, optFns ...func(*Options)) (*ListBackendEnvironmentsOutput, error) {
	if params == nil {
		params = &ListBackendEnvironmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBackendEnvironments", params, optFns, c.addOperationListBackendEnvironmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBackendEnvironmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the list backend environments request.
type ListBackendEnvironmentsInput struct {

	// The unique ID for an Amplify app.
	//
	// This member is required.
	AppId *string

	// The name of the backend environment
	EnvironmentName *string

	// The maximum number of records to list in a single response.
	MaxResults int32

	// A pagination token. Set to null to start listing backend environments from the
	// start. If a non-null pagination token is returned in a result, pass its value in
	// here to list more backend environments.
	NextToken *string

	noSmithyDocumentSerde
}

// The result structure for the list backend environments result.
type ListBackendEnvironmentsOutput struct {

	// The list of backend environments for an Amplify app.
	//
	// This member is required.
	BackendEnvironments []types.BackendEnvironment

	// A pagination token. If a non-null pagination token is returned in a result,
	// pass its value in another request to retrieve more entries.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBackendEnvironmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBackendEnvironments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBackendEnvironments{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListBackendEnvironments"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListBackendEnvironmentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBackendEnvironments(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListBackendEnvironments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListBackendEnvironments",
	}
}
