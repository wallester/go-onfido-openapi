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

// checks if the FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection{}

// FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection Asserts whether the live video is not a spoof (such as videos of digital screens).
type FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection struct {
	Result     *string                                                                               `json:"result,omitempty"`
	Properties *FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties `json:"properties,omitempty"`
}

// NewFacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection instantiates a new FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection() *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection {
	this := FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection{}
	return &this
}

// NewFacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetectionWithDefaults instantiates a new FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetectionWithDefaults() *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection {
	this := FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection) SetResult(v string) {
	o.Result = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection) GetProperties() FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties {
	if o == nil || IsNil(o.Properties) {
		var ret FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection) GetPropertiesOk() (*FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties and assigns it to the Properties field.
func (o *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection) SetProperties(v FacialSimilarityPhotoBreakdownVisualAuthenticityBreakdownSpoofingDetectionProperties) {
	o.Properties = &v
}

func (o FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableFacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection struct {
	value *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection
	isSet bool
}

func (v NullableFacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection) Get() *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection {
	return v.value
}

func (v *NullableFacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection) Set(val *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection) {
	v.value = val
	v.isSet = true
}

func (v NullableFacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection) IsSet() bool {
	return v.isSet
}

func (v *NullableFacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection(val *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection) *NullableFacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection {
	return &NullableFacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection{value: val, isSet: true}
}

func (v NullableFacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
