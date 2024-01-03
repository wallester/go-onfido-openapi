# ProofOfAddressProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | This property provides the address on the document. | [optional] 
**DocumentType** | Pointer to **string** | This property provides the document type according to the set of supported documents. | [optional] 
**FirstNames** | Pointer to **string** | This property provides the first names on the document, including any initials and middle names. | [optional] 
**LastNames** | Pointer to **string** | This property provided the last names on the document. | [optional] 
**IssueDate** | Pointer to **string** | This property provides the issue date of the document. | [optional] 
**Issuer** | Pointer to **string** | This property provides the document issuer (e.g. HSBC, British Gas). | [optional] 
**SummaryPeriodStart** | Pointer to **string** | This property provides the summary period start date. | [optional] 
**SummaryPeriodEnd** | Pointer to **string** | This property provides the summary period end date. | [optional] 

## Methods

### NewProofOfAddressProperties

`func NewProofOfAddressProperties() *ProofOfAddressProperties`

NewProofOfAddressProperties instantiates a new ProofOfAddressProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProofOfAddressPropertiesWithDefaults

`func NewProofOfAddressPropertiesWithDefaults() *ProofOfAddressProperties`

NewProofOfAddressPropertiesWithDefaults instantiates a new ProofOfAddressProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ProofOfAddressProperties) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ProofOfAddressProperties) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ProofOfAddressProperties) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ProofOfAddressProperties) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDocumentType

`func (o *ProofOfAddressProperties) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *ProofOfAddressProperties) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *ProofOfAddressProperties) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *ProofOfAddressProperties) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetFirstNames

`func (o *ProofOfAddressProperties) GetFirstNames() string`

GetFirstNames returns the FirstNames field if non-nil, zero value otherwise.

### GetFirstNamesOk

`func (o *ProofOfAddressProperties) GetFirstNamesOk() (*string, bool)`

GetFirstNamesOk returns a tuple with the FirstNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNames

`func (o *ProofOfAddressProperties) SetFirstNames(v string)`

SetFirstNames sets FirstNames field to given value.

### HasFirstNames

`func (o *ProofOfAddressProperties) HasFirstNames() bool`

HasFirstNames returns a boolean if a field has been set.

### GetLastNames

`func (o *ProofOfAddressProperties) GetLastNames() string`

GetLastNames returns the LastNames field if non-nil, zero value otherwise.

### GetLastNamesOk

`func (o *ProofOfAddressProperties) GetLastNamesOk() (*string, bool)`

GetLastNamesOk returns a tuple with the LastNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNames

`func (o *ProofOfAddressProperties) SetLastNames(v string)`

SetLastNames sets LastNames field to given value.

### HasLastNames

`func (o *ProofOfAddressProperties) HasLastNames() bool`

HasLastNames returns a boolean if a field has been set.

### GetIssueDate

`func (o *ProofOfAddressProperties) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *ProofOfAddressProperties) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *ProofOfAddressProperties) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *ProofOfAddressProperties) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetIssuer

`func (o *ProofOfAddressProperties) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ProofOfAddressProperties) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ProofOfAddressProperties) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *ProofOfAddressProperties) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSummaryPeriodStart

`func (o *ProofOfAddressProperties) GetSummaryPeriodStart() string`

GetSummaryPeriodStart returns the SummaryPeriodStart field if non-nil, zero value otherwise.

### GetSummaryPeriodStartOk

`func (o *ProofOfAddressProperties) GetSummaryPeriodStartOk() (*string, bool)`

GetSummaryPeriodStartOk returns a tuple with the SummaryPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryPeriodStart

`func (o *ProofOfAddressProperties) SetSummaryPeriodStart(v string)`

SetSummaryPeriodStart sets SummaryPeriodStart field to given value.

### HasSummaryPeriodStart

`func (o *ProofOfAddressProperties) HasSummaryPeriodStart() bool`

HasSummaryPeriodStart returns a boolean if a field has been set.

### GetSummaryPeriodEnd

`func (o *ProofOfAddressProperties) GetSummaryPeriodEnd() string`

GetSummaryPeriodEnd returns the SummaryPeriodEnd field if non-nil, zero value otherwise.

### GetSummaryPeriodEndOk

`func (o *ProofOfAddressProperties) GetSummaryPeriodEndOk() (*string, bool)`

GetSummaryPeriodEndOk returns a tuple with the SummaryPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryPeriodEnd

`func (o *ProofOfAddressProperties) SetSummaryPeriodEnd(v string)`

SetSummaryPeriodEnd sets SummaryPeriodEnd field to given value.

### HasSummaryPeriodEnd

`func (o *ProofOfAddressProperties) HasSummaryPeriodEnd() bool`

HasSummaryPeriodEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


