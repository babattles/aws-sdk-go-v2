// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches predefined attributes that meet certain criteria. Predefined
// attributes are attributes in an Amazon Connect instance that can be used to
// route contacts to an agent or pools of agents within a queue. For more
// information, see [Create predefined attributes for routing contacts to agents].
//
// [Create predefined attributes for routing contacts to agents]: https://docs.aws.amazon.com/connect/latest/adminguide/predefined-attributes.html
func (c *Client) SearchPredefinedAttributes(ctx context.Context, params *SearchPredefinedAttributesInput, optFns ...func(*Options)) (*SearchPredefinedAttributesOutput, error) {
	if params == nil {
		params = &SearchPredefinedAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchPredefinedAttributes", params, optFns, c.addOperationSearchPredefinedAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchPredefinedAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchPredefinedAttributesInput struct {

	// The identifier of the Amazon Connect instance. You can find the instance ID in
	// the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// The search criteria to be used to return predefined attributes.
	SearchCriteria *types.PredefinedAttributeSearchCriteria

	noSmithyDocumentSerde
}

type SearchPredefinedAttributesOutput struct {

	// The approximate number of predefined attributes which matched your search query.
	ApproximateTotalCount *int64

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// Predefined attributes matched by the search criteria.
	PredefinedAttributes []types.PredefinedAttribute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchPredefinedAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchPredefinedAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchPredefinedAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchPredefinedAttributes"); err != nil {
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
	if err = addOpSearchPredefinedAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchPredefinedAttributes(options.Region), middleware.Before); err != nil {
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

// SearchPredefinedAttributesPaginatorOptions is the paginator options for
// SearchPredefinedAttributes
type SearchPredefinedAttributesPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchPredefinedAttributesPaginator is a paginator for
// SearchPredefinedAttributes
type SearchPredefinedAttributesPaginator struct {
	options   SearchPredefinedAttributesPaginatorOptions
	client    SearchPredefinedAttributesAPIClient
	params    *SearchPredefinedAttributesInput
	nextToken *string
	firstPage bool
}

// NewSearchPredefinedAttributesPaginator returns a new
// SearchPredefinedAttributesPaginator
func NewSearchPredefinedAttributesPaginator(client SearchPredefinedAttributesAPIClient, params *SearchPredefinedAttributesInput, optFns ...func(*SearchPredefinedAttributesPaginatorOptions)) *SearchPredefinedAttributesPaginator {
	if params == nil {
		params = &SearchPredefinedAttributesInput{}
	}

	options := SearchPredefinedAttributesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchPredefinedAttributesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchPredefinedAttributesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchPredefinedAttributes page.
func (p *SearchPredefinedAttributesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchPredefinedAttributesOutput, error) {
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
	result, err := p.client.SearchPredefinedAttributes(ctx, &params, optFns...)
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

// SearchPredefinedAttributesAPIClient is a client that implements the
// SearchPredefinedAttributes operation.
type SearchPredefinedAttributesAPIClient interface {
	SearchPredefinedAttributes(context.Context, *SearchPredefinedAttributesInput, ...func(*Options)) (*SearchPredefinedAttributesOutput, error)
}

var _ SearchPredefinedAttributesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opSearchPredefinedAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchPredefinedAttributes",
	}
}
