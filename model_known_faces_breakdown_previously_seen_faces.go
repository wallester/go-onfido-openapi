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

// checks if the KnownFacesBreakdownPreviouslySeenFaces type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KnownFacesBreakdownPreviouslySeenFaces{}

// KnownFacesBreakdownPreviouslySeenFaces Asserts whether the applicant's most recent facial media (live photo or live video) matches any other live photos or live videos already in your Onfido account database.
type KnownFacesBreakdownPreviouslySeenFaces struct {
	Result *string `json:"result,omitempty"`
}

// NewKnownFacesBreakdownPreviouslySeenFaces instantiates a new KnownFacesBreakdownPreviouslySeenFaces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKnownFacesBreakdownPreviouslySeenFaces() *KnownFacesBreakdownPreviouslySeenFaces {
	this := KnownFacesBreakdownPreviouslySeenFaces{}
	return &this
}

// NewKnownFacesBreakdownPreviouslySeenFacesWithDefaults instantiates a new KnownFacesBreakdownPreviouslySeenFaces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKnownFacesBreakdownPreviouslySeenFacesWithDefaults() *KnownFacesBreakdownPreviouslySeenFaces {
	this := KnownFacesBreakdownPreviouslySeenFaces{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *KnownFacesBreakdownPreviouslySeenFaces) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KnownFacesBreakdownPreviouslySeenFaces) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *KnownFacesBreakdownPreviouslySeenFaces) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *KnownFacesBreakdownPreviouslySeenFaces) SetResult(v string) {
	o.Result = &v
}

func (o KnownFacesBreakdownPreviouslySeenFaces) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KnownFacesBreakdownPreviouslySeenFaces) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableKnownFacesBreakdownPreviouslySeenFaces struct {
	value *KnownFacesBreakdownPreviouslySeenFaces
	isSet bool
}

func (v NullableKnownFacesBreakdownPreviouslySeenFaces) Get() *KnownFacesBreakdownPreviouslySeenFaces {
	return v.value
}

func (v *NullableKnownFacesBreakdownPreviouslySeenFaces) Set(val *KnownFacesBreakdownPreviouslySeenFaces) {
	v.value = val
	v.isSet = true
}

func (v NullableKnownFacesBreakdownPreviouslySeenFaces) IsSet() bool {
	return v.isSet
}

func (v *NullableKnownFacesBreakdownPreviouslySeenFaces) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKnownFacesBreakdownPreviouslySeenFaces(val *KnownFacesBreakdownPreviouslySeenFaces) *NullableKnownFacesBreakdownPreviouslySeenFaces {
	return &NullableKnownFacesBreakdownPreviouslySeenFaces{value: val, isSet: true}
}

func (v NullableKnownFacesBreakdownPreviouslySeenFaces) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKnownFacesBreakdownPreviouslySeenFaces) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
