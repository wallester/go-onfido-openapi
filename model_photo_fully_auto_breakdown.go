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

// checks if the PhotoFullyAutoBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PhotoFullyAutoBreakdown{}

// PhotoFullyAutoBreakdown struct for PhotoFullyAutoBreakdown
type PhotoFullyAutoBreakdown struct {
	FaceComparison     *FacialSimilarityPhotoBreakdownFaceComparison     `json:"face_comparison,omitempty"`
	ImageIntegrity     *PhotoFullyAutoBreakdownImageIntegrity            `json:"image_integrity,omitempty"`
	VisualAuthenticity *FacialSimilarityPhotoBreakdownVisualAuthenticity `json:"visual_authenticity,omitempty"`
}

// NewPhotoFullyAutoBreakdown instantiates a new PhotoFullyAutoBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhotoFullyAutoBreakdown() *PhotoFullyAutoBreakdown {
	this := PhotoFullyAutoBreakdown{}
	return &this
}

// NewPhotoFullyAutoBreakdownWithDefaults instantiates a new PhotoFullyAutoBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhotoFullyAutoBreakdownWithDefaults() *PhotoFullyAutoBreakdown {
	this := PhotoFullyAutoBreakdown{}
	return &this
}

// GetFaceComparison returns the FaceComparison field value if set, zero value otherwise.
func (o *PhotoFullyAutoBreakdown) GetFaceComparison() FacialSimilarityPhotoBreakdownFaceComparison {
	if o == nil || IsNil(o.FaceComparison) {
		var ret FacialSimilarityPhotoBreakdownFaceComparison
		return ret
	}
	return *o.FaceComparison
}

// GetFaceComparisonOk returns a tuple with the FaceComparison field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhotoFullyAutoBreakdown) GetFaceComparisonOk() (*FacialSimilarityPhotoBreakdownFaceComparison, bool) {
	if o == nil || IsNil(o.FaceComparison) {
		return nil, false
	}
	return o.FaceComparison, true
}

// HasFaceComparison returns a boolean if a field has been set.
func (o *PhotoFullyAutoBreakdown) HasFaceComparison() bool {
	if o != nil && !IsNil(o.FaceComparison) {
		return true
	}

	return false
}

// SetFaceComparison gets a reference to the given FacialSimilarityPhotoBreakdownFaceComparison and assigns it to the FaceComparison field.
func (o *PhotoFullyAutoBreakdown) SetFaceComparison(v FacialSimilarityPhotoBreakdownFaceComparison) {
	o.FaceComparison = &v
}

// GetImageIntegrity returns the ImageIntegrity field value if set, zero value otherwise.
func (o *PhotoFullyAutoBreakdown) GetImageIntegrity() PhotoFullyAutoBreakdownImageIntegrity {
	if o == nil || IsNil(o.ImageIntegrity) {
		var ret PhotoFullyAutoBreakdownImageIntegrity
		return ret
	}
	return *o.ImageIntegrity
}

// GetImageIntegrityOk returns a tuple with the ImageIntegrity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhotoFullyAutoBreakdown) GetImageIntegrityOk() (*PhotoFullyAutoBreakdownImageIntegrity, bool) {
	if o == nil || IsNil(o.ImageIntegrity) {
		return nil, false
	}
	return o.ImageIntegrity, true
}

// HasImageIntegrity returns a boolean if a field has been set.
func (o *PhotoFullyAutoBreakdown) HasImageIntegrity() bool {
	if o != nil && !IsNil(o.ImageIntegrity) {
		return true
	}

	return false
}

// SetImageIntegrity gets a reference to the given PhotoFullyAutoBreakdownImageIntegrity and assigns it to the ImageIntegrity field.
func (o *PhotoFullyAutoBreakdown) SetImageIntegrity(v PhotoFullyAutoBreakdownImageIntegrity) {
	o.ImageIntegrity = &v
}

// GetVisualAuthenticity returns the VisualAuthenticity field value if set, zero value otherwise.
func (o *PhotoFullyAutoBreakdown) GetVisualAuthenticity() FacialSimilarityPhotoBreakdownVisualAuthenticity {
	if o == nil || IsNil(o.VisualAuthenticity) {
		var ret FacialSimilarityPhotoBreakdownVisualAuthenticity
		return ret
	}
	return *o.VisualAuthenticity
}

// GetVisualAuthenticityOk returns a tuple with the VisualAuthenticity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhotoFullyAutoBreakdown) GetVisualAuthenticityOk() (*FacialSimilarityPhotoBreakdownVisualAuthenticity, bool) {
	if o == nil || IsNil(o.VisualAuthenticity) {
		return nil, false
	}
	return o.VisualAuthenticity, true
}

// HasVisualAuthenticity returns a boolean if a field has been set.
func (o *PhotoFullyAutoBreakdown) HasVisualAuthenticity() bool {
	if o != nil && !IsNil(o.VisualAuthenticity) {
		return true
	}

	return false
}

// SetVisualAuthenticity gets a reference to the given FacialSimilarityPhotoBreakdownVisualAuthenticity and assigns it to the VisualAuthenticity field.
func (o *PhotoFullyAutoBreakdown) SetVisualAuthenticity(v FacialSimilarityPhotoBreakdownVisualAuthenticity) {
	o.VisualAuthenticity = &v
}

func (o PhotoFullyAutoBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PhotoFullyAutoBreakdown) ToMap() (map[string]interface{}, error) {
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

type NullablePhotoFullyAutoBreakdown struct {
	value *PhotoFullyAutoBreakdown
	isSet bool
}

func (v NullablePhotoFullyAutoBreakdown) Get() *PhotoFullyAutoBreakdown {
	return v.value
}

func (v *NullablePhotoFullyAutoBreakdown) Set(val *PhotoFullyAutoBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullablePhotoFullyAutoBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullablePhotoFullyAutoBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhotoFullyAutoBreakdown(val *PhotoFullyAutoBreakdown) *NullablePhotoFullyAutoBreakdown {
	return &NullablePhotoFullyAutoBreakdown{value: val, isSet: true}
}

func (v NullablePhotoFullyAutoBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhotoFullyAutoBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
