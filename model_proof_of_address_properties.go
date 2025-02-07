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

// checks if the ProofOfAddressProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProofOfAddressProperties{}

// ProofOfAddressProperties struct for ProofOfAddressProperties
type ProofOfAddressProperties struct {
	// This property provides the address on the document.
	Address *string `json:"address,omitempty"`
	// This property provides the document type according to the set of supported documents.
	DocumentType *string `json:"document_type,omitempty"`
	// This property provides the first names on the document, including any initials and middle names.
	FirstNames *string `json:"first_names,omitempty"`
	// This property provided the last names on the document.
	LastNames *string `json:"last_names,omitempty"`
	// This property provides the issue date of the document.
	IssueDate *string `json:"issue_date,omitempty"`
	// This property provides the document issuer (e.g. HSBC, British Gas).
	Issuer *string `json:"issuer,omitempty"`
	// This property provides the summary period start date.
	SummaryPeriodStart *string `json:"summary_period_start,omitempty"`
	// This property provides the summary period end date.
	SummaryPeriodEnd *string `json:"summary_period_end,omitempty"`
}

// NewProofOfAddressProperties instantiates a new ProofOfAddressProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProofOfAddressProperties() *ProofOfAddressProperties {
	this := ProofOfAddressProperties{}
	return &this
}

// NewProofOfAddressPropertiesWithDefaults instantiates a new ProofOfAddressProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProofOfAddressPropertiesWithDefaults() *ProofOfAddressProperties {
	this := ProofOfAddressProperties{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ProofOfAddressProperties) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProofOfAddressProperties) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ProofOfAddressProperties) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ProofOfAddressProperties) SetAddress(v string) {
	o.Address = &v
}

// GetDocumentType returns the DocumentType field value if set, zero value otherwise.
func (o *ProofOfAddressProperties) GetDocumentType() string {
	if o == nil || IsNil(o.DocumentType) {
		var ret string
		return ret
	}
	return *o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProofOfAddressProperties) GetDocumentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentType) {
		return nil, false
	}
	return o.DocumentType, true
}

// HasDocumentType returns a boolean if a field has been set.
func (o *ProofOfAddressProperties) HasDocumentType() bool {
	if o != nil && !IsNil(o.DocumentType) {
		return true
	}

	return false
}

// SetDocumentType gets a reference to the given string and assigns it to the DocumentType field.
func (o *ProofOfAddressProperties) SetDocumentType(v string) {
	o.DocumentType = &v
}

// GetFirstNames returns the FirstNames field value if set, zero value otherwise.
func (o *ProofOfAddressProperties) GetFirstNames() string {
	if o == nil || IsNil(o.FirstNames) {
		var ret string
		return ret
	}
	return *o.FirstNames
}

// GetFirstNamesOk returns a tuple with the FirstNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProofOfAddressProperties) GetFirstNamesOk() (*string, bool) {
	if o == nil || IsNil(o.FirstNames) {
		return nil, false
	}
	return o.FirstNames, true
}

// HasFirstNames returns a boolean if a field has been set.
func (o *ProofOfAddressProperties) HasFirstNames() bool {
	if o != nil && !IsNil(o.FirstNames) {
		return true
	}

	return false
}

// SetFirstNames gets a reference to the given string and assigns it to the FirstNames field.
func (o *ProofOfAddressProperties) SetFirstNames(v string) {
	o.FirstNames = &v
}

// GetLastNames returns the LastNames field value if set, zero value otherwise.
func (o *ProofOfAddressProperties) GetLastNames() string {
	if o == nil || IsNil(o.LastNames) {
		var ret string
		return ret
	}
	return *o.LastNames
}

// GetLastNamesOk returns a tuple with the LastNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProofOfAddressProperties) GetLastNamesOk() (*string, bool) {
	if o == nil || IsNil(o.LastNames) {
		return nil, false
	}
	return o.LastNames, true
}

// HasLastNames returns a boolean if a field has been set.
func (o *ProofOfAddressProperties) HasLastNames() bool {
	if o != nil && !IsNil(o.LastNames) {
		return true
	}

	return false
}

// SetLastNames gets a reference to the given string and assigns it to the LastNames field.
func (o *ProofOfAddressProperties) SetLastNames(v string) {
	o.LastNames = &v
}

// GetIssueDate returns the IssueDate field value if set, zero value otherwise.
func (o *ProofOfAddressProperties) GetIssueDate() string {
	if o == nil || IsNil(o.IssueDate) {
		var ret string
		return ret
	}
	return *o.IssueDate
}

