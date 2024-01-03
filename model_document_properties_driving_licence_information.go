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

// checks if the DocumentPropertiesDrivingLicenceInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentPropertiesDrivingLicenceInformation{}

// DocumentPropertiesDrivingLicenceInformation struct for DocumentPropertiesDrivingLicenceInformation
type DocumentPropertiesDrivingLicenceInformation struct {
	Category       *string `json:"category,omitempty"`
	ObtainmentDate *string `json:"obtainment_date,omitempty"`
	ExpiryDate     *string `json:"expiry_date,omitempty"`
	Codes          *string `json:"codes,omitempty"`
}

// NewDocumentPropertiesDrivingLicenceInformation instantiates a new DocumentPropertiesDrivingLicenceInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentPropertiesDrivingLicenceInformation() *DocumentPropertiesDrivingLicenceInformation {
	this := DocumentPropertiesDrivingLicenceInformation{}
	return &this
}

// NewDocumentPropertiesDrivingLicenceInformationWithDefaults instantiates a new DocumentPropertiesDrivingLicenceInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentPropertiesDrivingLicenceInformationWithDefaults() *DocumentPropertiesDrivingLicenceInformation {
	this := DocumentPropertiesDrivingLicenceInformation{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *DocumentPropertiesDrivingLicenceInformation) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesDrivingLicenceInformation) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *DocumentPropertiesDrivingLicenceInformation) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *DocumentPropertiesDrivingLicenceInformation) SetCategory(v string) {
	o.Category = &v
}

// GetObtainmentDate returns the ObtainmentDate field value if set, zero value otherwise.
func (o *DocumentPropertiesDrivingLicenceInformation) GetObtainmentDate() string {
	if o == nil || IsNil(o.ObtainmentDate) {
		var ret string
		return ret
	}
	return *o.ObtainmentDate
}

// GetObtainmentDateOk returns a tuple with the ObtainmentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesDrivingLicenceInformation) GetObtainmentDateOk() (*string, bool) {
	if o == nil || IsNil(o.ObtainmentDate) {
		return nil, false
	}
	return o.ObtainmentDate, true
}

// HasObtainmentDate returns a boolean if a field has been set.
func (o *DocumentPropertiesDrivingLicenceInformation) HasObtainmentDate() bool {
	if o != nil && !IsNil(o.ObtainmentDate) {
		return true
	}

	return false
}

// SetObtainmentDate gets a reference to the given string and assigns it to the ObtainmentDate field.
func (o *DocumentPropertiesDrivingLicenceInformation) SetObtainmentDate(v string) {
	o.ObtainmentDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *DocumentPropertiesDrivingLicenceInformation) GetExpiryDate() string {
	if o == nil || IsNil(o.ExpiryDate) {
		var ret string
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesDrivingLicenceInformation) GetExpiryDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiryDate) {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *DocumentPropertiesDrivingLicenceInformation) HasExpiryDate() bool {
	if o != nil && !IsNil(o.ExpiryDate) {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given string and assigns it to the ExpiryDate field.
func (o *DocumentPropertiesDrivingLicenceInformation) SetExpiryDate(v string) {
	o.ExpiryDate = &v
}

// GetCodes returns the Codes field value if set, zero value otherwise.
func (o *DocumentPropertiesDrivingLicenceInformation) GetCodes() string {
	if o == nil || IsNil(o.Codes) {
		var ret string
		return ret
	}
	return *o.Codes
}

// GetCodesOk returns a tuple with the Codes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesDrivingLicenceInformation) GetCodesOk() (*string, bool) {
	if o == nil || IsNil(o.Codes) {
		return nil, false
	}
	return o.Codes, true
}

// HasCodes returns a boolean if a field has been set.
func (o *DocumentPropertiesDrivingLicenceInformation) HasCodes() bool {
	if o != nil && !IsNil(o.Codes) {
		return true
	}

	return false
}

// SetCodes gets a reference to the given string and assigns it to the Codes field.
func (o *DocumentPropertiesDrivingLicenceInformation) SetCodes(v string) {
	o.Codes = &v
}

func (o DocumentPropertiesDrivingLicenceInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentPropertiesDrivingLicenceInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.ObtainmentDate) {
		toSerialize["obtainment_date"] = o.ObtainmentDate
	}
	if !IsNil(o.ExpiryDate) {
		toSerialize["expiry_date"] = o.ExpiryDate
	}
	if !IsNil(o.Codes) {
		toSerialize["codes"] = o.Codes
	}
	return toSerialize, nil
}

type NullableDocumentPropertiesDrivingLicenceInformation struct {
	value *DocumentPropertiesDrivingLicenceInformation
	isSet bool
}

func (v NullableDocumentPropertiesDrivingLicenceInformation) Get() *DocumentPropertiesDrivingLicenceInformation {
	return v.value
}

func (v *NullableDocumentPropertiesDrivingLicenceInformation) Set(val *DocumentPropertiesDrivingLicenceInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentPropertiesDrivingLicenceInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentPropertiesDrivingLicenceInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentPropertiesDrivingLicenceInformation(val *DocumentPropertiesDrivingLicenceInformation) *NullableDocumentPropertiesDrivingLicenceInformation {
	return &NullableDocumentPropertiesDrivingLicenceInformation{value: val, isSet: true}
}

func (v NullableDocumentPropertiesDrivingLicenceInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentPropertiesDrivingLicenceInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
