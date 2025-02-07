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

// checks if the IdentityEnhancedBreakdownAddressBreakdownCreditAgencies type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityEnhancedBreakdownAddressBreakdownCreditAgencies{}

// IdentityEnhancedBreakdownAddressBreakdownCreditAgencies The number of address matches against credit agencies.
type IdentityEnhancedBreakdownAddressBreakdownCreditAgencies struct {
	Result     *string                                                            `json:"result,omitempty"`
	Properties *IdentityEnhancedBreakdownAddressBreakdownCreditAgenciesProperties `json:"properties,omitempty"`
}

// NewIdentityEnhancedBreakdownAddressBreakdownCreditAgencies instantiates a new IdentityEnhancedBreakdownAddressBreakdownCreditAgencies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityEnhancedBreakdownAddressBreakdownCreditAgencies() *IdentityEnhancedBreakdownAddressBreakdownCreditAgencies {
	this := IdentityEnhancedBreakdownAddressBreakdownCreditAgencies{}
	return &this
}

// NewIdentityEnhancedBreakdownAddressBreakdownCreditAgenciesWithDefaults instantiates a new IdentityEnhancedBreakdownAddressBreakdownCreditAgencies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityEnhancedBreakdownAddressBreakdownCreditAgenciesWithDefaults() *IdentityEnhancedBreakdownAddressBreakdownCreditAgencies {
	this := IdentityEnhancedBreakdownAddressBreakdownCreditAgencies{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *IdentityEnhancedBreakdownAddressBreakdownCreditAgencies) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityEnhancedBreakdownAddressBreakdownCreditAgencies) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *IdentityEnhancedBreakdownAddressBreakdownCreditAgencies) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *IdentityEnhancedBreakdownAddressBreakdownCreditAgencies) SetResult(v string) {
	o.Result = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *IdentityEnhancedBreakdownAddressBreakdownCreditAgencies) GetProperties() IdentityEnhancedBreakdownAddressBreakdownCreditAgenciesProperties {
	if o == nil || IsNil(o.Properties) {
		var ret IdentityEnhancedBreakdownAddressBreakdownCreditAgenciesProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityEnhancedBreakdownAddressBreakdownCreditAgencies) GetPropertiesOk() (*IdentityEnhancedBreakdownAddressBreakdownCreditAgenciesProperties, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *IdentityEnhancedBreakdownAddressBreakdownCreditAgencies) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given IdentityEnhancedBreakdownAddressBreakdownCreditAgenciesProperties and assigns it to the Properties field.
func (o *IdentityEnhancedBreakdownAddressBreakdownCreditAgencies) SetProperties(v IdentityEnhancedBreakdownAddressBreakdownCreditAgenciesProperties) {
	o.Properties = &v
}

func (o IdentityEnhancedBreakdownAddressBreakdownCreditAgencies) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityEnhancedBreakdownAddressBreakdownCreditAgencies) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableIdentityEnhancedBreakdownAddressBreakdownCreditAgencies struct {
	value *IdentityEnhancedBreakdownAddressBreakdownCreditAgencies
	isSet bool
}

func (v NullableIdentityEnhancedBreakdownAddressBreakdownCreditAgencies) Get() *IdentityEnhancedBreakdownAddressBreakdownCreditAgencies {
	return v.value
}

func (v *NullableIdentityEnhancedBreakdownAddressBreakdownCreditAgencies) Set(val *IdentityEnhancedBreakdownAddressBreakdownCreditAgencies) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityEnhancedBreakdownAddressBreakdownCreditAgencies) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityEnhancedBreakdownAddressBreakdownCreditAgencies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityEnhancedBreakdownAddressBreakdownCreditAgencies(val *IdentityEnhancedBreakdownAddressBreakdownCreditAgencies) *NullableIdentityEnhancedBreakdownAddressBreakdownCreditAgencies {
	return &NullableIdentityEnhancedBreakdownAddressBreakdownCreditAgencies{value: val, isSet: true}
}

func (v NullableIdentityEnhancedBreakdownAddressBreakdownCreditAgencies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityEnhancedBreakdownAddressBreakdownCreditAgencies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
