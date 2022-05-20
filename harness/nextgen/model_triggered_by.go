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

type TriggeredBy struct {
	UnknownFields             *UnknownFieldSet       `json:"unknownFields,omitempty"`
	Identifier                string                 `json:"identifier,omitempty"`
	Initialized               bool                   `json:"initialized,omitempty"`
	UuidBytes                 *ByteString            `json:"uuidBytes,omitempty"`
	IdentifierBytes           *ByteString            `json:"identifierBytes,omitempty"`
	ExtraInfoCount            int32                  `json:"extraInfoCount,omitempty"`
	ExtraInfo                 map[string]string      `json:"extraInfo,omitempty"`
	ExtraInfoMap              map[string]string      `json:"extraInfoMap,omitempty"`
	ParserForType             *ParserTriggeredBy     `json:"parserForType,omitempty"`
	SerializedSize            int32                  `json:"serializedSize,omitempty"`
	DefaultInstanceForType    *TriggeredBy           `json:"defaultInstanceForType,omitempty"`
	Uuid                      string                 `json:"uuid,omitempty"`
	AllFields                 map[string]interface{} `json:"allFields,omitempty"`
	DescriptorForType         *Descriptor            `json:"descriptorForType,omitempty"`
	InitializationErrorString string                 `json:"initializationErrorString,omitempty"`
	MemoizedSerializedSize    int32                  `json:"memoizedSerializedSize,omitempty"`
}
