# KnownFacesPropertiesMatchesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicantId** | Pointer to **string** | The applicant ID of the matched applicant. | [optional] 
**Score** | Pointer to **float32** | A floating point number between 0 and 1 that expresses how similar the two faces are, where 1 is a perfect match. | [optional] 
**MediaId** | Pointer to **string** | The corresponding UUID for the media type. | [optional] 
**MediaType** | Pointer to **string** | The media type (for example live_photos or live_videos). | [optional] 
**Suspected** | Pointer to **bool** | Indicates if match is suspected based on fuzzy name matching feature. Dependent on feature being active for account, otherwise defaults to true. | [optional] 

## Methods

### NewKnownFacesPropertiesMatchesInner

`func NewKnownFacesPropertiesMatchesInner() *KnownFacesPropertiesMatchesInner`

NewKnownFacesPropertiesMatchesInner instantiates a new KnownFacesPropertiesMatchesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnownFacesPropertiesMatchesInnerWithDefaults

`func NewKnownFacesPropertiesMatchesInnerWithDefaults() *KnownFacesPropertiesMatchesInner`

NewKnownFacesPropertiesMatchesInnerWithDefaults instantiates a new KnownFacesPropertiesMatchesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicantId

`func (o *KnownFacesPropertiesMatchesInner) GetApplicantId() string`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *KnownFacesPropertiesMatchesInner) GetApplicantIdOk() (*string, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *KnownFacesPropertiesMatchesInner) SetApplicantId(v string)`

SetApplicantId sets ApplicantId field to given value.

### HasApplicantId

`func (o *KnownFacesPropertiesMatchesInner) HasApplicantId() bool`

HasApplicantId returns a boolean if a field has been set.

### GetScore

`func (o *KnownFacesPropertiesMatchesInner) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *KnownFacesPropertiesMatchesInner) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *KnownFacesPropertiesMatchesInner) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *KnownFacesPropertiesMatchesInner) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetMediaId

`func (o *KnownFacesPropertiesMatchesInner) GetMediaId() string`

GetMediaId returns the MediaId field if non-nil, zero value otherwise.

### GetMediaIdOk

`func (o *KnownFacesPropertiesMatchesInner) GetMediaIdOk() (*string, bool)`

GetMediaIdOk returns a tuple with the MediaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaId

`func (o *KnownFacesPropertiesMatchesInner) SetMediaId(v string)`

SetMediaId sets MediaId field to given value.

### HasMediaId

`func (o *KnownFacesPropertiesMatchesInner) HasMediaId() bool`

HasMediaId returns a boolean if a field has been set.

### GetMediaType

`func (o *KnownFacesPropertiesMatchesInner) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *KnownFacesPropertiesMatchesInner) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *KnownFacesPropertiesMatchesInner) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *KnownFacesPropertiesMatchesInner) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetSuspected

`func (o *KnownFacesPropertiesMatchesInner) GetSuspected() bool`

GetSuspected returns the Suspected field if non-nil, zero value otherwise.

### GetSuspectedOk

`func (o *KnownFacesPropertiesMatchesInner) GetSuspectedOk() (*bool, bool)`

GetSuspectedOk returns a tuple with the Suspected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspected

`func (o *KnownFacesPropertiesMatchesInner) SetSuspected(v bool)`

SetSuspected sets Suspected field to given value.

### HasSuspected

`func (o *KnownFacesPropertiesMatchesInner) HasSuspected() bool`

HasSuspected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


