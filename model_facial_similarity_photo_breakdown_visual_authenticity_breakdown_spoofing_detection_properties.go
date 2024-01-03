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

// checks if the FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties{}

// FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties struct for FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties
type FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties struct {
	// A floating point number between 0 and 1. The closer the score is to 0, the more likely it is to be a spoof.
	Score *float32 `json:"score,omitempty"`
}

// NewFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties instantiates a new FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties() *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties {
	this := FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties{}
	return &this
}

// NewFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionPropertiesWithDefaults instantiates a new FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionPropertiesWithDefaults() *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties {
	this := FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties) GetScore() float32 {
	if o == nil || IsNil(o.Score) {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties) GetScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties) SetScore(v float32) {
	o.Score = &v
}

func (o FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	return toSerialize, nil
}

type NullableFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties struct {
	value *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties
	isSet bool
}

func (v NullableFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties) Get() *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties {
	return v.value
}

func (v *NullableFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties) Set(val *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties(val *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties) *NullableFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties {
	return &NullableFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties{value: val, isSet: true}
}

func (v NullableFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
