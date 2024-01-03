# VideoReasons

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FakeWebcam** | Pointer to **string** | Flags when evidence is found that a fake webcam was used. | [optional] 
**ChallengeReuse** | Pointer to **string** | Flags when evidence is found that the video was uploaded in an attempt to circumvent the randomness of the speaking and head turn challenges | [optional] 
**Emulator** | Pointer to **string** | Flags when evidence is found that an Android emulator was used. | [optional] 
**Reasons** | Pointer to **string** | Additional comma separated details such as the name of the fake webcam. | [optional] 

## Methods

### NewVideoReasons

`func NewVideoReasons() *VideoReasons`

NewVideoReasons instantiates a new VideoReasons object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoReasonsWithDefaults

`func NewVideoReasonsWithDefaults() *VideoReasons`

NewVideoReasonsWithDefaults instantiates a new VideoReasons object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFakeWebcam

`func (o *VideoReasons) GetFakeWebcam() string`

GetFakeWebcam returns the FakeWebcam field if non-nil, zero value otherwise.

### GetFakeWebcamOk

`func (o *VideoReasons) GetFakeWebcamOk() (*string, bool)`

GetFakeWebcamOk returns a tuple with the FakeWebcam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFakeWebcam

`func (o *VideoReasons) SetFakeWebcam(v string)`

SetFakeWebcam sets FakeWebcam field to given value.

### HasFakeWebcam

`func (o *VideoReasons) HasFakeWebcam() bool`

HasFakeWebcam returns a boolean if a field has been set.

### GetChallengeReuse

`func (o *VideoReasons) GetChallengeReuse() string`

GetChallengeReuse returns the ChallengeReuse field if non-nil, zero value otherwise.

### GetChallengeReuseOk

`func (o *VideoReasons) GetChallengeReuseOk() (*string, bool)`

GetChallengeReuseOk returns a tuple with the ChallengeReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeReuse

`func (o *VideoReasons) SetChallengeReuse(v string)`

SetChallengeReuse sets ChallengeReuse field to given value.

### HasChallengeReuse

`func (o *VideoReasons) HasChallengeReuse() bool`

HasChallengeReuse returns a boolean if a field has been set.

### GetEmulator

`func (o *VideoReasons) GetEmulator() string`

GetEmulator returns the Emulator field if non-nil, zero value otherwise.

### GetEmulatorOk

`func (o *VideoReasons) GetEmulatorOk() (*string, bool)`

GetEmulatorOk returns a tuple with the Emulator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmulator

`func (o *VideoReasons) SetEmulator(v string)`

SetEmulator sets Emulator field to given value.

### HasEmulator

`func (o *VideoReasons) HasEmulator() bool`

HasEmulator returns a boolean if a field has been set.

### GetReasons

`func (o *VideoReasons) GetReasons() string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *VideoReasons) GetReasonsOk() (*string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *VideoReasons) SetReasons(v string)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *VideoReasons) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


