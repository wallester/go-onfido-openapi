# DocumentCDQReasons

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObscuredDataPoints** | Pointer to **string** | When data points are obscured to the point that we cannot confirm if the fonts match the expected ones. | [optional] 
**ObscuredSecurityFeatures** | Pointer to **string** | When a critical security feature is obscured. This can also refer to when the holder&#39;s wet signature, necessary for the document to be valid, is not present. | [optional] 
**AbnormalDocumentFeatures** | Pointer to **string** | One of 3 reasons (1) OCR Assisted Scans (i.e. when you&#39;re able to move the mouse and highlight part of text), (2) Severely Washed out Background, (3) Overlapping Text. | [optional] 
**WatermarksDigitalTextOverlay** | Pointer to **string** | Any digital text or electronic watermarks on the document. | [optional] 
**CornerRemoved** | Pointer to **string** | If the corner has been physically cut off. This can be found on some documents that are no longer valid. | [optional] 
**PuncturedDocument** | Pointer to **string** | A punched hole is present. | [optional] 
**MissingBack** | Pointer to **string** | When the back of a document is needed for processing, but is not available. | [optional] 
**DigitalDocument** | Pointer to **string** | When a document has been published digitally, there aren’t enough security features to review so we cannot perform a full fraud assessment. | [optional] 

## Methods

### NewDocumentCDQReasons

`func NewDocumentCDQReasons() *DocumentCDQReasons`

NewDocumentCDQReasons instantiates a new DocumentCDQReasons object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentCDQReasonsWithDefaults

`func NewDocumentCDQReasonsWithDefaults() *DocumentCDQReasons`

NewDocumentCDQReasonsWithDefaults instantiates a new DocumentCDQReasons object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObscuredDataPoints

`func (o *DocumentCDQReasons) GetObscuredDataPoints() string`

GetObscuredDataPoints returns the ObscuredDataPoints field if non-nil, zero value otherwise.

### GetObscuredDataPointsOk

`func (o *DocumentCDQReasons) GetObscuredDataPointsOk() (*string, bool)`

GetObscuredDataPointsOk returns a tuple with the ObscuredDataPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObscuredDataPoints

`func (o *DocumentCDQReasons) SetObscuredDataPoints(v string)`

SetObscuredDataPoints sets ObscuredDataPoints field to given value.

### HasObscuredDataPoints

`func (o *DocumentCDQReasons) HasObscuredDataPoints() bool`

HasObscuredDataPoints returns a boolean if a field has been set.

### GetObscuredSecurityFeatures

`func (o *DocumentCDQReasons) GetObscuredSecurityFeatures() string`

GetObscuredSecurityFeatures returns the ObscuredSecurityFeatures field if non-nil, zero value otherwise.

### GetObscuredSecurityFeaturesOk

`func (o *DocumentCDQReasons) GetObscuredSecurityFeaturesOk() (*string, bool)`

GetObscuredSecurityFeaturesOk returns a tuple with the ObscuredSecurityFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObscuredSecurityFeatures

`func (o *DocumentCDQReasons) SetObscuredSecurityFeatures(v string)`

SetObscuredSecurityFeatures sets ObscuredSecurityFeatures field to given value.

### HasObscuredSecurityFeatures

`func (o *DocumentCDQReasons) HasObscuredSecurityFeatures() bool`

HasObscuredSecurityFeatures returns a boolean if a field has been set.

### GetAbnormalDocumentFeatures

`func (o *DocumentCDQReasons) GetAbnormalDocumentFeatures() string`

GetAbnormalDocumentFeatures returns the AbnormalDocumentFeatures field if non-nil, zero value otherwise.

### GetAbnormalDocumentFeaturesOk

`func (o *DocumentCDQReasons) GetAbnormalDocumentFeaturesOk() (*string, bool)`

GetAbnormalDocumentFeaturesOk returns a tuple with the AbnormalDocumentFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbnormalDocumentFeatures

`func (o *DocumentCDQReasons) SetAbnormalDocumentFeatures(v string)`

SetAbnormalDocumentFeatures sets AbnormalDocumentFeatures field to given value.

### HasAbnormalDocumentFeatures

`func (o *DocumentCDQReasons) HasAbnormalDocumentFeatures() bool`

HasAbnormalDocumentFeatures returns a boolean if a field has been set.

### GetWatermarksDigitalTextOverlay

`func (o *DocumentCDQReasons) GetWatermarksDigitalTextOverlay() string`

GetWatermarksDigitalTextOverlay returns the WatermarksDigitalTextOverlay field if non-nil, zero value otherwise.

### GetWatermarksDigitalTextOverlayOk

`func (o *DocumentCDQReasons) GetWatermarksDigitalTextOverlayOk() (*string, bool)`

GetWatermarksDigitalTextOverlayOk returns a tuple with the WatermarksDigitalTextOverlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermarksDigitalTextOverlay

`func (o *DocumentCDQReasons) SetWatermarksDigitalTextOverlay(v string)`

SetWatermarksDigitalTextOverlay sets WatermarksDigitalTextOverlay field to given value.

### HasWatermarksDigitalTextOverlay

`func (o *DocumentCDQReasons) HasWatermarksDigitalTextOverlay() bool`

HasWatermarksDigitalTextOverlay returns a boolean if a field has been set.

### GetCornerRemoved

`func (o *DocumentCDQReasons) GetCornerRemoved() string`

GetCornerRemoved returns the CornerRemoved field if non-nil, zero value otherwise.

### GetCornerRemovedOk

`func (o *DocumentCDQReasons) GetCornerRemovedOk() (*string, bool)`

GetCornerRemovedOk returns a tuple with the CornerRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCornerRemoved

`func (o *DocumentCDQReasons) SetCornerRemoved(v string)`

SetCornerRemoved sets CornerRemoved field to given value.

### HasCornerRemoved

`func (o *DocumentCDQReasons) HasCornerRemoved() bool`

HasCornerRemoved returns a boolean if a field has been set.

### GetPuncturedDocument

`func (o *DocumentCDQReasons) GetPuncturedDocument() string`

GetPuncturedDocument returns the PuncturedDocument field if non-nil, zero value otherwise.

### GetPuncturedDocumentOk

`func (o *DocumentCDQReasons) GetPuncturedDocumentOk() (*string, bool)`

GetPuncturedDocumentOk returns a tuple with the PuncturedDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuncturedDocument

`func (o *DocumentCDQReasons) SetPuncturedDocument(v string)`

SetPuncturedDocument sets PuncturedDocument field to given value.

### HasPuncturedDocument

`func (o *DocumentCDQReasons) HasPuncturedDocument() bool`

HasPuncturedDocument returns a boolean if a field has been set.

### GetMissingBack

`func (o *DocumentCDQReasons) GetMissingBack() string`

GetMissingBack returns the MissingBack field if non-nil, zero value otherwise.

### GetMissingBackOk

`func (o *DocumentCDQReasons) GetMissingBackOk() (*string, bool)`

GetMissingBackOk returns a tuple with the MissingBack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingBack

`func (o *DocumentCDQReasons) SetMissingBack(v string)`

SetMissingBack sets MissingBack field to given value.

### HasMissingBack

`func (o *DocumentCDQReasons) HasMissingBack() bool`

HasMissingBack returns a boolean if a field has been set.

### GetDigitalDocument

`func (o *DocumentCDQReasons) GetDigitalDocument() string`

GetDigitalDocument returns the DigitalDocument field if non-nil, zero value otherwise.

### GetDigitalDocumentOk

`func (o *DocumentCDQReasons) GetDigitalDocumentOk() (*string, bool)`

GetDigitalDocumentOk returns a tuple with the DigitalDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalDocument

`func (o *DocumentCDQReasons) SetDigitalDocument(v string)`

SetDigitalDocument sets DigitalDocument field to given value.

### HasDigitalDocument

`func (o *DocumentCDQReasons) HasDigitalDocument() bool`

HasDigitalDocument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


