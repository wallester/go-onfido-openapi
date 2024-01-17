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

// checks if the DocumentBreakdownImageIntegrity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentBreakdownImageIntegrity{}

// DocumentBreakdownImageIntegrity Asserts if the document is of sufficient quality to verify.
type DocumentBreakdownImageIntegrity struct {
	Result    *string                                   `json:"result,omitempty"`
	Breakdown *DocumentBreakdownImageIntegrityBreakdown `json:"breakdown,omitempty"`
}

// NewDocumentBreakdownImageIntegrity instantiates a new DocumentBreakdownImageIntegrity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentBreakdownImageIntegrity() *DocumentBreakdownImageIntegrity {
	this := DocumentBreakdownImageIntegrity{}
	return &this
}

// NewDocumentBreakdownImageIntegrityWithDefaults instantiates a new DocumentBreakdownImageIntegrity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentBreakdownImageIntegrityWithDefaults() *DocumentBreakdownImageIntegrity {
	this := DocumentBreakdownImageIntegrity{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DocumentBreakdownImageIntegrity) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentBreakdownImageIntegrity) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DocumentBreakdownImageIntegrity) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *DocumentBreakdownImageIntegrity) SetResult(v string) {
	o.Result = &v
}

// GetBreakdown returns the Breakdown field value if set, zero value otherwise.
func (o *DocumentBreakdownImageIntegrity) GetBreakdown() DocumentBreakdownImageIntegrityBreakdown {
	if o == nil || IsNil(o.Breakdown) {
		var ret DocumentBreakdownImageIntegrityBreakdown
		return ret
	}
	return *o.Breakdown
}

// GetBreakdownOk returns a tuple with the Breakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentBreakdownImageIntegrity) GetBreakdownOk() (*DocumentBreakdownImageIntegrityBreakdown, bool) {
	if o == nil || IsNil(o.Breakdown) {
		return nil, false
	}
	return o.Breakdown, true
}

// HasBreakdown returns a boolean if a field has been set.
func (o *DocumentBreakdownImageIntegrity) HasBreakdown() bool {
	if o != nil && !IsNil(o.Breakdown) {
		return true
	}

	return false
}

// SetBreakdown gets a reference to the given DocumentBreakdownImageIntegrityBreakdown and assigns it to the Breakdown field.
func (o *DocumentBreakdownImageIntegrity) SetBreakdown(v DocumentBreakdownImageIntegrityBreakdown) {
	o.Breakdown = &v
}

func (o DocumentBreakdownImageIntegrity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentBreakdownImageIntegrity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Breakdown) {
		toSerialize["breakdown"] = o.Breakdown
	}
	return toSerialize, nil
}

type NullableDocumentBreakdownImageIntegrity struct {
	value *DocumentBreakdownImageIntegrity
	isSet bool
}

func (v NullableDocumentBreakdownImageIntegrity) Get() *DocumentBreakdownImageIntegrity {
	return v.value
}

func (v *NullableDocumentBreakdownImageIntegrity) Set(val *DocumentBreakdownImageIntegrity) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentBreakdownImageIntegrity) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentBreakdownImageIntegrity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentBreakdownImageIntegrity(val *DocumentBreakdownImageIntegrity) *NullableDocumentBreakdownImageIntegrity {
	return &NullableDocumentBreakdownImageIntegrity{value: val, isSet: true}
}

func (v NullableDocumentBreakdownImageIntegrity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentBreakdownImageIntegrity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
