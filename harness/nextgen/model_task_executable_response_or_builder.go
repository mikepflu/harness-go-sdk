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

type TaskExecutableResponseOrBuilder struct {
	TaskId                    string                 `json:"taskId,omitempty"`
	TaskName                  string                 `json:"taskName,omitempty"`
	TaskCategoryValue         int32                  `json:"taskCategoryValue,omitempty"`
	LogKeysCount              int32                  `json:"logKeysCount,omitempty"`
	TaskCategory              string                 `json:"taskCategory,omitempty"`
	TaskIdBytes               *ByteString            `json:"taskIdBytes,omitempty"`
	LogKeysList               []string               `json:"logKeysList,omitempty"`
	TaskNameBytes             *ByteString            `json:"taskNameBytes,omitempty"`
	UnitsList                 []string               `json:"unitsList,omitempty"`
	UnitsCount                int32                  `json:"unitsCount,omitempty"`
	AllFields                 map[string]interface{} `json:"allFields,omitempty"`
	DescriptorForType         *Descriptor            `json:"descriptorForType,omitempty"`
	UnknownFields             *UnknownFieldSet       `json:"unknownFields,omitempty"`
	InitializationErrorString string                 `json:"initializationErrorString,omitempty"`
	DefaultInstanceForType    *Message               `json:"defaultInstanceForType,omitempty"`
	Initialized               bool                   `json:"initialized,omitempty"`
}
