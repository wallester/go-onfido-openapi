# WorkflowRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for the Workflow Run. | [optional] 
**ApplicantId** | Pointer to **string** | The unique identifier for the Applicant. | [optional] 
**WorkflowId** | Pointer to **string** | The unique identifier for the Workflow. | [optional] 
**WorkflowVersionId** | Pointer to **string** | The identifier for the Workflow version. | [optional] 
**DashboardUrl** | Pointer to **string** | The URL for viewing the Workflow Run results on your Onfido Dashboard. | [optional] 
**Status** | Pointer to **string** | The status of the Workflow Run. Possible values are &#39;processing&#39;, &#39;awaiting_input&#39;, &#39;approved&#39;, &#39;declined&#39;, &#39;review&#39;, &#39;abandoned&#39; and &#39;error&#39;. | [optional] 
**Output** | Pointer to **map[string]interface{}** | Output object contains all of the properties configured on the Workflow version. | [optional] 
**Reasons** | Pointer to **[]string** | The reasons the Workflow Run outcome was reached. Configurable when creating the Workflow version. | [optional] 
**Error** | Pointer to [**WorkflowRunError**](WorkflowRunError.md) |  | [optional] 
**Tags** | Pointer to **[]string** | Tags or labels assigned to the workflow run. | [optional] 
**Link** | Pointer to [**WorkflowRunLink**](WorkflowRunLink.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when the Workflow Run was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when the Workflow Run was last updated. | [optional] 

## Methods

### NewWorkflowRun

`func NewWorkflowRun() *WorkflowRun`

NewWorkflowRun instantiates a new WorkflowRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRunWithDefaults

`func NewWorkflowRunWithDefaults() *WorkflowRun`

NewWorkflowRunWithDefaults instantiates a new WorkflowRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowRun) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowRun) HasId() bool`

HasId returns a boolean if a field has been set.

### GetApplicantId

`func (o *WorkflowRun) GetApplicantId() string`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *WorkflowRun) GetApplicantIdOk() (*string, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *WorkflowRun) SetApplicantId(v string)`

SetApplicantId sets ApplicantId field to given value.

### HasApplicantId

`func (o *WorkflowRun) HasApplicantId() bool`

HasApplicantId returns a boolean if a field has been set.

### GetWorkflowId

`func (o *WorkflowRun) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowRun) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowRun) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *WorkflowRun) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowVersionId

`func (o *WorkflowRun) GetWorkflowVersionId() string`

GetWorkflowVersionId returns the WorkflowVersionId field if non-nil, zero value otherwise.

### GetWorkflowVersionIdOk

`func (o *WorkflowRun) GetWorkflowVersionIdOk() (*string, bool)`

GetWorkflowVersionIdOk returns a tuple with the WorkflowVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowVersionId

`func (o *WorkflowRun) SetWorkflowVersionId(v string)`

SetWorkflowVersionId sets WorkflowVersionId field to given value.

### HasWorkflowVersionId

`func (o *WorkflowRun) HasWorkflowVersionId() bool`

HasWorkflowVersionId returns a boolean if a field has been set.

### GetDashboardUrl

`func (o *WorkflowRun) GetDashboardUrl() string`

GetDashboardUrl returns the DashboardUrl field if non-nil, zero value otherwise.

### GetDashboardUrlOk

`func (o *WorkflowRun) GetDashboardUrlOk() (*string, bool)`

GetDashboardUrlOk returns a tuple with the DashboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardUrl

`func (o *WorkflowRun) SetDashboardUrl(v string)`

SetDashboardUrl sets DashboardUrl field to given value.

### HasDashboardUrl

`func (o *WorkflowRun) HasDashboardUrl() bool`

HasDashboardUrl returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowRun) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowRun) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOutput

`func (o *WorkflowRun) GetOutput() map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WorkflowRun) GetOutputOk() (*map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WorkflowRun) SetOutput(v map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WorkflowRun) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetReasons

`func (o *WorkflowRun) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *WorkflowRun) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *WorkflowRun) SetReasons(v []string)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *WorkflowRun) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetError

`func (o *WorkflowRun) GetError() WorkflowRunError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WorkflowRun) GetErrorOk() (*WorkflowRunError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WorkflowRun) SetError(v WorkflowRunError)`

SetError sets Error field to given value.

### HasError

`func (o *WorkflowRun) HasError() bool`

HasError returns a boolean if a field has been set.

### GetTags

`func (o *WorkflowRun) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkflowRun) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkflowRun) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WorkflowRun) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLink

`func (o *WorkflowRun) GetLink() WorkflowRunLink`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *WorkflowRun) GetLinkOk() (*WorkflowRunLink, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *WorkflowRun) SetLink(v WorkflowRunLink)`

SetLink sets Link field to given value.

### HasLink

`func (o *WorkflowRun) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WorkflowRun) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkflowRun) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkflowRun) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WorkflowRun) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WorkflowRun) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkflowRun) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkflowRun) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WorkflowRun) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


