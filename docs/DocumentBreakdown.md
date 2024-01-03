# DocumentBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataComparison** | Pointer to [**DocumentBreakdownDataComparison**](DocumentBreakdownDataComparison.md) |  | [optional] 
**DataValidation** | Pointer to [**DocumentBreakdownDataValidation**](DocumentBreakdownDataValidation.md) |  | [optional] 
**ImageIntegrity** | Pointer to [**DocumentBreakdownImageIntegrity**](DocumentBreakdownImageIntegrity.md) |  | [optional] 
**VisualAuthenticity** | Pointer to [**DocumentBreakdownVisualAuthenticity**](DocumentBreakdownVisualAuthenticity.md) |  | [optional] 
**DataConsistency** | Pointer to [**DocumentBreakdownDataConsistency**](DocumentBreakdownDataConsistency.md) |  | [optional] 
**PoliceRecord** | Pointer to [**DocumentBreakdownPoliceRecord**](DocumentBreakdownPoliceRecord.md) |  | [optional] 
**CompromisedDocument** | Pointer to [**DocumentBreakdownCompromisedDocument**](DocumentBreakdownCompromisedDocument.md) |  | [optional] 
**AgeValidation** | Pointer to [**DocumentBreakdownAgeValidation**](DocumentBreakdownAgeValidation.md) |  | [optional] 
**IssuingAuthority** | Pointer to [**DocumentBreakdownIssuingAuthority**](DocumentBreakdownIssuingAuthority.md) |  | [optional] 

## Methods

### NewDocumentBreakdown

`func NewDocumentBreakdown() *DocumentBreakdown`

NewDocumentBreakdown instantiates a new DocumentBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentBreakdownWithDefaults

`func NewDocumentBreakdownWithDefaults() *DocumentBreakdown`

NewDocumentBreakdownWithDefaults instantiates a new DocumentBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataComparison

`func (o *DocumentBreakdown) GetDataComparison() DocumentBreakdownDataComparison`

GetDataComparison returns the DataComparison field if non-nil, zero value otherwise.

### GetDataComparisonOk

`func (o *DocumentBreakdown) GetDataComparisonOk() (*DocumentBreakdownDataComparison, bool)`

GetDataComparisonOk returns a tuple with the DataComparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataComparison

`func (o *DocumentBreakdown) SetDataComparison(v DocumentBreakdownDataComparison)`

SetDataComparison sets DataComparison field to given value.

### HasDataComparison

`func (o *DocumentBreakdown) HasDataComparison() bool`

HasDataComparison returns a boolean if a field has been set.

### GetDataValidation

`func (o *DocumentBreakdown) GetDataValidation() DocumentBreakdownDataValidation`

GetDataValidation returns the DataValidation field if non-nil, zero value otherwise.

### GetDataValidationOk

`func (o *DocumentBreakdown) GetDataValidationOk() (*DocumentBreakdownDataValidation, bool)`

GetDataValidationOk returns a tuple with the DataValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataValidation

`func (o *DocumentBreakdown) SetDataValidation(v DocumentBreakdownDataValidation)`

SetDataValidation sets DataValidation field to given value.

### HasDataValidation

`func (o *DocumentBreakdown) HasDataValidation() bool`

HasDataValidation returns a boolean if a field has been set.

### GetImageIntegrity

`func (o *DocumentBreakdown) GetImageIntegrity() DocumentBreakdownImageIntegrity`

GetImageIntegrity returns the ImageIntegrity field if non-nil, zero value otherwise.

### GetImageIntegrityOk

`func (o *DocumentBreakdown) GetImageIntegrityOk() (*DocumentBreakdownImageIntegrity, bool)`

GetImageIntegrityOk returns a tuple with the ImageIntegrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageIntegrity

`func (o *DocumentBreakdown) SetImageIntegrity(v DocumentBreakdownImageIntegrity)`

SetImageIntegrity sets ImageIntegrity field to given value.

### HasImageIntegrity

`func (o *DocumentBreakdown) HasImageIntegrity() bool`

HasImageIntegrity returns a boolean if a field has been set.

### GetVisualAuthenticity

`func (o *DocumentBreakdown) GetVisualAuthenticity() DocumentBreakdownVisualAuthenticity`

GetVisualAuthenticity returns the VisualAuthenticity field if non-nil, zero value otherwise.

### GetVisualAuthenticityOk

`func (o *DocumentBreakdown) GetVisualAuthenticityOk() (*DocumentBreakdownVisualAuthenticity, bool)`

