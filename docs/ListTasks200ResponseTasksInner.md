# ListTasks200ResponseTasksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The identifier for the Task. | [optional] 
**TaskDefId** | Pointer to **string** | The identifier for the Task Definition. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when the Task was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when the Task was last updated. | [optional] 

## Methods

### NewListTasks200ResponseTasksInner

`func NewListTasks200ResponseTasksInner() *ListTasks200ResponseTasksInner`

NewListTasks200ResponseTasksInner instantiates a new ListTasks200ResponseTasksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTasks200ResponseTasksInnerWithDefaults

`func NewListTasks200ResponseTasksInnerWithDefaults() *ListTasks200ResponseTasksInner`

NewListTasks200ResponseTasksInnerWithDefaults instantiates a new ListTasks200ResponseTasksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListTasks200ResponseTasksInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListTasks200ResponseTasksInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListTasks200ResponseTasksInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListTasks200ResponseTasksInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTaskDefId

`func (o *ListTasks200ResponseTasksInner) GetTaskDefId() string`

GetTaskDefId returns the TaskDefId field if non-nil, zero value otherwise.

### GetTaskDefIdOk

`func (o *ListTasks200ResponseTasksInner) GetTaskDefIdOk() (*string, bool)`

GetTaskDefIdOk returns a tuple with the TaskDefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefId

`func (o *ListTasks200ResponseTasksInner) SetTaskDefId(v string)`

SetTaskDefId sets TaskDefId field to given value.

### HasTaskDefId

`func (o *ListTasks200ResponseTasksInner) HasTaskDefId() bool`

HasTaskDefId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ListTasks200ResponseTasksInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListTasks200ResponseTasksInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListTasks200ResponseTasksInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListTasks200ResponseTasksInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ListTasks200ResponseTasksInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListTasks200ResponseTasksInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListTasks200ResponseTasksInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ListTasks200ResponseTasksInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


