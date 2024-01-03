# FacialSimilarityPhotoBreakdownFaceComparisonBreakdownFaceMatchProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **float32** | A floating point number between 0 and 1 that expresses how similar the two faces are, where 1 is a perfect match. | [optional] 
**DocumentId** | Pointer to **string** | The UUID for the document containing the extracted face that was used for face matching. | [optional] 

## Methods

### NewFacialSimilarityPhotoBreakdownFaceComparisonBreakdownFaceMatchProperties

`func NewFacialSimilarityPhotoBreakdownFaceComparisonBreakdownFaceMatchProperties() *FacialSimilarityPhotoBreakdownFaceComparisonBreakdownFaceMatchProperties`

NewFacialSimilarityPhotoBreakdownFaceComparisonBreakdownFaceMatchProperties instantiates a new FacialSimilarityPhotoBreakdownFaceComparisonBreakdownFaceMatchProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacialSimilarityPhotoBreakdownFaceComparisonBreakdownFaceMatchPropertiesWithDefaults

`func NewFacialSimilarityPhotoBreakdownFaceComparisonBreakdownFaceMatchPropertiesWithDefaults() *FacialSimilarityPhotoBreakdownFaceComparisonBreakdownFaceMatchProperties`

NewFacialSimilarityPhotoBreakdownFaceComparisonBreakdownFaceMatchPropertiesWithDefaults instantiates a new FacialSimilarityPhotoBreakdownFaceComparisonBreakdownFaceMatchProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *FacialSimilarityPhotoBreakdownFaceComparisonBreakdownFaceMatchProperties) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *FacialSimilarityPhotoBreakdownFaceComparisonBreakdownFaceMatchProperties) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *FacialSimilarityPhotoBreakdownFaceComparisonBreakdownFaceMatchProperties) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *FacialSimilarityPhotoBreakdownFaceComparisonBreakdownFaceMatchProperties) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetDocumentId

`func (o *FacialSimilarityPhotoBreakdownFaceComparisonBreakdownFaceMatchProperties) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *FacialSimilarityPhotoBreakdownFaceComparisonBreakdownFaceMatchProperties) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *FacialSimilarityPhotoBreakdownFaceComparisonBreakdownFaceMatchProperties) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *FacialSimilarityPhotoBreakdownFaceComparisonBreakdownFaceMatchProperties) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


