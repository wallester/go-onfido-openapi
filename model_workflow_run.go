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
	"time"
)

// checks if the WorkflowRun type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowRun{}

// WorkflowRun struct for WorkflowRun
type WorkflowRun struct {
	// The unique identifier for the Workflow Run.
	Id *string `json:"id,omitempty"`
	// The unique identifier for the Applicant.
	ApplicantId *string `json:"applicant_id,omitempty"`
	// The unique identifier for the Workflow.
	WorkflowId *string `json:"workflow_id,omitempty"`
	// The identifier for the Workflow version.
	WorkflowVersionId *string `json:"workflow_version_id,omitempty"`
	// The URL for viewing the Workflow Run results on your Onfido Dashboard.
	DashboardUrl *string `json:"dashboard_url,omitempty"`
	// The status of the Workflow Run. Possible values are 'processing', 'awaiting_input', 'approved', 'declined', 'review', 'abandoned' and 'error'.
	Status *string `json:"status,omitempty"`
	// Output object contains all of the properties configured on the Workflow version.
	Output map[string]interface{} `json:"output,omitempty"`
	// The reasons the Workflow Run outcome was reached. Configurable when creating the Workflow version.
	Reasons []string          `json:"reasons,omitempty"`
	Error   *WorkflowRunError `json:"error,omitempty"`
	// Tags or labels assigned to the workflow run.
	Tags []string         `json:"tags,omitempty"`
	Link *WorkflowRunLink `json:"link,omitempty"`
	// The date and time when the Workflow Run was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The date and time when the Workflow Run was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewWorkflowRun instantiates a new WorkflowRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRun() *WorkflowRun {
	this := WorkflowRun{}
	return &this
}

// NewWorkflowRunWithDefaults instantiates a new WorkflowRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRunWithDefaults() *WorkflowRun {
	this := WorkflowRun{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowRun) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowRun) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowRun) SetId(v string) {
	o.Id = &v
}

// GetApplicantId returns the ApplicantId field value if set, zero value otherwise.
func (o *WorkflowRun) GetApplicantId() string {
	if o == nil || IsNil(o.ApplicantId) {
		var ret string
		return ret
	}
	return *o.ApplicantId
}

// GetApplicantIdOk returns a tuple with the ApplicantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetApplicantIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicantId) {
		return nil, false
	}
	return o.ApplicantId, true
}

// HasApplicantId returns a boolean if a field has been set.
func (o *WorkflowRun) HasApplicantId() bool {
	if o != nil && !IsNil(o.ApplicantId) {
		return true
	}

	return false
}

// SetApplicantId gets a reference to the given string and assigns it to the ApplicantId field.
func (o *WorkflowRun) SetApplicantId(v string) {
	o.ApplicantId = &v
}

// GetWorkflowId returns the WorkflowId field value if set, zero value otherwise.
func (o *WorkflowRun) GetWorkflowId() string {
	if o == nil || IsNil(o.WorkflowId) {
		var ret string
		return ret
	}
	return *o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetWorkflowIdOk() (*string, bool) {
	if o == nil || IsNil(o.WorkflowId) {
		return nil, false
	}
	return o.WorkflowId, true
}

// HasWorkflowId returns a boolean if a field has been set.
func (o *WorkflowRun) HasWorkflowId() bool {
	if o != nil && !IsNil(o.WorkflowId) {
		return true
	}

	return false
}

// SetWorkflowId gets a reference to the given string and assigns it to the WorkflowId field.
func (o *WorkflowRun) SetWorkflowId(v string) {
	o.WorkflowId = &v
}

// GetWorkflowVersionId returns the WorkflowVersionId field value if set, zero value otherwise.
func (o *WorkflowRun) GetWorkflowVersionId() string {
	if o == nil || IsNil(o.WorkflowVersionId) {
		var ret string
		return ret
	}
	return *o.WorkflowVersionId
}

// GetWorkflowVersionIdOk returns a tuple with the WorkflowVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetWorkflowVersionIdOk() (*string, bool) {
	if o == nil || IsNil(o.WorkflowVersionId) {
		return nil, false
	}
	return o.WorkflowVersionId, true
}

