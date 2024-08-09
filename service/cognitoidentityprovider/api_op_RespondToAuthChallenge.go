// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Some API operations in a user pool generate a challenge, like a prompt for an
// MFA code, for device authentication that bypasses MFA, or for a custom
// authentication challenge. A RespondToAuthChallenge API request provides the
// answer to that challenge, like a code or a secure remote password (SRP). The
// parameters of a response to an authentication challenge vary with the type of
// challenge.
//
// For more information about custom authentication challenges, see [Custom authentication challenge Lambda triggers].
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with [Amazon Pinpoint]. Amazon Cognito uses the
// registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in.
//
// If you have never used SMS text messages with Amazon Cognito or any other
// Amazon Web Services service, Amazon Simple Notification Service might place your
// account in the SMS sandbox. In [sandbox mode], you can send messages only to verified phone
// numbers. After you test your app while in the sandbox environment, you can move
// out of the sandbox and into production. For more information, see [SMS message settings for Amazon Cognito user pools]in the Amazon
// Cognito Developer Guide.
//
// [SMS message settings for Amazon Cognito user pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [sandbox mode]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
// [Custom authentication challenge Lambda triggers]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-challenge.html
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
func (c *Client) RespondToAuthChallenge(ctx context.Context, params *RespondToAuthChallengeInput, optFns ...func(*Options)) (*RespondToAuthChallengeOutput, error) {
	if params == nil {
		params = &RespondToAuthChallengeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RespondToAuthChallenge", params, optFns, c.addOperationRespondToAuthChallengeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RespondToAuthChallengeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to respond to an authentication challenge.
type RespondToAuthChallengeInput struct {

	// The challenge name. For more information, see [InitiateAuth].
	//
	// ADMIN_NO_SRP_AUTH isn't a valid value.
	//
	// [InitiateAuth]: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_InitiateAuth.html
	//
	// This member is required.
	ChallengeName types.ChallengeNameType

	// The app client ID.
	//
	// This member is required.
	ClientId *string

	// The Amazon Pinpoint analytics metadata that contributes to your metrics for
	// RespondToAuthChallenge calls.
	AnalyticsMetadata *types.AnalyticsMetadataType

	// The responses to the challenge that you received in the previous request. Each
	// challenge has its own required response parameters. The following examples are
	// partial JSON request bodies that highlight challenge-response parameters.
	//
	// You must provide a SECRET_HASH parameter in all challenge responses to an app
	// client that has a client secret.
	//
	// SMS_MFA "ChallengeName": "SMS_MFA", "ChallengeResponses": {"SMS_MFA_CODE":
	// "[SMS_code]", "USERNAME": "[username]"}
	//
	// PASSWORD_VERIFIER "ChallengeName": "PASSWORD_VERIFIER", "ChallengeResponses":
	// {"PASSWORD_CLAIM_SIGNATURE": "[claim_signature]", "PASSWORD_CLAIM_SECRET_BLOCK":
	// "[secret_block]", "TIMESTAMP": [timestamp], "USERNAME": "[username]"}
	//
	// Add "DEVICE_KEY" when you sign in with a remembered device.
	//
	// CUSTOM_CHALLENGE "ChallengeName": "CUSTOM_CHALLENGE", "ChallengeResponses":
	// {"USERNAME": "[username]", "ANSWER": "[challenge_answer]"}
	//
	// Add "DEVICE_KEY" when you sign in with a remembered device.
	//
	// NEW_PASSWORD_REQUIRED "ChallengeName": "NEW_PASSWORD_REQUIRED",
	// "ChallengeResponses": {"NEW_PASSWORD": "[new_password]", "USERNAME":
	// "[username]"}
	//
	// To set any required attributes that InitiateAuth returned in an
	// requiredAttributes parameter, add "userAttributes.[attribute_name]":
	// "[attribute_value]" . This parameter can also set values for writable attributes
	// that aren't required by your user pool.
	//
	// In a NEW_PASSWORD_REQUIRED challenge response, you can't modify a required
	// attribute that already has a value. In RespondToAuthChallenge , set a value for
	// any keys that Amazon Cognito returned in the requiredAttributes parameter, then
	// use the UpdateUserAttributes API operation to modify the value of any
	// additional attributes.
	//
	// SOFTWARE_TOKEN_MFA "ChallengeName": "SOFTWARE_TOKEN_MFA", "ChallengeResponses":
	// {"USERNAME": "[username]", "SOFTWARE_TOKEN_MFA_CODE": [authenticator_code]}
	//
	// DEVICE_SRP_AUTH "ChallengeName": "DEVICE_SRP_AUTH", "ChallengeResponses":
	// {"USERNAME": "[username]", "DEVICE_KEY": "[device_key]", "SRP_A": "[srp_a]"}
	//
	// DEVICE_PASSWORD_VERIFIER "ChallengeName": "DEVICE_PASSWORD_VERIFIER",
	// "ChallengeResponses": {"DEVICE_KEY": "[device_key]", "PASSWORD_CLAIM_SIGNATURE":
	// "[claim_signature]", "PASSWORD_CLAIM_SECRET_BLOCK": "[secret_block]",
	// "TIMESTAMP": [timestamp], "USERNAME": "[username]"}
	//
	// MFA_SETUP "ChallengeName": "MFA_SETUP", "ChallengeResponses": {"USERNAME":
	// "[username]"}, "SESSION": "[Session ID from VerifySoftwareToken]"
	//
	// SELECT_MFA_TYPE "ChallengeName": "SELECT_MFA_TYPE", "ChallengeResponses":
	// {"USERNAME": "[username]", "ANSWER": "[SMS_MFA or SOFTWARE_TOKEN_MFA]"}
	//
	// For more information about SECRET_HASH , see [Computing secret hash values]. For information about DEVICE_KEY
	// , see [Working with user devices in your user pool].
	//
	// [Computing secret hash values]: https://docs.aws.amazon.com/cognito/latest/developerguide/signing-up-users-in-your-app.html#cognito-user-pools-computing-secret-hash
	// [Working with user devices in your user pool]: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.html
	ChallengeResponses map[string]string

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers.
	//
	// You create custom workflows by assigning Lambda functions to user pool
	// triggers. When you use the RespondToAuthChallenge API action, Amazon Cognito
	// invokes any functions that are assigned to the following triggers: post
	// authentication, pre token generation, define auth challenge, create auth
	// challenge, and verify auth challenge. When Amazon Cognito invokes any of these
	// functions, it passes a JSON payload, which the function receives as input. This
	// payload contains a clientMetadata attribute, which provides the data that you
	// assigned to the ClientMetadata parameter in your RespondToAuthChallenge request.
	// In your function code in Lambda, you can process the clientMetadata value to
	// enhance your workflow for your specific needs.
	//
	// For more information, see [Customizing user pool Workflows with Lambda Triggers] in the Amazon Cognito Developer Guide.
	//
	// When you use the ClientMetadata parameter, remember that Amazon Cognito won't
	// do the following:
	//
	//   - Store the ClientMetadata value. This data is available only to Lambda
	//   triggers that are assigned to a user pool to support custom workflows. If your
	//   user pool configuration doesn't include triggers, the ClientMetadata parameter
	//   serves no purpose.
	//
	//   - Validate the ClientMetadata value.
	//
	//   - Encrypt the ClientMetadata value. Don't use Amazon Cognito to provide
	//   sensitive information.
	//
	// [Customizing user pool Workflows with Lambda Triggers]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html
	ClientMetadata map[string]string

	// The session that should be passed both ways in challenge-response calls to the
	// service. If InitiateAuth or RespondToAuthChallenge API call determines that the
	// caller must pass another challenge, they return a session with other challenge
	// parameters. This session should be passed as it is to the next
	// RespondToAuthChallenge API call.
	Session *string

	// Contextual data about your user session, such as the device fingerprint, IP
	// address, or location. Amazon Cognito advanced security evaluates the risk of an
	// authentication event based on the context that your app generates and passes to
	// Amazon Cognito when it makes API requests.
	UserContextData *types.UserContextDataType

	noSmithyDocumentSerde
}

// The response to respond to the authentication challenge.
type RespondToAuthChallengeOutput struct {

	// The result returned by the server in response to the request to respond to the
	// authentication challenge.
	AuthenticationResult *types.AuthenticationResultType

	// The challenge name. For more information, see [InitiateAuth].
	//
	// [InitiateAuth]: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_InitiateAuth.html
	ChallengeName types.ChallengeNameType

	// The challenge parameters. For more information, see [InitiateAuth].
	//
	// [InitiateAuth]: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_InitiateAuth.html
	ChallengeParameters map[string]string

	// The session that should be passed both ways in challenge-response calls to the
	// service. If the caller must pass another challenge, they return a session with
	// other challenge parameters. This session should be passed as it is to the next
	// RespondToAuthChallenge API call.
	Session *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRespondToAuthChallengeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRespondToAuthChallenge{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRespondToAuthChallenge{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RespondToAuthChallenge"); err != nil {
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
	if err = addOpRespondToAuthChallengeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRespondToAuthChallenge(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRespondToAuthChallenge(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RespondToAuthChallenge",
	}
}
