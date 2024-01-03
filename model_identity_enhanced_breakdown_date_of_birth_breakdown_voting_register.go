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

// checks if the IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister{}

// IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister Date of birth match against voting register.
type IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister struct {
	Result     *string                `json:"result,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// NewIdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister instantiates a new IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister() *IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister {
	this := IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister{}
	return &this
}

// NewIdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegisterWithDefaults instantiates a new IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegisterWithDefaults() *IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister {
	this := IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister) SetResult(v string) {
	o.Result = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister) GetProperties() map[string]interface{} {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Properties) {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

func (o IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableIdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister struct {
	value *IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister
	isSet bool
}

func (v NullableIdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister) Get() *IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister {
	return v.value
}

func (v *NullableIdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister) Set(val *IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister(val *IdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister) *NullableIdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister {
	return &NullableIdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister{value: val, isSet: true}
}

func (v NullableIdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityEnhancedBreakdownDateOfBirthBreakdownVotingRegister) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
