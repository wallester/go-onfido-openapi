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

// checks if the DocumentWithDrivingLicenceInformationReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentWithDrivingLicenceInformationReport{}

// DocumentWithDrivingLicenceInformationReport struct for DocumentWithDrivingLicenceInformationReport
type DocumentWithDrivingLicenceInformationReport struct {
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
	Name       string              `json:"name"`
	Breakdown  *DocumentBreakdown  `json:"breakdown,omitempty"`
	Properties *DocumentProperties `json:"properties,omitempty"`
}

type _DocumentWithDrivingLicenceInformationReport DocumentWithDrivingLicenceInformationReport

// NewDocumentWithDrivingLicenceInformationReport instantiates a new DocumentWithDrivingLicenceInformationReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentWithDrivingLicenceInformationReport(name string) *DocumentWithDrivingLicenceInformationReport {
	this := DocumentWithDrivingLicenceInformationReport{}
	this.Name = name
	return &this
}

// NewDocumentWithDrivingLicenceInformationReportWithDefaults instantiates a new DocumentWithDrivingLicenceInformationReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentWithDrivingLicenceInformationReportWithDefaults() *DocumentWithDrivingLicenceInformationReport {
	this := DocumentWithDrivingLicenceInformationReport{}
	return &this
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *DocumentWithDrivingLicenceInformationReport) GetReadOnly() map[string]interface{} {
	if o == nil || IsNil(o.ReadOnly) {
		var ret map[string]interface{}
		return ret
	}
	return o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentWithDrivingLicenceInformationReport) GetReadOnlyOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return map[string]interface{}{}, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *DocumentWithDrivingLicenceInformationReport) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given map[string]interface{} and assigns it to the ReadOnly field.
func (o *DocumentWithDrivingLicenceInformationReport) SetReadOnly(v map[string]interface{}) {
	o.ReadOnly = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DocumentWithDrivingLicenceInformationReport) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentWithDrivingLicenceInformationReport) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DocumentWithDrivingLicenceInformationReport) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DocumentWithDrivingLicenceInformationReport) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DocumentWithDrivingLicenceInformationReport) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentWithDrivingLicenceInformationReport) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DocumentWithDrivingLicenceInformationReport) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DocumentWithDrivingLicenceInformationReport) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *DocumentWithDrivingLicenceInformationReport) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentWithDrivingLicenceInformationReport) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *DocumentWithDrivingLicenceInformationReport) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *DocumentWithDrivingLicenceInformationReport) SetHref(v string) {
	o.Href = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DocumentWithDrivingLicenceInformationReport) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentWithDrivingLicenceInformationReport) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DocumentWithDrivingLicenceInformationReport) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DocumentWithDrivingLicenceInformationReport) SetStatus(v string) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DocumentWithDrivingLicenceInformationReport) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentWithDrivingLicenceInformationReport) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DocumentWithDrivingLicenceInformationReport) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *DocumentWithDrivingLicenceInformationReport) SetResult(v string) {
	o.Result = &v
}

// GetSubResult returns the SubResult field value if set, zero value otherwise.
func (o *DocumentWithDrivingLicenceInformationReport) GetSubResult() string {
	if o == nil || IsNil(o.SubResult) {
		var ret string
		return ret
	}
	return *o.SubResult
}

// GetSubResultOk returns a tuple with the SubResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentWithDrivingLicenceInformationReport) GetSubResultOk() (*string, bool) {
	if o == nil || IsNil(o.SubResult) {
		return nil, false
	}
	return o.SubResult, true
}

// HasSubResult returns a boolean if a field has been set.
func (o *DocumentWithDrivingLicenceInformationReport) HasSubResult() bool {
	if o != nil && !IsNil(o.SubResult) {
		return true
	}

	return false
}

// SetSubResult gets a reference to the given string and assigns it to the SubResult field.
func (o *DocumentWithDrivingLicenceInformationReport) SetSubResult(v string) {
	o.SubResult = &v
}

// GetCheckId returns the CheckId field value if set, zero value otherwise.
func (o *DocumentWithDrivingLicenceInformationReport) GetCheckId() string {
	if o == nil || IsNil(o.CheckId) {
		var ret string
		return ret
	}
	return *o.CheckId
}

// GetCheckIdOk returns a tuple with the CheckId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentWithDrivingLicenceInformationReport) GetCheckIdOk() (*string, bool) {
	if o == nil || IsNil(o.CheckId) {
		return nil, false
	}
	return o.CheckId, true
}

