# DocumentWithAddressInformationReport

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
**Breakdown** | Pointer to [**DocumentBreakdown**](DocumentBreakdown.md) |  | [optional] 
**Properties** | Pointer to [**DocumentProperties**](DocumentProperties.md) |  | [optional] 

## Methods

### NewDocumentWithAddressInformationReport

`func NewDocumentWithAddressInformationReport(name string, ) *DocumentWithAddressInformationReport`

NewDocumentWithAddressInformationReport instantiates a new DocumentWithAddressInformationReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentWithAddressInformationReportWithDefaults

`func NewDocumentWithAddressInformationReportWithDefaults() *DocumentWithAddressInformationReport`

NewDocumentWithAddressInformationReportWithDefaults instantiates a new DocumentWithAddressInformationReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadOnly

`func (o *DocumentWithAddressInformationReport) GetReadOnly() interface{}`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *DocumentWithAddressInformationReport) GetReadOnlyOk() (*interface{}, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *DocumentWithAddressInformationReport) SetReadOnly(v interface{})`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *DocumentWithAddressInformationReport) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### SetReadOnlyNil

`func (o *DocumentWithAddressInformationReport) SetReadOnlyNil(b bool)`

 SetReadOnlyNil sets the value for ReadOnly to be an explicit nil

### UnsetReadOnly
`func (o *DocumentWithAddressInformationReport) UnsetReadOnly()`

UnsetReadOnly ensures that no value is present for ReadOnly, not even an explicit nil
### GetId

`func (o *DocumentWithAddressInformationReport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentWithAddressInformationReport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentWithAddressInformationReport) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DocumentWithAddressInformationReport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DocumentWithAddressInformationReport) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DocumentWithAddressInformationReport) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DocumentWithAddressInformationReport) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DocumentWithAddressInformationReport) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHref

`func (o *DocumentWithAddressInformationReport) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DocumentWithAddressInformationReport) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DocumentWithAddressInformationReport) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *DocumentWithAddressInformationReport) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetStatus

`func (o *DocumentWithAddressInformationReport) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DocumentWithAddressInformationReport) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DocumentWithAddressInformationReport) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DocumentWithAddressInformationReport) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *DocumentWithAddressInformationReport) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DocumentWithAddressInformationReport) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DocumentWithAddressInformationReport) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *DocumentWithAddressInformationReport) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSubResult

`func (o *DocumentWithAddressInformationReport) GetSubResult() string`

GetSubResult returns the SubResult field if non-nil, zero value otherwise.

### GetSubResultOk

`func (o *DocumentWithAddressInformationReport) GetSubResultOk() (*string, bool)`

GetSubResultOk returns a tuple with the SubResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubResult

`func (o *DocumentWithAddressInformationReport) SetSubResult(v string)`

SetSubResult sets SubResult field to given value.

### HasSubResult

`func (o *DocumentWithAddressInformationReport) HasSubResult() bool`

HasSubResult returns a boolean if a field has been set.

### GetCheckId

`func (o *DocumentWithAddressInformationReport) GetCheckId() string`

GetCheckId returns the CheckId field if non-nil, zero value otherwise.

### GetCheckIdOk

`func (o *DocumentWithAddressInformationReport) GetCheckIdOk() (*string, bool)`

GetCheckIdOk returns a tuple with the CheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckId

`func (o *DocumentWithAddressInformationReport) SetCheckId(v string)`

SetCheckId sets CheckId field to given value.

### HasCheckId

`func (o *DocumentWithAddressInformationReport) HasCheckId() bool`

HasCheckId returns a boolean if a field has been set.

### GetDocuments

`func (o *DocumentWithAddressInformationReport) GetDocuments() []ReportDocument`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *DocumentWithAddressInformationReport) GetDocumentsOk() (*[]ReportDocument, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *DocumentWithAddressInformationReport) SetDocuments(v []ReportDocument)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *DocumentWithAddressInformationReport) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetIdPhotos

