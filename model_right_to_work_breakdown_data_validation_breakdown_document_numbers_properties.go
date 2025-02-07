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

// checks if the RightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties{}

// RightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties struct for RightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties
type RightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties struct {
	// Flags if the document number is not the expected length and format for document.
	DocumentNumber *string `json:"document_number,omitempty"`
}

// NewRightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties instantiates a new RightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties() *RightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties {
	this := RightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties{}
	return &this
}

// NewRightToWorkBreakdownDataValidationBreakdownDocumentNumbersPropertiesWithDefaults instantiates a new RightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRightToWorkBreakdownDataValidationBreakdownDocumentNumbersPropertiesWithDefaults() *RightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties {
	this := RightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties{}
	return &this
}

// GetDocumentNumber returns the DocumentNumber field value if set, zero value otherwise.
func (o *RightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties) GetDocumentNumber() string {
	if o == nil || IsNil(o.DocumentNumber) {
		var ret string
		return ret
	}
	return *o.DocumentNumber
}

// GetDocumentNumberOk returns a tuple with the DocumentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties) GetDocumentNumberOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentNumber) {
		return nil, false
	}
	return o.DocumentNumber, true
}

// HasDocumentNumber returns a boolean if a field has been set.
func (o *RightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties) HasDocumentNumber() bool {
	if o != nil && !IsNil(o.DocumentNumber) {
		return true
	}

	return false
}

// SetDocumentNumber gets a reference to the given string and assigns it to the DocumentNumber field.
func (o *RightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties) SetDocumentNumber(v string) {
	o.DocumentNumber = &v
}

func (o RightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DocumentNumber) {
		toSerialize["document_number"] = o.DocumentNumber
	}
	return toSerialize, nil
}

type NullableRightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties struct {
	value *RightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties
	isSet bool
}

func (v NullableRightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties) Get() *RightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties {
	return v.value
}

func (v *NullableRightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties) Set(val *RightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableRightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableRightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties(val *RightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties) *NullableRightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties {
	return &NullableRightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties{value: val, isSet: true}
}

func (v NullableRightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRightToWorkBreakdownDataValidationBreakdownDocumentNumbersProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