// HasWorkflowVersionId returns a boolean if a field has been set.
func (o *WorkflowRun) HasWorkflowVersionId() bool {
	if o != nil && !IsNil(o.WorkflowVersionId) {
		return true
	}

	return false
}

// SetWorkflowVersionId gets a reference to the given string and assigns it to the WorkflowVersionId field.
func (o *WorkflowRun) SetWorkflowVersionId(v string) {
	o.WorkflowVersionId = &v
}

// GetDashboardUrl returns the DashboardUrl field value if set, zero value otherwise.
func (o *WorkflowRun) GetDashboardUrl() string {
	if o == nil || IsNil(o.DashboardUrl) {
		var ret string
		return ret
	}
	return *o.DashboardUrl
}

// GetDashboardUrlOk returns a tuple with the DashboardUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetDashboardUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DashboardUrl) {
		return nil, false
	}
	return o.DashboardUrl, true
}

// HasDashboardUrl returns a boolean if a field has been set.
func (o *WorkflowRun) HasDashboardUrl() bool {
	if o != nil && !IsNil(o.DashboardUrl) {
		return true
	}

	return false
}

// SetDashboardUrl gets a reference to the given string and assigns it to the DashboardUrl field.
func (o *WorkflowRun) SetDashboardUrl(v string) {
	o.DashboardUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowRun) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowRun) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WorkflowRun) SetStatus(v string) {
	o.Status = &v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *WorkflowRun) GetOutput() map[string]interface{} {
	if o == nil || IsNil(o.Output) {
		var ret map[string]interface{}
		return ret
	}
	return o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetOutputOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Output) {
		return map[string]interface{}{}, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *WorkflowRun) HasOutput() bool {
	if o != nil && !IsNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given map[string]interface{} and assigns it to the Output field.
func (o *WorkflowRun) SetOutput(v map[string]interface{}) {
	o.Output = v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *WorkflowRun) GetReasons() []string {
	if o == nil || IsNil(o.Reasons) {
		var ret []string
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetReasonsOk() ([]string, bool) {
	if o == nil || IsNil(o.Reasons) {
		return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *WorkflowRun) HasReasons() bool {
	if o != nil && !IsNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []string and assigns it to the Reasons field.
func (o *WorkflowRun) SetReasons(v []string) {
	o.Reasons = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *WorkflowRun) GetError() WorkflowRunError {
	if o == nil || IsNil(o.Error) {
		var ret WorkflowRunError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetErrorOk() (*WorkflowRunError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *WorkflowRun) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given WorkflowRunError and assigns it to the Error field.
func (o *WorkflowRun) SetError(v WorkflowRunError) {
	o.Error = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WorkflowRun) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WorkflowRun) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *WorkflowRun) SetTags(v []string) {
	o.Tags = v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *WorkflowRun) GetLink() WorkflowRunLink {
	if o == nil || IsNil(o.Link) {
		var ret WorkflowRunLink
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetLinkOk() (*WorkflowRunLink, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *WorkflowRun) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given WorkflowRunLink and assigns it to the Link field.
func (o *WorkflowRun) SetLink(v WorkflowRunLink) {
	o.Link = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WorkflowRun) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WorkflowRun) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *WorkflowRun) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *WorkflowRun) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *WorkflowRun) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *WorkflowRun) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o WorkflowRun) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowRun) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ApplicantId) {
		toSerialize["applicant_id"] = o.ApplicantId
	}
	if !IsNil(o.WorkflowId) {
		toSerialize["workflow_id"] = o.WorkflowId
	}
	if !IsNil(o.WorkflowVersionId) {
		toSerialize["workflow_version_id"] = o.WorkflowVersionId
	}
	if !IsNil(o.DashboardUrl) {
		toSerialize["dashboard_url"] = o.DashboardUrl
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Output) {
		toSerialize["output"] = o.Output
	}
	if !IsNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableWorkflowRun struct {
	value *WorkflowRun
	isSet bool
}

func (v NullableWorkflowRun) Get() *WorkflowRun {
	return v.value
}

func (v *NullableWorkflowRun) Set(val *WorkflowRun) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRun) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRun(val *WorkflowRun) *NullableWorkflowRun {
	return &NullableWorkflowRun{value: val, isSet: true}
}

func (v NullableWorkflowRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
