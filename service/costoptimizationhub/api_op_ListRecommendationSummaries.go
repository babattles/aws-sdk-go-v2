// Code generated by smithy-go-codegen DO NOT EDIT.

package costoptimizationhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/costoptimizationhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a concise representation of savings estimates for resources. Also
// returns de-duped savings across different types of recommendations.
//
// The following filters are not supported for this API: recommendationIds ,
// resourceArns , and resourceIds .
func (c *Client) ListRecommendationSummaries(ctx context.Context, params *ListRecommendationSummariesInput, optFns ...func(*Options)) (*ListRecommendationSummariesOutput, error) {
	if params == nil {
		params = &ListRecommendationSummariesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRecommendationSummaries", params, optFns, c.addOperationListRecommendationSummariesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRecommendationSummariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRecommendationSummariesInput struct {

	// The grouping of recommendations by a dimension.
	//
	// This member is required.
	GroupBy *string

	// Describes a filter that returns a more specific list of recommendations.
	// Filters recommendations by different dimensions.
	Filter *types.Filter

	// The maximum number of recommendations to be returned for the request.
	MaxResults *int32

	// Additional metrics to be returned for the request. The only valid value is
	// savingsPercentage .
	Metrics []types.SummaryMetrics

	// The token to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRecommendationSummariesOutput struct {

	// The currency code used for the recommendation.
	CurrencyCode *string

	// The total overall savings for the aggregated view.
	EstimatedTotalDedupedSavings *float64

	// The dimension used to group the recommendations by.
	GroupBy *string

	// A list of all savings recommendations.
	Items []types.RecommendationSummary

	// The results or descriptions for the additional metrics, based on whether the
	// metrics were or were not requested.
	Metrics *types.SummaryMetricsResult

	// The token to retrieve the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRecommendationSummariesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListRecommendationSummaries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListRecommendationSummaries{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRecommendationSummaries"); err != nil {
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
	if err = addOpListRecommendationSummariesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRecommendationSummaries(options.Region), middleware.Before); err != nil {
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

// ListRecommendationSummariesPaginatorOptions is the paginator options for
// ListRecommendationSummaries
type ListRecommendationSummariesPaginatorOptions struct {
	// The maximum number of recommendations to be returned for the request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRecommendationSummariesPaginator is a paginator for
// ListRecommendationSummaries
type ListRecommendationSummariesPaginator struct {
	options   ListRecommendationSummariesPaginatorOptions
	client    ListRecommendationSummariesAPIClient
	params    *ListRecommendationSummariesInput
	nextToken *string
	firstPage bool
}

// NewListRecommendationSummariesPaginator returns a new
// ListRecommendationSummariesPaginator
func NewListRecommendationSummariesPaginator(client ListRecommendationSummariesAPIClient, params *ListRecommendationSummariesInput, optFns ...func(*ListRecommendationSummariesPaginatorOptions)) *ListRecommendationSummariesPaginator {
	if params == nil {
		params = &ListRecommendationSummariesInput{}
	}

	options := ListRecommendationSummariesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRecommendationSummariesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRecommendationSummariesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRecommendationSummaries page.
func (p *ListRecommendationSummariesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRecommendationSummariesOutput, error) {
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
	result, err := p.client.ListRecommendationSummaries(ctx, &params, optFns...)
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

// ListRecommendationSummariesAPIClient is a client that implements the
// ListRecommendationSummaries operation.
type ListRecommendationSummariesAPIClient interface {
	ListRecommendationSummaries(context.Context, *ListRecommendationSummariesInput, ...func(*Options)) (*ListRecommendationSummariesOutput, error)
}

var _ ListRecommendationSummariesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListRecommendationSummaries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRecommendationSummaries",
	}
}
