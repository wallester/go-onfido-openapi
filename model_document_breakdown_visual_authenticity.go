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

// checks if the DocumentBreakdownVisualAuthenticity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentBreakdownVisualAuthenticity{}

// DocumentBreakdownVisualAuthenticity Asserts whether visual, non-textual, elements are correct given the type of document.
type DocumentBreakdownVisualAuthenticity struct {
	Result    *string                                       `json:"result,omitempty"`
	Breakdown *DocumentBreakdownVisualAuthenticityBreakdown `json:"breakdown,omitempty"`
}

// NewDocumentBreakdownVisualAuthenticity instantiates a new DocumentBreakdownVisualAuthenticity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentBreakdownVisualAuthenticity() *DocumentBreakdownVisualAuthenticity {
	this := DocumentBreakdownVisualAuthenticity{}
	return &this
}

// NewDocumentBreakdownVisualAuthenticityWithDefaults instantiates a new DocumentBreakdownVisualAuthenticity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentBreakdownVisualAuthenticityWithDefaults() *DocumentBreakdownVisualAuthenticity {
	this := DocumentBreakdownVisualAuthenticity{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DocumentBreakdownVisualAuthenticity) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentBreakdownVisualAuthenticity) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DocumentBreakdownVisualAuthenticity) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *DocumentBreakdownVisualAuthenticity) SetResult(v string) {
	o.Result = &v
}

// GetBreakdown returns the Breakdown field value if set, zero value otherwise.
func (o *DocumentBreakdownVisualAuthenticity) GetBreakdown() DocumentBreakdownVisualAuthenticityBreakdown {
	if o == nil || IsNil(o.Breakdown) {
		var ret DocumentBreakdownVisualAuthenticityBreakdown
		return ret
	}
	return *o.Breakdown
}

// GetBreakdownOk returns a tuple with the Breakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentBreakdownVisualAuthenticity) GetBreakdownOk() (*DocumentBreakdownVisualAuthenticityBreakdown, bool) {
	if o == nil || IsNil(o.Breakdown) {
		return nil, false
	}
	return o.Breakdown, true
}

// HasBreakdown returns a boolean if a field has been set.
func (o *DocumentBreakdownVisualAuthenticity) HasBreakdown() bool {
	if o != nil && !IsNil(o.Breakdown) {
		return true
	}

	return false
}

// SetBreakdown gets a reference to the given DocumentBreakdownVisualAuthenticityBreakdown and assigns it to the Breakdown field.
func (o *DocumentBreakdownVisualAuthenticity) SetBreakdown(v DocumentBreakdownVisualAuthenticityBreakdown) {
	o.Breakdown = &v
}

func (o DocumentBreakdownVisualAuthenticity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentBreakdownVisualAuthenticity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Breakdown) {
		toSerialize["breakdown"] = o.Breakdown
	}
	return toSerialize, nil
}

type NullableDocumentBreakdownVisualAuthenticity struct {
	value *DocumentBreakdownVisualAuthenticity
	isSet bool
}

func (v NullableDocumentBreakdownVisualAuthenticity) Get() *DocumentBreakdownVisualAuthenticity {
	return v.value
}

func (v *NullableDocumentBreakdownVisualAuthenticity) Set(val *DocumentBreakdownVisualAuthenticity) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentBreakdownVisualAuthenticity) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentBreakdownVisualAuthenticity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentBreakdownVisualAuthenticity(val *DocumentBreakdownVisualAuthenticity) *NullableDocumentBreakdownVisualAuthenticity {
	return &NullableDocumentBreakdownVisualAuthenticity{value: val, isSet: true}
}

func (v NullableDocumentBreakdownVisualAuthenticity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentBreakdownVisualAuthenticity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