// GetIssueDateOk returns a tuple with the IssueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProofOfAddressProperties) GetIssueDateOk() (*string, bool) {
	if o == nil || IsNil(o.IssueDate) {
		return nil, false
	}
	return o.IssueDate, true
}

// HasIssueDate returns a boolean if a field has been set.
func (o *ProofOfAddressProperties) HasIssueDate() bool {
	if o != nil && !IsNil(o.IssueDate) {
		return true
	}

	return false
}

// SetIssueDate gets a reference to the given string and assigns it to the IssueDate field.
func (o *ProofOfAddressProperties) SetIssueDate(v string) {
	o.IssueDate = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *ProofOfAddressProperties) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProofOfAddressProperties) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *ProofOfAddressProperties) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *ProofOfAddressProperties) SetIssuer(v string) {
	o.Issuer = &v
}

// GetSummaryPeriodStart returns the SummaryPeriodStart field value if set, zero value otherwise.
func (o *ProofOfAddressProperties) GetSummaryPeriodStart() string {
	if o == nil || IsNil(o.SummaryPeriodStart) {
		var ret string
		return ret
	}
	return *o.SummaryPeriodStart
}

// GetSummaryPeriodStartOk returns a tuple with the SummaryPeriodStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProofOfAddressProperties) GetSummaryPeriodStartOk() (*string, bool) {
	if o == nil || IsNil(o.SummaryPeriodStart) {
		return nil, false
	}
	return o.SummaryPeriodStart, true
}

// HasSummaryPeriodStart returns a boolean if a field has been set.
func (o *ProofOfAddressProperties) HasSummaryPeriodStart() bool {
	if o != nil && !IsNil(o.SummaryPeriodStart) {
		return true
	}

	return false
}

// SetSummaryPeriodStart gets a reference to the given string and assigns it to the SummaryPeriodStart field.
func (o *ProofOfAddressProperties) SetSummaryPeriodStart(v string) {
	o.SummaryPeriodStart = &v
}

// GetSummaryPeriodEnd returns the SummaryPeriodEnd field value if set, zero value otherwise.
func (o *ProofOfAddressProperties) GetSummaryPeriodEnd() string {
	if o == nil || IsNil(o.SummaryPeriodEnd) {
		var ret string
		return ret
	}
	return *o.SummaryPeriodEnd
}

// GetSummaryPeriodEndOk returns a tuple with the SummaryPeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProofOfAddressProperties) GetSummaryPeriodEndOk() (*string, bool) {
	if o == nil || IsNil(o.SummaryPeriodEnd) {
		return nil, false
	}
	return o.SummaryPeriodEnd, true
}

// HasSummaryPeriodEnd returns a boolean if a field has been set.
func (o *ProofOfAddressProperties) HasSummaryPeriodEnd() bool {
	if o != nil && !IsNil(o.SummaryPeriodEnd) {
		return true
	}

	return false
}

// SetSummaryPeriodEnd gets a reference to the given string and assigns it to the SummaryPeriodEnd field.
func (o *ProofOfAddressProperties) SetSummaryPeriodEnd(v string) {
	o.SummaryPeriodEnd = &v
}

func (o ProofOfAddressProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProofOfAddressProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.DocumentType) {
		toSerialize["document_type"] = o.DocumentType
	}
	if !IsNil(o.FirstNames) {
		toSerialize["first_names"] = o.FirstNames
	}
	if !IsNil(o.LastNames) {
		toSerialize["last_names"] = o.LastNames
	}
	if !IsNil(o.IssueDate) {
		toSerialize["issue_date"] = o.IssueDate
	}
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !IsNil(o.SummaryPeriodStart) {
		toSerialize["summary_period_start"] = o.SummaryPeriodStart
	}
	if !IsNil(o.SummaryPeriodEnd) {
		toSerialize["summary_period_end"] = o.SummaryPeriodEnd
	}
	return toSerialize, nil
}

type NullableProofOfAddressProperties struct {
	value *ProofOfAddressProperties
	isSet bool
}

func (v NullableProofOfAddressProperties) Get() *ProofOfAddressProperties {
	return v.value
}

func (v *NullableProofOfAddressProperties) Set(val *ProofOfAddressProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableProofOfAddressProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableProofOfAddressProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProofOfAddressProperties(val *ProofOfAddressProperties) *NullableProofOfAddressProperties {
	return &NullableProofOfAddressProperties{value: val, isSet: true}
}

func (v NullableProofOfAddressProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProofOfAddressProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
