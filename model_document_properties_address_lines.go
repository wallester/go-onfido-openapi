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

// checks if the DocumentPropertiesAddressLines type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentPropertiesAddressLines{}

// DocumentPropertiesAddressLines struct for DocumentPropertiesAddressLines
type DocumentPropertiesAddressLines struct {
	StreetAddress *string `json:"street_address,omitempty"`
	State         *string `json:"state,omitempty"`
	PostalCode    *string `json:"postal_code,omitempty"`
	Country       *string `json:"country,omitempty"`
	City          *string `json:"city,omitempty"`
	CountryCode   *string `json:"country_code,omitempty"`
}

// NewDocumentPropertiesAddressLines instantiates a new DocumentPropertiesAddressLines object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentPropertiesAddressLines() *DocumentPropertiesAddressLines {
	this := DocumentPropertiesAddressLines{}
	return &this
}

// NewDocumentPropertiesAddressLinesWithDefaults instantiates a new DocumentPropertiesAddressLines object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentPropertiesAddressLinesWithDefaults() *DocumentPropertiesAddressLines {
	this := DocumentPropertiesAddressLines{}
	return &this
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *DocumentPropertiesAddressLines) GetStreetAddress() string {
	if o == nil || IsNil(o.StreetAddress) {
		var ret string
		return ret
	}
	return *o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesAddressLines) GetStreetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.StreetAddress) {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *DocumentPropertiesAddressLines) HasStreetAddress() bool {
	if o != nil && !IsNil(o.StreetAddress) {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given string and assigns it to the StreetAddress field.
func (o *DocumentPropertiesAddressLines) SetStreetAddress(v string) {
	o.StreetAddress = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DocumentPropertiesAddressLines) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesAddressLines) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DocumentPropertiesAddressLines) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *DocumentPropertiesAddressLines) SetState(v string) {
	o.State = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *DocumentPropertiesAddressLines) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesAddressLines) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *DocumentPropertiesAddressLines) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *DocumentPropertiesAddressLines) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *DocumentPropertiesAddressLines) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesAddressLines) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *DocumentPropertiesAddressLines) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *DocumentPropertiesAddressLines) SetCountry(v string) {
	o.Country = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *DocumentPropertiesAddressLines) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesAddressLines) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *DocumentPropertiesAddressLines) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *DocumentPropertiesAddressLines) SetCity(v string) {
	o.City = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *DocumentPropertiesAddressLines) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesAddressLines) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *DocumentPropertiesAddressLines) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *DocumentPropertiesAddressLines) SetCountryCode(v string) {
	o.CountryCode = &v
}

func (o DocumentPropertiesAddressLines) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentPropertiesAddressLines) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StreetAddress) {
		toSerialize["street_address"] = o.StreetAddress
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.CountryCode) {
		toSerialize["country_code"] = o.CountryCode
	}
	return toSerialize, nil
}

type NullableDocumentPropertiesAddressLines struct {
	value *DocumentPropertiesAddressLines
	isSet bool
}

func (v NullableDocumentPropertiesAddressLines) Get() *DocumentPropertiesAddressLines {
	return v.value
}

func (v *NullableDocumentPropertiesAddressLines) Set(val *DocumentPropertiesAddressLines) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentPropertiesAddressLines) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentPropertiesAddressLines) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentPropertiesAddressLines(val *DocumentPropertiesAddressLines) *NullableDocumentPropertiesAddressLines {
	return &NullableDocumentPropertiesAddressLines{value: val, isSet: true}
}

func (v NullableDocumentPropertiesAddressLines) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentPropertiesAddressLines) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
