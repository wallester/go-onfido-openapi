# PhotoReasons

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DigitalTampering** | Pointer to **string** | Flags when evidence is found that the image was manipulated by Photoshop, or other software. | [optional] 
**FakeWebcam** | Pointer to **string** | Flags when evidence is found that a fake webcam was used. | [optional] 
**TimeOfCapture** | Pointer to **string** | Flags when evidence is found that the live photo was taken more than 24 hours before live photo upload. | [optional] 
**Emulator** | Pointer to **string** | Flags when evidence is found that an Android emulator was used. | [optional] 
**Reasons** | Pointer to **string** | Additional comma separated details such as the exact digital tampering software used, or the name of the fake webcam. | [optional] 

## Methods

### NewPhotoReasons

`func NewPhotoReasons() *PhotoReasons`

NewPhotoReasons instantiates a new PhotoReasons object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhotoReasonsWithDefaults

`func NewPhotoReasonsWithDefaults() *PhotoReasons`

NewPhotoReasonsWithDefaults instantiates a new PhotoReasons object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigitalTampering

`func (o *PhotoReasons) GetDigitalTampering() string`

GetDigitalTampering returns the DigitalTampering field if non-nil, zero value otherwise.

### GetDigitalTamperingOk

`func (o *PhotoReasons) GetDigitalTamperingOk() (*string, bool)`

GetDigitalTamperingOk returns a tuple with the DigitalTampering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalTampering

`func (o *PhotoReasons) SetDigitalTampering(v string)`

SetDigitalTampering sets DigitalTampering field to given value.

### HasDigitalTampering

`func (o *PhotoReasons) HasDigitalTampering() bool`

HasDigitalTampering returns a boolean if a field has been set.

### GetFakeWebcam

`func (o *PhotoReasons) GetFakeWebcam() string`

GetFakeWebcam returns the FakeWebcam field if non-nil, zero value otherwise.

### GetFakeWebcamOk

`func (o *PhotoReasons) GetFakeWebcamOk() (*string, bool)`

GetFakeWebcamOk returns a tuple with the FakeWebcam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFakeWebcam

`func (o *PhotoReasons) SetFakeWebcam(v string)`

SetFakeWebcam sets FakeWebcam field to given value.

### HasFakeWebcam

`func (o *PhotoReasons) HasFakeWebcam() bool`

HasFakeWebcam returns a boolean if a field has been set.

### GetTimeOfCapture

`func (o *PhotoReasons) GetTimeOfCapture() string`

GetTimeOfCapture returns the TimeOfCapture field if non-nil, zero value otherwise.

### GetTimeOfCaptureOk

`func (o *PhotoReasons) GetTimeOfCaptureOk() (*string, bool)`

GetTimeOfCaptureOk returns a tuple with the TimeOfCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfCapture

`func (o *PhotoReasons) SetTimeOfCapture(v string)`

SetTimeOfCapture sets TimeOfCapture field to given value.

### HasTimeOfCapture

`func (o *PhotoReasons) HasTimeOfCapture() bool`

HasTimeOfCapture returns a boolean if a field has been set.

### GetEmulator

`func (o *PhotoReasons) GetEmulator() string`

GetEmulator returns the Emulator field if non-nil, zero value otherwise.

### GetEmulatorOk

`func (o *PhotoReasons) GetEmulatorOk() (*string, bool)`

GetEmulatorOk returns a tuple with the Emulator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmulator

`func (o *PhotoReasons) SetEmulator(v string)`

SetEmulator sets Emulator field to given value.

### HasEmulator

`func (o *PhotoReasons) HasEmulator() bool`

HasEmulator returns a boolean if a field has been set.

### GetReasons

`func (o *PhotoReasons) GetReasons() string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *PhotoReasons) GetReasonsOk() (*string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *PhotoReasons) SetReasons(v string)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *PhotoReasons) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


