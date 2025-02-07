/*
Onfido API v3.6

The Onfido API is used to submit check requests.

API version: 3.6.0
Contact: engineering@onfido.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onfido_openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the DocumentWithDriverVerificationReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentWithDriverVerificationReport{}

// DocumentWithDriverVerificationReport struct for DocumentWithDriverVerificationReport
type DocumentWithDriverVerificationReport struct {
	ReadOnly map[string]interface{} `json:"readOnly,omitempty"`
	// The unique identifier for the report. Read-only.
	Id *string `json:"id,omitempty"`
	// The date and time at which the report was first initiated. Read-only.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The API endpoint to retrieve the report. Read-only.
	Href *string `json:"href,omitempty"`
	// The current state of the report in the checking process. Read-only.
	Status *string `json:"status,omitempty"`
	// The result of the report. Read-only.
	Result *string `json:"result,omitempty"`
	// The sub_result of the report. It gives a more detailed result for document reports only, and will be null otherwise. Read-only.
	SubResult *string `json:"sub_result,omitempty"`
	// The ID of the check to which the report belongs. Read-only.
	CheckId *string `json:"check_id,omitempty"`
	// Array of objects with document ids that were used in the Onfido engine. [ONLY POPULATED FOR DOCUMENT AND FACIAL SIMILARITY REPORTS]
	Documents []ReportDocument `json:"documents,omitempty"`
	// The name of the report type.
	Name       string                                               `json:"name"`
	Breakdown  *DocumentBreakdown                                   `json:"breakdown,omitempty"`
	Properties *DocumentWithDriverVerificationReportAllOfProperties `json:"properties,omitempty"`
}

type _DocumentWithDriverVerificationReport DocumentWithDriverVerificationReport

// NewDocumentWithDriverVerificationReport instantiates a new DocumentWithDriverVerificationReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentWithDriverVerificationReport(name string) *DocumentWithDriverVerificationReport {
	this := DocumentWithDriverVerificationReport{}
	this.Name = name
	return &this
}

// NewDocumentWithDriverVerificationReportWithDefaults instantiates a new DocumentWithDriverVerificationReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentWithDriverVerificationReportWithDefaults() *DocumentWithDriverVerificationReport {
	this := DocumentWithDriverVerificationReport{}
	return &this
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *DocumentWithDriverVerificationReport) GetReadOnly() map[string]interface{} {
	if o == nil || IsNil(o.ReadOnly) {
		var ret map[string]interface{}
		return ret
	}
	return o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentWithDriverVerificationReport) GetReadOnlyOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return map[string]interface{}{}, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *DocumentWithDriverVerificationReport) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given map[string]interface{} and assigns it to the ReadOnly field.
func (o *DocumentWithDriverVerificationReport) SetReadOnly(v map[string]interface{}) {
	o.ReadOnly = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DocumentWithDriverVerificationReport) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentWithDriverVerificationReport) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DocumentWithDriverVerificationReport) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DocumentWithDriverVerificationReport) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DocumentWithDriverVerificationReport) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentWithDriverVerificationReport) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DocumentWithDriverVerificationReport) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DocumentWithDriverVerificationReport) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *DocumentWithDriverVerificationReport) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentWithDriverVerificationReport) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *DocumentWithDriverVerificationReport) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *DocumentWithDriverVerificationReport) SetHref(v string) {
	o.Href = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DocumentWithDriverVerificationReport) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentWithDriverVerificationReport) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DocumentWithDriverVerificationReport) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DocumentWithDriverVerificationReport) SetStatus(v string) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DocumentWithDriverVerificationReport) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentWithDriverVerificationReport) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DocumentWithDriverVerificationReport) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *DocumentWithDriverVerificationReport) SetResult(v string) {
	o.Result = &v
}

// GetSubResult returns the SubResult field value if set, zero value otherwise.
func (o *DocumentWithDriverVerificationReport) GetSubResult() string {
	if o == nil || IsNil(o.SubResult) {
		var ret string
		return ret
	}
	return *o.SubResult
}

// GetSubResultOk returns a tuple with the SubResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentWithDriverVerificationReport) GetSubResultOk() (*string, bool) {
	if o == nil || IsNil(o.SubResult) {
		return nil, false
	}
	return o.SubResult, true
}

// HasSubResult returns a boolean if a field has been set.
func (o *DocumentWithDriverVerificationReport) HasSubResult() bool {
	if o != nil && !IsNil(o.SubResult) {
		return true
	}

	return false
}

// SetSubResult gets a reference to the given string and assigns it to the SubResult field.
func (o *DocumentWithDriverVerificationReport) SetSubResult(v string) {
	o.SubResult = &v
}

// GetCheckId returns the CheckId field value if set, zero value otherwise.
func (o *DocumentWithDriverVerificationReport) GetCheckId() string {
	if o == nil || IsNil(o.CheckId) {
		var ret string
		return ret
	}
	return *o.CheckId
}

// GetCheckIdOk returns a tuple with the CheckId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentWithDriverVerificationReport) GetCheckIdOk() (*string, bool) {
	if o == nil || IsNil(o.CheckId) {
		return nil, false
	}
	return o.CheckId, true
}

// HasCheckId returns a boolean if a field has been set.
func (o *DocumentWithDriverVerificationReport) HasCheckId() bool {
	if o != nil && !IsNil(o.CheckId) {
		return true
	}

	return false
}

// SetCheckId gets a reference to the given string and assigns it to the CheckId field.
func (o *DocumentWithDriverVerificationReport) SetCheckId(v string) {
	o.CheckId = &v
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *DocumentWithDriverVerificationReport) GetDocuments() []ReportDocument {
	if o == nil || IsNil(o.Documents) {
		var ret []ReportDocument
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentWithDriverVerificationReport) GetDocumentsOk() ([]ReportDocument, bool) {
	if o == nil || IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *DocumentWithDriverVerificationReport) HasDocuments() bool {
	if o != nil && !IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []ReportDocument and assigns it to the Documents field.
func (o *DocumentWithDriverVerificationReport) SetDocuments(v []ReportDocument) {
	o.Documents = v
}

// GetName returns the Name field value
func (o *DocumentWithDriverVerificationReport) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DocumentWithDriverVerificationReport) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DocumentWithDriverVerificationReport) SetName(v string) {
	o.Name = v
}

// GetBreakdown returns the Breakdown field value if set, zero value otherwise.
func (o *DocumentWithDriverVerificationReport) GetBreakdown() DocumentBreakdown {
	if o == nil || IsNil(o.Breakdown) {
		var ret DocumentBreakdown
		return ret
	}
	return *o.Breakdown
}

// GetBreakdownOk returns a tuple with the Breakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentWithDriverVerificationReport) GetBreakdownOk() (*DocumentBreakdown, bool) {
	if o == nil || IsNil(o.Breakdown) {
		return nil, false
	}
	return o.Breakdown, true
}

// HasBreakdown returns a boolean if a field has been set.
func (o *DocumentWithDriverVerificationReport) HasBreakdown() bool {
	if o != nil && !IsNil(o.Breakdown) {
		return true
	}

	return false
}

// SetBreakdown gets a reference to the given DocumentBreakdown and assigns it to the Breakdown field.
func (o *DocumentWithDriverVerificationReport) SetBreakdown(v DocumentBreakdown) {
	o.Breakdown = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *DocumentWithDriverVerificationReport) GetProperties() DocumentWithDriverVerificationReportAllOfProperties {
	if o == nil || IsNil(o.Properties) {
		var ret DocumentWithDriverVerificationReportAllOfProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentWithDriverVerificationReport) GetPropertiesOk() (*DocumentWithDriverVerificationReportAllOfProperties, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *DocumentWithDriverVerificationReport) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given DocumentWithDriverVerificationReportAllOfProperties and assigns it to the Properties field.
func (o *DocumentWithDriverVerificationReport) SetProperties(v DocumentWithDriverVerificationReportAllOfProperties) {
	o.Properties = &v
}

func (o DocumentWithDriverVerificationReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentWithDriverVerificationReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.SubResult) {
		toSerialize["sub_result"] = o.SubResult
	}
	if !IsNil(o.CheckId) {
		toSerialize["check_id"] = o.CheckId
	}
	if !IsNil(o.Documents) {
		toSerialize["documents"] = o.Documents
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Breakdown) {
		toSerialize["breakdown"] = o.Breakdown
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

func (o *DocumentWithDriverVerificationReport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDocumentWithDriverVerificationReport := _DocumentWithDriverVerificationReport{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDocumentWithDriverVerificationReport)

	if err != nil {
		return err
	}

	*o = DocumentWithDriverVerificationReport(varDocumentWithDriverVerificationReport)

	return err
}

type NullableDocumentWithDriverVerificationReport struct {
	value *DocumentWithDriverVerificationReport
	isSet bool
}

func (v NullableDocumentWithDriverVerificationReport) Get() *DocumentWithDriverVerificationReport {
	return v.value
}

func (v *NullableDocumentWithDriverVerificationReport) Set(val *DocumentWithDriverVerificationReport) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentWithDriverVerificationReport) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentWithDriverVerificationReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentWithDriverVerificationReport(val *DocumentWithDriverVerificationReport) *NullableDocumentWithDriverVerificationReport {
	return &NullableDocumentWithDriverVerificationReport{value: val, isSet: true}
}

func (v NullableDocumentWithDriverVerificationReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentWithDriverVerificationReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
