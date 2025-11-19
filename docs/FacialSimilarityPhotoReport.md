# FacialSimilarityPhotoReport

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
**Breakdown** | Pointer to [**FacialSimilarityPhotoBreakdown**](FacialSimilarityPhotoBreakdown.md) |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** | The properties associated with the report, if any. Read-only. | [optional] [readonly] 

## Methods

### NewFacialSimilarityPhotoReport

`func NewFacialSimilarityPhotoReport(name string, ) *FacialSimilarityPhotoReport`

NewFacialSimilarityPhotoReport instantiates a new FacialSimilarityPhotoReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacialSimilarityPhotoReportWithDefaults

`func NewFacialSimilarityPhotoReportWithDefaults() *FacialSimilarityPhotoReport`

NewFacialSimilarityPhotoReportWithDefaults instantiates a new FacialSimilarityPhotoReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadOnly

`func (o *FacialSimilarityPhotoReport) GetReadOnly() interface{}`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *FacialSimilarityPhotoReport) GetReadOnlyOk() (*interface{}, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *FacialSimilarityPhotoReport) SetReadOnly(v interface{})`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *FacialSimilarityPhotoReport) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### SetReadOnlyNil

`func (o *FacialSimilarityPhotoReport) SetReadOnlyNil(b bool)`

 SetReadOnlyNil sets the value for ReadOnly to be an explicit nil

### UnsetReadOnly
`func (o *FacialSimilarityPhotoReport) UnsetReadOnly()`

UnsetReadOnly ensures that no value is present for ReadOnly, not even an explicit nil
### GetId

`func (o *FacialSimilarityPhotoReport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FacialSimilarityPhotoReport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FacialSimilarityPhotoReport) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FacialSimilarityPhotoReport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FacialSimilarityPhotoReport) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FacialSimilarityPhotoReport) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FacialSimilarityPhotoReport) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FacialSimilarityPhotoReport) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHref

`func (o *FacialSimilarityPhotoReport) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FacialSimilarityPhotoReport) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FacialSimilarityPhotoReport) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FacialSimilarityPhotoReport) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetStatus

`func (o *FacialSimilarityPhotoReport) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FacialSimilarityPhotoReport) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FacialSimilarityPhotoReport) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FacialSimilarityPhotoReport) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *FacialSimilarityPhotoReport) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *FacialSimilarityPhotoReport) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *FacialSimilarityPhotoReport) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *FacialSimilarityPhotoReport) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSubResult

`func (o *FacialSimilarityPhotoReport) GetSubResult() string`

GetSubResult returns the SubResult field if non-nil, zero value otherwise.

### GetSubResultOk

`func (o *FacialSimilarityPhotoReport) GetSubResultOk() (*string, bool)`

GetSubResultOk returns a tuple with the SubResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubResult

`func (o *FacialSimilarityPhotoReport) SetSubResult(v string)`

SetSubResult sets SubResult field to given value.

### HasSubResult

`func (o *FacialSimilarityPhotoReport) HasSubResult() bool`

HasSubResult returns a boolean if a field has been set.

### GetCheckId

`func (o *FacialSimilarityPhotoReport) GetCheckId() string`

GetCheckId returns the CheckId field if non-nil, zero value otherwise.

### GetCheckIdOk

`func (o *FacialSimilarityPhotoReport) GetCheckIdOk() (*string, bool)`

GetCheckIdOk returns a tuple with the CheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckId

`func (o *FacialSimilarityPhotoReport) SetCheckId(v string)`

SetCheckId sets CheckId field to given value.

### HasCheckId

`func (o *FacialSimilarityPhotoReport) HasCheckId() bool`

HasCheckId returns a boolean if a field has been set.

### GetDocuments

`func (o *FacialSimilarityPhotoReport) GetDocuments() []ReportDocument`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *FacialSimilarityPhotoReport) GetDocumentsOk() (*[]ReportDocument, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *FacialSimilarityPhotoReport) SetDocuments(v []ReportDocument)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *FacialSimilarityPhotoReport) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetName

`func (o *FacialSimilarityPhotoReport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FacialSimilarityPhotoReport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FacialSimilarityPhotoReport) SetName(v string)`

SetName sets Name field to given value.


### GetBreakdown

`func (o *FacialSimilarityPhotoReport) GetBreakdown() FacialSimilarityPhotoBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *FacialSimilarityPhotoReport) GetBreakdownOk() (*FacialSimilarityPhotoBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *FacialSimilarityPhotoReport) SetBreakdown(v FacialSimilarityPhotoBreakdown)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *FacialSimilarityPhotoReport) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.

### GetProperties

`func (o *FacialSimilarityPhotoReport) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *FacialSimilarityPhotoReport) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *FacialSimilarityPhotoReport) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *FacialSimilarityPhotoReport) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


