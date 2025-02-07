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

// checks if the RightToWorkBreakdownRightToWork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RightToWorkBreakdownRightToWork{}

// RightToWorkBreakdownRightToWork Asserts whether the applicant has the right to work.
type RightToWorkBreakdownRightToWork struct {
	Result    *string                                   `json:"result,omitempty"`
	Breakdown *RightToWorkBreakdownRightToWorkBreakdown `json:"breakdown,omitempty"`
}

// NewRightToWorkBreakdownRightToWork instantiates a new RightToWorkBreakdownRightToWork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRightToWorkBreakdownRightToWork() *RightToWorkBreakdownRightToWork {
	this := RightToWorkBreakdownRightToWork{}
	return &this
}

// NewRightToWorkBreakdownRightToWorkWithDefaults instantiates a new RightToWorkBreakdownRightToWork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRightToWorkBreakdownRightToWorkWithDefaults() *RightToWorkBreakdownRightToWork {
	this := RightToWorkBreakdownRightToWork{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *RightToWorkBreakdownRightToWork) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RightToWorkBreakdownRightToWork) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *RightToWorkBreakdownRightToWork) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *RightToWorkBreakdownRightToWork) SetResult(v string) {
	o.Result = &v
}

// GetBreakdown returns the Breakdown field value if set, zero value otherwise.
func (o *RightToWorkBreakdownRightToWork) GetBreakdown() RightToWorkBreakdownRightToWorkBreakdown {
	if o == nil || IsNil(o.Breakdown) {
		var ret RightToWorkBreakdownRightToWorkBreakdown
		return ret
	}
	return *o.Breakdown
}

// GetBreakdownOk returns a tuple with the Breakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RightToWorkBreakdownRightToWork) GetBreakdownOk() (*RightToWorkBreakdownRightToWorkBreakdown, bool) {
	if o == nil || IsNil(o.Breakdown) {
		return nil, false
	}
	return o.Breakdown, true
}

// HasBreakdown returns a boolean if a field has been set.
func (o *RightToWorkBreakdownRightToWork) HasBreakdown() bool {
	if o != nil && !IsNil(o.Breakdown) {
		return true
	}

	return false
}

// SetBreakdown gets a reference to the given RightToWorkBreakdownRightToWorkBreakdown and assigns it to the Breakdown field.
func (o *RightToWorkBreakdownRightToWork) SetBreakdown(v RightToWorkBreakdownRightToWorkBreakdown) {
	o.Breakdown = &v
}

func (o RightToWorkBreakdownRightToWork) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RightToWorkBreakdownRightToWork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Breakdown) {
		toSerialize["breakdown"] = o.Breakdown
	}
	return toSerialize, nil
}

type NullableRightToWorkBreakdownRightToWork struct {
	value *RightToWorkBreakdownRightToWork
	isSet bool
}

func (v NullableRightToWorkBreakdownRightToWork) Get() *RightToWorkBreakdownRightToWork {
	return v.value
}

func (v *NullableRightToWorkBreakdownRightToWork) Set(val *RightToWorkBreakdownRightToWork) {
	v.value = val
	v.isSet = true
}

func (v NullableRightToWorkBreakdownRightToWork) IsSet() bool {
	return v.isSet
}

func (v *NullableRightToWorkBreakdownRightToWork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRightToWorkBreakdownRightToWork(val *RightToWorkBreakdownRightToWork) *NullableRightToWorkBreakdownRightToWork {
	return &NullableRightToWorkBreakdownRightToWork{value: val, isSet: true}
}

func (v NullableRightToWorkBreakdownRightToWork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRightToWorkBreakdownRightToWork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
