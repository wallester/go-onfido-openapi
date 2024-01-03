# RightToWorkProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nationality** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**IssuingCountry** | Pointer to **string** |  | [optional] 
**Gender** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**DocumentType** | Pointer to **string** |  | [optional] 
**DocumentNumbers** | Pointer to [**[]DocumentPropertiesDocumentNumbersInner**](DocumentPropertiesDocumentNumbersInner.md) |  | [optional] 
**DateOfExpiry** | Pointer to **string** |  | [optional] 
**DateOfBirth** | Pointer to **string** |  | [optional] 
**DocumentVersion** | Pointer to **string** |  | [optional] 
**IssuingDate** | Pointer to **string** |  | [optional] 
**MrzLine1** | Pointer to **string** |  | [optional] 
**MrzLine2** | Pointer to **string** |  | [optional] 
**PlaceOfBirth** | Pointer to **string** |  | [optional] 
**IssuingAuthority** | Pointer to **string** |  | [optional] 

## Methods

### NewRightToWorkProperties

`func NewRightToWorkProperties() *RightToWorkProperties`

NewRightToWorkProperties instantiates a new RightToWorkProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRightToWorkPropertiesWithDefaults

`func NewRightToWorkPropertiesWithDefaults() *RightToWorkProperties`

NewRightToWorkPropertiesWithDefaults instantiates a new RightToWorkProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNationality

`func (o *RightToWorkProperties) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *RightToWorkProperties) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *RightToWorkProperties) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *RightToWorkProperties) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetLastName

`func (o *RightToWorkProperties) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *RightToWorkProperties) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *RightToWorkProperties) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *RightToWorkProperties) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetIssuingCountry

`func (o *RightToWorkProperties) GetIssuingCountry() string`

GetIssuingCountry returns the IssuingCountry field if non-nil, zero value otherwise.

### GetIssuingCountryOk

`func (o *RightToWorkProperties) GetIssuingCountryOk() (*string, bool)`

GetIssuingCountryOk returns a tuple with the IssuingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCountry

`func (o *RightToWorkProperties) SetIssuingCountry(v string)`

SetIssuingCountry sets IssuingCountry field to given value.

### HasIssuingCountry

`func (o *RightToWorkProperties) HasIssuingCountry() bool`

HasIssuingCountry returns a boolean if a field has been set.

### GetGender

`func (o *RightToWorkProperties) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *RightToWorkProperties) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *RightToWorkProperties) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *RightToWorkProperties) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetFirstName

`func (o *RightToWorkProperties) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *RightToWorkProperties) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *RightToWorkProperties) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *RightToWorkProperties) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetDocumentType

