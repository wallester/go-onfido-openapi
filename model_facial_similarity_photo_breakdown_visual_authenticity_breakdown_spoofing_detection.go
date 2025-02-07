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

// checks if the FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection{}

// FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection Contains a score value under the properties bag.
type FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection struct {
	Result     *string                                                                               `json:"result,omitempty"`
	Properties *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties `json:"properties,omitempty"`
}

// NewFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection instantiates a new FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection() *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection {
	this := FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection{}
	return &this
}

// NewFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionWithDefaults instantiates a new FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionWithDefaults() *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection {
	this := FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection) SetResult(v string) {
	o.Result = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection) GetProperties() FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties {
	if o == nil || IsNil(o.Properties) {
		var ret FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection) GetPropertiesOk() (*FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties and assigns it to the Properties field.
func (o *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection) SetProperties(v FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties) {
	o.Properties = &v
}

func (o FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection struct {
	value *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection
	isSet bool
}

func (v NullableFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection) Get() *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection {
	return v.value
}

func (v *NullableFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection) Set(val *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection) {
	v.value = val
	v.isSet = true
}

func (v NullableFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection) IsSet() bool {
	return v.isSet
}

func (v *NullableFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection(val *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection) *NullableFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection {
	return &NullableFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection{value: val, isSet: true}
}

func (v NullableFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
