/*
Onfido API v3.6

The Onfido API is used to submit check requests.

API version: 3.6.0
Contact: engineering@onfido.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onfido_openapi

import (
	"encoding/json"
)

// checks if the VideoReasons type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoReasons{}

// VideoReasons struct for VideoReasons
type VideoReasons struct {
	// Flags when evidence is found that a fake webcam was used.
	FakeWebcam *string `json:"fake_webcam,omitempty"`
	// Flags when evidence is found that the video was uploaded in an attempt to circumvent the randomness of the speaking and head turn challenges
	ChallengeReuse *string `json:"challenge_reuse,omitempty"`
	// Flags when evidence is found that an Android emulator was used.
	Emulator *string `json:"emulator,omitempty"`
	// Additional comma separated details such as the name of the fake webcam.
	Reasons *string `json:"reasons,omitempty"`
}

// NewVideoReasons instantiates a new VideoReasons object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoReasons() *VideoReasons {
	this := VideoReasons{}
	return &this
}

// NewVideoReasonsWithDefaults instantiates a new VideoReasons object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoReasonsWithDefaults() *VideoReasons {
	this := VideoReasons{}
	return &this
}

// GetFakeWebcam returns the FakeWebcam field value if set, zero value otherwise.
func (o *VideoReasons) GetFakeWebcam() string {
	if o == nil || IsNil(o.FakeWebcam) {
		var ret string
		return ret
	}
	return *o.FakeWebcam
}

// GetFakeWebcamOk returns a tuple with the FakeWebcam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoReasons) GetFakeWebcamOk() (*string, bool) {
	if o == nil || IsNil(o.FakeWebcam) {
		return nil, false
	}
	return o.FakeWebcam, true
}

// HasFakeWebcam returns a boolean if a field has been set.
func (o *VideoReasons) HasFakeWebcam() bool {
	if o != nil && !IsNil(o.FakeWebcam) {
		return true
	}

	return false
}

// SetFakeWebcam gets a reference to the given string and assigns it to the FakeWebcam field.
func (o *VideoReasons) SetFakeWebcam(v string) {
	o.FakeWebcam = &v
}

// GetChallengeReuse returns the ChallengeReuse field value if set, zero value otherwise.
func (o *VideoReasons) GetChallengeReuse() string {
	if o == nil || IsNil(o.ChallengeReuse) {
		var ret string
		return ret
	}
	return *o.ChallengeReuse
}

// GetChallengeReuseOk returns a tuple with the ChallengeReuse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoReasons) GetChallengeReuseOk() (*string, bool) {
	if o == nil || IsNil(o.ChallengeReuse) {
		return nil, false
	}
	return o.ChallengeReuse, true
}

// HasChallengeReuse returns a boolean if a field has been set.
func (o *VideoReasons) HasChallengeReuse() bool {
	if o != nil && !IsNil(o.ChallengeReuse) {
		return true
	}

	return false
}

// SetChallengeReuse gets a reference to the given string and assigns it to the ChallengeReuse field.
func (o *VideoReasons) SetChallengeReuse(v string) {
	o.ChallengeReuse = &v
}

// GetEmulator returns the Emulator field value if set, zero value otherwise.
func (o *VideoReasons) GetEmulator() string {
	if o == nil || IsNil(o.Emulator) {
		var ret string
		return ret
	}
	return *o.Emulator
}

// GetEmulatorOk returns a tuple with the Emulator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoReasons) GetEmulatorOk() (*string, bool) {
	if o == nil || IsNil(o.Emulator) {
		return nil, false
	}
	return o.Emulator, true
}

// HasEmulator returns a boolean if a field has been set.
func (o *VideoReasons) HasEmulator() bool {
	if o != nil && !IsNil(o.Emulator) {
		return true
	}

	return false
}

// SetEmulator gets a reference to the given string and assigns it to the Emulator field.
func (o *VideoReasons) SetEmulator(v string) {
	o.Emulator = &v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *VideoReasons) GetReasons() string {
	if o == nil || IsNil(o.Reasons) {
		var ret string
		return ret
	}
	return *o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoReasons) GetReasonsOk() (*string, bool) {
	if o == nil || IsNil(o.Reasons) {
		return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *VideoReasons) HasReasons() bool {
	if o != nil && !IsNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given string and assigns it to the Reasons field.
func (o *VideoReasons) SetReasons(v string) {
	o.Reasons = &v
}

func (o VideoReasons) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoReasons) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FakeWebcam) {
		toSerialize["fake_webcam"] = o.FakeWebcam
	}
	if !IsNil(o.ChallengeReuse) {
		toSerialize["challenge_reuse"] = o.ChallengeReuse
	}
	if !IsNil(o.Emulator) {
		toSerialize["emulator"] = o.Emulator
	}
	if !IsNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	return toSerialize, nil
}

type NullableVideoReasons struct {
	value *VideoReasons
	isSet bool
}

func (v NullableVideoReasons) Get() *VideoReasons {
	return v.value
}

func (v *NullableVideoReasons) Set(val *VideoReasons) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoReasons) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoReasons) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoReasons(val *VideoReasons) *NullableVideoReasons {
	return &NullableVideoReasons{value: val, isSet: true}
}

func (v NullableVideoReasons) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoReasons) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
