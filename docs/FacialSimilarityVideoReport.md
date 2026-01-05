# FacialSimilarityVideoReport

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
**IdPhotos** | Pointer to [**[]FacialSimilarityReportMedia**](FacialSimilarityReportMedia.md) | Array of objects with id photo ids that were used in the Onfido engine. | [optional] 
**LivePhotos** | Pointer to [**[]FacialSimilarityReportMedia**](FacialSimilarityReportMedia.md) | Array of objects with live photo ids that were used in the Onfido engine. | [optional] 
**LiveVideos** | Pointer to [**[]FacialSimilarityReportMedia**](FacialSimilarityReportMedia.md) | Array of objects with live video ids that were used in the Onfido engine. | [optional] 
**MotionCaptures** | Pointer to [**[]FacialSimilarityReportMedia**](FacialSimilarityReportMedia.md) | Array of objects with motion capture ids that were used in the Onfido engine. | [optional] 
**Name** | **string** | The name of the report type. | 
**Breakdown** | Pointer to [**FacialSimilarityVideoBreakdown**](FacialSimilarityVideoBreakdown.md) |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** | The properties associated with the report, if any. Read-only. | [optional] [readonly] 

## Methods

### NewFacialSimilarityVideoReport

`func NewFacialSimilarityVideoReport(name string, ) *FacialSimilarityVideoReport`

NewFacialSimilarityVideoReport instantiates a new FacialSimilarityVideoReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacialSimilarityVideoReportWithDefaults

`func NewFacialSimilarityVideoReportWithDefaults() *FacialSimilarityVideoReport`

NewFacialSimilarityVideoReportWithDefaults instantiates a new FacialSimilarityVideoReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadOnly

`func (o *FacialSimilarityVideoReport) GetReadOnly() interface{}`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *FacialSimilarityVideoReport) GetReadOnlyOk() (*interface{}, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *FacialSimilarityVideoReport) SetReadOnly(v interface{})`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *FacialSimilarityVideoReport) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### SetReadOnlyNil

`func (o *FacialSimilarityVideoReport) SetReadOnlyNil(b bool)`

 SetReadOnlyNil sets the value for ReadOnly to be an explicit nil

### UnsetReadOnly
`func (o *FacialSimilarityVideoReport) UnsetReadOnly()`

UnsetReadOnly ensures that no value is present for ReadOnly, not even an explicit nil
### GetId

`func (o *FacialSimilarityVideoReport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FacialSimilarityVideoReport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FacialSimilarityVideoReport) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FacialSimilarityVideoReport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FacialSimilarityVideoReport) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FacialSimilarityVideoReport) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FacialSimilarityVideoReport) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FacialSimilarityVideoReport) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHref

`func (o *FacialSimilarityVideoReport) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FacialSimilarityVideoReport) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FacialSimilarityVideoReport) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FacialSimilarityVideoReport) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetStatus

`func (o *FacialSimilarityVideoReport) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FacialSimilarityVideoReport) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FacialSimilarityVideoReport) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FacialSimilarityVideoReport) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *FacialSimilarityVideoReport) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *FacialSimilarityVideoReport) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *FacialSimilarityVideoReport) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *FacialSimilarityVideoReport) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSubResult

`func (o *FacialSimilarityVideoReport) GetSubResult() string`

GetSubResult returns the SubResult field if non-nil, zero value otherwise.

### GetSubResultOk

`func (o *FacialSimilarityVideoReport) GetSubResultOk() (*string, bool)`

GetSubResultOk returns a tuple with the SubResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubResult

`func (o *FacialSimilarityVideoReport) SetSubResult(v string)`

SetSubResult sets SubResult field to given value.

### HasSubResult

`func (o *FacialSimilarityVideoReport) HasSubResult() bool`

HasSubResult returns a boolean if a field has been set.

### GetCheckId

`func (o *FacialSimilarityVideoReport) GetCheckId() string`

GetCheckId returns the CheckId field if non-nil, zero value otherwise.

### GetCheckIdOk

`func (o *FacialSimilarityVideoReport) GetCheckIdOk() (*string, bool)`

GetCheckIdOk returns a tuple with the CheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckId

`func (o *FacialSimilarityVideoReport) SetCheckId(v string)`

SetCheckId sets CheckId field to given value.

### HasCheckId

`func (o *FacialSimilarityVideoReport) HasCheckId() bool`

