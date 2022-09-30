// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhuborchestrator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/migrationhuborchestrator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the step groups in a template.
func (c *Client) ListTemplateStepGroups(ctx context.Context, params *ListTemplateStepGroupsInput, optFns ...func(*Options)) (*ListTemplateStepGroupsOutput, error) {
	if params == nil {
		params = &ListTemplateStepGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTemplateStepGroups", params, optFns, c.addOperationListTemplateStepGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTemplateStepGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTemplateStepGroupsInput struct {

	// The ID of the template.
	//
	// This member is required.
	TemplateId *string

	// The maximum number of results that can be returned.
	MaxResults int32

	// The pagination token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListTemplateStepGroupsOutput struct {

	// The summary of the step group in the template.
	//
	// This member is required.
	TemplateStepGroupSummary []types.TemplateStepGroupSummary

	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTemplateStepGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTemplateStepGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTemplateStepGroups{}, middleware.After)
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
	if err = addOpListTemplateStepGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTemplateStepGroups(options.Region), middleware.Before); err != nil {
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

// ListTemplateStepGroupsAPIClient is a client that implements the
// ListTemplateStepGroups operation.
type ListTemplateStepGroupsAPIClient interface {
	ListTemplateStepGroups(context.Context, *ListTemplateStepGroupsInput, ...func(*Options)) (*ListTemplateStepGroupsOutput, error)
}

var _ ListTemplateStepGroupsAPIClient = (*Client)(nil)

// ListTemplateStepGroupsPaginatorOptions is the paginator options for
// ListTemplateStepGroups
type ListTemplateStepGroupsPaginatorOptions struct {
	// The maximum number of results that can be returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTemplateStepGroupsPaginator is a paginator for ListTemplateStepGroups
type ListTemplateStepGroupsPaginator struct {
	options   ListTemplateStepGroupsPaginatorOptions
	client    ListTemplateStepGroupsAPIClient
	params    *ListTemplateStepGroupsInput
	nextToken *string
	firstPage bool
}

// NewListTemplateStepGroupsPaginator returns a new ListTemplateStepGroupsPaginator
func NewListTemplateStepGroupsPaginator(client ListTemplateStepGroupsAPIClient, params *ListTemplateStepGroupsInput, optFns ...func(*ListTemplateStepGroupsPaginatorOptions)) *ListTemplateStepGroupsPaginator {
	if params == nil {
		params = &ListTemplateStepGroupsInput{}
	}

	options := ListTemplateStepGroupsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTemplateStepGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTemplateStepGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTemplateStepGroups page.
func (p *ListTemplateStepGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTemplateStepGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListTemplateStepGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTemplateStepGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "migrationhub-orchestrator",
		OperationName: "ListTemplateStepGroups",
	}
}