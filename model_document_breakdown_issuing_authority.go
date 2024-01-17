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

// checks if the DocumentBreakdownIssuingAuthority type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentBreakdownIssuingAuthority{}

// DocumentBreakdownIssuingAuthority Asserts whether data on the document matches the issuing authority data.
type DocumentBreakdownIssuingAuthority struct {
	Result    *string                                     `json:"result,omitempty"`
	Breakdown *DocumentBreakdownIssuingAuthorityBreakdown `json:"breakdown,omitempty"`
}

// NewDocumentBreakdownIssuingAuthority instantiates a new DocumentBreakdownIssuingAuthority object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentBreakdownIssuingAuthority() *DocumentBreakdownIssuingAuthority {
	this := DocumentBreakdownIssuingAuthority{}
	return &this
}

// NewDocumentBreakdownIssuingAuthorityWithDefaults instantiates a new DocumentBreakdownIssuingAuthority object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentBreakdownIssuingAuthorityWithDefaults() *DocumentBreakdownIssuingAuthority {
	this := DocumentBreakdownIssuingAuthority{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DocumentBreakdownIssuingAuthority) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentBreakdownIssuingAuthority) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DocumentBreakdownIssuingAuthority) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *DocumentBreakdownIssuingAuthority) SetResult(v string) {
	o.Result = &v
}

// GetBreakdown returns the Breakdown field value if set, zero value otherwise.
func (o *DocumentBreakdownIssuingAuthority) GetBreakdown() DocumentBreakdownIssuingAuthorityBreakdown {
	if o == nil || IsNil(o.Breakdown) {
		var ret DocumentBreakdownIssuingAuthorityBreakdown
		return ret
	}
	return *o.Breakdown
}

// GetBreakdownOk returns a tuple with the Breakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentBreakdownIssuingAuthority) GetBreakdownOk() (*DocumentBreakdownIssuingAuthorityBreakdown, bool) {
	if o == nil || IsNil(o.Breakdown) {
		return nil, false
	}
	return o.Breakdown, true
}

// HasBreakdown returns a boolean if a field has been set.
func (o *DocumentBreakdownIssuingAuthority) HasBreakdown() bool {
	if o != nil && !IsNil(o.Breakdown) {
		return true
	}

	return false
}

// SetBreakdown gets a reference to the given DocumentBreakdownIssuingAuthorityBreakdown and assigns it to the Breakdown field.
func (o *DocumentBreakdownIssuingAuthority) SetBreakdown(v DocumentBreakdownIssuingAuthorityBreakdown) {
	o.Breakdown = &v
}

func (o DocumentBreakdownIssuingAuthority) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentBreakdownIssuingAuthority) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Breakdown) {
		toSerialize["breakdown"] = o.Breakdown
	}
	return toSerialize, nil
}

type NullableDocumentBreakdownIssuingAuthority struct {
	value *DocumentBreakdownIssuingAuthority
	isSet bool
}

func (v NullableDocumentBreakdownIssuingAuthority) Get() *DocumentBreakdownIssuingAuthority {
	return v.value
}

func (v *NullableDocumentBreakdownIssuingAuthority) Set(val *DocumentBreakdownIssuingAuthority) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentBreakdownIssuingAuthority) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentBreakdownIssuingAuthority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentBreakdownIssuingAuthority(val *DocumentBreakdownIssuingAuthority) *NullableDocumentBreakdownIssuingAuthority {
	return &NullableDocumentBreakdownIssuingAuthority{value: val, isSet: true}
}

func (v NullableDocumentBreakdownIssuingAuthority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentBreakdownIssuingAuthority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
