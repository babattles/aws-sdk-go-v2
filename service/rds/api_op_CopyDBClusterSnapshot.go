// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	presignedurlcust "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Copies a snapshot of a DB cluster.
//
// To copy a DB cluster snapshot from a shared manual DB cluster snapshot,
// SourceDBClusterSnapshotIdentifier must be the Amazon Resource Name (ARN) of the
// shared DB cluster snapshot.
//
// You can copy an encrypted DB cluster snapshot from another Amazon Web Services
// Region. In that case, the Amazon Web Services Region where you call the
// CopyDBClusterSnapshot operation is the destination Amazon Web Services Region
// for the encrypted DB cluster snapshot to be copied to. To copy an encrypted DB
// cluster snapshot from another Amazon Web Services Region, you must provide the
// following values:
//
//   - KmsKeyId - The Amazon Web Services Key Management System (Amazon Web
//     Services KMS) key identifier for the key to use to encrypt the copy of the DB
//     cluster snapshot in the destination Amazon Web Services Region.
//
//   - TargetDBClusterSnapshotIdentifier - The identifier for the new copy of the
//     DB cluster snapshot in the destination Amazon Web Services Region.
//
//   - SourceDBClusterSnapshotIdentifier - The DB cluster snapshot identifier for
//     the encrypted DB cluster snapshot to be copied. This identifier must be in the
//     ARN format for the source Amazon Web Services Region and is the same value as
//     the SourceDBClusterSnapshotIdentifier in the presigned URL.
//
// To cancel the copy operation once it is in progress, delete the target DB
// cluster snapshot identified by TargetDBClusterSnapshotIdentifier while that DB
// cluster snapshot is in "copying" status.
//
// For more information on copying encrypted Amazon Aurora DB cluster snapshots
// from one Amazon Web Services Region to another, see [Copying a Snapshot]in the Amazon Aurora User
// Guide.
//
// For more information on Amazon Aurora DB clusters, see [What is Amazon Aurora?] in the Amazon Aurora
// User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User Guide.
//
// [Copying a Snapshot]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_CopySnapshot.html
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
func (c *Client) CopyDBClusterSnapshot(ctx context.Context, params *CopyDBClusterSnapshotInput, optFns ...func(*Options)) (*CopyDBClusterSnapshotOutput, error) {
	if params == nil {
		params = &CopyDBClusterSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopyDBClusterSnapshot", params, optFns, c.addOperationCopyDBClusterSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopyDBClusterSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopyDBClusterSnapshotInput struct {

	// The identifier of the DB cluster snapshot to copy. This parameter isn't
	// case-sensitive.
	//
	// You can't copy an encrypted, shared DB cluster snapshot from one Amazon Web
	// Services Region to another.
	//
	// Constraints:
	//
	//   - Must specify a valid system snapshot in the "available" state.
	//
	//   - If the source snapshot is in the same Amazon Web Services Region as the
	//   copy, specify a valid DB snapshot identifier.
	//
	//   - If the source snapshot is in a different Amazon Web Services Region than
	//   the copy, specify a valid DB cluster snapshot ARN. For more information, go to [Copying Snapshots Across Amazon Web Services Regions]
	//   in the Amazon Aurora User Guide.
	//
	// Example: my-cluster-snapshot1
	//
	// [Copying Snapshots Across Amazon Web Services Regions]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_CopySnapshot.html#USER_CopySnapshot.AcrossRegions
	//
	// This member is required.
	SourceDBClusterSnapshotIdentifier *string

	// The identifier of the new DB cluster snapshot to create from the source DB
	// cluster snapshot. This parameter isn't case-sensitive.
	//
	// Constraints:
	//
	//   - Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	//   - First character must be a letter.
	//
	//   - Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Example: my-cluster-snapshot2
	//
	// This member is required.
	TargetDBClusterSnapshotIdentifier *string

	// Specifies whether to copy all tags from the source DB cluster snapshot to the
	// target DB cluster snapshot. By default, tags are not copied.
	CopyTags *bool

	// The Amazon Web Services KMS key identifier for an encrypted DB cluster
	// snapshot. The Amazon Web Services KMS key identifier is the key ARN, key ID,
	// alias ARN, or alias name for the Amazon Web Services KMS key.
	//
	// If you copy an encrypted DB cluster snapshot from your Amazon Web Services
	// account, you can specify a value for KmsKeyId to encrypt the copy with a new
	// KMS key. If you don't specify a value for KmsKeyId , then the copy of the DB
	// cluster snapshot is encrypted with the same KMS key as the source DB cluster
	// snapshot.
	//
	// If you copy an encrypted DB cluster snapshot that is shared from another Amazon
	// Web Services account, then you must specify a value for KmsKeyId .
	//
	// To copy an encrypted DB cluster snapshot to another Amazon Web Services Region,
	// you must set KmsKeyId to the Amazon Web Services KMS key identifier you want to
	// use to encrypt the copy of the DB cluster snapshot in the destination Amazon Web
	// Services Region. KMS keys are specific to the Amazon Web Services Region that
	// they are created in, and you can't use KMS keys from one Amazon Web Services
	// Region in another Amazon Web Services Region.
	//
	// If you copy an unencrypted DB cluster snapshot and specify a value for the
	// KmsKeyId parameter, an error is returned.
	KmsKeyId *string

	// When you are copying a DB cluster snapshot from one Amazon Web Services
	// GovCloud (US) Region to another, the URL that contains a Signature Version 4
	// signed request for the CopyDBClusterSnapshot API operation in the Amazon Web
	// Services Region that contains the source DB cluster snapshot to copy. Use the
	// PreSignedUrl parameter when copying an encrypted DB cluster snapshot from
	// another Amazon Web Services Region. Don't specify PreSignedUrl when copying an
	// encrypted DB cluster snapshot in the same Amazon Web Services Region.
	//
	// This setting applies only to Amazon Web Services GovCloud (US) Regions. It's
	// ignored in other Amazon Web Services Regions.
	//
	// The presigned URL must be a valid request for the CopyDBClusterSnapshot API
	// operation that can run in the source Amazon Web Services Region that contains
	// the encrypted DB cluster snapshot to copy. The presigned URL request must
	// contain the following parameter values:
	//
	//   - KmsKeyId - The KMS key identifier for the KMS key to use to encrypt the copy
	//   of the DB cluster snapshot in the destination Amazon Web Services Region. This
	//   is the same identifier for both the CopyDBClusterSnapshot operation that is
	//   called in the destination Amazon Web Services Region, and the operation
	//   contained in the presigned URL.
	//
	//   - DestinationRegion - The name of the Amazon Web Services Region that the DB
	//   cluster snapshot is to be created in.
	//
	//   - SourceDBClusterSnapshotIdentifier - The DB cluster snapshot identifier for
	//   the encrypted DB cluster snapshot to be copied. This identifier must be in the
	//   Amazon Resource Name (ARN) format for the source Amazon Web Services Region. For
	//   example, if you are copying an encrypted DB cluster snapshot from the us-west-2
	//   Amazon Web Services Region, then your SourceDBClusterSnapshotIdentifier looks
	//   like the following example:
	//   arn:aws:rds:us-west-2:123456789012:cluster-snapshot:aurora-cluster1-snapshot-20161115
	//   .
	//
	// To learn how to generate a Signature Version 4 signed request, see [Authenticating Requests: Using Query Parameters (Amazon Web Services Signature Version 4)] and [Signature Version 4 Signing Process].
	//
	// If you are using an Amazon Web Services SDK tool or the CLI, you can specify
	// SourceRegion (or --source-region for the CLI) instead of specifying PreSignedUrl
	// manually. Specifying SourceRegion autogenerates a presigned URL that is a valid
	// request for the operation that can run in the source Amazon Web Services Region.
	//
	// [Authenticating Requests: Using Query Parameters (Amazon Web Services Signature Version 4)]: https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html
	// [Signature Version 4 Signing Process]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
	PreSignedUrl *string

	// The AWS region the resource is in. The presigned URL will be created with this
	// region, if the PresignURL member is empty set.
	SourceRegion *string

	// A list of tags.
	//
	// For more information, see [Tagging Amazon RDS resources] in the Amazon RDS User Guide or [Tagging Amazon Aurora and Amazon RDS resources] in the Amazon
	// Aurora User Guide.
	//
	// [Tagging Amazon RDS resources]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html
	// [Tagging Amazon Aurora and Amazon RDS resources]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Tagging.html
	Tags []types.Tag

	// Used by the SDK's PresignURL autofill customization to specify the region the
	// of the client's request.
	destinationRegion *string

	noSmithyDocumentSerde
}

type CopyDBClusterSnapshotOutput struct {

	// Contains the details for an Amazon RDS DB cluster snapshot
	//
	// This data type is used as a response element in the DescribeDBClusterSnapshots
	// action.
	DBClusterSnapshot *types.DBClusterSnapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCopyDBClusterSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCopyDBClusterSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCopyDBClusterSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CopyDBClusterSnapshot"); err != nil {
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
	if err = addCopyDBClusterSnapshotPresignURLMiddleware(stack, options); err != nil {
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
	if err = addOpCopyDBClusterSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCopyDBClusterSnapshot(options.Region), middleware.Before); err != nil {
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

func copyCopyDBClusterSnapshotInputForPresign(params interface{}) (interface{}, error) {
	input, ok := params.(*CopyDBClusterSnapshotInput)
	if !ok {
		return nil, fmt.Errorf("expect *CopyDBClusterSnapshotInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func getCopyDBClusterSnapshotPreSignedUrl(params interface{}) (string, bool, error) {
	input, ok := params.(*CopyDBClusterSnapshotInput)
	if !ok {
		return ``, false, fmt.Errorf("expect *CopyDBClusterSnapshotInput type, got %T", params)
	}
	if input.PreSignedUrl == nil || len(*input.PreSignedUrl) == 0 {
		return ``, false, nil
	}
	return *input.PreSignedUrl, true, nil
}
func getCopyDBClusterSnapshotSourceRegion(params interface{}) (string, bool, error) {
	input, ok := params.(*CopyDBClusterSnapshotInput)
	if !ok {
		return ``, false, fmt.Errorf("expect *CopyDBClusterSnapshotInput type, got %T", params)
	}
	if input.SourceRegion == nil || len(*input.SourceRegion) == 0 {
		return ``, false, nil
	}
	return *input.SourceRegion, true, nil
}
func setCopyDBClusterSnapshotPreSignedUrl(params interface{}, value string) error {
	input, ok := params.(*CopyDBClusterSnapshotInput)
	if !ok {
		return fmt.Errorf("expect *CopyDBClusterSnapshotInput type, got %T", params)
	}
	input.PreSignedUrl = &value
	return nil
}
func setCopyDBClusterSnapshotdestinationRegion(params interface{}, value string) error {
	input, ok := params.(*CopyDBClusterSnapshotInput)
	if !ok {
		return fmt.Errorf("expect *CopyDBClusterSnapshotInput type, got %T", params)
	}
	input.destinationRegion = &value
	return nil
}
func addCopyDBClusterSnapshotPresignURLMiddleware(stack *middleware.Stack, options Options) error {
	return presignedurlcust.AddMiddleware(stack, presignedurlcust.Options{
		Accessor: presignedurlcust.ParameterAccessor{
			GetPresignedURL: getCopyDBClusterSnapshotPreSignedUrl,

			GetSourceRegion: getCopyDBClusterSnapshotSourceRegion,

			CopyInput: copyCopyDBClusterSnapshotInputForPresign,

			SetDestinationRegion: setCopyDBClusterSnapshotdestinationRegion,

			SetPresignedURL: setCopyDBClusterSnapshotPreSignedUrl,
		},
		Presigner: &presignAutoFillCopyDBClusterSnapshotClient{client: NewPresignClient(New(options))},
	})
}

type presignAutoFillCopyDBClusterSnapshotClient struct {
	client *PresignClient
}

// PresignURL is a middleware accessor that satisfies URLPresigner interface.
func (c *presignAutoFillCopyDBClusterSnapshotClient) PresignURL(ctx context.Context, srcRegion string, params interface{}) (*v4.PresignedHTTPRequest, error) {
	input, ok := params.(*CopyDBClusterSnapshotInput)
	if !ok {
		return nil, fmt.Errorf("expect *CopyDBClusterSnapshotInput type, got %T", params)
	}
	optFn := func(o *Options) {
		o.Region = srcRegion
		o.APIOptions = append(o.APIOptions, presignedurlcust.RemoveMiddleware)
	}
	presignOptFn := WithPresignClientFromClientOptions(optFn)
	return c.client.PresignCopyDBClusterSnapshot(ctx, input, presignOptFn)
}

func newServiceMetadataMiddleware_opCopyDBClusterSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CopyDBClusterSnapshot",
	}
}

// PresignCopyDBClusterSnapshot is used to generate a presigned HTTP Request which
// contains presigned URL, signed headers and HTTP method used.
func (c *PresignClient) PresignCopyDBClusterSnapshot(ctx context.Context, params *CopyDBClusterSnapshotInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	if params == nil {
		params = &CopyDBClusterSnapshotInput{}
	}
	options := c.options.copy()
	for _, fn := range optFns {
		fn(&options)
	}
	clientOptFns := append(options.ClientOptions, withNopHTTPClientAPIOption)

	result, _, err := c.client.invokeOperation(ctx, "CopyDBClusterSnapshot", params, clientOptFns,
		c.client.addOperationCopyDBClusterSnapshotMiddlewares,
		presignConverter(options).convertToPresignMiddleware,
	)
	if err != nil {
		return nil, err
	}

	out := result.(*v4.PresignedHTTPRequest)
	return out, nil
}
