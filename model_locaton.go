package onfido_openapi

// Location struct for Location
type Location struct {
	// The IP address of the applicant
	IPAddress *string `json:"ip_address,omitempty"`
	// The 3 character ISO country code of the country where the applicant resides
	CountryOfResidence *string `json:"country_of_residence,omitempty"`
}

// NewLocation instantiates a new Location object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocation() *Location {
	this := Location{}
	return &this
}

// NewLocationWithDefaults instantiates a new Location object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationWithDefaults() *Location {
	this := Location{}
	return &this
}

// GetIPAddress returns the IPAddress field value if set, zero value otherwise.
func (o *Location) GetIPAddress() string {
	if o == nil || o.IPAddress == nil {
		var ret string
		return ret
	}
	return *o.IPAddress
}

// GetIPAddressOk returns a tuple with the IPAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetIPAddressOk() (*string, bool) {
	if o == nil || o.IPAddress == nil {
		return nil, false
	}
	return o.IPAddress, true
}

// HasIPAddress returns a boolean if a field has been set.
func (o *Location) HasIPAddress() bool {
	if o != nil && o.IPAddress != nil {
		return true
	}

	return false
}

// SetIPAddress gets a reference to the given string and assigns it to the IPAddress field.
func (o *Location) SetIPAddress(v string) {
	o.IPAddress = &v
}

// GetCountryOfResidence returns the CountryOfResidence field value if set, zero value otherwise.
func (o *Location) GetCountryOfResidence() string {
	if o == nil || o.CountryOfResidence == nil {
		var ret string
		return ret
	}
	return *o.CountryOfResidence
}

// GetCountryOfResidenceOk returns a tuple with the CountryOfResidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetCountryOfResidenceOk() (*string, bool) {
	if o == nil || o.CountryOfResidence == nil {
		return nil, false
	}
	return o.CountryOfResidence, true
}

// HasCountryOfResidence returns a boolean if a field has been set.
func (o *Location) HasCountryOfResidence() bool {
	if o != nil && o.CountryOfResidence != nil {
		return true
	}

	return false
}

// SetCountryOfResidence gets a reference to the given string and assigns it to the CountryOfResidence field.
func (o *Location) SetCountryOfResidence(v string) {
	o.CountryOfResidence = &v
}