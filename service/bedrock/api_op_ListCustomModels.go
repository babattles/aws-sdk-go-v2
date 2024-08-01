// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrock

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrock/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a list of the custom models that you have created with the
// CreateModelCustomizationJob operation.
//
// For more information, see [Custom models] in the [Amazon Bedrock User Guide].
//
// [Custom models]: https://docs.aws.amazon.com/bedrock/latest/userguide/custom-models.html
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func (c *Client) ListCustomModels(ctx context.Context, params *ListCustomModelsInput, optFns ...func(*Options)) (*ListCustomModelsOutput, error) {
	if params == nil {
		params = &ListCustomModelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCustomModels", params, optFns, c.addOperationListCustomModelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCustomModelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCustomModelsInput struct {

	// Return custom models only if the base model Amazon Resource Name (ARN) matches
	// this parameter.
	BaseModelArnEquals *string

	// Return custom models created after the specified time.
	CreationTimeAfter *time.Time

	// Return custom models created before the specified time.
	CreationTimeBefore *time.Time

	// Return custom models only if the foundation model Amazon Resource Name (ARN)
	// matches this parameter.
	FoundationModelArnEquals *string

	// Return custom models depending on if the current account owns them ( true ) or
	// if they were shared with the current account ( false ).
	IsOwned *bool

	// The maximum number of results to return in the response. If the total number of
	// results is greater than this value, use the token returned in the response in
	// the nextToken field when making another request to return the next batch of
	// results.
	MaxResults *int32

	// Return custom models only if the job name contains these characters.
	NameContains *string

	// If the total number of results is greater than the maxResults value provided in
	// the request, enter the token returned in the nextToken field in the response in
	// this field to return the next batch of results.
	NextToken *string

	// The field to sort by in the returned list of models.
	SortBy types.SortModelsBy

	// The sort order of the results.
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListCustomModelsOutput struct {

	// Model summaries.
	ModelSummaries []types.CustomModelSummary

	// If the total number of results is greater than the maxResults value provided in
	// the request, use this token when making another request in the nextToken field
	// to return the next batch of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCustomModelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCustomModels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCustomModels{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCustomModels"); err != nil {
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
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCustomModels(options.Region), middleware.Before); err != nil {
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

// ListCustomModelsPaginatorOptions is the paginator options for ListCustomModels
type ListCustomModelsPaginatorOptions struct {
	// The maximum number of results to return in the response. If the total number of
	// results is greater than this value, use the token returned in the response in
	// the nextToken field when making another request to return the next batch of
	// results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCustomModelsPaginator is a paginator for ListCustomModels
type ListCustomModelsPaginator struct {
	options   ListCustomModelsPaginatorOptions
	client    ListCustomModelsAPIClient
	params    *ListCustomModelsInput
	nextToken *string
	firstPage bool
}

// NewListCustomModelsPaginator returns a new ListCustomModelsPaginator
func NewListCustomModelsPaginator(client ListCustomModelsAPIClient, params *ListCustomModelsInput, optFns ...func(*ListCustomModelsPaginatorOptions)) *ListCustomModelsPaginator {
	if params == nil {
		params = &ListCustomModelsInput{}
	}

	options := ListCustomModelsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCustomModelsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCustomModelsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCustomModels page.
func (p *ListCustomModelsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCustomModelsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListCustomModels(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListCustomModelsAPIClient is a client that implements the ListCustomModels
// operation.
type ListCustomModelsAPIClient interface {
	ListCustomModels(context.Context, *ListCustomModelsInput, ...func(*Options)) (*ListCustomModelsOutput, error)
}

var _ ListCustomModelsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListCustomModels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCustomModels",
	}
}
