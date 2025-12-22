# DeviceIntelligencePropertiesDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SdkVersion** | Pointer to **string** | The SDK version that was used. | [optional] 
**SdkSource** | Pointer to **string** | The SDK used to upload the media. | [optional] 
**AuthenticationType** | Pointer to **string** | The token used to authenticate the request. | [optional] 
**RawModel** | Pointer to **string** | The model as set by the phone manufacturer (for Android and iOS) or the browser manufacturer (for Web). | [optional] 
**Os** | Pointer to **string** | The operating system of the device. | [optional] 
**Browser** | Pointer to **string** | The browser name reported by the browser&#39;s user agent. | [optional] 
**Emulator** | Pointer to **bool** | Whether the device is an emulator. | [optional] 
**RandomizedDevice** | Pointer to **bool** | Whether the device is providing false randomized device and network information. | [optional] 
**FakeNetworkRequest** | Pointer to **bool** | Whether device is using stolen security tokens to send the network information. | [optional] 
**IpReputation** | Pointer to **string** | Whether there is highly suspicious traffic related to the IP address. | [optional] 
**DeviceFingerprintReuse** | Pointer to **float32** | The number of times the device was used to create a report for a new applicant. | [optional] 
**SingleDeviceUsed** | Pointer to **NullableBool** | Whether the document or biometric media were uploaded from a single device. | [optional] 
**DocumentCapture** | Pointer to **string** | Whether the document media were live captured from the device camera. | [optional] 
**BiometricCapture** | Pointer to **string** | Whether the biometric media were live captured from the device camera. | [optional] 

## Methods

### NewDeviceIntelligencePropertiesDevice

`func NewDeviceIntelligencePropertiesDevice() *DeviceIntelligencePropertiesDevice`

NewDeviceIntelligencePropertiesDevice instantiates a new DeviceIntelligencePropertiesDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceIntelligencePropertiesDeviceWithDefaults

`func NewDeviceIntelligencePropertiesDeviceWithDefaults() *DeviceIntelligencePropertiesDevice`

NewDeviceIntelligencePropertiesDeviceWithDefaults instantiates a new DeviceIntelligencePropertiesDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSdkVersion

`func (o *DeviceIntelligencePropertiesDevice) GetSdkVersion() string`

GetSdkVersion returns the SdkVersion field if non-nil, zero value otherwise.

### GetSdkVersionOk

`func (o *DeviceIntelligencePropertiesDevice) GetSdkVersionOk() (*string, bool)`

GetSdkVersionOk returns a tuple with the SdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkVersion

`func (o *DeviceIntelligencePropertiesDevice) SetSdkVersion(v string)`

SetSdkVersion sets SdkVersion field to given value.

### HasSdkVersion

`func (o *DeviceIntelligencePropertiesDevice) HasSdkVersion() bool`

HasSdkVersion returns a boolean if a field has been set.

### GetSdkSource

`func (o *DeviceIntelligencePropertiesDevice) GetSdkSource() string`

GetSdkSource returns the SdkSource field if non-nil, zero value otherwise.

### GetSdkSourceOk

`func (o *DeviceIntelligencePropertiesDevice) GetSdkSourceOk() (*string, bool)`

GetSdkSourceOk returns a tuple with the SdkSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkSource

`func (o *DeviceIntelligencePropertiesDevice) SetSdkSource(v string)`

SetSdkSource sets SdkSource field to given value.

### HasSdkSource

`func (o *DeviceIntelligencePropertiesDevice) HasSdkSource() bool`

HasSdkSource returns a boolean if a field has been set.

### GetAuthenticationType

