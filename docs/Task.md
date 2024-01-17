# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The identifier for the Task. | [optional] 
**TaskDefId** | Pointer to **string** | The identifier for the Task Definition. | [optional] 
**Input** | Pointer to **map[string]interface{}** | Input object with the fields used by the Task to execute. | [optional] 
**Output** | Pointer to **map[string]interface{}** | Output object with the fields produced by the Task execution. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when the Task was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when the Task was last updated. | [optional] 

## Methods

### NewTask

`func NewTask() *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Task) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Task) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Task) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Task) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTaskDefId

`func (o *Task) GetTaskDefId() string`

GetTaskDefId returns the TaskDefId field if non-nil, zero value otherwise.

### GetTaskDefIdOk

`func (o *Task) GetTaskDefIdOk() (*string, bool)`

GetTaskDefIdOk returns a tuple with the TaskDefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefId

`func (o *Task) SetTaskDefId(v string)`

SetTaskDefId sets TaskDefId field to given value.

### HasTaskDefId

`func (o *Task) HasTaskDefId() bool`

HasTaskDefId returns a boolean if a field has been set.

### GetInput

`func (o *Task) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *Task) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *Task) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *Task) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetOutput

`func (o *Task) GetOutput() map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *Task) GetOutputOk() (*map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *Task) SetOutput(v map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *Task) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Task) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Task) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Task) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Task) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Task) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Task) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Task) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Task) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


