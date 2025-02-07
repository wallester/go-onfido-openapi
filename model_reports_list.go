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

// checks if the ReportsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportsList{}

// ReportsList struct for ReportsList
type ReportsList struct {
	Reports []Report `json:"reports,omitempty"`
}

// NewReportsList instantiates a new ReportsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportsList() *ReportsList {
	this := ReportsList{}
	return &this
}

// NewReportsListWithDefaults instantiates a new ReportsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportsListWithDefaults() *ReportsList {
	this := ReportsList{}
	return &this
}

// GetReports returns the Reports field value if set, zero value otherwise.
func (o *ReportsList) GetReports() []Report {
	if o == nil || IsNil(o.Reports) {
		var ret []Report
		return ret
	}
	return o.Reports
}

// GetReportsOk returns a tuple with the Reports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsList) GetReportsOk() ([]Report, bool) {
	if o == nil || IsNil(o.Reports) {
		return nil, false
	}
	return o.Reports, true
}

// HasReports returns a boolean if a field has been set.
func (o *ReportsList) HasReports() bool {
	if o != nil && !IsNil(o.Reports) {
		return true
	}

	return false
}

// SetReports gets a reference to the given []Report and assigns it to the Reports field.
func (o *ReportsList) SetReports(v []Report) {
	o.Reports = v
}

func (o ReportsList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reports) {
		toSerialize["reports"] = o.Reports
	}
	return toSerialize, nil
}

type NullableReportsList struct {
	value *ReportsList
	isSet bool
}

func (v NullableReportsList) Get() *ReportsList {
	return v.value
}

func (v *NullableReportsList) Set(val *ReportsList) {
	v.value = val
	v.isSet = true
}

func (v NullableReportsList) IsSet() bool {
	return v.isSet
}

func (v *NullableReportsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportsList(val *ReportsList) *NullableReportsList {
	return &NullableReportsList{value: val, isSet: true}
}

func (v NullableReportsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
