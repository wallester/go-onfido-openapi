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

// checks if the DocumentBreakdownIssuingAuthorityBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentBreakdownIssuingAuthorityBreakdown{}

// DocumentBreakdownIssuingAuthorityBreakdown struct for DocumentBreakdownIssuingAuthorityBreakdown
type DocumentBreakdownIssuingAuthorityBreakdown struct {
	NfcActiveAuthentication  *DocumentBreakdownIssuingAuthorityBreakdownNfcActiveAuthentication  `json:"nfc_active_authentication,omitempty"`
	NfcPassiveAuthentication *DocumentBreakdownIssuingAuthorityBreakdownNfcPassiveAuthentication `json:"nfc_passive_authentication,omitempty"`
}

// NewDocumentBreakdownIssuingAuthorityBreakdown instantiates a new DocumentBreakdownIssuingAuthorityBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentBreakdownIssuingAuthorityBreakdown() *DocumentBreakdownIssuingAuthorityBreakdown {
	this := DocumentBreakdownIssuingAuthorityBreakdown{}
	return &this
}

// NewDocumentBreakdownIssuingAuthorityBreakdownWithDefaults instantiates a new DocumentBreakdownIssuingAuthorityBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentBreakdownIssuingAuthorityBreakdownWithDefaults() *DocumentBreakdownIssuingAuthorityBreakdown {
	this := DocumentBreakdownIssuingAuthorityBreakdown{}
	return &this
}

// GetNfcActiveAuthentication returns the NfcActiveAuthentication field value if set, zero value otherwise.
func (o *DocumentBreakdownIssuingAuthorityBreakdown) GetNfcActiveAuthentication() DocumentBreakdownIssuingAuthorityBreakdownNfcActiveAuthentication {
	if o == nil || IsNil(o.NfcActiveAuthentication) {
		var ret DocumentBreakdownIssuingAuthorityBreakdownNfcActiveAuthentication
		return ret
	}
	return *o.NfcActiveAuthentication
}

// GetNfcActiveAuthenticationOk returns a tuple with the NfcActiveAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentBreakdownIssuingAuthorityBreakdown) GetNfcActiveAuthenticationOk() (*DocumentBreakdownIssuingAuthorityBreakdownNfcActiveAuthentication, bool) {
	if o == nil || IsNil(o.NfcActiveAuthentication) {
		return nil, false
	}
	return o.NfcActiveAuthentication, true
}

// HasNfcActiveAuthentication returns a boolean if a field has been set.
func (o *DocumentBreakdownIssuingAuthorityBreakdown) HasNfcActiveAuthentication() bool {
	if o != nil && !IsNil(o.NfcActiveAuthentication) {
		return true
	}

	return false
}

// SetNfcActiveAuthentication gets a reference to the given DocumentBreakdownIssuingAuthorityBreakdownNfcActiveAuthentication and assigns it to the NfcActiveAuthentication field.
func (o *DocumentBreakdownIssuingAuthorityBreakdown) SetNfcActiveAuthentication(v DocumentBreakdownIssuingAuthorityBreakdownNfcActiveAuthentication) {
	o.NfcActiveAuthentication = &v
}

// GetNfcPassiveAuthentication returns the NfcPassiveAuthentication field value if set, zero value otherwise.
func (o *DocumentBreakdownIssuingAuthorityBreakdown) GetNfcPassiveAuthentication() DocumentBreakdownIssuingAuthorityBreakdownNfcPassiveAuthentication {
	if o == nil || IsNil(o.NfcPassiveAuthentication) {
		var ret DocumentBreakdownIssuingAuthorityBreakdownNfcPassiveAuthentication
		return ret
	}
	return *o.NfcPassiveAuthentication
}

// GetNfcPassiveAuthenticationOk returns a tuple with the NfcPassiveAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentBreakdownIssuingAuthorityBreakdown) GetNfcPassiveAuthenticationOk() (*DocumentBreakdownIssuingAuthorityBreakdownNfcPassiveAuthentication, bool) {
	if o == nil || IsNil(o.NfcPassiveAuthentication) {
		return nil, false
	}
	return o.NfcPassiveAuthentication, true
}

// HasNfcPassiveAuthentication returns a boolean if a field has been set.
func (o *DocumentBreakdownIssuingAuthorityBreakdown) HasNfcPassiveAuthentication() bool {
	if o != nil && !IsNil(o.NfcPassiveAuthentication) {
		return true
	}

	return false
}

// SetNfcPassiveAuthentication gets a reference to the given DocumentBreakdownIssuingAuthorityBreakdownNfcPassiveAuthentication and assigns it to the NfcPassiveAuthentication field.
func (o *DocumentBreakdownIssuingAuthorityBreakdown) SetNfcPassiveAuthentication(v DocumentBreakdownIssuingAuthorityBreakdownNfcPassiveAuthentication) {
	o.NfcPassiveAuthentication = &v
}

func (o DocumentBreakdownIssuingAuthorityBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentBreakdownIssuingAuthorityBreakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NfcActiveAuthentication) {
		toSerialize["nfc_active_authentication"] = o.NfcActiveAuthentication
	}
	if !IsNil(o.NfcPassiveAuthentication) {
		toSerialize["nfc_passive_authentication"] = o.NfcPassiveAuthentication
	}
	return toSerialize, nil
}

type NullableDocumentBreakdownIssuingAuthorityBreakdown struct {
	value *DocumentBreakdownIssuingAuthorityBreakdown
	isSet bool
}

func (v NullableDocumentBreakdownIssuingAuthorityBreakdown) Get() *DocumentBreakdownIssuingAuthorityBreakdown {
	return v.value
}

func (v *NullableDocumentBreakdownIssuingAuthorityBreakdown) Set(val *DocumentBreakdownIssuingAuthorityBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentBreakdownIssuingAuthorityBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentBreakdownIssuingAuthorityBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentBreakdownIssuingAuthorityBreakdown(val *DocumentBreakdownIssuingAuthorityBreakdown) *NullableDocumentBreakdownIssuingAuthorityBreakdown {
	return &NullableDocumentBreakdownIssuingAuthorityBreakdown{value: val, isSet: true}
}

func (v NullableDocumentBreakdownIssuingAuthorityBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentBreakdownIssuingAuthorityBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
