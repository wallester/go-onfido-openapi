# CreateWorkflowRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The unique identifier for the Workflow. | 
**ApplicantId** | **string** | The unique identifier for the Applicant. | 
**CustomData** | Pointer to **map[string]interface{}** | Optional object with custom input data to be used in the Workflow Run. | [optional] 
**Tags** | Pointer to **[]string** | Tags or labels to assign to the workflow run. | [optional] 
**Link** | Pointer to **map[string]interface{}** | Optional object for the configuration of the Workflow Run link. | [optional] 

## Methods

### NewCreateWorkflowRunRequest

`func NewCreateWorkflowRunRequest(workflowId string, applicantId string, ) *CreateWorkflowRunRequest`

NewCreateWorkflowRunRequest instantiates a new CreateWorkflowRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkflowRunRequestWithDefaults

`func NewCreateWorkflowRunRequestWithDefaults() *CreateWorkflowRunRequest`

NewCreateWorkflowRunRequestWithDefaults instantiates a new CreateWorkflowRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *CreateWorkflowRunRequest) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *CreateWorkflowRunRequest) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *CreateWorkflowRunRequest) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetApplicantId

`func (o *CreateWorkflowRunRequest) GetApplicantId() string`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *CreateWorkflowRunRequest) GetApplicantIdOk() (*string, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *CreateWorkflowRunRequest) SetApplicantId(v string)`

SetApplicantId sets ApplicantId field to given value.


### GetCustomData

`func (o *CreateWorkflowRunRequest) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *CreateWorkflowRunRequest) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *CreateWorkflowRunRequest) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *CreateWorkflowRunRequest) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### GetTags

`func (o *CreateWorkflowRunRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateWorkflowRunRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateWorkflowRunRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateWorkflowRunRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLink

`func (o *CreateWorkflowRunRequest) GetLink() map[string]interface{}`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *CreateWorkflowRunRequest) GetLinkOk() (*map[string]interface{}, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *CreateWorkflowRunRequest) SetLink(v map[string]interface{})`

SetLink sets Link field to given value.

### HasLink

`func (o *CreateWorkflowRunRequest) HasLink() bool`

HasLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


