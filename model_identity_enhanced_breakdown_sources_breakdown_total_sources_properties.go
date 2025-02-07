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

// checks if the IdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties{}

// IdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties struct for IdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties
type IdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties struct {
	TotalNumberOfSources *int32 `json:"total_number_of_sources,omitempty"`
}

// NewIdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties instantiates a new IdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties() *IdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties {
	this := IdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties{}
	return &this
}

// NewIdentityEnhancedBreakdownSourcesBreakdownTotalSourcesPropertiesWithDefaults instantiates a new IdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityEnhancedBreakdownSourcesBreakdownTotalSourcesPropertiesWithDefaults() *IdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties {
	this := IdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties{}
	return &this
}

// GetTotalNumberOfSources returns the TotalNumberOfSources field value if set, zero value otherwise.
func (o *IdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties) GetTotalNumberOfSources() int32 {
	if o == nil || IsNil(o.TotalNumberOfSources) {
		var ret int32
		return ret
	}
	return *o.TotalNumberOfSources
}

// GetTotalNumberOfSourcesOk returns a tuple with the TotalNumberOfSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties) GetTotalNumberOfSourcesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalNumberOfSources) {
		return nil, false
	}
	return o.TotalNumberOfSources, true
}

// HasTotalNumberOfSources returns a boolean if a field has been set.
func (o *IdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties) HasTotalNumberOfSources() bool {
	if o != nil && !IsNil(o.TotalNumberOfSources) {
		return true
	}

	return false
}

// SetTotalNumberOfSources gets a reference to the given int32 and assigns it to the TotalNumberOfSources field.
func (o *IdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties) SetTotalNumberOfSources(v int32) {
	o.TotalNumberOfSources = &v
}

func (o IdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalNumberOfSources) {
		toSerialize["total_number_of_sources"] = o.TotalNumberOfSources
	}
	return toSerialize, nil
}

type NullableIdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties struct {
	value *IdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties
	isSet bool
}

func (v NullableIdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties) Get() *IdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties {
	return v.value
}

func (v *NullableIdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties) Set(val *IdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties(val *IdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties) *NullableIdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties {
	return &NullableIdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties{value: val, isSet: true}
}

func (v NullableIdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityEnhancedBreakdownSourcesBreakdownTotalSourcesProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
