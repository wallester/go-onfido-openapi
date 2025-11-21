# ProofOfAddressReport

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
**Breakdown** | Pointer to [**ProofOfAddressBreakdown**](ProofOfAddressBreakdown.md) |  | [optional] 
**Properties** | Pointer to [**ProofOfAddressProperties**](ProofOfAddressProperties.md) |  | [optional] 

## Methods

### NewProofOfAddressReport

`func NewProofOfAddressReport(name string, ) *ProofOfAddressReport`

NewProofOfAddressReport instantiates a new ProofOfAddressReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProofOfAddressReportWithDefaults

`func NewProofOfAddressReportWithDefaults() *ProofOfAddressReport`

NewProofOfAddressReportWithDefaults instantiates a new ProofOfAddressReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadOnly

`func (o *ProofOfAddressReport) GetReadOnly() interface{}`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ProofOfAddressReport) GetReadOnlyOk() (*interface{}, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ProofOfAddressReport) SetReadOnly(v interface{})`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ProofOfAddressReport) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### SetReadOnlyNil

`func (o *ProofOfAddressReport) SetReadOnlyNil(b bool)`

 SetReadOnlyNil sets the value for ReadOnly to be an explicit nil

### UnsetReadOnly
`func (o *ProofOfAddressReport) UnsetReadOnly()`

UnsetReadOnly ensures that no value is present for ReadOnly, not even an explicit nil
### GetId

`func (o *ProofOfAddressReport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProofOfAddressReport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProofOfAddressReport) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProofOfAddressReport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProofOfAddressReport) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProofOfAddressReport) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProofOfAddressReport) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProofOfAddressReport) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHref

`func (o *ProofOfAddressReport) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ProofOfAddressReport) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ProofOfAddressReport) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ProofOfAddressReport) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetStatus

`func (o *ProofOfAddressReport) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProofOfAddressReport) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProofOfAddressReport) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProofOfAddressReport) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *ProofOfAddressReport) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ProofOfAddressReport) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ProofOfAddressReport) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *ProofOfAddressReport) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSubResult

`func (o *ProofOfAddressReport) GetSubResult() string`

GetSubResult returns the SubResult field if non-nil, zero value otherwise.

### GetSubResultOk

`func (o *ProofOfAddressReport) GetSubResultOk() (*string, bool)`

GetSubResultOk returns a tuple with the SubResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubResult

`func (o *ProofOfAddressReport) SetSubResult(v string)`

SetSubResult sets SubResult field to given value.

### HasSubResult

`func (o *ProofOfAddressReport) HasSubResult() bool`

HasSubResult returns a boolean if a field has been set.

### GetCheckId

`func (o *ProofOfAddressReport) GetCheckId() string`

GetCheckId returns the CheckId field if non-nil, zero value otherwise.

### GetCheckIdOk

`func (o *ProofOfAddressReport) GetCheckIdOk() (*string, bool)`

GetCheckIdOk returns a tuple with the CheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckId

`func (o *ProofOfAddressReport) SetCheckId(v string)`

SetCheckId sets CheckId field to given value.

### HasCheckId

`func (o *ProofOfAddressReport) HasCheckId() bool`

HasCheckId returns a boolean if a field has been set.

### GetDocuments

`func (o *ProofOfAddressReport) GetDocuments() []ReportDocument`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *ProofOfAddressReport) GetDocumentsOk() (*[]ReportDocument, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *ProofOfAddressReport) SetDocuments(v []ReportDocument)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *ProofOfAddressReport) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetName

`func (o *ProofOfAddressReport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProofOfAddressReport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProofOfAddressReport) SetName(v string)`

SetName sets Name field to given value.


### GetBreakdown

`func (o *ProofOfAddressReport) GetBreakdown() ProofOfAddressBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *ProofOfAddressReport) GetBreakdownOk() (*ProofOfAddressBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *ProofOfAddressReport) SetBreakdown(v ProofOfAddressBreakdown)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *ProofOfAddressReport) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.

### GetProperties

`func (o *ProofOfAddressReport) GetProperties() ProofOfAddressProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ProofOfAddressReport) GetPropertiesOk() (*ProofOfAddressProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ProofOfAddressReport) SetProperties(v ProofOfAddressProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ProofOfAddressReport) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


