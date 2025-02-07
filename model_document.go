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
	"time"
)

// checks if the Document type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Document{}

// Document struct for Document
type Document struct {
	// The unique identifier for the document
	Id *string `json:"id,omitempty"`
	// The date and time at which the document was uploaded
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The name of the uploaded file
	FileName *string `json:"file_name,omitempty"`
	// The size of the file in bytes
	FileSize *int32 `json:"file_size,omitempty"`
	// The file type of the uploaded file
	FileType *string `json:"file_type,omitempty"`
	// The type of document
	Type *string `json:"type,omitempty"`
	// The side of the document, if applicable. The possible values are front and back
	Side *string `json:"side,omitempty"`
	// The issuing country of the document, a 3-letter ISO code.
	IssuingCountry *string `json:"issuing_country,omitempty"`
	// The uri of this resource
	Href *string `json:"href,omitempty"`
	// The uri that can be used to download the document
	DownloadHref *string `json:"download_href,omitempty"`
}

// NewDocument instantiates a new Document object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocument() *Document {
	this := Document{}
	return &this
}

// NewDocumentWithDefaults instantiates a new Document object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentWithDefaults() *Document {
	this := Document{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Document) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Document) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Document) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Document) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Document) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Document) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Document) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Document) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *Document) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Document) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *Document) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *Document) SetFileName(v string) {
	o.FileName = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *Document) GetFileSize() int32 {
	if o == nil || IsNil(o.FileSize) {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Document) GetFileSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *Document) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *Document) SetFileSize(v int32) {
	o.FileSize = &v
}

// GetFileType returns the FileType field value if set, zero value otherwise.
func (o *Document) GetFileType() string {
	if o == nil || IsNil(o.FileType) {
		var ret string
		return ret
	}
	return *o.FileType
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Document) GetFileTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FileType) {
		return nil, false
	}
	return o.FileType, true
}

// HasFileType returns a boolean if a field has been set.
func (o *Document) HasFileType() bool {
	if o != nil && !IsNil(o.FileType) {
		return true
	}

	return false
}

// SetFileType gets a reference to the given string and assigns it to the FileType field.
func (o *Document) SetFileType(v string) {
	o.FileType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Document) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Document) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Document) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Document) SetType(v string) {
	o.Type = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *Document) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Document) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *Document) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *Document) SetSide(v string) {
	o.Side = &v
}

// GetIssuingCountry returns the IssuingCountry field value if set, zero value otherwise.
func (o *Document) GetIssuingCountry() string {
	if o == nil || IsNil(o.IssuingCountry) {
		var ret string
		return ret
	}
	return *o.IssuingCountry
}

// GetIssuingCountryOk returns a tuple with the IssuingCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Document) GetIssuingCountryOk() (*string, bool) {
	if o == nil || IsNil(o.IssuingCountry) {
		return nil, false
	}
	return o.IssuingCountry, true
}

// HasIssuingCountry returns a boolean if a field has been set.
func (o *Document) HasIssuingCountry() bool {
	if o != nil && !IsNil(o.IssuingCountry) {
		return true
	}

	return false
}

// SetIssuingCountry gets a reference to the given string and assigns it to the IssuingCountry field.
func (o *Document) SetIssuingCountry(v string) {
	o.IssuingCountry = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Document) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Document) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Document) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Document) SetHref(v string) {
	o.Href = &v
}

// GetDownloadHref returns the DownloadHref field value if set, zero value otherwise.
func (o *Document) GetDownloadHref() string {
	if o == nil || IsNil(o.DownloadHref) {
		var ret string
		return ret
	}
	return *o.DownloadHref
}

// GetDownloadHrefOk returns a tuple with the DownloadHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Document) GetDownloadHrefOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadHref) {
		return nil, false
	}
	return o.DownloadHref, true
}

// HasDownloadHref returns a boolean if a field has been set.
func (o *Document) HasDownloadHref() bool {
	if o != nil && !IsNil(o.DownloadHref) {
		return true
	}

	return false
}

// SetDownloadHref gets a reference to the given string and assigns it to the DownloadHref field.
func (o *Document) SetDownloadHref(v string) {
	o.DownloadHref = &v
}

func (o Document) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Document) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.FileName) {
		toSerialize["file_name"] = o.FileName
	}
	if !IsNil(o.FileSize) {
		toSerialize["file_size"] = o.FileSize
	}
	if !IsNil(o.FileType) {
		toSerialize["file_type"] = o.FileType
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !IsNil(o.IssuingCountry) {
		toSerialize["issuing_country"] = o.IssuingCountry
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.DownloadHref) {
		toSerialize["download_href"] = o.DownloadHref
	}
	return toSerialize, nil
}

type NullableDocument struct {
	value *Document
	isSet bool
}

func (v NullableDocument) Get() *Document {
	return v.value
}

func (v *NullableDocument) Set(val *Document) {
	v.value = val
	v.isSet = true
}

func (v NullableDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocument(val *Document) *NullableDocument {
	return &NullableDocument{value: val, isSet: true}
}

func (v NullableDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
