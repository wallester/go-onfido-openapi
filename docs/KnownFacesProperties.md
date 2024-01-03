# KnownFacesProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Matches** | Pointer to [**[]KnownFacesPropertiesMatchesInner**](KnownFacesPropertiesMatchesInner.md) | Returns any matching applicant IDs as entries inside a matches array under a properties bag. | [optional] 

## Methods

### NewKnownFacesProperties

`func NewKnownFacesProperties() *KnownFacesProperties`

NewKnownFacesProperties instantiates a new KnownFacesProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnownFacesPropertiesWithDefaults

`func NewKnownFacesPropertiesWithDefaults() *KnownFacesProperties`

NewKnownFacesPropertiesWithDefaults instantiates a new KnownFacesProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatches

`func (o *KnownFacesProperties) GetMatches() []KnownFacesPropertiesMatchesInner`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *KnownFacesProperties) GetMatchesOk() (*[]KnownFacesPropertiesMatchesInner, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *KnownFacesProperties) SetMatches(v []KnownFacesPropertiesMatchesInner)`

SetMatches sets Matches field to given value.

### HasMatches

`func (o *KnownFacesProperties) HasMatches() bool`

HasMatches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


