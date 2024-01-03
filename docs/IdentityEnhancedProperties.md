# IdentityEnhancedProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchedAddress** | Pointer to **float32** | Returns address number which has been matched. | [optional] 
**MatchedAddresses** | Pointer to [**[]IdentityEnhancedPropertiesMatchedAddressesInner**](IdentityEnhancedPropertiesMatchedAddressesInner.md) | Returns array of sources which contain matched addresses for the corresponding address number. | [optional] 

## Methods

### NewIdentityEnhancedProperties

`func NewIdentityEnhancedProperties() *IdentityEnhancedProperties`

NewIdentityEnhancedProperties instantiates a new IdentityEnhancedProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityEnhancedPropertiesWithDefaults

`func NewIdentityEnhancedPropertiesWithDefaults() *IdentityEnhancedProperties`

NewIdentityEnhancedPropertiesWithDefaults instantiates a new IdentityEnhancedProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchedAddress

`func (o *IdentityEnhancedProperties) GetMatchedAddress() float32`

GetMatchedAddress returns the MatchedAddress field if non-nil, zero value otherwise.

### GetMatchedAddressOk

`func (o *IdentityEnhancedProperties) GetMatchedAddressOk() (*float32, bool)`

GetMatchedAddressOk returns a tuple with the MatchedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedAddress

`func (o *IdentityEnhancedProperties) SetMatchedAddress(v float32)`

SetMatchedAddress sets MatchedAddress field to given value.

### HasMatchedAddress

`func (o *IdentityEnhancedProperties) HasMatchedAddress() bool`

HasMatchedAddress returns a boolean if a field has been set.

### GetMatchedAddresses

`func (o *IdentityEnhancedProperties) GetMatchedAddresses() []IdentityEnhancedPropertiesMatchedAddressesInner`

GetMatchedAddresses returns the MatchedAddresses field if non-nil, zero value otherwise.

### GetMatchedAddressesOk

`func (o *IdentityEnhancedProperties) GetMatchedAddressesOk() (*[]IdentityEnhancedPropertiesMatchedAddressesInner, bool)`

GetMatchedAddressesOk returns a tuple with the MatchedAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedAddresses

`func (o *IdentityEnhancedProperties) SetMatchedAddresses(v []IdentityEnhancedPropertiesMatchedAddressesInner)`

SetMatchedAddresses sets MatchedAddresses field to given value.

### HasMatchedAddresses

`func (o *IdentityEnhancedProperties) HasMatchedAddresses() bool`

HasMatchedAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


