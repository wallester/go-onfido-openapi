# DocumentIQReasons

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DarkPhoto** | Pointer to **string** | When an image of the document is too dark to be able to see data points. | [optional] 
**GlareOnPhoto** | Pointer to **string** | When there is light reflecting on the document causing glare to obstruct data points. | [optional] 
**BlurredPhoto** | Pointer to **string** | When data points are blurred and no reference can be made elsewhere in the document or if the data points are too blurry and &#39;they could be something else&#39;. | [optional] 
**CoveredPhoto** | Pointer to **string** | When data points have been covered either by the applicant or by another object such as a sticker. | [optional] 
**OtherPhotoIssue** | Pointer to **string** | Any other reason not listed, such as when holograms are obscuring data points. | [optional] 
**DamagedDocument** | Pointer to **string** | When a document is damaged and we are unable to make out data points. | [optional] 
**IncorrectSide** | Pointer to **string** | When the incorrect side of a document has been uploaded, and we have not received the front. | [optional] 
**CutOffDocument** | Pointer to **string** | When data points are not included in the image due to the document being cut off. | [optional] 
**NoDocumentInImage** | Pointer to **string** | If no document has been uploaded or there is a blank image. | [optional] 
**TwoDocumentsUploaded** | Pointer to **string** | When 2 different documents are submitted in the same check. | [optional] 

## Methods

### NewDocumentIQReasons

`func NewDocumentIQReasons() *DocumentIQReasons`

NewDocumentIQReasons instantiates a new DocumentIQReasons object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentIQReasonsWithDefaults

`func NewDocumentIQReasonsWithDefaults() *DocumentIQReasons`

NewDocumentIQReasonsWithDefaults instantiates a new DocumentIQReasons object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDarkPhoto

`func (o *DocumentIQReasons) GetDarkPhoto() string`

GetDarkPhoto returns the DarkPhoto field if non-nil, zero value otherwise.

### GetDarkPhotoOk

`func (o *DocumentIQReasons) GetDarkPhotoOk() (*string, bool)`

GetDarkPhotoOk returns a tuple with the DarkPhoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkPhoto

`func (o *DocumentIQReasons) SetDarkPhoto(v string)`

SetDarkPhoto sets DarkPhoto field to given value.

### HasDarkPhoto

`func (o *DocumentIQReasons) HasDarkPhoto() bool`

HasDarkPhoto returns a boolean if a field has been set.

### GetGlareOnPhoto

`func (o *DocumentIQReasons) GetGlareOnPhoto() string`

GetGlareOnPhoto returns the GlareOnPhoto field if non-nil, zero value otherwise.

### GetGlareOnPhotoOk

`func (o *DocumentIQReasons) GetGlareOnPhotoOk() (*string, bool)`

GetGlareOnPhotoOk returns a tuple with the GlareOnPhoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlareOnPhoto

`func (o *DocumentIQReasons) SetGlareOnPhoto(v string)`

SetGlareOnPhoto sets GlareOnPhoto field to given value.

### HasGlareOnPhoto

`func (o *DocumentIQReasons) HasGlareOnPhoto() bool`

HasGlareOnPhoto returns a boolean if a field has been set.

### GetBlurredPhoto

`func (o *DocumentIQReasons) GetBlurredPhoto() string`

GetBlurredPhoto returns the BlurredPhoto field if non-nil, zero value otherwise.

### GetBlurredPhotoOk

`func (o *DocumentIQReasons) GetBlurredPhotoOk() (*string, bool)`

GetBlurredPhotoOk returns a tuple with the BlurredPhoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlurredPhoto

`func (o *DocumentIQReasons) SetBlurredPhoto(v string)`

SetBlurredPhoto sets BlurredPhoto field to given value.

### HasBlurredPhoto

`func (o *DocumentIQReasons) HasBlurredPhoto() bool`

HasBlurredPhoto returns a boolean if a field has been set.

### GetCoveredPhoto

`func (o *DocumentIQReasons) GetCoveredPhoto() string`

GetCoveredPhoto returns the CoveredPhoto field if non-nil, zero value otherwise.

### GetCoveredPhotoOk

`func (o *DocumentIQReasons) GetCoveredPhotoOk() (*string, bool)`

GetCoveredPhotoOk returns a tuple with the CoveredPhoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoveredPhoto

`func (o *DocumentIQReasons) SetCoveredPhoto(v string)`

SetCoveredPhoto sets CoveredPhoto field to given value.

### HasCoveredPhoto

`func (o *DocumentIQReasons) HasCoveredPhoto() bool`

