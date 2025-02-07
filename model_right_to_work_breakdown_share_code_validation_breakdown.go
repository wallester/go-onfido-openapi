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

// checks if the RightToWorkBreakdownShareCodeValidationBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RightToWorkBreakdownShareCodeValidationBreakdown{}

// RightToWorkBreakdownShareCodeValidationBreakdown struct for RightToWorkBreakdownShareCodeValidationBreakdown
type RightToWorkBreakdownShareCodeValidationBreakdown struct {
	DocumentId                 *RightToWorkBreakdownShareCodeValidationBreakdownDocumentId                 `json:"document_id,omitempty"`
	ApplicantHasValidShareCode *RightToWorkBreakdownShareCodeValidationBreakdownApplicantHasValidShareCode `json:"applicant_has_valid_share_code,omitempty"`
	NameMatched                *RightToWorkBreakdownShareCodeValidationBreakdownNameMatched                `json:"name_matched,omitempty"`
}

// NewRightToWorkBreakdownShareCodeValidationBreakdown instantiates a new RightToWorkBreakdownShareCodeValidationBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRightToWorkBreakdownShareCodeValidationBreakdown() *RightToWorkBreakdownShareCodeValidationBreakdown {
	this := RightToWorkBreakdownShareCodeValidationBreakdown{}
	return &this
}

// NewRightToWorkBreakdownShareCodeValidationBreakdownWithDefaults instantiates a new RightToWorkBreakdownShareCodeValidationBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRightToWorkBreakdownShareCodeValidationBreakdownWithDefaults() *RightToWorkBreakdownShareCodeValidationBreakdown {
	this := RightToWorkBreakdownShareCodeValidationBreakdown{}
	return &this
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *RightToWorkBreakdownShareCodeValidationBreakdown) GetDocumentId() RightToWorkBreakdownShareCodeValidationBreakdownDocumentId {
	if o == nil || IsNil(o.DocumentId) {
		var ret RightToWorkBreakdownShareCodeValidationBreakdownDocumentId
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RightToWorkBreakdownShareCodeValidationBreakdown) GetDocumentIdOk() (*RightToWorkBreakdownShareCodeValidationBreakdownDocumentId, bool) {
	if o == nil || IsNil(o.DocumentId) {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *RightToWorkBreakdownShareCodeValidationBreakdown) HasDocumentId() bool {
	if o != nil && !IsNil(o.DocumentId) {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given RightToWorkBreakdownShareCodeValidationBreakdownDocumentId and assigns it to the DocumentId field.
func (o *RightToWorkBreakdownShareCodeValidationBreakdown) SetDocumentId(v RightToWorkBreakdownShareCodeValidationBreakdownDocumentId) {
	o.DocumentId = &v
}

// GetApplicantHasValidShareCode returns the ApplicantHasValidShareCode field value if set, zero value otherwise.
func (o *RightToWorkBreakdownShareCodeValidationBreakdown) GetApplicantHasValidShareCode() RightToWorkBreakdownShareCodeValidationBreakdownApplicantHasValidShareCode {
	if o == nil || IsNil(o.ApplicantHasValidShareCode) {
		var ret RightToWorkBreakdownShareCodeValidationBreakdownApplicantHasValidShareCode
		return ret
	}
	return *o.ApplicantHasValidShareCode
}

// GetApplicantHasValidShareCodeOk returns a tuple with the ApplicantHasValidShareCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RightToWorkBreakdownShareCodeValidationBreakdown) GetApplicantHasValidShareCodeOk() (*RightToWorkBreakdownShareCodeValidationBreakdownApplicantHasValidShareCode, bool) {
	if o == nil || IsNil(o.ApplicantHasValidShareCode) {
		return nil, false
	}
	return o.ApplicantHasValidShareCode, true
}

// HasApplicantHasValidShareCode returns a boolean if a field has been set.
func (o *RightToWorkBreakdownShareCodeValidationBreakdown) HasApplicantHasValidShareCode() bool {
	if o != nil && !IsNil(o.ApplicantHasValidShareCode) {
		return true
	}

	return false
}

// SetApplicantHasValidShareCode gets a reference to the given RightToWorkBreakdownShareCodeValidationBreakdownApplicantHasValidShareCode and assigns it to the ApplicantHasValidShareCode field.
func (o *RightToWorkBreakdownShareCodeValidationBreakdown) SetApplicantHasValidShareCode(v RightToWorkBreakdownShareCodeValidationBreakdownApplicantHasValidShareCode) {
	o.ApplicantHasValidShareCode = &v
}

// GetNameMatched returns the NameMatched field value if set, zero value otherwise.
func (o *RightToWorkBreakdownShareCodeValidationBreakdown) GetNameMatched() RightToWorkBreakdownShareCodeValidationBreakdownNameMatched {
	if o == nil || IsNil(o.NameMatched) {
		var ret RightToWorkBreakdownShareCodeValidationBreakdownNameMatched
		return ret
	}
	return *o.NameMatched
}

// GetNameMatchedOk returns a tuple with the NameMatched field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RightToWorkBreakdownShareCodeValidationBreakdown) GetNameMatchedOk() (*RightToWorkBreakdownShareCodeValidationBreakdownNameMatched, bool) {
	if o == nil || IsNil(o.NameMatched) {
		return nil, false
	}
	return o.NameMatched, true
}

// HasNameMatched returns a boolean if a field has been set.
func (o *RightToWorkBreakdownShareCodeValidationBreakdown) HasNameMatched() bool {
	if o != nil && !IsNil(o.NameMatched) {
		return true
	}

	return false
}

// SetNameMatched gets a reference to the given RightToWorkBreakdownShareCodeValidationBreakdownNameMatched and assigns it to the NameMatched field.
func (o *RightToWorkBreakdownShareCodeValidationBreakdown) SetNameMatched(v RightToWorkBreakdownShareCodeValidationBreakdownNameMatched) {
	o.NameMatched = &v
}

func (o RightToWorkBreakdownShareCodeValidationBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RightToWorkBreakdownShareCodeValidationBreakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DocumentId) {
		toSerialize["document_id"] = o.DocumentId
	}
	if !IsNil(o.ApplicantHasValidShareCode) {
		toSerialize["applicant_has_valid_share_code"] = o.ApplicantHasValidShareCode
	}
	if !IsNil(o.NameMatched) {
		toSerialize["name_matched"] = o.NameMatched
	}
	return toSerialize, nil
}

type NullableRightToWorkBreakdownShareCodeValidationBreakdown struct {
	value *RightToWorkBreakdownShareCodeValidationBreakdown
	isSet bool
}

func (v NullableRightToWorkBreakdownShareCodeValidationBreakdown) Get() *RightToWorkBreakdownShareCodeValidationBreakdown {
	return v.value
}

func (v *NullableRightToWorkBreakdownShareCodeValidationBreakdown) Set(val *RightToWorkBreakdownShareCodeValidationBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableRightToWorkBreakdownShareCodeValidationBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableRightToWorkBreakdownShareCodeValidationBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRightToWorkBreakdownShareCodeValidationBreakdown(val *RightToWorkBreakdownShareCodeValidationBreakdown) *NullableRightToWorkBreakdownShareCodeValidationBreakdown {
	return &NullableRightToWorkBreakdownShareCodeValidationBreakdown{value: val, isSet: true}
}

func (v NullableRightToWorkBreakdownShareCodeValidationBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRightToWorkBreakdownShareCodeValidationBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
