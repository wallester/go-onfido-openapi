/*
 * Onfido API
 *
 * The Onfido API is used to submit check requests.
 *
 * API version: 3.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onfido_openapi

import (
	"encoding/json"
	"time"
)

// LivePhoto struct for LivePhoto
type LivePhoto struct {
	// The unique identifier for the photo.
	Id *string `json:"id,omitempty"`
	// The date and time at which the photo was uploaded.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The uri of this resource.
	Href *string `json:"href,omitempty"`
	// The uri that can be used to download the photo.
	DownloadHref *string `json:"download_href,omitempty"`
	// The name of the uploaded file.
	FileName *string `json:"file_name,omitempty"`
	// The size of the file in bytes.
	FileSize *int32 `json:"file_size,omitempty"`
	// The file type of the uploaded file.
	FileType *string `json:"file_type,omitempty"`
}

// NewLivePhoto instantiates a new LivePhoto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLivePhoto() *LivePhoto {
	this := LivePhoto{}
	return &this
}

// NewLivePhotoWithDefaults instantiates a new LivePhoto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLivePhotoWithDefaults() *LivePhoto {
	this := LivePhoto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LivePhoto) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LivePhoto) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LivePhoto) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LivePhoto) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LivePhoto) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LivePhoto) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LivePhoto) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *LivePhoto) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *LivePhoto) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LivePhoto) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *LivePhoto) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *LivePhoto) SetHref(v string) {
	o.Href = &v
}

// GetDownloadHref returns the DownloadHref field value if set, zero value otherwise.
func (o *LivePhoto) GetDownloadHref() string {
	if o == nil || o.DownloadHref == nil {
		var ret string
		return ret
	}
	return *o.DownloadHref
}

// GetDownloadHrefOk returns a tuple with the DownloadHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LivePhoto) GetDownloadHrefOk() (*string, bool) {
	if o == nil || o.DownloadHref == nil {
		return nil, false
	}
	return o.DownloadHref, true
}

// HasDownloadHref returns a boolean if a field has been set.
func (o *LivePhoto) HasDownloadHref() bool {
	if o != nil && o.DownloadHref != nil {
		return true
	}

	return false
}

// SetDownloadHref gets a reference to the given string and assigns it to the DownloadHref field.
func (o *LivePhoto) SetDownloadHref(v string) {
	o.DownloadHref = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *LivePhoto) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LivePhoto) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *LivePhoto) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *LivePhoto) SetFileName(v string) {
	o.FileName = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *LivePhoto) GetFileSize() int32 {
	if o == nil || o.FileSize == nil {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LivePhoto) GetFileSizeOk() (*int32, bool) {
	if o == nil || o.FileSize == nil {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *LivePhoto) HasFileSize() bool {
	if o != nil && o.FileSize != nil {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *LivePhoto) SetFileSize(v int32) {
	o.FileSize = &v
}

// GetFileType returns the FileType field value if set, zero value otherwise.
func (o *LivePhoto) GetFileType() string {
	if o == nil || o.FileType == nil {
		var ret string
		return ret
	}
	return *o.FileType
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LivePhoto) GetFileTypeOk() (*string, bool) {
	if o == nil || o.FileType == nil {
		return nil, false
	}
	return o.FileType, true
}

// HasFileType returns a boolean if a field has been set.
func (o *LivePhoto) HasFileType() bool {
	if o != nil && o.FileType != nil {
		return true
	}

	return false
}

// SetFileType gets a reference to the given string and assigns it to the FileType field.
func (o *LivePhoto) SetFileType(v string) {
	o.FileType = &v
}

func (o LivePhoto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.DownloadHref != nil {
		toSerialize["download_href"] = o.DownloadHref
	}
	if o.FileName != nil {
		toSerialize["file_name"] = o.FileName
	}
	if o.FileSize != nil {
		toSerialize["file_size"] = o.FileSize
	}
	if o.FileType != nil {
		toSerialize["file_type"] = o.FileType
	}
	return json.Marshal(toSerialize)
}

type NullableLivePhoto struct {
	value *LivePhoto
	isSet bool
}

func (v NullableLivePhoto) Get() *LivePhoto {
	return v.value
}

func (v *NullableLivePhoto) Set(val *LivePhoto) {
	v.value = val
	v.isSet = true
}

func (v NullableLivePhoto) IsSet() bool {
	return v.isSet
}

func (v *NullableLivePhoto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLivePhoto(val *LivePhoto) *NullableLivePhoto {
	return &NullableLivePhoto{value: val, isSet: true}
}

func (v NullableLivePhoto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLivePhoto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


