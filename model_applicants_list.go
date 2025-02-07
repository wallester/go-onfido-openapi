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

// checks if the ApplicantsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicantsList{}

// ApplicantsList struct for ApplicantsList
type ApplicantsList struct {
	Applicants []Applicant `json:"applicants,omitempty"`
}

// NewApplicantsList instantiates a new ApplicantsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicantsList() *ApplicantsList {
	this := ApplicantsList{}
	return &this
}

// NewApplicantsListWithDefaults instantiates a new ApplicantsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicantsListWithDefaults() *ApplicantsList {
	this := ApplicantsList{}
	return &this
}

// GetApplicants returns the Applicants field value if set, zero value otherwise.
func (o *ApplicantsList) GetApplicants() []Applicant {
	if o == nil || IsNil(o.Applicants) {
		var ret []Applicant
		return ret
	}
	return o.Applicants
}

// GetApplicantsOk returns a tuple with the Applicants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicantsList) GetApplicantsOk() ([]Applicant, bool) {
	if o == nil || IsNil(o.Applicants) {
		return nil, false
	}
	return o.Applicants, true
}

// HasApplicants returns a boolean if a field has been set.
func (o *ApplicantsList) HasApplicants() bool {
	if o != nil && !IsNil(o.Applicants) {
		return true
	}

	return false
}

// SetApplicants gets a reference to the given []Applicant and assigns it to the Applicants field.
func (o *ApplicantsList) SetApplicants(v []Applicant) {
	o.Applicants = v
}

func (o ApplicantsList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicantsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Applicants) {
		toSerialize["applicants"] = o.Applicants
	}
	return toSerialize, nil
}

type NullableApplicantsList struct {
	value *ApplicantsList
	isSet bool
}

func (v NullableApplicantsList) Get() *ApplicantsList {
	return v.value
}

func (v *NullableApplicantsList) Set(val *ApplicantsList) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicantsList) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicantsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicantsList(val *ApplicantsList) *NullableApplicantsList {
	return &NullableApplicantsList{value: val, isSet: true}
}

func (v NullableApplicantsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicantsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