GetVisualAuthenticityOk returns a tuple with the VisualAuthenticity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualAuthenticity

`func (o *DocumentBreakdown) SetVisualAuthenticity(v DocumentBreakdownVisualAuthenticity)`

SetVisualAuthenticity sets VisualAuthenticity field to given value.

### HasVisualAuthenticity

`func (o *DocumentBreakdown) HasVisualAuthenticity() bool`

HasVisualAuthenticity returns a boolean if a field has been set.

### GetDataConsistency

`func (o *DocumentBreakdown) GetDataConsistency() DocumentBreakdownDataConsistency`

GetDataConsistency returns the DataConsistency field if non-nil, zero value otherwise.

### GetDataConsistencyOk

`func (o *DocumentBreakdown) GetDataConsistencyOk() (*DocumentBreakdownDataConsistency, bool)`

GetDataConsistencyOk returns a tuple with the DataConsistency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataConsistency

`func (o *DocumentBreakdown) SetDataConsistency(v DocumentBreakdownDataConsistency)`

SetDataConsistency sets DataConsistency field to given value.

### HasDataConsistency

`func (o *DocumentBreakdown) HasDataConsistency() bool`

HasDataConsistency returns a boolean if a field has been set.

### GetPoliceRecord

`func (o *DocumentBreakdown) GetPoliceRecord() DocumentBreakdownPoliceRecord`

GetPoliceRecord returns the PoliceRecord field if non-nil, zero value otherwise.

### GetPoliceRecordOk

`func (o *DocumentBreakdown) GetPoliceRecordOk() (*DocumentBreakdownPoliceRecord, bool)`

GetPoliceRecordOk returns a tuple with the PoliceRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoliceRecord

`func (o *DocumentBreakdown) SetPoliceRecord(v DocumentBreakdownPoliceRecord)`

SetPoliceRecord sets PoliceRecord field to given value.

### HasPoliceRecord

`func (o *DocumentBreakdown) HasPoliceRecord() bool`

HasPoliceRecord returns a boolean if a field has been set.

### GetCompromisedDocument

`func (o *DocumentBreakdown) GetCompromisedDocument() DocumentBreakdownCompromisedDocument`

GetCompromisedDocument returns the CompromisedDocument field if non-nil, zero value otherwise.

### GetCompromisedDocumentOk

`func (o *DocumentBreakdown) GetCompromisedDocumentOk() (*DocumentBreakdownCompromisedDocument, bool)`

GetCompromisedDocumentOk returns a tuple with the CompromisedDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompromisedDocument

`func (o *DocumentBreakdown) SetCompromisedDocument(v DocumentBreakdownCompromisedDocument)`

SetCompromisedDocument sets CompromisedDocument field to given value.

### HasCompromisedDocument

`func (o *DocumentBreakdown) HasCompromisedDocument() bool`

HasCompromisedDocument returns a boolean if a field has been set.

### GetAgeValidation

`func (o *DocumentBreakdown) GetAgeValidation() DocumentBreakdownAgeValidation`

GetAgeValidation returns the AgeValidation field if non-nil, zero value otherwise.

### GetAgeValidationOk

`func (o *DocumentBreakdown) GetAgeValidationOk() (*DocumentBreakdownAgeValidation, bool)`

GetAgeValidationOk returns a tuple with the AgeValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeValidation

`func (o *DocumentBreakdown) SetAgeValidation(v DocumentBreakdownAgeValidation)`

SetAgeValidation sets AgeValidation field to given value.

### HasAgeValidation

`func (o *DocumentBreakdown) HasAgeValidation() bool`

HasAgeValidation returns a boolean if a field has been set.

### GetIssuingAuthority

`func (o *DocumentBreakdown) GetIssuingAuthority() DocumentBreakdownIssuingAuthority`

GetIssuingAuthority returns the IssuingAuthority field if non-nil, zero value otherwise.

### GetIssuingAuthorityOk

`func (o *DocumentBreakdown) GetIssuingAuthorityOk() (*DocumentBreakdownIssuingAuthority, bool)`

GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAuthority

`func (o *DocumentBreakdown) SetIssuingAuthority(v DocumentBreakdownIssuingAuthority)`

SetIssuingAuthority sets IssuingAuthority field to given value.

### HasIssuingAuthority

`func (o *DocumentBreakdown) HasIssuingAuthority() bool`

HasIssuingAuthority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


