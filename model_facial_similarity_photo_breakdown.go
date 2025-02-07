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

// checks if the FacialSimilarityPhotoBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FacialSimilarityPhotoBreakdown{}

// FacialSimilarityPhotoBreakdown struct for FacialSimilarityPhotoBreakdown
type FacialSimilarityPhotoBreakdown struct {
	FaceComparison     *FacialSimilarityPhotoBreakdownFaceComparison     `json:"face_comparison,omitempty"`
	ImageIntegrity     *FacialSimilarityPhotoBreakdownImageIntegrity     `json:"image_integrity,omitempty"`
	VisualAuthenticity *FacialSimilarityPhotoBreakdownVisualAuthenticity `json:"visual_authenticity,omitempty"`
}

// NewFacialSimilarityPhotoBreakdown instantiates a new FacialSimilarityPhotoBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFacialSimilarityPhotoBreakdown() *FacialSimilarityPhotoBreakdown {
	this := FacialSimilarityPhotoBreakdown{}
	return &this
}

// NewFacialSimilarityPhotoBreakdownWithDefaults instantiates a new FacialSimilarityPhotoBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFacialSimilarityPhotoBreakdownWithDefaults() *FacialSimilarityPhotoBreakdown {
	this := FacialSimilarityPhotoBreakdown{}
	return &this
}

// GetFaceComparison returns the FaceComparison field value if set, zero value otherwise.
func (o *FacialSimilarityPhotoBreakdown) GetFaceComparison() FacialSimilarityPhotoBreakdownFaceComparison {
	if o == nil || IsNil(o.FaceComparison) {
		var ret FacialSimilarityPhotoBreakdownFaceComparison
		return ret
	}
	return *o.FaceComparison
}

// GetFaceComparisonOk returns a tuple with the FaceComparison field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacialSimilarityPhotoBreakdown) GetFaceComparisonOk() (*FacialSimilarityPhotoBreakdownFaceComparison, bool) {
	if o == nil || IsNil(o.FaceComparison) {
		return nil, false
	}
	return o.FaceComparison, true
}

// HasFaceComparison returns a boolean if a field has been set.
func (o *FacialSimilarityPhotoBreakdown) HasFaceComparison() bool {
	if o != nil && !IsNil(o.FaceComparison) {
		return true
	}

	return false
}

// SetFaceComparison gets a reference to the given FacialSimilarityPhotoBreakdownFaceComparison and assigns it to the FaceComparison field.
func (o *FacialSimilarityPhotoBreakdown) SetFaceComparison(v FacialSimilarityPhotoBreakdownFaceComparison) {
	o.FaceComparison = &v
}

// GetImageIntegrity returns the ImageIntegrity field value if set, zero value otherwise.
func (o *FacialSimilarityPhotoBreakdown) GetImageIntegrity() FacialSimilarityPhotoBreakdownImageIntegrity {
	if o == nil || IsNil(o.ImageIntegrity) {
		var ret FacialSimilarityPhotoBreakdownImageIntegrity
		return ret
	}
	return *o.ImageIntegrity
}

// GetImageIntegrityOk returns a tuple with the ImageIntegrity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacialSimilarityPhotoBreakdown) GetImageIntegrityOk() (*FacialSimilarityPhotoBreakdownImageIntegrity, bool) {
	if o == nil || IsNil(o.ImageIntegrity) {
		return nil, false
	}
	return o.ImageIntegrity, true
}

// HasImageIntegrity returns a boolean if a field has been set.
func (o *FacialSimilarityPhotoBreakdown) HasImageIntegrity() bool {
	if o != nil && !IsNil(o.ImageIntegrity) {
		return true
	}

	return false
}

// SetImageIntegrity gets a reference to the given FacialSimilarityPhotoBreakdownImageIntegrity and assigns it to the ImageIntegrity field.
func (o *FacialSimilarityPhotoBreakdown) SetImageIntegrity(v FacialSimilarityPhotoBreakdownImageIntegrity) {
	o.ImageIntegrity = &v
}

// GetVisualAuthenticity returns the VisualAuthenticity field value if set, zero value otherwise.
func (o *FacialSimilarityPhotoBreakdown) GetVisualAuthenticity() FacialSimilarityPhotoBreakdownVisualAuthenticity {
	if o == nil || IsNil(o.VisualAuthenticity) {
		var ret FacialSimilarityPhotoBreakdownVisualAuthenticity
		return ret
	}
	return *o.VisualAuthenticity
}

// GetVisualAuthenticityOk returns a tuple with the VisualAuthenticity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacialSimilarityPhotoBreakdown) GetVisualAuthenticityOk() (*FacialSimilarityPhotoBreakdownVisualAuthenticity, bool) {
	if o == nil || IsNil(o.VisualAuthenticity) {
		return nil, false
	}
	return o.VisualAuthenticity, true
}

// HasVisualAuthenticity returns a boolean if a field has been set.
func (o *FacialSimilarityPhotoBreakdown) HasVisualAuthenticity() bool {
	if o != nil && !IsNil(o.VisualAuthenticity) {
		return true
	}

	return false
}

// SetVisualAuthenticity gets a reference to the given FacialSimilarityPhotoBreakdownVisualAuthenticity and assigns it to the VisualAuthenticity field.
func (o *FacialSimilarityPhotoBreakdown) SetVisualAuthenticity(v FacialSimilarityPhotoBreakdownVisualAuthenticity) {
	o.VisualAuthenticity = &v
}

func (o FacialSimilarityPhotoBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FacialSimilarityPhotoBreakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FaceComparison) {
		toSerialize["face_comparison"] = o.FaceComparison
	}
	if !IsNil(o.ImageIntegrity) {
		toSerialize["image_integrity"] = o.ImageIntegrity
	}
	if !IsNil(o.VisualAuthenticity) {
		toSerialize["visual_authenticity"] = o.VisualAuthenticity
	}
	return toSerialize, nil
}

type NullableFacialSimilarityPhotoBreakdown struct {
	value *FacialSimilarityPhotoBreakdown
	isSet bool
}

func (v NullableFacialSimilarityPhotoBreakdown) Get() *FacialSimilarityPhotoBreakdown {
	return v.value
}

func (v *NullableFacialSimilarityPhotoBreakdown) Set(val *FacialSimilarityPhotoBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableFacialSimilarityPhotoBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableFacialSimilarityPhotoBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacialSimilarityPhotoBreakdown(val *FacialSimilarityPhotoBreakdown) *NullableFacialSimilarityPhotoBreakdown {
	return &NullableFacialSimilarityPhotoBreakdown{value: val, isSet: true}
}

func (v NullableFacialSimilarityPhotoBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacialSimilarityPhotoBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
