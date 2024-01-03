/*
Onfido API v3.6

The Onfido API is used to submit check requests.

API version: 3.6.0
Contact: engineering@onfido.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onfido_openapi

import (
	"encoding/json"
)

// checks if the WorkflowRunError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowRunError{}

// WorkflowRunError Error object. Only set when the Workflow Run status is 'error'.
type WorkflowRunError struct {
	// The type of error.
	Type *string `json:"type,omitempty"`
	// A textual description of the error.
	Message *string `json:"message,omitempty"`
}

// NewWorkflowRunError instantiates a new WorkflowRunError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRunError() *WorkflowRunError {
	this := WorkflowRunError{}
	return &this
}

// NewWorkflowRunErrorWithDefaults instantiates a new WorkflowRunError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRunErrorWithDefaults() *WorkflowRunError {
	this := WorkflowRunError{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowRunError) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRunError) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowRunError) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowRunError) SetType(v string) {
	o.Type = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *WorkflowRunError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRunError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *WorkflowRunError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *WorkflowRunError) SetMessage(v string) {
	o.Message = &v
}

func (o WorkflowRunError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowRunError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableWorkflowRunError struct {
	value *WorkflowRunError
	isSet bool
}

func (v NullableWorkflowRunError) Get() *WorkflowRunError {
	return v.value
}

func (v *NullableWorkflowRunError) Set(val *WorkflowRunError) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRunError) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRunError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRunError(val *WorkflowRunError) *NullableWorkflowRunError {
	return &NullableWorkflowRunError{value: val, isSet: true}
}

func (v NullableWorkflowRunError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRunError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
