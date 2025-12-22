# DeviceIntelligenceReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadOnly** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **string** | The unique identifier for the report. Read-only. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | The date and time at which the report was first initiated. Read-only. | [optional] [readonly] 
**Href** | Pointer to **string** | The API endpoint to retrieve the report. Read-only. | [optional] [readonly] 
**Status** | Pointer to **string** | The current state of the report in the checking process. Read-only. | [optional] [readonly] 
**Result** | Pointer to **string** | The result of the report. Read-only. | [optional] [readonly] 
**SubResult** | Pointer to **string** | The sub_result of the report. It gives a more detailed result for document reports only, and will be null otherwise. Read-only. | [optional] [readonly] 
**CheckId** | Pointer to **string** | The ID of the check to which the report belongs. Read-only. | [optional] [readonly] 
**Documents** | Pointer to [**[]ReportDocument**](ReportDocument.md) | Array of objects with document ids that were used in the Onfido engine. [ONLY POPULATED FOR DOCUMENT AND FACIAL SIMILARITY REPORTS] | [optional] 
**Name** | **string** | The name of the report type. | 
**Breakdown** | Pointer to [**DeviceIntelligenceBreakdown**](DeviceIntelligenceBreakdown.md) |  | [optional] 
**Properties** | Pointer to [**DeviceIntelligenceProperties**](DeviceIntelligenceProperties.md) |  | [optional] 

## Methods

### NewDeviceIntelligenceReport

`func NewDeviceIntelligenceReport(name string, ) *DeviceIntelligenceReport`

NewDeviceIntelligenceReport instantiates a new DeviceIntelligenceReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceIntelligenceReportWithDefaults

`func NewDeviceIntelligenceReportWithDefaults() *DeviceIntelligenceReport`

NewDeviceIntelligenceReportWithDefaults instantiates a new DeviceIntelligenceReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadOnly

`func (o *DeviceIntelligenceReport) GetReadOnly() interface{}`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *DeviceIntelligenceReport) GetReadOnlyOk() (*interface{}, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *DeviceIntelligenceReport) SetReadOnly(v interface{})`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *DeviceIntelligenceReport) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### SetReadOnlyNil

`func (o *DeviceIntelligenceReport) SetReadOnlyNil(b bool)`

 SetReadOnlyNil sets the value for ReadOnly to be an explicit nil

### UnsetReadOnly
`func (o *DeviceIntelligenceReport) UnsetReadOnly()`

UnsetReadOnly ensures that no value is present for ReadOnly, not even an explicit nil
### GetId

`func (o *DeviceIntelligenceReport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceIntelligenceReport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceIntelligenceReport) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceIntelligenceReport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DeviceIntelligenceReport) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeviceIntelligenceReport) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeviceIntelligenceReport) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DeviceIntelligenceReport) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHref

`func (o *DeviceIntelligenceReport) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DeviceIntelligenceReport) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DeviceIntelligenceReport) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *DeviceIntelligenceReport) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceIntelligenceReport) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceIntelligenceReport) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceIntelligenceReport) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceIntelligenceReport) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *DeviceIntelligenceReport) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DeviceIntelligenceReport) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DeviceIntelligenceReport) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *DeviceIntelligenceReport) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSubResult

`func (o *DeviceIntelligenceReport) GetSubResult() string`

GetSubResult returns the SubResult field if non-nil, zero value otherwise.

### GetSubResultOk

`func (o *DeviceIntelligenceReport) GetSubResultOk() (*string, bool)`

GetSubResultOk returns a tuple with the SubResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubResult

`func (o *DeviceIntelligenceReport) SetSubResult(v string)`

SetSubResult sets SubResult field to given value.

### HasSubResult

`func (o *DeviceIntelligenceReport) HasSubResult() bool`

HasSubResult returns a boolean if a field has been set.

### GetCheckId

`func (o *DeviceIntelligenceReport) GetCheckId() string`

GetCheckId returns the CheckId field if non-nil, zero value otherwise.

### GetCheckIdOk

`func (o *DeviceIntelligenceReport) GetCheckIdOk() (*string, bool)`

GetCheckIdOk returns a tuple with the CheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckId

`func (o *DeviceIntelligenceReport) SetCheckId(v string)`

SetCheckId sets CheckId field to given value.

### HasCheckId

`func (o *DeviceIntelligenceReport) HasCheckId() bool`

HasCheckId returns a boolean if a field has been set.

### GetDocuments

`func (o *DeviceIntelligenceReport) GetDocuments() []ReportDocument`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *DeviceIntelligenceReport) GetDocumentsOk() (*[]ReportDocument, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *DeviceIntelligenceReport) SetDocuments(v []ReportDocument)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *DeviceIntelligenceReport) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetName

`func (o *DeviceIntelligenceReport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceIntelligenceReport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceIntelligenceReport) SetName(v string)`

SetName sets Name field to given value.


### GetBreakdown

`func (o *DeviceIntelligenceReport) GetBreakdown() DeviceIntelligenceBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *DeviceIntelligenceReport) GetBreakdownOk() (*DeviceIntelligenceBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *DeviceIntelligenceReport) SetBreakdown(v DeviceIntelligenceBreakdown)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *DeviceIntelligenceReport) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.

### GetProperties

`func (o *DeviceIntelligenceReport) GetProperties() DeviceIntelligenceProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DeviceIntelligenceReport) GetPropertiesOk() (*DeviceIntelligenceProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DeviceIntelligenceReport) SetProperties(v DeviceIntelligenceProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DeviceIntelligenceReport) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


