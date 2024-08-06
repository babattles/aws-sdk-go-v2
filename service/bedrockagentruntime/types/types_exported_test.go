// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime/document"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime/types"
)

func ExampleFlowInputContent_outputUsage() {
	var union types.FlowInputContent
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.FlowInputContentMemberDocument:
		_ = v.Value // Value is document.Interface

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ document.Interface

func ExampleFlowOutputContent_outputUsage() {
	var union types.FlowOutputContent
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.FlowOutputContentMemberDocument:
		_ = v.Value // Value is document.Interface

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ document.Interface

func ExampleFlowResponseStream_outputUsage() {
	var union types.FlowResponseStream
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.FlowResponseStreamMemberFlowCompletionEvent:
		_ = v.Value // Value is types.FlowCompletionEvent

	case *types.FlowResponseStreamMemberFlowOutputEvent:
		_ = v.Value // Value is types.FlowOutputEvent

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.FlowOutputEvent
var _ *types.FlowCompletionEvent

func ExampleInvocationInputMember_outputUsage() {
	var union types.InvocationInputMember
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.InvocationInputMemberMemberApiInvocationInput:
		_ = v.Value // Value is types.ApiInvocationInput

	case *types.InvocationInputMemberMemberFunctionInvocationInput:
		_ = v.Value // Value is types.FunctionInvocationInput

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ApiInvocationInput
var _ *types.FunctionInvocationInput

func ExampleInvocationResultMember_outputUsage() {
	var union types.InvocationResultMember
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.InvocationResultMemberMemberApiResult:
		_ = v.Value // Value is types.ApiResult

	case *types.InvocationResultMemberMemberFunctionResult:
		_ = v.Value // Value is types.FunctionResult

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ApiResult
var _ *types.FunctionResult

func ExampleMemory_outputUsage() {
	var union types.Memory
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.MemoryMemberSessionSummary:
		_ = v.Value // Value is types.MemorySessionSummary

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.MemorySessionSummary

func ExampleOrchestrationTrace_outputUsage() {
	var union types.OrchestrationTrace
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.OrchestrationTraceMemberInvocationInput:
		_ = v.Value // Value is types.InvocationInput

	case *types.OrchestrationTraceMemberModelInvocationInput:
		_ = v.Value // Value is types.ModelInvocationInput

	case *types.OrchestrationTraceMemberModelInvocationOutput:
		_ = v.Value // Value is types.OrchestrationModelInvocationOutput

	case *types.OrchestrationTraceMemberObservation:
		_ = v.Value // Value is types.Observation

	case *types.OrchestrationTraceMemberRationale:
		_ = v.Value // Value is types.Rationale

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.InvocationInput
var _ *types.ModelInvocationInput
var _ *types.OrchestrationModelInvocationOutput
var _ *types.Rationale
var _ *types.Observation

func ExamplePostProcessingTrace_outputUsage() {
	var union types.PostProcessingTrace
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PostProcessingTraceMemberModelInvocationInput:
		_ = v.Value // Value is types.ModelInvocationInput

	case *types.PostProcessingTraceMemberModelInvocationOutput:
		_ = v.Value // Value is types.PostProcessingModelInvocationOutput

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ModelInvocationInput
var _ *types.PostProcessingModelInvocationOutput

func ExamplePreProcessingTrace_outputUsage() {
	var union types.PreProcessingTrace
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PreProcessingTraceMemberModelInvocationInput:
		_ = v.Value // Value is types.ModelInvocationInput

	case *types.PreProcessingTraceMemberModelInvocationOutput:
		_ = v.Value // Value is types.PreProcessingModelInvocationOutput

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ModelInvocationInput
var _ *types.PreProcessingModelInvocationOutput

func ExampleResponseStream_outputUsage() {
	var union types.ResponseStream
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ResponseStreamMemberChunk:
		_ = v.Value // Value is types.PayloadPart

	case *types.ResponseStreamMemberFiles:
		_ = v.Value // Value is types.FilePart

	case *types.ResponseStreamMemberReturnControl:
		_ = v.Value // Value is types.ReturnControlPayload

	case *types.ResponseStreamMemberTrace:
		_ = v.Value // Value is types.TracePart

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ReturnControlPayload
var _ *types.PayloadPart
var _ *types.TracePart
var _ *types.FilePart

func ExampleRetrievalFilter_outputUsage() {
	var union types.RetrievalFilter
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.RetrievalFilterMemberAndAll:
		_ = v.Value // Value is []types.RetrievalFilter

	case *types.RetrievalFilterMemberEquals:
		_ = v.Value // Value is types.FilterAttribute

	case *types.RetrievalFilterMemberGreaterThan:
		_ = v.Value // Value is types.FilterAttribute

	case *types.RetrievalFilterMemberGreaterThanOrEquals:
		_ = v.Value // Value is types.FilterAttribute

	case *types.RetrievalFilterMemberIn:
		_ = v.Value // Value is types.FilterAttribute

	case *types.RetrievalFilterMemberLessThan:
		_ = v.Value // Value is types.FilterAttribute

	case *types.RetrievalFilterMemberLessThanOrEquals:
		_ = v.Value // Value is types.FilterAttribute

	case *types.RetrievalFilterMemberListContains:
		_ = v.Value // Value is types.FilterAttribute

	case *types.RetrievalFilterMemberNotEquals:
		_ = v.Value // Value is types.FilterAttribute

	case *types.RetrievalFilterMemberNotIn:
		_ = v.Value // Value is types.FilterAttribute

	case *types.RetrievalFilterMemberOrAll:
		_ = v.Value // Value is []types.RetrievalFilter

	case *types.RetrievalFilterMemberStartsWith:
		_ = v.Value // Value is types.FilterAttribute

	case *types.RetrievalFilterMemberStringContains:
		_ = v.Value // Value is types.FilterAttribute

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ []types.RetrievalFilter
var _ *types.FilterAttribute

func ExampleTrace_outputUsage() {
	var union types.Trace
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.TraceMemberFailureTrace:
		_ = v.Value // Value is types.FailureTrace

	case *types.TraceMemberGuardrailTrace:
		_ = v.Value // Value is types.GuardrailTrace

	case *types.TraceMemberOrchestrationTrace:
		_ = v.Value // Value is types.OrchestrationTrace

	case *types.TraceMemberPostProcessingTrace:
		_ = v.Value // Value is types.PostProcessingTrace

	case *types.TraceMemberPreProcessingTrace:
		_ = v.Value // Value is types.PreProcessingTrace

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.FailureTrace
var _ *types.GuardrailTrace
var _ types.PreProcessingTrace
var _ types.PostProcessingTrace
var _ types.OrchestrationTrace
