// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagent/types"
)

func ExampleActionGroupExecutor_outputUsage() {
	var union types.ActionGroupExecutor
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ActionGroupExecutorMemberCustomControl:
		_ = v.Value // Value is types.CustomControlMethod

	case *types.ActionGroupExecutorMemberLambda:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ types.CustomControlMethod

func ExampleAPISchema_outputUsage() {
	var union types.APISchema
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.APISchemaMemberPayload:
		_ = v.Value // Value is string

	case *types.APISchemaMemberS3:
		_ = v.Value // Value is types.S3Identifier

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *types.S3Identifier

func ExampleFlowConnectionConfiguration_outputUsage() {
	var union types.FlowConnectionConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.FlowConnectionConfigurationMemberConditional:
		_ = v.Value // Value is types.FlowConditionalConnectionConfiguration

	case *types.FlowConnectionConfigurationMemberData:
		_ = v.Value // Value is types.FlowDataConnectionConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.FlowConditionalConnectionConfiguration
var _ *types.FlowDataConnectionConfiguration

func ExampleFlowNodeConfiguration_outputUsage() {
	var union types.FlowNodeConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.FlowNodeConfigurationMemberAgent:
		_ = v.Value // Value is types.AgentFlowNodeConfiguration

	case *types.FlowNodeConfigurationMemberCollector:
		_ = v.Value // Value is types.CollectorFlowNodeConfiguration

	case *types.FlowNodeConfigurationMemberCondition:
		_ = v.Value // Value is types.ConditionFlowNodeConfiguration

	case *types.FlowNodeConfigurationMemberInput:
		_ = v.Value // Value is types.InputFlowNodeConfiguration

	case *types.FlowNodeConfigurationMemberIterator:
		_ = v.Value // Value is types.IteratorFlowNodeConfiguration

	case *types.FlowNodeConfigurationMemberKnowledgeBase:
		_ = v.Value // Value is types.KnowledgeBaseFlowNodeConfiguration

	case *types.FlowNodeConfigurationMemberLambdaFunction:
		_ = v.Value // Value is types.LambdaFunctionFlowNodeConfiguration

	case *types.FlowNodeConfigurationMemberLex:
		_ = v.Value // Value is types.LexFlowNodeConfiguration

	case *types.FlowNodeConfigurationMemberOutput:
		_ = v.Value // Value is types.OutputFlowNodeConfiguration

	case *types.FlowNodeConfigurationMemberPrompt:
		_ = v.Value // Value is types.PromptFlowNodeConfiguration

	case *types.FlowNodeConfigurationMemberRetrieval:
		_ = v.Value // Value is types.RetrievalFlowNodeConfiguration

	case *types.FlowNodeConfigurationMemberStorage:
		_ = v.Value // Value is types.StorageFlowNodeConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.IteratorFlowNodeConfiguration
var _ *types.LambdaFunctionFlowNodeConfiguration
var _ *types.ConditionFlowNodeConfiguration
var _ *types.CollectorFlowNodeConfiguration
var _ *types.OutputFlowNodeConfiguration
var _ *types.AgentFlowNodeConfiguration
var _ *types.RetrievalFlowNodeConfiguration
var _ *types.StorageFlowNodeConfiguration
var _ *types.KnowledgeBaseFlowNodeConfiguration
var _ *types.LexFlowNodeConfiguration
var _ *types.PromptFlowNodeConfiguration
var _ *types.InputFlowNodeConfiguration

func ExampleFunctionSchema_outputUsage() {
	var union types.FunctionSchema
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.FunctionSchemaMemberFunctions:
		_ = v.Value // Value is []types.Function

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ []types.Function

func ExamplePromptFlowNodeSourceConfiguration_outputUsage() {
	var union types.PromptFlowNodeSourceConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PromptFlowNodeSourceConfigurationMemberInline:
		_ = v.Value // Value is types.PromptFlowNodeInlineConfiguration

	case *types.PromptFlowNodeSourceConfigurationMemberResource:
		_ = v.Value // Value is types.PromptFlowNodeResourceConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.PromptFlowNodeInlineConfiguration
var _ *types.PromptFlowNodeResourceConfiguration

func ExamplePromptInferenceConfiguration_outputUsage() {
	var union types.PromptInferenceConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PromptInferenceConfigurationMemberText:
		_ = v.Value // Value is types.PromptModelInferenceConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.PromptModelInferenceConfiguration

func ExamplePromptTemplateConfiguration_outputUsage() {
	var union types.PromptTemplateConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PromptTemplateConfigurationMemberText:
		_ = v.Value // Value is types.TextPromptTemplateConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.TextPromptTemplateConfiguration

func ExampleRetrievalFlowNodeServiceConfiguration_outputUsage() {
	var union types.RetrievalFlowNodeServiceConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.RetrievalFlowNodeServiceConfigurationMemberS3:
		_ = v.Value // Value is types.RetrievalFlowNodeS3Configuration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.RetrievalFlowNodeS3Configuration

func ExampleStorageFlowNodeServiceConfiguration_outputUsage() {
	var union types.StorageFlowNodeServiceConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.StorageFlowNodeServiceConfigurationMemberS3:
		_ = v.Value // Value is types.StorageFlowNodeS3Configuration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.StorageFlowNodeS3Configuration