`func (o *DocumentWithAddressInformationReport) GetIdPhotos() []FacialSimilarityReportMedia`

GetIdPhotos returns the IdPhotos field if non-nil, zero value otherwise.

### GetIdPhotosOk

`func (o *DocumentWithAddressInformationReport) GetIdPhotosOk() (*[]FacialSimilarityReportMedia, bool)`

GetIdPhotosOk returns a tuple with the IdPhotos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdPhotos

`func (o *DocumentWithAddressInformationReport) SetIdPhotos(v []FacialSimilarityReportMedia)`

SetIdPhotos sets IdPhotos field to given value.

### HasIdPhotos

`func (o *DocumentWithAddressInformationReport) HasIdPhotos() bool`

HasIdPhotos returns a boolean if a field has been set.

### GetLivePhotos

`func (o *DocumentWithAddressInformationReport) GetLivePhotos() []FacialSimilarityReportMedia`

GetLivePhotos returns the LivePhotos field if non-nil, zero value otherwise.

### GetLivePhotosOk

`func (o *DocumentWithAddressInformationReport) GetLivePhotosOk() (*[]FacialSimilarityReportMedia, bool)`

GetLivePhotosOk returns a tuple with the LivePhotos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivePhotos

`func (o *DocumentWithAddressInformationReport) SetLivePhotos(v []FacialSimilarityReportMedia)`

SetLivePhotos sets LivePhotos field to given value.

### HasLivePhotos

`func (o *DocumentWithAddressInformationReport) HasLivePhotos() bool`

HasLivePhotos returns a boolean if a field has been set.

### GetLiveVideos

`func (o *DocumentWithAddressInformationReport) GetLiveVideos() []FacialSimilarityReportMedia`

GetLiveVideos returns the LiveVideos field if non-nil, zero value otherwise.

### GetLiveVideosOk

`func (o *DocumentWithAddressInformationReport) GetLiveVideosOk() (*[]FacialSimilarityReportMedia, bool)`

GetLiveVideosOk returns a tuple with the LiveVideos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveVideos

`func (o *DocumentWithAddressInformationReport) SetLiveVideos(v []FacialSimilarityReportMedia)`

SetLiveVideos sets LiveVideos field to given value.

### HasLiveVideos

`func (o *DocumentWithAddressInformationReport) HasLiveVideos() bool`

HasLiveVideos returns a boolean if a field has been set.

### GetMotionCaptures

`func (o *DocumentWithAddressInformationReport) GetMotionCaptures() []FacialSimilarityReportMedia`

GetMotionCaptures returns the MotionCaptures field if non-nil, zero value otherwise.

### GetMotionCapturesOk

`func (o *DocumentWithAddressInformationReport) GetMotionCapturesOk() (*[]FacialSimilarityReportMedia, bool)`

GetMotionCapturesOk returns a tuple with the MotionCaptures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotionCaptures

`func (o *DocumentWithAddressInformationReport) SetMotionCaptures(v []FacialSimilarityReportMedia)`

SetMotionCaptures sets MotionCaptures field to given value.

### HasMotionCaptures

`func (o *DocumentWithAddressInformationReport) HasMotionCaptures() bool`

HasMotionCaptures returns a boolean if a field has been set.

### GetName

`func (o *DocumentWithAddressInformationReport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentWithAddressInformationReport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentWithAddressInformationReport) SetName(v string)`

SetName sets Name field to given value.


### GetBreakdown

`func (o *DocumentWithAddressInformationReport) GetBreakdown() DocumentBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *DocumentWithAddressInformationReport) GetBreakdownOk() (*DocumentBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *DocumentWithAddressInformationReport) SetBreakdown(v DocumentBreakdown)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *DocumentWithAddressInformationReport) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.

### GetProperties

`func (o *DocumentWithAddressInformationReport) GetProperties() DocumentProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DocumentWithAddressInformationReport) GetPropertiesOk() (*DocumentProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DocumentWithAddressInformationReport) SetProperties(v DocumentProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DocumentWithAddressInformationReport) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