`func (o *DeviceIntelligencePropertiesDevice) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *DeviceIntelligencePropertiesDevice) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *DeviceIntelligencePropertiesDevice) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *DeviceIntelligencePropertiesDevice) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetRawModel

`func (o *DeviceIntelligencePropertiesDevice) GetRawModel() string`

GetRawModel returns the RawModel field if non-nil, zero value otherwise.

### GetRawModelOk

`func (o *DeviceIntelligencePropertiesDevice) GetRawModelOk() (*string, bool)`

GetRawModelOk returns a tuple with the RawModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawModel

`func (o *DeviceIntelligencePropertiesDevice) SetRawModel(v string)`

SetRawModel sets RawModel field to given value.

### HasRawModel

`func (o *DeviceIntelligencePropertiesDevice) HasRawModel() bool`

HasRawModel returns a boolean if a field has been set.

### GetOs

`func (o *DeviceIntelligencePropertiesDevice) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *DeviceIntelligencePropertiesDevice) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *DeviceIntelligencePropertiesDevice) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *DeviceIntelligencePropertiesDevice) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetBrowser

`func (o *DeviceIntelligencePropertiesDevice) GetBrowser() string`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *DeviceIntelligencePropertiesDevice) GetBrowserOk() (*string, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowser

`func (o *DeviceIntelligencePropertiesDevice) SetBrowser(v string)`

SetBrowser sets Browser field to given value.

### HasBrowser

`func (o *DeviceIntelligencePropertiesDevice) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.

### GetEmulator

`func (o *DeviceIntelligencePropertiesDevice) GetEmulator() bool`

GetEmulator returns the Emulator field if non-nil, zero value otherwise.

### GetEmulatorOk

`func (o *DeviceIntelligencePropertiesDevice) GetEmulatorOk() (*bool, bool)`

GetEmulatorOk returns a tuple with the Emulator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmulator

`func (o *DeviceIntelligencePropertiesDevice) SetEmulator(v bool)`

SetEmulator sets Emulator field to given value.

### HasEmulator

`func (o *DeviceIntelligencePropertiesDevice) HasEmulator() bool`

HasEmulator returns a boolean if a field has been set.

### GetRandomizedDevice

`func (o *DeviceIntelligencePropertiesDevice) GetRandomizedDevice() bool`

GetRandomizedDevice returns the RandomizedDevice field if non-nil, zero value otherwise.

### GetRandomizedDeviceOk

`func (o *DeviceIntelligencePropertiesDevice) GetRandomizedDeviceOk() (*bool, bool)`

GetRandomizedDeviceOk returns a tuple with the RandomizedDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizedDevice

`func (o *DeviceIntelligencePropertiesDevice) SetRandomizedDevice(v bool)`

SetRandomizedDevice sets RandomizedDevice field to given value.

### HasRandomizedDevice

`func (o *DeviceIntelligencePropertiesDevice) HasRandomizedDevice() bool`

HasRandomizedDevice returns a boolean if a field has been set.

### GetFakeNetworkRequest

`func (o *DeviceIntelligencePropertiesDevice) GetFakeNetworkRequest() bool`

GetFakeNetworkRequest returns the FakeNetworkRequest field if non-nil, zero value otherwise.

### GetFakeNetworkRequestOk

`func (o *DeviceIntelligencePropertiesDevice) GetFakeNetworkRequestOk() (*bool, bool)`

GetFakeNetworkRequestOk returns a tuple with the FakeNetworkRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFakeNetworkRequest

`func (o *DeviceIntelligencePropertiesDevice) SetFakeNetworkRequest(v bool)`

SetFakeNetworkRequest sets FakeNetworkRequest field to given value.

### HasFakeNetworkRequest

`func (o *DeviceIntelligencePropertiesDevice) HasFakeNetworkRequest() bool`

HasFakeNetworkRequest returns a boolean if a field has been set.

### GetIpReputation

`func (o *DeviceIntelligencePropertiesDevice) GetIpReputation() string`

GetIpReputation returns the IpReputation field if non-nil, zero value otherwise.

### GetIpReputationOk

`func (o *DeviceIntelligencePropertiesDevice) GetIpReputationOk() (*string, bool)`

GetIpReputationOk returns a tuple with the IpReputation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpReputation

`func (o *DeviceIntelligencePropertiesDevice) SetIpReputation(v string)`

SetIpReputation sets IpReputation field to given value.

### HasIpReputation

`func (o *DeviceIntelligencePropertiesDevice) HasIpReputation() bool`

HasIpReputation returns a boolean if a field has been set.

### GetDeviceFingerprintReuse

`func (o *DeviceIntelligencePropertiesDevice) GetDeviceFingerprintReuse() float32`

GetDeviceFingerprintReuse returns the DeviceFingerprintReuse field if non-nil, zero value otherwise.

### GetDeviceFingerprintReuseOk

`func (o *DeviceIntelligencePropertiesDevice) GetDeviceFingerprintReuseOk() (*float32, bool)`

GetDeviceFingerprintReuseOk returns a tuple with the DeviceFingerprintReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFingerprintReuse

`func (o *DeviceIntelligencePropertiesDevice) SetDeviceFingerprintReuse(v float32)`

SetDeviceFingerprintReuse sets DeviceFingerprintReuse field to given value.

### HasDeviceFingerprintReuse

`func (o *DeviceIntelligencePropertiesDevice) HasDeviceFingerprintReuse() bool`

HasDeviceFingerprintReuse returns a boolean if a field has been set.

### GetSingleDeviceUsed

`func (o *DeviceIntelligencePropertiesDevice) GetSingleDeviceUsed() bool`

GetSingleDeviceUsed returns the SingleDeviceUsed field if non-nil, zero value otherwise.

### GetSingleDeviceUsedOk

`func (o *DeviceIntelligencePropertiesDevice) GetSingleDeviceUsedOk() (*bool, bool)`

GetSingleDeviceUsedOk returns a tuple with the SingleDeviceUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleDeviceUsed

`func (o *DeviceIntelligencePropertiesDevice) SetSingleDeviceUsed(v bool)`

SetSingleDeviceUsed sets SingleDeviceUsed field to given value.

### HasSingleDeviceUsed

`func (o *DeviceIntelligencePropertiesDevice) HasSingleDeviceUsed() bool`

HasSingleDeviceUsed returns a boolean if a field has been set.

### SetSingleDeviceUsedNil

`func (o *DeviceIntelligencePropertiesDevice) SetSingleDeviceUsedNil(b bool)`

 SetSingleDeviceUsedNil sets the value for SingleDeviceUsed to be an explicit nil

### UnsetSingleDeviceUsed
`func (o *DeviceIntelligencePropertiesDevice) UnsetSingleDeviceUsed()`

UnsetSingleDeviceUsed ensures that no value is present for SingleDeviceUsed, not even an explicit nil
### GetDocumentCapture

`func (o *DeviceIntelligencePropertiesDevice) GetDocumentCapture() string`

GetDocumentCapture returns the DocumentCapture field if non-nil, zero value otherwise.

### GetDocumentCaptureOk

`func (o *DeviceIntelligencePropertiesDevice) GetDocumentCaptureOk() (*string, bool)`

GetDocumentCaptureOk returns a tuple with the DocumentCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentCapture

`func (o *DeviceIntelligencePropertiesDevice) SetDocumentCapture(v string)`

SetDocumentCapture sets DocumentCapture field to given value.

### HasDocumentCapture

`func (o *DeviceIntelligencePropertiesDevice) HasDocumentCapture() bool`

HasDocumentCapture returns a boolean if a field has been set.

### GetBiometricCapture

`func (o *DeviceIntelligencePropertiesDevice) GetBiometricCapture() string`

GetBiometricCapture returns the BiometricCapture field if non-nil, zero value otherwise.

### GetBiometricCaptureOk

`func (o *DeviceIntelligencePropertiesDevice) GetBiometricCaptureOk() (*string, bool)`

GetBiometricCaptureOk returns a tuple with the BiometricCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiometricCapture

`func (o *DeviceIntelligencePropertiesDevice) SetBiometricCapture(v string)`

SetBiometricCapture sets BiometricCapture field to given value.

### HasBiometricCapture

`func (o *DeviceIntelligencePropertiesDevice) HasBiometricCapture() bool`

HasBiometricCapture returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


