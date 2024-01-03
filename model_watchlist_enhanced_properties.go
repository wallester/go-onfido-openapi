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

// checks if the WatchlistEnhancedProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WatchlistEnhancedProperties{}

// WatchlistEnhancedProperties struct for WatchlistEnhancedProperties
type WatchlistEnhancedProperties struct {
	// Returns any matches including, but not limited to, name and date of birth of match, aliases and associates, and relevant events and sources.
	Records []string `json:"records,omitempty"`
}

// NewWatchlistEnhancedProperties instantiates a new WatchlistEnhancedProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistEnhancedProperties() *WatchlistEnhancedProperties {
	this := WatchlistEnhancedProperties{}
	return &this
}

// NewWatchlistEnhancedPropertiesWithDefaults instantiates a new WatchlistEnhancedProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistEnhancedPropertiesWithDefaults() *WatchlistEnhancedProperties {
	this := WatchlistEnhancedProperties{}
	return &this
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *WatchlistEnhancedProperties) GetRecords() []string {
	if o == nil || IsNil(o.Records) {
		var ret []string
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistEnhancedProperties) GetRecordsOk() ([]string, bool) {
	if o == nil || IsNil(o.Records) {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *WatchlistEnhancedProperties) HasRecords() bool {
	if o != nil && !IsNil(o.Records) {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []string and assigns it to the Records field.
func (o *WatchlistEnhancedProperties) SetRecords(v []string) {
	o.Records = v
}

func (o WatchlistEnhancedProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WatchlistEnhancedProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Records) {
		toSerialize["records"] = o.Records
	}
	return toSerialize, nil
}

type NullableWatchlistEnhancedProperties struct {
	value *WatchlistEnhancedProperties
	isSet bool
}

func (v NullableWatchlistEnhancedProperties) Get() *WatchlistEnhancedProperties {
	return v.value
}

func (v *NullableWatchlistEnhancedProperties) Set(val *WatchlistEnhancedProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistEnhancedProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistEnhancedProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistEnhancedProperties(val *WatchlistEnhancedProperties) *NullableWatchlistEnhancedProperties {
	return &NullableWatchlistEnhancedProperties{value: val, isSet: true}
}

func (v NullableWatchlistEnhancedProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistEnhancedProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
