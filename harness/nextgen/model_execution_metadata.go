/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type ExecutionMetadata struct {
	UnknownFields             *UnknownFieldSet                 `json:"unknownFields,omitempty"`
	Initialized               bool                             `json:"initialized,omitempty"`
	PrincipalInfo             *ExecutionPrincipalInfo          `json:"principalInfo,omitempty"`
	RunSequence               int32                            `json:"runSequence,omitempty"`
	TriggerInfo               *ExecutionTriggerInfo            `json:"triggerInfo,omitempty"`
	TriggerInfoOrBuilder      *ExecutionTriggerInfoOrBuilder   `json:"triggerInfoOrBuilder,omitempty"`
	PipelineIdentifier        string                           `json:"pipelineIdentifier,omitempty"`
	PipelineIdentifierBytes   *ByteString                      `json:"pipelineIdentifierBytes,omitempty"`
	ExecutionUuid             string                           `json:"executionUuid,omitempty"`
	ExecutionUuidBytes        *ByteString                      `json:"executionUuidBytes,omitempty"`
	PrincipalInfoOrBuilder    *ExecutionPrincipalInfoOrBuilder `json:"principalInfoOrBuilder,omitempty"`
	GitSyncBranchContext      *ByteString                      `json:"gitSyncBranchContext,omitempty"`
	ModuleType                string                           `json:"moduleType,omitempty"`
	ModuleTypeBytes           *ByteString                      `json:"moduleTypeBytes,omitempty"`
	RetryInfo                 *RetryExecutionInfo              `json:"retryInfo,omitempty"`
	RetryInfoOrBuilder        *RetryExecutionInfoOrBuilder     `json:"retryInfoOrBuilder,omitempty"`
	IsNotificationConfigured  bool                             `json:"isNotificationConfigured,omitempty"`
	ParserForType             *ParserExecutionMetadata         `json:"parserForType,omitempty"`
	SerializedSize            int32                            `json:"serializedSize,omitempty"`
	DefaultInstanceForType    *ExecutionMetadata               `json:"defaultInstanceForType,omitempty"`
	AllFields                 map[string]interface{}           `json:"allFields,omitempty"`
	DescriptorForType         *Descriptor                      `json:"descriptorForType,omitempty"`
	InitializationErrorString string                           `json:"initializationErrorString,omitempty"`
	MemoizedSerializedSize    int32                            `json:"memoizedSerializedSize,omitempty"`
}