`func (o *RightToWorkProperties) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *RightToWorkProperties) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *RightToWorkProperties) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *RightToWorkProperties) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetDocumentNumbers

`func (o *RightToWorkProperties) GetDocumentNumbers() []DocumentPropertiesDocumentNumbersInner`

GetDocumentNumbers returns the DocumentNumbers field if non-nil, zero value otherwise.

### GetDocumentNumbersOk

`func (o *RightToWorkProperties) GetDocumentNumbersOk() (*[]DocumentPropertiesDocumentNumbersInner, bool)`

GetDocumentNumbersOk returns a tuple with the DocumentNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumbers

`func (o *RightToWorkProperties) SetDocumentNumbers(v []DocumentPropertiesDocumentNumbersInner)`

SetDocumentNumbers sets DocumentNumbers field to given value.

### HasDocumentNumbers

`func (o *RightToWorkProperties) HasDocumentNumbers() bool`

HasDocumentNumbers returns a boolean if a field has been set.

### GetDateOfExpiry

`func (o *RightToWorkProperties) GetDateOfExpiry() string`

GetDateOfExpiry returns the DateOfExpiry field if non-nil, zero value otherwise.

### GetDateOfExpiryOk

`func (o *RightToWorkProperties) GetDateOfExpiryOk() (*string, bool)`

GetDateOfExpiryOk returns a tuple with the DateOfExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfExpiry

`func (o *RightToWorkProperties) SetDateOfExpiry(v string)`

SetDateOfExpiry sets DateOfExpiry field to given value.

### HasDateOfExpiry

`func (o *RightToWorkProperties) HasDateOfExpiry() bool`

HasDateOfExpiry returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *RightToWorkProperties) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *RightToWorkProperties) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *RightToWorkProperties) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *RightToWorkProperties) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetDocumentVersion

`func (o *RightToWorkProperties) GetDocumentVersion() string`

GetDocumentVersion returns the DocumentVersion field if non-nil, zero value otherwise.

### GetDocumentVersionOk

`func (o *RightToWorkProperties) GetDocumentVersionOk() (*string, bool)`

GetDocumentVersionOk returns a tuple with the DocumentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentVersion

`func (o *RightToWorkProperties) SetDocumentVersion(v string)`

SetDocumentVersion sets DocumentVersion field to given value.

### HasDocumentVersion

`func (o *RightToWorkProperties) HasDocumentVersion() bool`

HasDocumentVersion returns a boolean if a field has been set.

### GetIssuingDate

`func (o *RightToWorkProperties) GetIssuingDate() string`

GetIssuingDate returns the IssuingDate field if non-nil, zero value otherwise.

### GetIssuingDateOk

`func (o *RightToWorkProperties) GetIssuingDateOk() (*string, bool)`

GetIssuingDateOk returns a tuple with the IssuingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingDate

`func (o *RightToWorkProperties) SetIssuingDate(v string)`

SetIssuingDate sets IssuingDate field to given value.

### HasIssuingDate

`func (o *RightToWorkProperties) HasIssuingDate() bool`

HasIssuingDate returns a boolean if a field has been set.

### GetMrzLine1

`func (o *RightToWorkProperties) GetMrzLine1() string`

GetMrzLine1 returns the MrzLine1 field if non-nil, zero value otherwise.

### GetMrzLine1Ok

`func (o *RightToWorkProperties) GetMrzLine1Ok() (*string, bool)`

GetMrzLine1Ok returns a tuple with the MrzLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrzLine1

`func (o *RightToWorkProperties) SetMrzLine1(v string)`

SetMrzLine1 sets MrzLine1 field to given value.

### HasMrzLine1

`func (o *RightToWorkProperties) HasMrzLine1() bool`

HasMrzLine1 returns a boolean if a field has been set.

### GetMrzLine2

`func (o *RightToWorkProperties) GetMrzLine2() string`

GetMrzLine2 returns the MrzLine2 field if non-nil, zero value otherwise.

### GetMrzLine2Ok

`func (o *RightToWorkProperties) GetMrzLine2Ok() (*string, bool)`

GetMrzLine2Ok returns a tuple with the MrzLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrzLine2

`func (o *RightToWorkProperties) SetMrzLine2(v string)`

SetMrzLine2 sets MrzLine2 field to given value.

### HasMrzLine2

`func (o *RightToWorkProperties) HasMrzLine2() bool`

HasMrzLine2 returns a boolean if a field has been set.

### GetPlaceOfBirth

`func (o *RightToWorkProperties) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *RightToWorkProperties) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *RightToWorkProperties) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *RightToWorkProperties) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### GetIssuingAuthority

`func (o *RightToWorkProperties) GetIssuingAuthority() string`

GetIssuingAuthority returns the IssuingAuthority field if non-nil, zero value otherwise.

### GetIssuingAuthorityOk

`func (o *RightToWorkProperties) GetIssuingAuthorityOk() (*string, bool)`

GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAuthority

`func (o *RightToWorkProperties) SetIssuingAuthority(v string)`

SetIssuingAuthority sets IssuingAuthority field to given value.

### HasIssuingAuthority

`func (o *RightToWorkProperties) HasIssuingAuthority() bool`

HasIssuingAuthority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


