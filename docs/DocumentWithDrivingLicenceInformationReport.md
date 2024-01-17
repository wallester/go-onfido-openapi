# DocumentWithDrivingLicenceInformationReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadOnly** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | Pointer to **string** | The unique identifier for the report. Read-only. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | The date and time at which the report was first initiated. Read-only. | [optional] [readonly] 
**Href** | Pointer to **string** | The API endpoint to retrieve the report. Read-only. | [optional] [readonly] 
**Status** | Pointer to **string** | The current state of the report in the checking process. Read-only. | [optional] [readonly] 
**Result** | Pointer to **string** | The result of the report. Read-only. | [optional] [readonly] 
**SubResult** | Pointer to **string** | The sub_result of the report. It gives a more detailed result for document reports only, and will be null otherwise. Read-only. | [optional] [readonly] 
**CheckId** | Pointer to **string** | The ID of the check to which the report belongs. Read-only. | [optional] [readonly] 
**Documents** | Pointer to [**[]ReportDocument**](ReportDocument.md) | Array of objects with document ids that were used in the Onfido engine. [ONLY POPULATED FOR DOCUMENT AND FACIAL SIMILARITY REPORTS] | [optional] 
**Name** | **string** | The name of the report type. | 
**Breakdown** | Pointer to [**DocumentBreakdown**](DocumentBreakdown.md) |  | [optional] 
**Properties** | Pointer to [**DocumentProperties**](DocumentProperties.md) |  | [optional] 

## Methods

### NewDocumentWithDrivingLicenceInformationReport

`func NewDocumentWithDrivingLicenceInformationReport(name string, ) *DocumentWithDrivingLicenceInformationReport`

NewDocumentWithDrivingLicenceInformationReport instantiates a new DocumentWithDrivingLicenceInformationReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentWithDrivingLicenceInformationReportWithDefaults

`func NewDocumentWithDrivingLicenceInformationReportWithDefaults() *DocumentWithDrivingLicenceInformationReport`

NewDocumentWithDrivingLicenceInformationReportWithDefaults instantiates a new DocumentWithDrivingLicenceInformationReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadOnly

`func (o *DocumentWithDrivingLicenceInformationReport) GetReadOnly() map[string]interface{}`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *DocumentWithDrivingLicenceInformationReport) GetReadOnlyOk() (*map[string]interface{}, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *DocumentWithDrivingLicenceInformationReport) SetReadOnly(v map[string]interface{})`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *DocumentWithDrivingLicenceInformationReport) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetId

`func (o *DocumentWithDrivingLicenceInformationReport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentWithDrivingLicenceInformationReport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentWithDrivingLicenceInformationReport) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DocumentWithDrivingLicenceInformationReport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DocumentWithDrivingLicenceInformationReport) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DocumentWithDrivingLicenceInformationReport) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DocumentWithDrivingLicenceInformationReport) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DocumentWithDrivingLicenceInformationReport) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHref

`func (o *DocumentWithDrivingLicenceInformationReport) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DocumentWithDrivingLicenceInformationReport) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DocumentWithDrivingLicenceInformationReport) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *DocumentWithDrivingLicenceInformationReport) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetStatus

`func (o *DocumentWithDrivingLicenceInformationReport) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DocumentWithDrivingLicenceInformationReport) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DocumentWithDrivingLicenceInformationReport) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DocumentWithDrivingLicenceInformationReport) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *DocumentWithDrivingLicenceInformationReport) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DocumentWithDrivingLicenceInformationReport) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DocumentWithDrivingLicenceInformationReport) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *DocumentWithDrivingLicenceInformationReport) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSubResult

`func (o *DocumentWithDrivingLicenceInformationReport) GetSubResult() string`

GetSubResult returns the SubResult field if non-nil, zero value otherwise.

### GetSubResultOk

`func (o *DocumentWithDrivingLicenceInformationReport) GetSubResultOk() (*string, bool)`

GetSubResultOk returns a tuple with the SubResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubResult

`func (o *DocumentWithDrivingLicenceInformationReport) SetSubResult(v string)`

SetSubResult sets SubResult field to given value.

### HasSubResult

`func (o *DocumentWithDrivingLicenceInformationReport) HasSubResult() bool`

HasSubResult returns a boolean if a field has been set.

### GetCheckId

`func (o *DocumentWithDrivingLicenceInformationReport) GetCheckId() string`

GetCheckId returns the CheckId field if non-nil, zero value otherwise.

### GetCheckIdOk

`func (o *DocumentWithDrivingLicenceInformationReport) GetCheckIdOk() (*string, bool)`

GetCheckIdOk returns a tuple with the CheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckId

`func (o *DocumentWithDrivingLicenceInformationReport) SetCheckId(v string)`

SetCheckId sets CheckId field to given value.

### HasCheckId

`func (o *DocumentWithDrivingLicenceInformationReport) HasCheckId() bool`

HasCheckId returns a boolean if a field has been set.

### GetDocuments

`func (o *DocumentWithDrivingLicenceInformationReport) GetDocuments() []ReportDocument`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *DocumentWithDrivingLicenceInformationReport) GetDocumentsOk() (*[]ReportDocument, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *DocumentWithDrivingLicenceInformationReport) SetDocuments(v []ReportDocument)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *DocumentWithDrivingLicenceInformationReport) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetName

`func (o *DocumentWithDrivingLicenceInformationReport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentWithDrivingLicenceInformationReport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentWithDrivingLicenceInformationReport) SetName(v string)`

SetName sets Name field to given value.


### GetBreakdown

`func (o *DocumentWithDrivingLicenceInformationReport) GetBreakdown() DocumentBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *DocumentWithDrivingLicenceInformationReport) GetBreakdownOk() (*DocumentBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *DocumentWithDrivingLicenceInformationReport) SetBreakdown(v DocumentBreakdown)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *DocumentWithDrivingLicenceInformationReport) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.

### GetProperties

`func (o *DocumentWithDrivingLicenceInformationReport) GetProperties() DocumentProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DocumentWithDrivingLicenceInformationReport) GetPropertiesOk() (*DocumentProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DocumentWithDrivingLicenceInformationReport) SetProperties(v DocumentProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DocumentWithDrivingLicenceInformationReport) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