// HasCheckId returns a boolean if a field has been set.
func (o *DocumentWithDrivingLicenceInformationReport) HasCheckId() bool {
	if o != nil && !IsNil(o.CheckId) {
		return true
	}

	return false
}

// SetCheckId gets a reference to the given string and assigns it to the CheckId field.
func (o *DocumentWithDrivingLicenceInformationReport) SetCheckId(v string) {
	o.CheckId = &v
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *DocumentWithDrivingLicenceInformationReport) GetDocuments() []ReportDocument {
	if o == nil || IsNil(o.Documents) {
		var ret []ReportDocument
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentWithDrivingLicenceInformationReport) GetDocumentsOk() ([]ReportDocument, bool) {
	if o == nil || IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *DocumentWithDrivingLicenceInformationReport) HasDocuments() bool {
	if o != nil && !IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []ReportDocument and assigns it to the Documents field.
func (o *DocumentWithDrivingLicenceInformationReport) SetDocuments(v []ReportDocument) {
	o.Documents = v
}

// GetName returns the Name field value
func (o *DocumentWithDrivingLicenceInformationReport) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DocumentWithDrivingLicenceInformationReport) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DocumentWithDrivingLicenceInformationReport) SetName(v string) {
	o.Name = v
}

// GetBreakdown returns the Breakdown field value if set, zero value otherwise.
func (o *DocumentWithDrivingLicenceInformationReport) GetBreakdown() DocumentBreakdown {
	if o == nil || IsNil(o.Breakdown) {
		var ret DocumentBreakdown
		return ret
	}
	return *o.Breakdown
}

// GetBreakdownOk returns a tuple with the Breakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentWithDrivingLicenceInformationReport) GetBreakdownOk() (*DocumentBreakdown, bool) {
	if o == nil || IsNil(o.Breakdown) {
		return nil, false
	}
	return o.Breakdown, true
}

// HasBreakdown returns a boolean if a field has been set.
func (o *DocumentWithDrivingLicenceInformationReport) HasBreakdown() bool {
	if o != nil && !IsNil(o.Breakdown) {
		return true
	}

	return false
}

// SetBreakdown gets a reference to the given DocumentBreakdown and assigns it to the Breakdown field.
func (o *DocumentWithDrivingLicenceInformationReport) SetBreakdown(v DocumentBreakdown) {
	o.Breakdown = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *DocumentWithDrivingLicenceInformationReport) GetProperties() DocumentProperties {
	if o == nil || IsNil(o.Properties) {
		var ret DocumentProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentWithDrivingLicenceInformationReport) GetPropertiesOk() (*DocumentProperties, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *DocumentWithDrivingLicenceInformationReport) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given DocumentProperties and assigns it to the Properties field.
func (o *DocumentWithDrivingLicenceInformationReport) SetProperties(v DocumentProperties) {
	o.Properties = &v
}

func (o DocumentWithDrivingLicenceInformationReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentWithDrivingLicenceInformationReport) ToMap() (map[string]interface{}, error) {
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

func (o *DocumentWithDrivingLicenceInformationReport) UnmarshalJSON(data []byte) (err error) {
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

	varDocumentWithDrivingLicenceInformationReport := _DocumentWithDrivingLicenceInformationReport{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDocumentWithDrivingLicenceInformationReport)

	if err != nil {
		return err
	}

	*o = DocumentWithDrivingLicenceInformationReport(varDocumentWithDrivingLicenceInformationReport)

	return err
}

type NullableDocumentWithDrivingLicenceInformationReport struct {
	value *DocumentWithDrivingLicenceInformationReport
	isSet bool
}

func (v NullableDocumentWithDrivingLicenceInformationReport) Get() *DocumentWithDrivingLicenceInformationReport {
	return v.value
}

func (v *NullableDocumentWithDrivingLicenceInformationReport) Set(val *DocumentWithDrivingLicenceInformationReport) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentWithDrivingLicenceInformationReport) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentWithDrivingLicenceInformationReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentWithDrivingLicenceInformationReport(val *DocumentWithDrivingLicenceInformationReport) *NullableDocumentWithDrivingLicenceInformationReport {
	return &NullableDocumentWithDrivingLicenceInformationReport{value: val, isSet: true}
}

func (v NullableDocumentWithDrivingLicenceInformationReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentWithDrivingLicenceInformationReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