HasCheckId returns a boolean if a field has been set.

### GetDocuments

`func (o *FacialSimilarityVideoReport) GetDocuments() []ReportDocument`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *FacialSimilarityVideoReport) GetDocumentsOk() (*[]ReportDocument, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *FacialSimilarityVideoReport) SetDocuments(v []ReportDocument)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *FacialSimilarityVideoReport) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetIdPhotos

`func (o *FacialSimilarityVideoReport) GetIdPhotos() []FacialSimilarityReportMedia`

GetIdPhotos returns the IdPhotos field if non-nil, zero value otherwise.

### GetIdPhotosOk

`func (o *FacialSimilarityVideoReport) GetIdPhotosOk() (*[]FacialSimilarityReportMedia, bool)`

GetIdPhotosOk returns a tuple with the IdPhotos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdPhotos

`func (o *FacialSimilarityVideoReport) SetIdPhotos(v []FacialSimilarityReportMedia)`

SetIdPhotos sets IdPhotos field to given value.

### HasIdPhotos

`func (o *FacialSimilarityVideoReport) HasIdPhotos() bool`

HasIdPhotos returns a boolean if a field has been set.

### GetLivePhotos

`func (o *FacialSimilarityVideoReport) GetLivePhotos() []FacialSimilarityReportMedia`

GetLivePhotos returns the LivePhotos field if non-nil, zero value otherwise.

### GetLivePhotosOk

`func (o *FacialSimilarityVideoReport) GetLivePhotosOk() (*[]FacialSimilarityReportMedia, bool)`

GetLivePhotosOk returns a tuple with the LivePhotos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivePhotos

`func (o *FacialSimilarityVideoReport) SetLivePhotos(v []FacialSimilarityReportMedia)`

SetLivePhotos sets LivePhotos field to given value.

### HasLivePhotos

`func (o *FacialSimilarityVideoReport) HasLivePhotos() bool`

HasLivePhotos returns a boolean if a field has been set.

### GetLiveVideos

`func (o *FacialSimilarityVideoReport) GetLiveVideos() []FacialSimilarityReportMedia`

GetLiveVideos returns the LiveVideos field if non-nil, zero value otherwise.

### GetLiveVideosOk

`func (o *FacialSimilarityVideoReport) GetLiveVideosOk() (*[]FacialSimilarityReportMedia, bool)`

GetLiveVideosOk returns a tuple with the LiveVideos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveVideos

`func (o *FacialSimilarityVideoReport) SetLiveVideos(v []FacialSimilarityReportMedia)`

SetLiveVideos sets LiveVideos field to given value.

### HasLiveVideos

`func (o *FacialSimilarityVideoReport) HasLiveVideos() bool`

HasLiveVideos returns a boolean if a field has been set.

### GetMotionCaptures

`func (o *FacialSimilarityVideoReport) GetMotionCaptures() []FacialSimilarityReportMedia`

GetMotionCaptures returns the MotionCaptures field if non-nil, zero value otherwise.

### GetMotionCapturesOk

`func (o *FacialSimilarityVideoReport) GetMotionCapturesOk() (*[]FacialSimilarityReportMedia, bool)`

GetMotionCapturesOk returns a tuple with the MotionCaptures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotionCaptures

`func (o *FacialSimilarityVideoReport) SetMotionCaptures(v []FacialSimilarityReportMedia)`

SetMotionCaptures sets MotionCaptures field to given value.

### HasMotionCaptures

`func (o *FacialSimilarityVideoReport) HasMotionCaptures() bool`

HasMotionCaptures returns a boolean if a field has been set.

### GetName

`func (o *FacialSimilarityVideoReport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FacialSimilarityVideoReport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FacialSimilarityVideoReport) SetName(v string)`

SetName sets Name field to given value.


### GetBreakdown

`func (o *FacialSimilarityVideoReport) GetBreakdown() FacialSimilarityVideoBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *FacialSimilarityVideoReport) GetBreakdownOk() (*FacialSimilarityVideoBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *FacialSimilarityVideoReport) SetBreakdown(v FacialSimilarityVideoBreakdown)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *FacialSimilarityVideoReport) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.

### GetProperties

`func (o *FacialSimilarityVideoReport) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *FacialSimilarityVideoReport) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *FacialSimilarityVideoReport) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *FacialSimilarityVideoReport) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


