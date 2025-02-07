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

// checks if the RightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties{}

// RightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties struct for RightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties
type RightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties struct {
	// The document ID to retrieve the GOV.UK evidence document of the applicant's right to work.
	DocumentId *string `json:"document_id,omitempty"`
}

// NewRightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties instantiates a new RightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties() *RightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties {
	this := RightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties{}
	return &this
}

// NewRightToWorkBreakdownShareCodeValidationBreakdownDocumentIdPropertiesWithDefaults instantiates a new RightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRightToWorkBreakdownShareCodeValidationBreakdownDocumentIdPropertiesWithDefaults() *RightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties {
	this := RightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties{}
	return &this
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *RightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties) GetDocumentId() string {
	if o == nil || IsNil(o.DocumentId) {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties) GetDocumentIdOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentId) {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *RightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties) HasDocumentId() bool {
	if o != nil && !IsNil(o.DocumentId) {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *RightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties) SetDocumentId(v string) {
	o.DocumentId = &v
}

func (o RightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DocumentId) {
		toSerialize["document_id"] = o.DocumentId
	}
	return toSerialize, nil
}

type NullableRightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties struct {
	value *RightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties
	isSet bool
}

func (v NullableRightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties) Get() *RightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties {
	return v.value
}

func (v *NullableRightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties) Set(val *RightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableRightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableRightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties(val *RightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties) *NullableRightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties {
	return &NullableRightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties{value: val, isSet: true}
}

func (v NullableRightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRightToWorkBreakdownShareCodeValidationBreakdownDocumentIdProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