HasCoveredPhoto returns a boolean if a field has been set.

### GetOtherPhotoIssue

`func (o *DocumentIQReasons) GetOtherPhotoIssue() string`

GetOtherPhotoIssue returns the OtherPhotoIssue field if non-nil, zero value otherwise.

### GetOtherPhotoIssueOk

`func (o *DocumentIQReasons) GetOtherPhotoIssueOk() (*string, bool)`

GetOtherPhotoIssueOk returns a tuple with the OtherPhotoIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherPhotoIssue

`func (o *DocumentIQReasons) SetOtherPhotoIssue(v string)`

SetOtherPhotoIssue sets OtherPhotoIssue field to given value.

### HasOtherPhotoIssue

`func (o *DocumentIQReasons) HasOtherPhotoIssue() bool`

HasOtherPhotoIssue returns a boolean if a field has been set.

### GetDamagedDocument

`func (o *DocumentIQReasons) GetDamagedDocument() string`

GetDamagedDocument returns the DamagedDocument field if non-nil, zero value otherwise.

### GetDamagedDocumentOk

`func (o *DocumentIQReasons) GetDamagedDocumentOk() (*string, bool)`

GetDamagedDocumentOk returns a tuple with the DamagedDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamagedDocument

`func (o *DocumentIQReasons) SetDamagedDocument(v string)`

SetDamagedDocument sets DamagedDocument field to given value.

### HasDamagedDocument

`func (o *DocumentIQReasons) HasDamagedDocument() bool`

HasDamagedDocument returns a boolean if a field has been set.

### GetIncorrectSide

`func (o *DocumentIQReasons) GetIncorrectSide() string`

GetIncorrectSide returns the IncorrectSide field if non-nil, zero value otherwise.

### GetIncorrectSideOk

`func (o *DocumentIQReasons) GetIncorrectSideOk() (*string, bool)`

GetIncorrectSideOk returns a tuple with the IncorrectSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncorrectSide

`func (o *DocumentIQReasons) SetIncorrectSide(v string)`

SetIncorrectSide sets IncorrectSide field to given value.

### HasIncorrectSide

`func (o *DocumentIQReasons) HasIncorrectSide() bool`

HasIncorrectSide returns a boolean if a field has been set.

### GetCutOffDocument

`func (o *DocumentIQReasons) GetCutOffDocument() string`

GetCutOffDocument returns the CutOffDocument field if non-nil, zero value otherwise.

### GetCutOffDocumentOk

`func (o *DocumentIQReasons) GetCutOffDocumentOk() (*string, bool)`

GetCutOffDocumentOk returns a tuple with the CutOffDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutOffDocument

`func (o *DocumentIQReasons) SetCutOffDocument(v string)`

SetCutOffDocument sets CutOffDocument field to given value.

### HasCutOffDocument

`func (o *DocumentIQReasons) HasCutOffDocument() bool`

HasCutOffDocument returns a boolean if a field has been set.

### GetNoDocumentInImage

`func (o *DocumentIQReasons) GetNoDocumentInImage() string`

GetNoDocumentInImage returns the NoDocumentInImage field if non-nil, zero value otherwise.

### GetNoDocumentInImageOk

`func (o *DocumentIQReasons) GetNoDocumentInImageOk() (*string, bool)`

GetNoDocumentInImageOk returns a tuple with the NoDocumentInImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDocumentInImage

`func (o *DocumentIQReasons) SetNoDocumentInImage(v string)`

SetNoDocumentInImage sets NoDocumentInImage field to given value.

### HasNoDocumentInImage

`func (o *DocumentIQReasons) HasNoDocumentInImage() bool`

HasNoDocumentInImage returns a boolean if a field has been set.

### GetTwoDocumentsUploaded

`func (o *DocumentIQReasons) GetTwoDocumentsUploaded() string`

GetTwoDocumentsUploaded returns the TwoDocumentsUploaded field if non-nil, zero value otherwise.

### GetTwoDocumentsUploadedOk

`func (o *DocumentIQReasons) GetTwoDocumentsUploadedOk() (*string, bool)`

GetTwoDocumentsUploadedOk returns a tuple with the TwoDocumentsUploaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoDocumentsUploaded

`func (o *DocumentIQReasons) SetTwoDocumentsUploaded(v string)`

SetTwoDocumentsUploaded sets TwoDocumentsUploaded field to given value.

### HasTwoDocumentsUploaded

`func (o *DocumentIQReasons) HasTwoDocumentsUploaded() bool`

HasTwoDocumentsUploaded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


